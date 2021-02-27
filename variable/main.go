package main

import (
	"fmt"
)

//Global Variable Declaration
var z int = 20

func main() {
	//ShortHand of Declare a Variable
	x := 10
	fmt.Println(x)

	y := "Hello Go Language"
	fmt.Println(y)

	sum(z)
	types()
}

func sum(a int) {
	fmt.Println("The total sum is:", a+10)
}

func types() {
	a := "Is is a string?"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := 10
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	c := 12.14700
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
