package main

import (
	"fmt"
	"go-calculate/calculator"
)

func main() {
	result := calculator.Add(2, 4)

	fmt.Println(result)
}
