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
	dst := flag.String("of", "flamegraph.svg", "output file")
	flag.Parse()

	if *src == "" {
		fmt.Println("source file cannot be empty")
		return
	}
	if *dst == "" {
		fmt.Println("output file cannot be empty")
		return
	}

	srcFile := *src

	ext := filepath.Ext(srcFile)

	var err error
	switch ext {
	case ".json":
		err = jsonDraw(srcFile, *dst)
	default:
		err = defaultDraw(srcFile, *dst)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func defaultDraw(srcFile string, dstFile string) (err error) {
	cmd := fmt.Sprintf("cat %s | flamegraph.pl > %s", srcFile, dstFile)
	err = exec.Command("bash", "-c", cmd).Run()
	return
}

func jsonDraw(srcFile string, dstFile string) (err error) {
	buf, err := os.ReadFile(srcFile)
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

	err = defaultDraw("tmp.log", dstFile)

	return
}
