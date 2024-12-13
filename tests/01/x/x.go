package x

import (
	"context"
	"time"

	"github.com/eaglexiang/costwhere-go"
	"github.com/eaglexiang/costwhere-go/tests/01/y"
)

func F2(ctx context.Context) {
	defer costwhere.Mark(ctx)()

	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 3; i++ {
		y.F3(ctx)
	}
}
