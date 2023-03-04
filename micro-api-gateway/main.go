package main

import (
	apiCmd "github.com/LCY2013/thinking-in-go/paas/micro-api-gateway/cmd"
	"github.com/go-acme/lego/v4/log"
)

func main() {
	err := apiCmd.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
}
