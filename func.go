package main

import (
	"fmt"
	"math"
)

func circleArea() {
	var circleRadius int64
	fmt.Print("enter a number: ")
	fmt.Scan(&circleRadius)
	result := math.Pow(float64(circleRadius), float64(2)) * math.Pi
	fmt.Println("circle radius :", result)
}

func newCircleArea(radius int) float64 {
	return math.Pow(float64(radius), float64(2)) * math.Pi
}
func printCircleArea(radius int) {
	result := newCircleArea(5)
	fmt.Println("circle radius :", result)
}

func main() {
	circleArea()
	printCircleArea(5)
}
