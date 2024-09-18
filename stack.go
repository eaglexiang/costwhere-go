package costwhere

import (
	"strconv"
	"sync"
	"time"
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
	if head != "" {
		head += ";"
	}
	head += s.name
	line := head + " " + strconv.FormatInt(s.times.Cost().Milliseconds(), 10)
	lines = append(lines, line)

	for _, chiled := range s.children {
		lines = append(lines, chiled.Format(head)...)
	}

	return
}
