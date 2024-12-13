package costwhere

import (
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
	for f, ok := frames.Next(); ok; f, ok = frames.Next() {
		stacks = append(stacks, f)
	}

	path := make([]string, 0, len(stacks))
	for _, frame := range stacks {
		path = append(path, frame.Function)
	}

	slices.Reverse(path)

	pathText = strings.Join(path, ";")

	return
}
