package costwhere

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

func Init(ctx context.Context, topic string) (newCtx context.Context, s *CostWhere) {
	// 构造根帧
	newLayer := newStackLayer(topic)
	// 构造栈
	s = newCostWhere(newLayer)

	// 将栈帧写入 ctx
	newCtx = writeThis(ctx, newLayer)

	return
}

type CostWhere struct {
	Root *StackLayer
}

func newCostWhere(root *StackLayer) *CostWhere {
	s := &CostWhere{
		Root: root,
	}
	return s
}

func (s *CostWhere) EndWithJSON() (j []byte, err error) {
	stacks := s.End()
	j, err = json.Marshal(stacks)
	err = errors.WithStack(err)
	return
}

func (s *CostWhere) End() (stacks []string) {
	s.Root.Stop()
	stacks = s.Root.Format("")
	return
}

func Mark(ctx *context.Context, topic string) (end func()) {
	if ctx == nil {
		end = func() {}
		return
	}
	newCtx, end := Begin(*ctx, topic)
	*ctx = newCtx
	return
}

// StartStack 开始一个栈帧
func Begin(ctx context.Context, topic string) (newCtx context.Context, end func()) {
	// 读取父级栈帧
	parent, ok := readThis(ctx)
	if !ok {
		newCtx = ctx
		end = func() {}
		return
	}

	// 写入本级栈帧
	newLayer := newStackLayer(topic)
	parent.AddChild(newLayer)

	// 将本级栈帧写入 ctx
	newCtx = writeThis(ctx, newLayer)
	end = func() { endStack(newCtx) }

	return
}

func endStack(ctx context.Context) {
	// 读取栈帧
	stackLayer, ok := readThis(ctx)
	if !ok {
		return
	}

	// 结束栈帧
	stackLayer.Stop()
}
