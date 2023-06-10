package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome To Go!")

	fmt.Println("Welcome to Pizza app! Please rate the pizza in the scale of 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating the pizza! Rated: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("there is an error in numRating variable...", err)
	} else {
		fmt.Println("Added 1 to it.. ", numRating+1)
	}

}
