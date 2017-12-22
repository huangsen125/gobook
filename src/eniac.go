package main

import "fmt"

func main() {
	a := 100
	b := 200
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	x, y := swap2(&a, &b)
	fmt.Println("a:", *x)
	fmt.Println("b:", y)
}

func swap2(a, b *int) (x, y *int) {
	fmt.Println("a2:", a)
	fmt.Println("b2:", b)

	temp := a
	x = b
	y = temp
	fmt.Println("a3:", a)
	fmt.Println("b3:", b)
	return
}
