package lib

import "fmt"

func init() {
	fmt.Printf("[+] in lib/b.go\n")
	fmt.Printf("[+] dummyList before modification: %+v\n", dummyList)
	dummyList = append(dummyList, "THIS IS MODIFIED")
}
