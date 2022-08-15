package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type mapResp struct {
	Address  string `json:"address,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串 stu01
		value := rand.Intn(100)          // 生成0~99的随机整数
		scoreMap[key] = value
	}
	// 取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	var map01 mapResp
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
		if value != nil {
			marshal, err := json.Marshal(value)
			if err != nil {
				panic(err)
				return
			}
			fmt.Println("--->marshal", marshal)
			json.Unmarshal(marshal, &map01)
			fmt.Println("---->map01", map01)
		}
	}
	println("==============array==============")
	var map02 = make(map[[2]int]int, 10)
	map02[[2]int{1, 1}] = 11
	map02[[2]int{1, 2}] = 12
	fmt.Println("---map02---", map02)

	fmt.Println("==============value ==============")
	m0 := make(map[string][]int, 2)
	m0["1"] = []int{1, 2, 3, 4, 5}
	m0["2"] = []int{2, 3, 4, 5, 6, 7}
	fmt.Println("m0", m0)

	m1 := make(map[int]map[int]int, 10)
	fmt.Println("---->m1", m1)
	m1[1] = map[int]int{1: 1}
	fmt.Println("---->m1", m1)
	fmt.Println("---->m1", m1)

	println("===============复制===============")

}
