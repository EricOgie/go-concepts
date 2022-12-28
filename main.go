package main

import "fmt"


/* Arraay follow the general concepts of array for other languages. The onliy difference that may exist for some other languages is that
 *In Go, array has fixed defined length (bound). Once the length of a list is defined it is immediately taken as an array in Go and the 
 * size or length can not be changed
*/

/*
 * Slice, on the other hand, correspond to a list with no definit size bound. i.e, items can be appended and removed from the list
*/
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
	fmt.Println("Welcome to Eric's Goland")

	name := "Eric Ogie"
	score := 25.67
	age := 24

	fmt.Printf("My name is %v, I am %v years old and my ave score on Slam is %v point average \n", name, age, score)
	//fmt.Printf("name as used here is of type %T \n", name)
	// fmt.Printf("score as used is of type %T /n", score)
	fmt.Printf("My last score is %0.5F \n", 98.221)

	statement := fmt.Sprintf("My name is %v, I am %v years old and my ave score on Slam is %v point average \n", name, age, score)

	fmt.Println("My last state is")
	fmt.Println(statement)

	playaarray()

}
