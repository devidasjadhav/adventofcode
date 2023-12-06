package main

import "fmt"

func main() {
    for true {
	var input string
	fmt.Scanln(&input)
	if len(input) == 0 {
	    break
	}
    	fmt.Println(input)
    }
}
