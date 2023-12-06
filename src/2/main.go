package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var sum int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var n int
        var cn int
        var colorName string
        s := scanner.Text()
        //for _,txt := range strings.Split(s, ":") {
        txt := strings.Split(s, ":") 
        fmt.Sscanf(txt[0],"Game %d",&n)
        fmt.Printf("--> %d ",n)
        for _,i := range strings.Split(txt[1],";") {
            for _,j := range strings.Split(i,",") {
                fmt.Sscanf(j,"%d %s",&cn,&colorName)
                switch colorName {
                case "red":
                    if cn > 12 {
                        fmt.Printf("RED balls more %d GAME %d\n",cn,n)
                        sum -= n
                        goto UP
                    }
                case "green":
                    if cn > 13 {
                        fmt.Printf("GREEN balls more %d GAME %d\n",cn,n)
                        sum -= n
                        goto UP
                    }
                case "blue":
                    if cn > 14 {
                        fmt.Printf("BLUE balls more %d GAME %d\n",cn,n)
                        sum -= n
                        goto UP
                    }
                }
                fmt.Printf("%d -> %s|\n", cn, colorName)
            }
             fmt.Println("=========================================")
        }
        UP: sum += n
    }
    fmt.Println("Total Sum is %d\n",sum)
    if err := scanner.Err(); err != nil {
        os.Exit(1)
    }
}
