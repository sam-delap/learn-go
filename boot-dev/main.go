package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		multipleOf3 := i % 3 == 0
		multipleOf5 := i % 5 == 0
		if multipleOf3 && multipleOf5 {
			fmt.Println("fizzbuzz")
		} else if multipleOf3 {
			fmt.Println("fizz")
		} else if multipleOf5 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}

