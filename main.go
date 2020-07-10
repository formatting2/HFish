package main

import (
	"HFish/core/rpc/client"
	"HFish/utils/setting"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args
	go func() {
		time.Sleep(60 * time.Second)
		client.SEND_MESSAGE = true
	}()
	if args == nil || len(args) < 2 {
		setting.Help()
	} else {
		if args[1] == "help" || args[1] == "--help" {
			setting.Help()
		} else if args[1] == "init" || args[1] == "--init" {
			setting.Init()
		} else if args[1] == "version" || args[1] == "--version" {
			fmt.Println("v0.6.1")
		} else if args[1] == "run" || args[1] == "--run" {
			setting.Run()
		} else {
			setting.Help()
		}
	}
}
