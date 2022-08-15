package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
)

func main() {
	// 直接生成二维码
	_ = qrcode.WriteFile("http://www.baidu.com/?id = 2", qrcode.Medium, 256, "./temp/png.png")
	// 一般保存到库里，用io流的方式
	code, filename, err := generateQRCode(111, "www.baidu.com/")
	if err != nil {
		panic(err)
		return
	}
	err = save("./temp/", code, filename)
	if err != nil {
		log.Fatal(err)
	}
}

// 保存到某一路径
func save(path string, code io.Reader, filename string) error {
	// 二维码都是小文件，不需要bufio
	data, err := ioutil.ReadAll(code)
	if err != nil {
		return errors.Wrap(err, "readAll failed")
	}
	err = ioutil.WriteFile(path+filename, data, 0666)
	if err != nil {
		return errors.Wrap(err, "writeFile failed")
	}
	return nil
}

// 生成二维码
func generateQRCode(suffix int64, url string) (io.Reader, string, error) {
	var err error
	id := strconv.Itoa(int(suffix))
	encode, err := qrcode.Encode(url+id, qrcode.High, 256)
	if err != nil {
		return nil, "", err
	}
	reader := bytes.NewBuffer(encode)
	obj := md5.New()
	_, err = obj.Write(encode)
	if err != nil {
		return nil, "", err
	}
	return reader, hex.EncodeToString(obj.Sum(nil)) + ".png", nil
}
