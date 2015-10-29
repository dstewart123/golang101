package main

import (
	"fmt"
)

func main() {
	fmt.Println("Cycles -> Result")

	fmt.Println(0, " -> ", calculateGrowth(0))
	
	fmt.Println(1, " -> ", calculateGrowth(1))
	
	fmt.Println(4, " -> ", calculateGrowth(4))
}

func calculateGrowth(cycles int) int {
	growth := 1
	
	for i := 0; i < cycles; i++ {
		if i%2 == 0 {
			// Spring Growth
			growth = growth * 2
		} else {
			// Summer Growth
			growth++
		}
	}

	return growth
}
