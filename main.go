package randompick

import (
	"math/rand"
)

func Result(i map[string]int) string {
	totalWeight := 0
	for _, weight := range i {
		totalWeight += weight
	}
	r := rand.Intn(totalWeight)
	totalWeight = 0
	for res, weight := range i {
		totalWeight += weight
		if totalWeight > r {
			return res
		}
	}
	return ""
}
