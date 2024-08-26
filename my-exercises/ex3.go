package main

import "fmt"

func main3() {
	grades := []float64{85.5, 90.0, 78.5, 92.0, 88.0, 76.0, 95.5}

	var total float64 = 0
	for _, grade := range grades {
		total += grade
	}

	average := total / float64(len(grades))
	fmt.Printf("ortalama not: %.2f\n", average)
}
