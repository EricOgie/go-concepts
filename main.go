package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func createNewBill(userInput string) bill {
	// Create a plane Bill
	newBill := getNewBill(userInput)
	// Let user know a bill has been created
	fmt.Println(newBill.name, "bill has been successfully created")
	// Return bill
	return newBill
}

// Helper function
func getUserInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	userInput, error := reader.ReadString('\n')
	return strings.TrimSpace(userInput), error
}

func promptOptions(b bill, reader *bufio.Reader) {
	userOption, _ := getUserInput("Choose Option \na) - Add Item\nb) - Save Bill\nc) - Add Tip", reader)
	switch userOption {
	case "a":
		itemName, _ := getUserInput("Enter Item name: ", reader)
		itemPriceAsString, _ := getUserInput("Enter Item price in $: ", reader)
		price, er := strconv.ParseFloat(itemPriceAsString, 64)
		if er != nil {
			fmt.Println("Amount can only be numbers")
			promptOptions(b, reader)
		}
		b.items[itemName] = price
		promptOptions(b, reader)
	case "b":
		tipAmount, _ := getUserInput("Enter tip amount $: ", reader)
		tipAmountf, er := strconv.ParseFloat(tipAmount, 64)
		if er != nil {
			fmt.Println("Amount can only be numbers")
		}
		b.tip = tipAmountf
		promptOptions(b, reader)
	case "c":
		b.saveBill()
	default:
		fmt.Println("That was an Invalid Option")
		getUserInput("Choose Option \na) - Add Item\nb) - Save Bill\nc) - Add Tip", reader)
	}
}

func runBillMaker() {
	reader := bufio.NewReader(os.Stdin)

	userName, _ := getUserInput("Hi! My name is Cynthia from Bill maker... Please who am I speaking to?", reader)
	time.Sleep(1 * time.Second)

	billName, _ := getUserInput("Hello "+userName+" Welcome To bill maker.... What is the name of the bill you wish to make?\n", reader)

	// Sleep for a second
	fmt.Printf("... Creating %v Bill... This will only take a few seconds \n", billName)
	time.Sleep(4 * time.Second)

	userBill := createNewBill(billName)
	fmt.Println(userBill)

	fmt.Println("---------------------")
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------")
	promptOptions(userBill, reader)

}

func main() {
	runBillMaker()
}
