package costwhere

import (
	"context"
)

type ctxKeyType string

const (
	thisLayerKey ctxKeyType = "thisLayerCtx"
)

func readThis(ctx context.Context) (c *Costs, ok bool) {
	val := ctx.Value(thisLayerKey)
	c, ok = val.(*Costs)
	return
}

func writeThis(ctx context.Context, c *Costs) (dst context.Context) {
	dst = context.WithValue(ctx, thisLayerKey, c)
	return
}
