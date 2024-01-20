package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 1. accept arguments
	args := os.Args[1:]
	if len(args) != 0 {
		copyToClipboard(strings.Join(args, " "))
		fmt.Println("Copied to the clipboard!")
		os.Exit(0)
	}

	// 2. accept piped input
	piped, ioErr := io.ReadAll(os.Stdin)
	if ioErr != nil {
		fmt.Println("Error reading from stdin:", ioErr)
		os.Exit(1)
	}

	if len(piped) != 0 {
		copyToClipboard(string(piped))
		fmt.Println("Copied to the clipboard!")
		os.Exit(0)
	}

	fmt.Println("Nothing to copy!")
	os.Exit(1)
}

func copyToClipboard(text string) {
	cmd := exec.Command("clip.exe")
	in, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		os.Exit(1)
	}

	go func() {
		defer in.Close()
		io.WriteString(in, text)
	}()

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running clip. Sorry, This command only works for WSL:", err)
		os.Exit(1)
	}
}
