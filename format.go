package costwhere

import (
	"fmt"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func formatCosts(c *Costs) (costs []string) {
	arr := lo.Values(c.costs)
	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i].Path) < len(arr[j].Path) {
			return true
		}
		return arr[i].Path < arr[j].Path
	})

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if strings.HasPrefix(arr[j].Path, arr[i].Path) {
				arr[i].Cost -= arr[j].Cost
				break
			}
		}
	}

	for _, cost := range arr {
		text := fmt.Sprintf("%s %d", cost.Path, cost.Cost.Milliseconds())
		costs = append(costs, text)
	}

	return
}
