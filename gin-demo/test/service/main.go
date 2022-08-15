package main

import (
	"fmt"
	"os"
	"time"

	_ "gin-demo/test/service/service2"
	"gin-demo/utils"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Println(utils.GetSelfFuncName())

	for {
		if CheckReliabilityResult() {
			log.Warn("", "CheckReliabilityResult ok, exit")
			os.Exit(0)
			return
		} else {
			log.Warn("", "CheckReliabilityResult failed , wait.... ")
		}
		time.Sleep(time.Duration(300) * time.Second)
	}
}
func CheckReliabilityResult() bool {
	return true
}
