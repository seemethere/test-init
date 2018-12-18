package lib

import "fmt"

var dummyList = []string{"ello", "pappet"}

func init() {
	fmt.Println("[+] in lib/a.go")
}

func DummyFunction() {
	fmt.Println("[!] executing dummyFunction")
	fmt.Printf("[!] dummyList: %+v\n", dummyList)
}
