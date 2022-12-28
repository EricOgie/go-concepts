package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := []int{45, 24, 15, 100, 33, 61}

	boys := []string{"Eric", "James", "Greg", "Josh", "Samuel"}

	// Search the index of a particular slice before sorting
	fmt.Println(sort.SearchInts(ages, 100))
	fmt.Println(sort.SearchStrings(boys, "Greg"))
	// Sort ages in increasing order
	sort.Ints(ages)
	// Sort boys in Alphabetical order
	sort.Strings(boys)
	// print sorted list
	fmt.Println(ages)
	fmt.Println(boys)

	// Search the index of a particular slice after sorting
	fmt.Println(sort.SearchInts(ages, 100))
	fmt.Println(sort.SearchStrings(boys, "Greg"))
}
