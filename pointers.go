package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x)

	z := x
	fmt.Println(&z, z)
	z = 20
	fmt.Println(&z, z, x)

	p := &x
	fmt.Println(p)
	*p = 15
	fmt.Println(*p)

	var n int64 = 30
	fmt.Println(n)
	increment(&n)
	fmt.Println(n)

	var l *int
	fmt.Println(l)

	a := 1
	b := new(int)
	fmt.Println(a, &a, b, *b)
	b = &a
	fmt.Println(a, &a, b, *b)
}

func increment(m *int64) {
	*m += 1

}
