package costwhere

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

func Init(ctx context.Context, topic string) (newCtx context.Context, c *CostWhere) {
	costs := newCosts()

	newCtx = writeThis(ctx, costs)

	startAt := time.Now()
	end := func() {
		cost := time.Since(startAt)
		costs.addCost(5, cost)
	}
	c = newCostWhere(costs, end)

	return
}

type CostWhere struct {
	costs *Costs
	end   func()
}

func newCostWhere(costs *Costs, end func()) *CostWhere {
	s := &CostWhere{
		costs: costs,
		end:   end,
	}
	return s
}

func (s *CostWhere) EndWithJSON() (j []byte, err error) {
	s.end()
	j, err = s.ExportJSON()
	return
}

func (s *CostWhere) ExportJSON() (j []byte, err error) {
	stacks := formatCosts(s.costs)
	j, err = json.Marshal(stacks)
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
