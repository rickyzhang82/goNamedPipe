package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var execerPath string

func main() {
	flag.StringVar(&execerPath, "execer", "./execer", "path to execer")
	flag.Parse()

	tmpDir, err := ioutil.TempDir("", uuid.New().String())
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}

	// Create named pipe
	pipeFileName := filepath.Join(tmpDir, "pipeFileName")
	syscall.Mkfifo(pipeFileName, 0600)

	wait := make(chan string)
	go func() {
		cmd := exec.Command(execerPath, pipeFileName)
		// Just to forward the namedPiep
		cmd.Stdout = os.Stdout
		cmd.Run()
		wait <- "done"
	}()

	// Open named pipe for reading
	fmt.Printf("Opening named pipe for reading: %s\n", pipeFileName)
	namedPiep, _ := os.OpenFile(pipeFileName, os.O_RDONLY, 0600)
	fmt.Println("Reading")

	var buff bytes.Buffer
	fmt.Println("Waiting for someone to write something")
	io.Copy(&buff, namedPiep)
	namedPiep.Close()
	fmt.Printf("Data: %s\n", buff.String())
	fmt.Printf("Waiting for exercise is done...\n")
	<-wait
	fmt.Printf("exercise is done!\n")
	os.RemoveAll(tmpDir)
}