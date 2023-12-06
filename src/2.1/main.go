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
        var r,g,b int
        var n int
        var cn int
        var colorName string
        r = 0
        g = 0
        b = 0
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
                    if r < cn {
                        r = cn
                    }
                case "green":
                    if g < cn {
                        g = cn
                    }
                case "blue":
                    if b < cn {
                        b = cn
                    }
                }
                fmt.Printf("%d -> %s|\n", cn, colorName)
            }
             fmt.Println("=========================================")
        }
        fmt.Printf("%d[%d] --> %d %d %d\n",n,sum,r,g,b)
        sum += r*g*b
    }
    fmt.Println("Total Sum is %d\n",sum)
    if err := scanner.Err(); err != nil {
        os.Exit(1)
    }
}
