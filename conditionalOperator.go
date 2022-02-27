package main

import (
	"errors"
	"fmt"
	"math"
)

func printCircleArea(radius int) {
	result, err := circleRadius(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("circle radius :", result)

}

func circleRadius(radius int) (float64, error) {
	if radius <= 0 {
		return float64(radius), errors.New("the radius of the circle cannot be negative")
	}
	return math.Pow(float64(radius), float64(2)) * math.Pi, nil
}

func main() {
	printCircleArea(-2)
}
