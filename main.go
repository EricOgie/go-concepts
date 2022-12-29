package main

import "fmt"

// Pointer wrapper Values: Slice, Map, functions

func updateSampleOne(x string) {
	// Attempt to change the passed  value of x to Column
	x = "Column"
}

func changeSampleOne(x *string) {
	*x = "Grade solid"
}

func main() {

	sample := "Samsung"

	updateSampleOne(sample)

	fmt.Println(sample)

	fmt.Println("memory Address for sameple is: ", &sample)

	// We can assign the memory add for a variable to another variable.
	sampleMemAdd := &sample // Now holds the memory address for variable, sample.

	fmt.Printf("Value of sampleMemAddress is: %v. It should be the same value as memory address, %v printed above \n", sampleMemAdd, &sample)

	// De-Referencing a Memory pointer
	// If we are sure we are dealing with memory pointer, like sampleMemAdd above, we can access the true value in the mem address by de-referencing
	fmt.Printf("True value at %v address is %v \n", sampleMemAdd, *sampleMemAdd)

	// To modified the original value passed to a function, we have to pass the pointer to that value as the argument to the function as did changeSampleOne
	changeSampleOne(&sample)
	fmt.Println("Changed value of sample is: ", sample)

}
