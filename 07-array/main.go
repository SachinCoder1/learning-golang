package main

import "fmt"

func main() {
	fmt.Println("Let's see array in golang")

	var newArr [5]string

	fmt.Println("array:", newArr, "length of the array", len(newArr))

	// newArr[] = "item0"
	newArr[1] = "item1"
	newArr[2] = "item2"
	newArr[3] = "item3"
	newArr[4] = "item4"

	fmt.Println("array:", newArr, "length of the array", len(newArr))

}
