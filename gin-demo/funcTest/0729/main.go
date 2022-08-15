package main

import (
	"errors"
	"fmt"
	"strings"
)

type PersonInfo struct {
	code          string
	unifyID       string "json:zjh" // 学号/工作证号
	internetID    string // 网络号
	name          string
	personType    string
	serviceCenter string
	email         string
}


func main() {
	fmt.Println("test start")
	//var data []string
	var item item
	strs := []string{"code=0", "zjh=2017920018", "yhm=wangjianmei", "xm=王健美", "yhlb=J0000", "dw=训练中心", "email=wangjianmei@mail.tsinghua.edu.cn"}
	for _, str := range strs {
		item = append(item, strings.Split(str, "=")[1])
	}
	personInfo := PersonInfo{
		code:          item.getItem(0),
		unifyID:       item.getItem(1),
		internetID:    item.getItem(2),
		name:          item.getItem(3),
		personType:    item.getItem(4),
		serviceCenter: item.getItem(5),
		email:         item.getItem(6),
	}
	fmt.Println("======personInfo=====", personInfo)
}
type item []string
func (i item) getItem(idx int) string {
	if len(i)-1 < idx {
		panic(errors.New("array out of bounds"))
	}
	return i[idx]
}
