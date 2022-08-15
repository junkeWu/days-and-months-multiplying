package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func GetSelfFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return cleanUpfuncName(runtime.FuncForPC(pc).Name())
}
func cleanUpfuncName(funcName string) string {
	fmt.Println(funcName)
	end := strings.LastIndex(funcName, ".")
	if end == -1 {
		return ""
	}
	return funcName[end+1:]
}
