package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flag.Parse()
	namedPipe := flag.Args()[0]

	fmt.Println("Opening named pipe for writing")
	stdout, _ := os.OpenFile(namedPipe, os.O_RDWR, 0600)
	fmt.Println("Writing")
	stdout.Write([]byte("hello, world"))
	stdout.Write([]byte("hello, you"))
	stdout.Close()
	time.Sleep(time.Second * 10)
	os.Exit(1)
}

