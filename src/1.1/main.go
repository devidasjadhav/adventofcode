package main

import (
	"fmt"
)

//import "strings"
//import "regexp"

func isDigit( num byte) bool {
    if num >= '0' && num <= '9' {
	return true
    }
    return false
}

func isSubStrFrom( i int, str string, sub_str string) bool {
    if len(str) < i + len(sub_str) {
	return false
    }

    if str[i:i+len(sub_str)] == sub_str {
	return true
    }
    return false
}

func isCharDigit(i int, l byte, str string) int{
    switch l {
    case 'o' :
	if isSubStrFrom(i, str, "one") {
	    return 1
	}
    case 't' :
	if isSubStrFrom(i, str, "two") {
	    return 2
	} else if isSubStrFrom(i, str, "three") {
	    return 3
	}
    case 'f' :
	if isSubStrFrom(i, str, "four") {
	    return 4
	} else if isSubStrFrom(i, str, "five") {
	    return 5
	}
    case 's' :
	if isSubStrFrom(i, str, "six") {
	    return 6
	} else if isSubStrFrom(i, str, "seven") {
	    return 7
	}
    case 'e' :
	if isSubStrFrom(i, str, "eight") {
	    return 8
	}
    case 'n' :
	if isSubStrFrom(i, str, "nine") {
	    return 9
	}
    default:
	return -1
    }
    return -1
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
	for i,num := range []byte(input) {
	   if isDigit(num) {
		if first == -1 {
		first = int(num) - int('0') 
		}
	   last = int(num) - int('0') 
	   } else {
		n := isCharDigit(i, num,input)
		if n != -1 {
		    if first == -1 {
			first = n;
		    }
		    last = n
		}
	   }
	}

	fmt.Printf("%d%d \n %s\n",first,last, input)
	d := 10 * first + last
	sum += d
    }
    fmt.Printf("\n Sum %d \n",sum)
}




