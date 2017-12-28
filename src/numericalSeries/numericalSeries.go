package main

import (
	"fmt"
	"os"
	"strconv"
)

//Show numbers between valueA and valueB in pairs
func main() {
	valueA, _ := strconv.Atoi(os.Args[1])
	valueB, _ := strconv.Atoi(os.Args[2])
	for i := valueA; i <= valueB; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
