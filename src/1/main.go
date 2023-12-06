package main

import "fmt"
//import "strings"
//import "regexp"

func isDigit( num byte) bool {
    if num >= '0' && num <= '9' {
	return true
    }
    return false
}
func main() {
    first := -1
    sum := 0
    for true {
	var input string
	fmt.Scanln(&input)
	if len(input) == 0 {
	    break
	}
	first = -1
	last := -1
	//for _,i := range strings.Split(input,"") {
	for _,num := range []byte(input) {
	   if isDigit(num) {
		if first == -1 {
		first = int(num) - int('0') 
		}
	   last = int(num) - int('0') 
	   }
	}
	d := 10 * first + last
	sum += d
    }
    fmt.Printf("\n Sum %d \n",sum)
}




