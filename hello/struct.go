package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 21})

	fmt.Println(person{name: "Fred"})

	fmt.Println(person{age: 15})

	fmt.Println(&person{name: "Ane", age: 40})

	s := person{"Sean", 50}
	fmt.Println(s)
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
}