package costwhere

import (
	"sync"
	"time"
)

type Costs struct {
	mu    *sync.Mutex
	costs map[string]*Cost

	compressPath bool
	pathDict     map[string]string
}

func newCosts(compressPath bool, pathDict map[string]string) *Costs {
	if pathDict == nil {
		pathDict = make(map[string]string)
	}

	return &Costs{
		mu:    new(sync.Mutex),
		costs: make(map[string]*Cost),

		compressPath: compressPath,
		pathDict:     pathDict,
	}
}

func (c *Costs) addCost(skip int, cost time.Duration) {
	path := getStackInfo(skip)
	text := formatStackInfo(path, c.compressPath, c.pathDict)
	c.addCostWithPath(text, cost)
}

func (c *Costs) addCostWithPath(path string, cost time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	costObj, ok := c.costs[path]
	if !ok {
		costObj = &Cost{
			Path: path,
		}
		c.costs[path] = costObj
	}
	costObj.CalledCount++
	costObj.Cost += cost
}

type Cost struct {
	Path        string
	CalledCount int
	Cost        time.Duration
}
