package main

import (
	"fmt"
	"strings"
)

/* It's important to note that that Go comes as a light weight language/tool by default. Onlike many other languages where you have
 * many in-built functions and methods loaded into the program by default, Go seperates her utility functionionalities in packages that must be
 * explicitly imported in order to use the public functions in these packages. Typical examples include fmt and string package.
 */

var myHello string = "Welcome to Oneticha's world!"

func playaarray() {
	first := []string{"yoshi", "luigi", "peach", "bowser"}

	second := []int{1, 2, 3, 4, 5, 6, 7}

	thrid := append(second, 9)

	fmt.Println(first, len(first))
	fmt.Println(second, len(second))
	fmt.Println(thrid, len(thrid))

	rr := first[1:3]

	fmt.Println(rr)
	fmt.Println(first[2:])
	fmt.Println(first[:2])

}

func main() {

	/*
	* To manipulate the string, myHello, we can import "string" package
	* and test our hands on a number of functions built-in to the string package.
	 */

	fmt.Println(strings.Split(myHello, " "))
	fmt.Println(strings.Split(myHello, " ")[0])
	fmt.Println(strings.Contains(myHello, "Welcome"))
	fmt.Println(strings.ToUpper(myHello))
	fmt.Println(strings.Index(myHello, "world"))

	// NB: the manipulation of these function is never on the original string, instead, it create a copy of the original string and only
	// manipulate the copy of the string it created.

	newMyHello := strings.ReplaceAll(myHello, "Oneticha", "Aghahowa")
	fmt.Println(newMyHello)

	fmt.Println("Altered MyHello = " + strings.ReplaceAll(myHello, "Oneticha", "Aghahowa"))
	fmt.Println("Original MyHello = " + myHello)
}
