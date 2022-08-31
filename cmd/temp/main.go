package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: temp number")
		os.Exit(1)
	}
	inValue, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("NAN")
		os.Exit(2)
	}

	celsius := (inValue - 32) * 5 / 9
	fahrenheit := inValue*9/5 + 32

	fmt.Printf("Celsius: %d\n", celsius)
	fmt.Printf("Fahrenheit: %d\n", fahrenheit)

}
