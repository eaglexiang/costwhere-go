package costwhere

import "context"

type ctxKeyType string

const (
	thisLayerKey ctxKeyType = "thisLayerCtx"
)

func readThis(ctx context.Context) (s *StackLayer, ok bool) {
	val := ctx.Value(thisLayerKey)
	s, ok = val.(*StackLayer)
	return
}

func writeThis(ctx context.Context, s *StackLayer) (dst context.Context) {
	dst = context.WithValue(ctx, thisLayerKey, s)
	return
}
