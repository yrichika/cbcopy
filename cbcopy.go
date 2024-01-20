package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	clipErr := clipboard.Init()
	if clipErr != nil {
		fmt.Println("Error initializing clipboard", clipErr)
		os.Exit(1)
	}

	// 1. accept arguments
	args := os.Args[1:]
	if (len(args)) != 0 {
		clipboard.Write(clipboard.FmtText, []byte(strings.Join(args, " ")))
		fmt.Println("Copied to the clipboard!")
		os.Exit(0)
	}

	piped, ioErr := io.ReadAll(os.Stdin)
	if ioErr != nil {
		fmt.Println("Error reading from stdin:", ioErr)
		os.Exit(1)
	}

	// 2. accept piped input
	if len(piped) != 0 {
		clipboard.Write(clipboard.FmtText, piped)
		fmt.Println("Copied to the clipboard!")
		os.Exit(0)
	}

	fmt.Println("Nothing to copy!")
	os.Exit(1)

}
