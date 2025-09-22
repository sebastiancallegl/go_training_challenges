package main

import (
	"fmt"
	"os"
)

func main() {
	var reversedString string
	for i := len(os.Args[1]) - 1; i >= 0; i-- {
		reversedString += string(os.Args[1][i])
	}
	fmt.Println(reversedString)
}
