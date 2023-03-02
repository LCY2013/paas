package main

import (
	apiCmd "api/cmd"
	"github.com/go-acme/lego/v4/log"
)

func main() {
	err := apiCmd.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
}
