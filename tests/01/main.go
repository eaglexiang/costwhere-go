package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eaglexiang/costwhere-go"
)

func main() {
	ctx := context.Background()
	ctx, cw := costwhere.Init(ctx, "main")
	defer func() {
		stacks := cw.End()
		buf := bytes.NewBuffer(nil)
		for _, line := range stacks {
			buf.WriteString(line + "\n")
		}
		err := os.WriteFile("costwhere.log", buf.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
		}
	}()

	F0(ctx)
	F1(ctx)
}

func F0(ctx context.Context) {
	defer costwhere.Mark(&ctx, "F0")()

	time.Sleep(100 * time.Millisecond)
	F2(ctx)
}

func F1(ctx context.Context) {
	defer costwhere.Mark(&ctx, "F1")()

	time.Sleep(1 * time.Second)
	F2(ctx)
}

func F2(ctx context.Context) {
	defer costwhere.Mark(&ctx, "F2")()

	time.Sleep(100 * time.Millisecond)
	F3(ctx)
	F3(ctx)
	F3(ctx)
}

func F3(ctx context.Context) {
	defer costwhere.Mark(&ctx, "F3")()

	time.Sleep(300 * time.Millisecond)
}
