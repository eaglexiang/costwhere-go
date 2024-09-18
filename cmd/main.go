package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	src := flag.String("if", "", "source file")
	flag.Parse()

	if *src == "" {
		fmt.Println("source file cannot be empty")
		return
	}

	cmd := fmt.Sprintf("cat %s | flamegraph.pl > flamegraph.svg", *src)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		fmt.Println(err)
	}
}
