package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "ww_junke@mails.tsinghua.edu.cn"
	str2 := "junke@mail.tsinghua.edu.cn"
	matched, err := regexp.MatchString("\\w[-\\w.+]*@mails.tsinghua.edu.cn", str)
	matched2, err := regexp.MatchString("\\w[-\\w.+]*@mails.tsinghua.edu.cn", str2)
	fmt.Println("后缀固定邮箱的正则正确否", matched, err)
	fmt.Println("后缀固定邮箱的正则正确否", matched2, err)

	str3 := "53010219200508011X"
	matched3, err := regexp.MatchString("^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$", str3)
	fmt.Println("身份证的正则正确否：", matched3, err)

	reg := "^[A-Za-z]?[\\d]{10}$"

	str4 := "S1133332222"
	matched4, err := regexp.MatchString(reg, str4)
	fmt.Println("S1133332222学号的正则正确否：", matched4, err)
	str5 := "1133332222"
	matched5, err := regexp.MatchString(reg, str5)
	fmt.Println("1133332222学号的正则正确否：", matched5, err)

	fmt.Println("===========错误的case")
	str6 := "S11133332222"
	matched6, err := regexp.MatchString(reg, str6)
	fmt.Println("S11133332222学号的正则正确否：", matched6, err)
	str7 := "1333322"
	matched7, err := regexp.MatchString(reg, str7)
	fmt.Println("133332222学号的正则正确否：", matched7, err)
}
