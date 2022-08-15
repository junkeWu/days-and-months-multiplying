package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"gin-demo/redis/test02/cache"
	"github.com/go-redis/redis"
)

type User struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Phone string `json:"phone,omitempty"`
}
type Z struct {
	Score  int
	Member string
}

/*
- 如果只需存储简单的键值对，或者对数字进行递增递减操作，就可以使用**String**存储。
- 如果需要一个简单的分布式队列服务，**List**就可以满足需求。
- 如果出了需要存储键值数据，还想单独对某一字段进行操作，使用**Hash**就非常方便了。
- 如果想得到一个不重复的集合，就可以用**Set**，而且它还可以做并集、差集和交集。
- 如果想实现一个带权重的评论、排行榜列表，可以使用**Sorted Set**。
*/

// 底层数据结构一共有 6 种，分别是简单动态字符串、双向链表、压缩列表、哈希表、跳表和整数数组。
// string -> 简单动态字符串
// List -> 双向链表          List -> 压缩列表
// Hash -> 哈希表            Hash -> 压缩列表
// Sorted Set -> 跳表       Sorted Set -> 压缩列表
// Set -> 整数数组 Set-> 哈希表
// 1.Redis使用一个哈希表保存所有键值对 2.哈希桶中的元素保存的不是值的本身，而是指向具体元素的指针 具体元素都是RedisObject
// 压缩列表实际上类似于一个数组，数组中的每一个元素都对应保存一个数据。 在压缩列表中，如果我们要查找定位第一个元素和最后一个元素，可以通过表头三个字段的长度直接定位，复杂度是 O(1)。而查找其他元素时，就没有这么高效了，只能逐个查找，此时的复杂度就是 O(N) 了。
// 和数组不同的是，压缩列表在表头有三个字段 zlbytes、zltail 和 zllen，分别表示列表长度、列表尾的偏移量和列表中的 entry 个数；压缩列表在表尾还有一个 zlend，表示列表结束。
// 在压缩列表中，如果我们要查找定位第一个元素和最后一个元素，可以通过表头三个字段的长度直接定位，复杂度是 O(1)。而查找其他元素时，就没有这么高效了，只能逐个查找，此时的复杂度就是 O(N) 了。
func main() {

	fmt.Println("===========General string============")
	key := "string:key"
	set := cache.Set(key, "字符串作为Redis最简单的类型，其底层实现只有一种数据结构，就是简单动态字符串（SDS）。")
	if set != nil {
		fmt.Println("缓存设置错误")
	}
	value, err := cache.Get(key)
	if err != nil {
		fmt.Println("get 缓存出错")
	}
	fmt.Printf("获取到缓存值: %s\n", value)
	fmt.Println("===========struct string============")
	user := User{
		Name:  "xiaobai",
		Age:   19,
		Phone: "138828281",
	}
	data, _ := json.Marshal(user)

	hash := sha1.New()
	hash.Write(data)
	hashed := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("hashed", hashed)
	if err := cache.Set(hashed, string(data)); err != nil {
		fmt.Println("缓存出错")
	}
	value, err = cache.Get(hashed)
	if err != nil {
		fmt.Println("get cache error")
	}
	fmt.Printf("获取到缓存值: %s \n", value)
	var userResp User
	json.Unmarshal([]byte(value), &userResp)
	fmt.Println("user Get", userResp)

	// List本身是按先进先出的顺序对数据进行存取的，可以通过Lpush从一端存入数据，通过RPop从一端消费数据。
	// 同时为了解决RPop在消费数据解决while(1)循环，导致消费者CPU一直消耗，Redis引入了新的方法BRPop，及阻塞式读取，客户端在没有读取到队列数据时，
	// 自动阻塞，直到有新的数据写入队列，在开始读取新数据。
	// 我们在使用List类型时需要注意一个问题，及生产速度大于消费速度，这样会导致List中的数据越来越多，给Redis的内存带来很大压力，所以我们在使用List类型时需要考虑生产消费的能力。
	// 这里我们重点将几个常用的方法Lpush、Rpop、BRpop、LLen、LRange;
	fmt.Println("===========List============")

	keyList := "string:list"
	err = cache.LPush(keyList, "A", "B", "C", 20, "D", "E", "F")
	if err != nil {
		fmt.Println("缓存设置错误", err)
	}
	lLen := cache.RedisCache.LLen(keyList)
	lRange := cache.RedisCache.LRange(keyList, 0, 5)

	// Redis Lrem 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素。
	// COUNT 的值可以是以下几种：
	// count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
	// count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
	// count = 0 : 移除表中所有与 VALUE 相等的值。
	cache.RedisCache.LRem(keyList, 0, "D")
	fmt.Println("len", lLen)
	fmt.Println("lRange", lRange)

	value, err = cache.RPop(keyList)
	if err != nil {
		fmt.Println("get 缓存出错")
	}
	// for {
	// 	// 阻塞式读取
	// 	value := cache.RedisCache.BRPop(time.Second*5, keyList)
	// 	if value.Err() != nil {
	// 		fmt.Println("get 缓存出错", err)
	// 		break
	// 	}
	fmt.Printf("获取到缓存值: %s\n", value)
	// }
	fmt.Println("over")

	// 用于同时将多个 field-value (字段-值)对设置到哈希表中,此方法会覆盖哈希表中已存在的字段。如果哈希表不存在，会创建一个空哈希表。
	fmt.Println("===========Hash============")
	hKey := "string:hash"
	cache.RedisCache.HSet(hKey, "name", "xiaohong")
	cache.RedisCache.HSet(hKey, "age", 19)
	cache.RedisCache.HSet(hKey, "phone", "119")
	all := cache.RedisCache.HGetAll(hKey)
	fmt.Println("hash get:", all)
	// 修改已存在的字段
	cache.RedisCache.HSet(hKey, "name", "李四")
	// 获取指定字段
	name := cache.RedisCache.HGet(hKey, "name")
	fmt.Println("获取指定字段", name.String())
	existsName := cache.RedisCache.HExists(hKey, "name")
	existsId := cache.RedisCache.HExists(key, "id")
	result, _ := existsName.Result()
	result2, _ := existsId.Result()
	fmt.Printf("name 字段是否存在 %v\n", result)
	fmt.Printf("id 字段是否存在 %v\n", result2)
	fmt.Println("执行移除name")
	cache.RedisCache.HDel(hKey, "name")
	existsName2 := cache.RedisCache.HExists(hKey, "name")
	result3, _ := existsName2.Result()
	fmt.Printf("name 字段是否存在 %v\n", result3)

	// Set 是 String 类型的无序集合。集合成员是唯一的，这就意味着集合中不能出现重复的数据。
	// SET常用方法：
	// SADD：向集合添加一个或多个成员
	// SCard: 获取集合的成员数
	// SMembers：获取集合的所有成员
	// SRem: 移除集合里的某个元素
	// SPop: 移除并返回set的一个随机元素(SET是无序的)
	// SDiff: 返回第一个集合与其他集合之间的差异。
	// SDIFFSTORE： 返回给定所有集合的差集并存储在 destination 中
	// SInter: 返回所有给定集合的交集
	// Sunion: 返回所有给定集合的并集
	fmt.Println("===========Set============")

	keySet := "string:set"
	cache.RedisCache.SAdd(keySet, "phone")
	err2 := cache.RedisCache.SAdd(keySet, "hahh").Err()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 获取全部hash对象
	allSet := cache.RedisCache.SCard(keySet)
	fmt.Println(allSet.Result())
	members := cache.RedisCache.SMembers(keySet)
	fmt.Println(members)

	// Redis 有序集合和集合一样也是 string 类型元素的集合,且不允许重复的成员。
	// 不同的是每个元素都会关联一个 double 类型的分数。redis 正是通过分数来为集合中的成员进行从小到大的排序。
	// 有序集合的成员是唯一的,但分数(score)却可以重复。
	// sorted set 函数有很多，这里我们主要演示几个函数，通过sorted set 查看编程语言排行榜热度。
	fmt.Println("=========== Sorted Set============")
	keySortSet := "string:zset"
	z := []redis.Z{
		{Score: 80, Member: "Java"},
		{Score: 90, Member: "Python"},
		{Score: 95, Member: "Golang"},
		{Score: 98, Member: "PHP"},
	}
	zAddResp := cache.RedisCache.ZAdd(keySortSet, z...)
	if err != nil {
		fmt.Println(err)
	}
	scores := cache.RedisCache.ZRevRangeWithScores(keySortSet, 0, 2)
	fmt.Println(zAddResp)
	cache.RedisCache.ZIncrBy(keySortSet, 5, "Golang")
	scores = cache.RedisCache.ZRevRangeWithScores(keySortSet, 0, 2)
	fmt.Println("加分后----")
	fmt.Println(scores)
}
