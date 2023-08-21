package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int64][]float64)

	for _, temp := range temperatures {
		groupKey :=  int64(temp - math.Mod(temp, 10))
		groups[groupKey] = append(groups[groupKey], temp)
	}

	for groupKey, groupTemps := range groups {
		fmt.Printf("%d: %v\n", groupKey, groupTemps)
	}
}