package main

import "github.com/eaglexiang/costwhere-go/tests/01/a"

func main() {
	sig := make(chan struct{})
	go a.A(sig)
	<-sig
}
