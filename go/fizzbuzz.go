package main

import "fmt"
import "strings"

func main() {
    for i := 1; i <= 30; i++ {
        m := [] string{}
        if i % 3 == 0 {
            m = append(m, "fizz")
        }
        if i % 5 == 0 {
            m = append(m, "buzz")
        }

        if len(m) > 0 {
            fmt.Println(strings.Join(m, ""))
        } else {
            fmt.Println(i)
        }
    }
}
