package main

import "fmt"

func main() {
	fmt.Println("Learning Pointer")

	myValue := 100
	var myPointer = &myValue
	fmt.Println("myPointer:", myPointer)
	fmt.Println("myPointer value", *myPointer)

	*myPointer = *myPointer + *myPointer

	fmt.Println("myValue:", myValue, "myPointerValue", *myPointer)

}
