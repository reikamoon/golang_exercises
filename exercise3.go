// Declares a public struct that represents a house.

// The struct should contain variables numberOfRooms, city,  address, and price, each assigned their appropriate data type (string, int, etc)
//
// When your program runs, allow users to enter information for one to many houses.
//
// After the user submits all their information, print a list of the houses for sale in the following format:
// 123 Address Road   San Francisco   8 Rooms   $123,456
// 456 Another Ln     Pittsburgh      4 Rooms   $90,000

package main

import (
    "fmt"
    "os"

  )


type House struct {
   numberOfRooms int
   city string
   address string
   price int

}
func main() {
   var myhouse House

   fmt.Println("Enter how many rooms are in the house. ")

     // var then variable name then variable type
     var numberOfRooms int

     // Taking input from user
     fmt.Scanln(&numberOfRooms)
     fmt.Println("Enter the city: ")
     var city string
     fmt.Scanln(&city)
     fmt.Println("Enter the address: ")
     var address string
     fmt.Scanln(&city)
     fmt.Println("Enter the price: ")
     var price int

     fmt.PrintLn("My house has " + numberOfRooms. "My house is in " + address, "in", + city, "and costs" + price)
}
