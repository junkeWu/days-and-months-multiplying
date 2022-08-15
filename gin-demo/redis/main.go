package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type User struct {
	name  string
	age   int
	class string
}

func main() {
	fmt.Println("===========string=========")
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	conn.Do("set", "a1", "hello world")
	// send会存进缓存区
	conn.Send("get", "a2")
	// conn.Send("set", "a2", "bbb") // 这一步把命令发送到缓冲区，并没有执行。
	conn.Flush()
	rb, _ := redis.String(conn.Receive()) // 强转操作
	fmt.Println("receive", rb)
	fmt.Println("===========set=========")
	var user User
	resp, _ := redis.Values(conn.Do("mget", "name", "age", "class"))
	scan, _ := redis.Scan(resp, &user.name, &user.age, &user.class)
	fmt.Println("scan", scan)
	fmt.Println("mget", user)
	EncodeFunc()
}

// golang 操作数据库分三类函数：1. 连接  2. 执行（do、send） 3. 强转操作（bool、bytes、int64map）
// 使用字节流进行编码和解码
func EncodeFunc() {
	var src string
	dest := "hello world"
	var buffer bytes.Buffer        // 容器
	enc := gob.NewEncoder(&buffer) // 编码器
	enc.Encode(dest)
	fmt.Println("dest", buffer)

	// 解码
	dec := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	dec.Decode(&src)
	fmt.Println("src", src)
}
