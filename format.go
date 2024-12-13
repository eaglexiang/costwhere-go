package costwhere

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
)

func formatCosts(c *Costs) (costs []string) {
	// 重建栈结构
	s := newCostStack()
	for _, cost := range c.costs {
		path := strings.Split(cost.Path, ";")
		s.add(path, cost.Cost)
	}

	costs = s.root.format("")

	return
}

type costStack struct {
	root *stackFrame
}

func newCostStack() *costStack {
	return &costStack{}
}

func (c *costStack) add(path []string, cost time.Duration) {
	if len(path) == 0 {
		return
	}

	root := path[0]

	if c.root == nil {
		c.root = newStackFrame(root)
		return
	}

	c.root.add(path[1:], cost)
}

type stackFrame struct {
	path     string
	cost     time.Duration
	children []*stackFrame
}

func newStackFrame(path string) *stackFrame {
	return &stackFrame{
		path:     path,
		children: make([]*stackFrame, 0),
	}
}

func (s *stackFrame) add(childPath []string, cost time.Duration) {
	if len(childPath) == 0 {
		s.cost += cost
		return
	}

	childRoot := childPath[0]

	child, ok := lo.Find(s.children, func(item *stackFrame) bool {
		return item.path == childRoot
	})
	if ok {
		child.add(childPath[1:], cost)
		return
	}

	newChild := newStackFrame(childRoot)
	s.children = append(s.children, newChild)
	newChild.add(childPath[1:], cost)

	if len(childPath) == 1 {
		s.cost -= cost
	}
}

func (s *stackFrame) format(head string) (costs []string) {
	path := s.path
	if head != "" {
		path = head + ";" + path
	}

	cost := fmt.Sprintf("%s %d", head, s.cost.Milliseconds())
	costs = append(costs, cost)

	for _, child := range s.children {
		childCosts := child.format(path)
		costs = append(costs, childCosts...)
	}

	return
}
