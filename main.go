package main

import (
	"fmt"
	"os"
)

func main() {
	msg, code := run()
	fmt.Fprint(os.Stderr, msg)
	os.Exit(code)
}

func run() (string, int) {
	return "", 0
}
