package main

import (
	"fmt"
)

func main() {
	mySentence := "Hello this is a sentence."

	for _, value := range mySentence {
		fmt.Println(string(value))
		// if index % 2 == 0 {
		// 	fmt.Println(string(value))
		// }
	}
}
