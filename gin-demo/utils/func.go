package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ComputeAge(PbBirthday time.Time) string {
	if PbBirthday.IsZero() || time.Now().Sub(PbBirthday) < 0 {
		return ""
	}
	year := time.Now().Year() - PbBirthday.Year()
	month := time.Now().Month() - PbBirthday.Month()
	day := time.Now().Day() - PbBirthday.Day()
	if day < 0 {
		month = month - 1
	}
	if month < 0 {
		year = year - 1
	}
	if year < 0 {
		year = 0
	}
	return strconv.Itoa(year)
}

func main1() {
	// dt := time.Date(2018, 7, 10, 0, 0, 1, 100, time.Local)
	// fmt.Println(ComputeAge(dt))
	// fmt.Println(dt)
	dt1 := time.Date(2019, 8, 1, 0, 0, 0, 100, time.Local)
	dt2 := time.Date(2019, 7, 27, 0, 0, 0, 100, time.Local)
	// dt2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	fmt.Println("==>", dt1.Sub(dt2).Hours()/24)
	dt2.YearDay()
	// 不用关注时区，go会转换成时间戳进行计算
	fmt.Sprintf("%s-%s", strconv.Itoa(20), dt1.Month())
	// fmt.Println(ComputeAge(dt1))
	fmt.Println(fmt.Sprintf("%s-%d", strconv.Itoa(20), dt1.Month()))
	// fmt.Println( time.Now().Sub(dt1) < 0)
}
