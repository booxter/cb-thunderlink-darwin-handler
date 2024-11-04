package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "thunderlink.h"
*/
import "C"

import (
	"log"
	"os"
	"os/exec"
	"time"
)

var urlListener chan string = make(chan string)
var commandPath = "cb_thunderlink"

func main() {
	go func() {
		timeout := time.After(4 * time.Second)
		f, err := os.Create("/tmp/file.txt")
		if err != nil {
			log.Fatal(err)
		}
		select {
		case url := <-urlListener:
			cmd := exec.Command(commandPath, url)
			f.WriteString(commandPath + " " + url + "\n")

			cmd.Run()
			f.WriteString("success\n")
			os.Exit(0)
		case <-timeout:
			f.WriteString("failed\n")
			os.Exit(1)
		}
	}()

	C.RunApp()
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
