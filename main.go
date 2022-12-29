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

	bill  := getNewBill("smartBill")

	fmt.Println(bill.format())

}
