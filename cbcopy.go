package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func main() {

	// 1. accept piped input
	data, ioErr := io.ReadAll(os.Stdin)
	if ioErr != nil {
		fmt.Println("Error reading from stdin:", ioErr)
		os.Exit(1)
	}
	if len(data) != 0 {
		clipboard.Write(clipboard.FmtText, data)
		fmt.Println("Copied to the clipboard!")
		os.Exit(0)
	}

	// 2. accept arguments
	args := os.Args[1:]
	if (len(args)) != 0 {
		clipErr := clipboard.Init()
		if clipErr != nil {
			fmt.Println("Error initializing clipboard", clipErr)
			os.Exit(1)
		}
		clipboard.Write(clipboard.FmtText, []byte(strings.Join(args, " ")))
		fmt.Println("Copied to the clipboard!")
		os.Exit(0)
	}

	fmt.Println("Nothing to copy!")
	os.Exit(1)

}
