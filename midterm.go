// Question 1

// Write a program that calculates the year using a provided date of birth and age.
// HINT: Get the date of birth and age from stdin!
//
// Write a program that calculates the average weight of 5 people.

package main

import (
        "bufio"
        "fmt"
	      "strconv"
	      "os"
)

func findbirthyear(x int, y int) int {
  return x - y
}

func main() {
  currentyear := 2021
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter your birth date: ")
  birthday, _ := reader.ReadString('\n')
  fmt.Print("Your birthday is" + birthday)
  fmt.Print("Enter your age: ")
  age, _ := reader.ReadString('\n')
  i, err := strconv.ParseInt("age", 10, 64)

  fmt.Println(findbirthyear(currentyear, age))

}
