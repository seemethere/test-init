# test-init

Have you ever wondered if you can monkey patch other variables in a go project by just dropping a file in
and modifying the variable with an `init()` function?

Well wonder no more because this is a proof of concept that does exactly that!

Expected output:
```
❯ go run main.go
[+] in lib/a.go
[+] in lib/b.go
[+] dummyList before modification: [ello pappet]
[+] in main.go
[!] main executed
[!] executing dummyFunction
[!] dummyList: [ello pappet THIS IS MODIFIED]
```
