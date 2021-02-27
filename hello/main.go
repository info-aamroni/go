package main

import "fmt"

func main() {
	fmt.Println(foo())
	iterator()
	declared()
	bar()
}

func iterator() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Is", i, "a even number?")
		}
	}
}

func bar() {
	fmt.Println("It's time to exit")
}

func foo() string {
	return "Hello Web Application Developer, How are you?"
}

func declared() {
	a := 10
	b := 10
	c := a + b
	fmt.Println(a, "+", b, "; The total equation is =", c)
}
