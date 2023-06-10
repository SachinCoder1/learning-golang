package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Let's jump and see time in golang")

	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	formattedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("formattedTime:", formattedTime)

	customDate := time.Date(2022, time.February, 20, 18, 05, 01, 00, time.Local)
	fmt.Println("Custom Date:", customDate.Format("01-02-2006 Monday"))
}
