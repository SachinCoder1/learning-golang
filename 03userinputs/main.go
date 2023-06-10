package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	input, err := reader.ReadString('\n')
	fmt.Println(err)
	fmt.Println("Hello, ", input, "Nice to meet you!")
}
