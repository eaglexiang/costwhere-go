package costwhere

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

func Init(ctx context.Context, opts ...InitOption) (newCtx context.Context, c *CostWhere) {
	compressPath := true
	var pathDict map[string]string
	if len(opts) > 0 {
		opt := opts[0]

		if opt.CompressPath != nil {
			compressPath = *opt.CompressPath
		}
		pathDict = opt.PathDict
	}

	costs := newCosts(compressPath, pathDict)

	newCtx = writeThis(ctx, costs)

	startAt := time.Now()
	path := getStackInfo(3)
	parentPath := getStackInfo(4)

	end := func() {
		cost := time.Since(startAt)
		text := formatStackInfo(path, compressPath, pathDict)
		costs.addCostWithPath(text, cost)
	}

	text := formatStackInfo(parentPath, compressPath, pathDict)
	c = newCostWhere(costs, end, text)

	return
}

type CostWhere struct {
	costs      *Costs
	end        func()
	parentPath string
}

func newCostWhere(costs *Costs, end func(), parentPath string) *CostWhere {
	s := &CostWhere{
		costs:      costs,
		end:        end,
		parentPath: parentPath,
	}
	return s
}

func (s *CostWhere) EndWithJSON() (j []byte, err error) {
	s.end()
	j, err = s.ExportJSON()
	return
}

func (s *CostWhere) Export() (output Output) {
	stacks := formatCosts(s.costs, s.parentPath)

	output = Output{
		Stacks: stacks,
	}

	return
}

func (s *CostWhere) ExportJSON() (j []byte, err error) {
	output := s.Export()

	j, err = json.Marshal(output)
	err = errors.WithStack(err)

	return
}

func Mark(ctx context.Context) (end func()) {
	if ctx == nil {
		end = func() {}
		return
	}

	costs, ok := readThis(ctx)
	if !ok {
		end = func() {}
		return
	}

	startAt := time.Now()
	end = func() {
		cost := time.Since(startAt)
		costs.addCost(4, cost)
	}

	return
}

type InitOption struct {
	CompressPath *bool // 是否压缩路径，默认为 true
	PathDict     map[string]string
}

type Output struct {
	Stacks []string
}
