package y

import (
	"context"
	"time"

	"github.com/eaglexiang/costwhere-go"
)

func F3(ctx context.Context) {
	defer costwhere.Mark(ctx)()

	time.Sleep(300 * time.Millisecond)
}
