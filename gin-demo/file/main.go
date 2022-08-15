package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	DownloadFile2()
	// WriteFile()
	// ReadFile()
}

// io.copy()可以直接copy指针，避免大量数据读到内存。
func DownloadFile2() {
	url := "https://junkewu.github.io/images/slice01.png"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}
	defer resp.Body.Close()

	os.Mkdir("./temp", 0777)
	newFile, err := os.Create("./temp/slice01.png")
	if err != nil {
		panic(err)
	}
	// wt := bufio.NewWriter(newFile)
	defer newFile.Close()
	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		panic(err)
	}
	// wt.Flush()
}

func DownloadFile1() {
	url := "https://junkewu.github.io/images/slice01.png"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", url)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile("./temp/slice01.png", data, 0755)
}

func WriteFile() {
	file, err := os.OpenFile("./temp/hello.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteRune('我')
		writer.WriteRune('爱')
		writer.WriteRune('你')
		writer.WriteString("\n")
	}
	writer.Flush()
}
func ReadFile() {
	file, err := os.Open("./temp/hello.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println("read", readString)
	}
}
