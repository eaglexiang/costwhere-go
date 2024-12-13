package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/eaglexiang/costwhere-go"
)

func main() {
	sig := make(chan struct{})
	go f(sig)
	<-sig
}

func f(sig chan struct{}) {
	defer func() { sig <- struct{}{} }()

	ctx := context.Background()
	ctx, cw := costwhere.Init(ctx, "main")
	defer func() {
		stacks, err := cw.EndWithJSON()
		if err != nil {
			log.Printf("%+v", err)
			return
		}
		err = os.WriteFile("costwhere.json", stacks, 0644)
		if err != nil {
			log.Printf("%+v", err)
		}
	}()

	F0(ctx)
	F1(ctx)
}

func F0(ctx context.Context) {
	defer costwhere.Mark(ctx)()

	time.Sleep(100 * time.Millisecond)
	F2(ctx)
}

func F1(ctx context.Context) {
	defer costwhere.Mark(ctx)()

	time.Sleep(1 * time.Second)
	F2(ctx)
}

func F2(ctx context.Context) {
	defer costwhere.Mark(ctx)()

	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 3; i++ {
		F3(ctx)
	}
}

func F3(ctx context.Context) {
	defer costwhere.Mark(ctx)()

	time.Sleep(300 * time.Millisecond)
}
