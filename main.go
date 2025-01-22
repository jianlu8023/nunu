package main

import (
	"fmt"

	"gitee.com/jianlu8023/nunu/cmd/nunu"
)

func main() {
	err := nunu.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
