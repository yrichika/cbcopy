package main

import (
	"fmt"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	args := os.Args[1:]

	if (len(args)) == 0 {
		fmt.Println("Nothing to copy!")
		os.Exit(1)
	}

	fmt.Println("Copied to the clipboard!")
	err := clipboard.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	clipboard.Write(clipboard.FmtText, []byte(strings.Join(args, " ")))

	os.Exit(0)

}
