package main

import "testing"

func TestCalculateGrowth(t *testing.T) {
        result := calculateGrowth(1)
        if result != 2 {
                t.Fatal("calculateGrowth != 2")
        }
}
