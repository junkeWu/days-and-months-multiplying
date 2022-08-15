package main

import "fmt"

func main() {
	s01 := []int{1, 2, 3}

	s02 := append(s01, []int{4, 5, 6, 7}...)

	fmt.Println("cap :", cap(s02), "len:", len(s02))
	// 四种定义方式
	var str01 = []int{1, 2, 3, 4}
	str02 := []string{"1", "2", "3"}
	str03 := make([]int, 3)
	var str04 []int
	str04 = append(str04, 1)
	fmt.Println("str01", str01)
	fmt.Println("str02", str02)
	fmt.Println("str03", len(str03))
	fmt.Println("str03", cap(str03))
	fmt.Println("str03", str03)
	fmt.Println("str04", str04)
	slice05 := []int{1, 2, 3, 4}
	fmt.Printf("slice05: %p, len: %d", &slice05[0], len(slice05))
	fmt.Println()
	slice05 = append(slice05, 4, 5)
	fmt.Printf("slice05: %p", &slice05[0])
	fmt.Println()
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("--->a:%p", &a)
	fmt.Println("a 的值", a)
	// s1 := a[:] // len=3 cap=5
	// // s1 = append(s1, 1)
	// fmt.Printf("--->s1:%p", &s1[0])
	// // fmt.Println("--->s2", s2)
	fmt.Println("---------------------")
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(&arr[0])

	println("=====================")

	var s1 = arr[:4]
	fmt.Println(&s1[0], " len: ", len(s1), " cap: ", cap(s1))

	s1 = append(s1, 1, 2, 3, 4, 5, 6)
	fmt.Println(&s1[0], " len: ", len(s1), " cap: ", cap(s1))

	println("=====================")
	fmt.Println(arr)
	println("........二维数组.....")

	res := make([][]int, 4)
	for i := range res {
		res[i] = make([]int, 4)
	}
	println("res", res)
	for _, val := range res {
		for _, val2 := range val {
			fmt.Print(val2)
		}
		fmt.Println()
	}
	var (
		k, n = 0, 3
	)
	for _, val := range res {
		for j, _ := range val {
			if j == k {
				val[j] = 1
			}
			if j == n {
				val[j] = 1
			}
		}
		k++
		n--
	}
	for _, val := range res {
		for _, val2 := range val {
			fmt.Print(val2)
		}
		fmt.Println()
	}
}

// 函数传递为值传递
func add1(s []int, x int) []int {
	s = append(s, x)
	fmt.Printf("add %p", &s)
	fmt.Println()
	return s
}
func add2(s []int, x int) {
	s = append(s, x)
}
