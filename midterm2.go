// Question 2
//
// Create a program that calculates the average weight of 5 people.

package main

import (
    "fmt"
)

func main() {
    p1 := 50
    p2 := 30
    p3 := 20
    p4 := 25
    p5 := 35

    avg := (p1+p2+p3+p4+p5) / 5

    fmt.Println(avg)
}
