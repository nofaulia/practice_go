package main

import "fmt"

func main() {
	fmt.Println("Test 1")

	var a string = "Hello"
	var c, d string = "All", "The World"

	var e int = 1
	var f bool = true

	g := 1.5

	fmt.Println(a + " " + c + " " + d + " ")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	f = false

	if f {
		fmt.Println("F is true")
	} else {
		fmt.Println("F is false")
	}

	var h int = 10

	for i := e; i <= h; i++ {
		fmt.Println(i)
	}

	const k = 20
	fmt.Println(k)

	switch {
	case e == 1:
		fmt.Println("e : 1")
	default:
		fmt.Println("default")
	}
}
