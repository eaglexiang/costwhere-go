package costwhere

import (
	"fmt"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func getStackInfo(skip int) (pathText string) {
	const maxDepth = 50

	pcs := make([]uintptr, maxDepth)
	depth := runtime.Callers(skip, pcs[:])
	frames := runtime.CallersFrames(pcs[:depth])

	stacks := make([]runtime.Frame, 0)
	for f, again := frames.Next(); again; f, again = frames.Next() {
		stacks = append(stacks, f)
	}

	path := make([]string, 0, len(stacks))
	for _, stack := range stacks {
		filename := filepath.Base(stack.File)
		frame := fmt.Sprintf("%s/%s", filename, stack.Function)
		path = append(path, frame)
	}

	slices.Reverse(path)

	pathText = strings.Join(path, ";")

	return
}
