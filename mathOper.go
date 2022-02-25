package main

import (
	"fmt"
	"math"
)

var a1 int32 = 20

func main() {
	a := 10
	res := int32(a) + a1
	fmt.Println("the sum of two numbers:", res)
	res = int32(a) - a1
	fmt.Println("difference of two numbers:", res)
	res = int32(a) * a1
	fmt.Println("multiplication of two numbers:", res)
	var c float32 = float32(a1) / float32(a) //type conversion
	fmt.Println("division of a number by a number:", c)

	circleRadius := 5
	circleArea := (math.Pow(float64(circleRadius), float64(2))) * math.Pi //type conversion
	fmt.Println(circleArea)
	var age int32
	fmt.Print("how old are you?: ")
	fmt.Scanln(&age)
	fmt.Printf("I'm %d./n", age)

	//creating an error description
	const name, id = "Venk", 23
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

}
