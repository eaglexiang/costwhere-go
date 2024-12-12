package costwhere

import (
	"strconv"
	"sync"
	"time"

	"github.com/samber/lo"
)

type StackLayer struct {
	l *sync.Mutex

	name  string
	times TimeSeg

	children []*StackLayer
}

func newStackLayer(name string) *StackLayer {
	return &StackLayer{
		l: new(sync.Mutex),

		name:  name,
		times: newTimeSeg(),

		children: make([]*StackLayer, 0),
	}
}

func (s *StackLayer) AddChild(child *StackLayer) {
	s.l.Lock()
	defer s.l.Unlock()
	s.children = append(s.children, child)
}

func (s *StackLayer) Stop() (end time.Time) {
	end = time.Now()
	s.times.End = end
	return
}

func (s *StackLayer) Format(head string) (lines []string) {
	cost := s.buildCost()
	lines = cost.Format(head)
	return
}

func (s *StackLayer) buildCost() *CostLayer {
	cost := &CostLayer{
		name:     s.name,
		cost:     s.times.Cost(),
		children: make([]*CostLayer, 0, len(s.children)),
	}

	for _, child := range s.children {
		costChild := child.buildCost()
		cost.AddChild(costChild)
	}

	return cost
}

type CostLayer struct {
	name string
	cost time.Duration

	children []*CostLayer
}

func (c *CostLayer) AddChild(child *CostLayer) {
	old, ok := lo.Find(c.children, func(item *CostLayer) bool {
		return item.name == child.name
	})
	if ok {
		old.cost += child.cost
		for _, subChild := range child.children {
			old.AddChild(subChild)
		}
		return
	}

	c.children = append(c.children, child)
}

func (c *CostLayer) Format(head string) (lines []string) {
	if head != "" {
		head += ";"
	}
	head += c.name
	line := head + " " + strconv.FormatInt(c.cost.Milliseconds(), 10)
	lines = append(lines, line)

	for _, chiled := range c.children {
		lines = append(lines, chiled.Format(head)...)
	}

	return
}
