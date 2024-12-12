package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	src := flag.String("if", "", "source file")
	flag.Parse()

	if *src == "" {
		fmt.Println("source file cannot be empty")
		return
	}
	filename := *src

	ext := filepath.Ext(filename)

	var err error
	switch ext {
	case ".json":
		err = jsonDraw(filename)
	default:
		err = defaultDraw(filename)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func defaultDraw(filename string) (err error) {
	cmd := fmt.Sprintf("cat %s | flamegraph.pl > flamegraph.svg", filename)
	err = exec.Command("bash", "-c", cmd).Run()
	return
}

func jsonDraw(filename string) (err error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	lines := []string{}
	err = json.Unmarshal(buf, &lines)
	if err != nil {
		return
	}

	text := strings.Join(lines, "\n")
	err = os.WriteFile("tmp.log", []byte(text), os.ModePerm)
	if err != nil {
		return
	}
	defer os.Remove("tmp.log")

	err = defaultDraw("tmp.log")

	return
}
