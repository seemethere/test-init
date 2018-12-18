package main

import (
	"fmt"
	"github.com/seemethere/test-init/lib"
)

func init() {
	fmt.Println("[+] in main.go")
}

func main() {
	fmt.Println("[!] main executed")
	lib.DummyFunction()
}
