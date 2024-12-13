package costwhere

import (
	"fmt"
)

func formatCost(c *Cost) (cost string) {
	cost = fmt.Sprintf("%s(%d) %d", c.Path, c.CalledCount, c.TotalCost.Milliseconds())
	return
}

func formatCosts(c *Costs) (costs []string) {
	costs = make([]string, 0, len(c.costs))
	for _, cost := range c.costs {
		costs = append(costs, formatCost(cost))
	}
	return
}
