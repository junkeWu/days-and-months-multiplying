package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

type Data struct {
	Data string `json:"data"`
}

func main() {
	type Person struct {
		name string
		age  int
	}
	p1 := &Person{name: "wbw", age: 18}
	fmt.Println((*p1).name)
	fmt.Println(p1.name) // 隐式解引用

	p2 := Person{name: "wbw", age: 18}
	fmt.Println(p2.name)

	p3 := p1
	fmt.Println(p3.name)
	p4 := p2
	fmt.Println(p4.name)

	p3.name = "sdfsd"
	fmt.Println((*p1).name)
	fmt.Println(p1.name)
	fmt.Println(p3.name)

	p4.name = "sdfsdss"
	fmt.Println(p2.name)
	fmt.Println(p4.name)
	// bodyBytes, err := json.Marshal(map[string]string{
	// 	"condition": "{\"op\":\"&\",\"c\":[{\"l\":{\"f\":\"identity_id\",\"op\":\"=\",\"v\":\""+"22111"+"\"}}]}",
	// })
	// fmt.Println(string(bodyBytes),err)
	// fmt.Println(time.Now().Format("_060102-030405"))
	// fmt.Println(time.Now().Format("2006/01/02"))
	//
	parse, _ := time.Parse("2006", "2018")
	fmt.Println("p", parse.Format("2006-01-02"))
	var (
		data        []*Data
		changedData Data
	)
	s := `[{"data": "wenjuuan"},{"data": "https://data.sigs.tsinghua.edu.cn/minio-cubic-survey/private/2022/02/13/6G6QwqFyJMwnNZXV8pyjVf.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=admin%2F20220213%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20220213T034715Z&X-Amz-Expires=43199&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3D%224C6D61A4-802F-4549-8670-336191D8C783.png%22%3B%20filename%2A%3DUTF-8%27%274C6D61A4-802F-4549-8670-336191D8C783.png&X-Amz-Signature=9564d6a83afa21568c98919b3eda4d181cd58b51b2d94d2dd7a5be5899c9778d"},{"data": "https://data.sigs.tsinghua.edu.cn/minio-cubic-survey/private/2022/02/13/14124124124.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=admin%2F20220213%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20220213T034715Z&X-Amz-Expires=43199&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3D%224C6D61A4-802F-4549-8670-336191D8C783.png%22%3B%20filename%2A%3DUTF-8%27%274C6D61A4-802F-4549-8670-336191D8C783.png&X-Amz-Signature=9564d6a83afa21568c98919b3eda4d181cd58b51b2d94d2dd7a5be5899c9778d"}]`
	_ = json.Unmarshal([]byte(s), &data)
	fmt.Println("-=--->", data)
	var re = regexp.MustCompile(`(?m)(20..\/..\/..\/[^\?]+)`)
	for index, i := range data {
		temp := i
		for j, match := range re.FindAllString(temp.Data, -1) {
			fmt.Println(match, "found at index", j)
			str := "heep"
			if index == 1 {
				str = "xxx"
			}
			changedData.Data = str
			data[index] = &changedData
			fmt.Println(index, "changed", data[index])
		}
	}

}
