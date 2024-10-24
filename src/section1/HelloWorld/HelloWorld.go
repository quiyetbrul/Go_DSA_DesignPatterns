package main

import "fmt"

// if func name starts with a capital letter, it is public
// if func name starts with a lower case letter, it is private
func hello(name string) {
	fmt.Println("Hello", name)
}

func sum(a int, b int) int {
	return a + b
}

// a function can return multiple values
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	name := "quiyet"
	hello(name)

	// variables and elementary types

	var age int = 12 // only way to define a package level variable
	age = 40

	fmt.Println("Age:", age)

	fmt.Println("Sum:", sum(1, 2))

	a, b := 3, 4
	fmt.Println("Before swap:", a, b)
	a, b = swap(a, b)
	fmt.Println("After swap:", a, b)

	// anonymous function
	// save a function into a variable
	// this is called a closure
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println("Sum:", sum(1, 2))

	// control structures
	// typical for loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("evem", i)
		}
	}

	// while loop
	i := 1
	for i < 4 {
		switch i {
		case 1:
			fmt.Println("One")
			// golang doesn't require a break statement in a switch
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Default")
		}
		i++
	}

	// arrays and slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// append to a slice
	numbers = append(numbers, 6, 7, 5, 1)
	fmt.Println("Numbers:", numbers)

	// for range
	for i, number := range numbers {
		fmt.Println("Index:", i, "Number:", number)
	}
	// if you don't need the index, you can use _
	for _, number := range numbers {
		fmt.Println("Number:", number)
	}

	// maps
	spanishNumbers := map[string]string{
		"one": "uno",
		"two": "dos",
		"three": "tres",
		"four": "cuatro",
	}
	fmt.Println("Spanish numbers:", spanishNumbers)
	for key, value := range spanishNumbers {
		fmt.Println("Key:", key, "Value:", value)
	}

	four, ok := spanishNumbers["four"]
	if ok {
		fmt.Println("Four in Spanish:", four)
	} else {
		fmt.Println("Four not found")
	}
}
