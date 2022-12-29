package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Getter finction for a bill
func getNewBill(billName string) bill {
	return bill{
		name:  billName,
		items: map[string]float64{"pie": 5.99, "cake": 35.12},
		tip:   55.00,
	}
}

// Receiver function
func (b bill) format() string {
	formattedBill := "Bill Breakdown: \n"
	total := 0.00

	for key, value := range b.items {
		formattedBill += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	// Now Add the total to the formattedBill
	formattedBill += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total)
	// NB. %-25 is to pad up with 25charracter space just to make it the text more formatted
	return formattedBill
}
