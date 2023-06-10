package main

import "fmt"

const LoginToken string = "Reusable Login Token"

func main() {
	var username string = "this is variable"
	fmt.Println("username: ", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isUsername bool = false
	fmt.Println("username: ", isUsername)
	fmt.Printf("Variable is of type: %T \n", isUsername)

	var smallVal uint8 = 255
	fmt.Println("smallVal: ", smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var floatVal float64 = 255.02029292
	fmt.Println("floatVal: ", floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)

	var emptyVal uint64
	fmt.Println("emptyVal: ", emptyVal)
	fmt.Printf("Variable is of type: %T \n", emptyVal)

	// short hand declaration

	shortHand := "this is short hand"
	fmt.Println("shortHand: ", shortHand)
	fmt.Printf("Variable is of type: %T \n", shortHand)

	fmt.Println("LoginToken: ", LoginToken)
}
