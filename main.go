package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "thunderlink.h"
*/
import "C"

import (
	"os"
	"os/exec"
	"time"
)

var urlListener chan string = make(chan string)
var commandPath = "cb_thunderlink"

func main() {
	go func() {
		timeout := time.After(4 * time.Second)
		select {
		case url := <-urlListener:
			cmd := exec.Command(commandPath, url)

			cmd.Run()
			os.Exit(0)
		case <-timeout:
			os.Exit(1)
		}
	}()

	C.RunApp()
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
