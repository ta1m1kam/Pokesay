package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if terminal.IsTerminal(syscall.Stdin) {
		// Execute: go run main.go
		fmt.Print("Type something then press the enter key: ")
		var stdin string
		fmt.Scan(&stdin)
		fmt.Printf("Result: %s\n", stdin)
		return
	}

	// Execute: echo "foo" | go run main.go
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))
}
