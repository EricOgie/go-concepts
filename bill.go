package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Getter finction for a bill
func getNewBill(billName string) bill {
	return bill{
		name:  billName,
		items: map[string]float64{},
		tip:   0.00,
	}
}

// ---------  Receiver functions
func (b bill) format() string {
	formattedBill := "Bill Breakdown: \n"
	total := 0.00

	for key, value := range b.items {
		formattedBill += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	total += b.tip
	// Add Tip
	formattedBill += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	// Now Add the total to the formattedBill
	formattedBill += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total)
	// NB. %-25 is to pad up with 25charracter space just to make it the text more formatted
	return formattedBill
}

// Added functions that can update and modify the original bill passed
func (b *bill) addItem(itemName string, itemPrice float64) {
	(*b).items[itemName] = itemPrice // Note that (*b) is not really needed to de-reference b since struct are de-reference by default
}

func (b *bill) updateBillTip(newtip float64) {
	b.tip = newtip
}

func (b *bill) saveBill() {
	datatoSave := []byte(b.format())
	filepath := "bill/" + b.name + ".txt"
	er := os.WriteFile(filepath, datatoSave, 0644)

	if er != nil {
		panic(er)
	}

	fmt.Println("Bill saved successfully")
}
