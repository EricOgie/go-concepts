package main

import (
	"fmt"
)

func main() {

	boys := []string{"Eric", "James", "Greg", "Josh", "Samuel"}

	// Loops like mare conditional block
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	// Normal for loop
	for x := 0; x < 5; x++ {
		fmt.Println("Value of x is: ", x)
	}

	for x := 0; x < len(boys); x++ {
		fmt.Println(boys[x])
	}

	// Foreach loop with rang

	for i, v := range boys {
		fmt.Printf("value at index %v is : %v \n", i, v)
	}

	// Since Go won't allow you declare unused variables, the foreach loop can be adapted for cases where either index or the value is not needed

	for _, v := range boys {
		fmt.Printf("value is : %v \n", v)
	}
}
