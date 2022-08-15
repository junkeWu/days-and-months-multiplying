package utils

import (
	"fmt"
	"strconv"
	"time"
)

const (
	TimeOffset = 8 * 3600  // 8 hour offset
	HalfOffset = 12 * 3600 // Half-day hourly offset
)

func GetCurDayZeroTimestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix() - TimeOffset
}

func main() {
	fmt.Println("-->", GetCurDayZeroTimestamp())
	itoa := strconv.Itoa(int(GetCurDayZeroTimestamp()))
	parse, _ := time.Parse("2006-01-02", itoa)
	fmt.Println("---东八", parse)
	fmt.Println("---东八", time.Now())
}
