// https://www.hackerrank.com/challenges/30-data-types

package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
  var i uint64 = 4
  var d float64 = 4.0
  var s string = "HackerRank "

  scanner := bufio.NewScanner(os.Stdin)
  
  // Declare constants for type conversion.
  const (
    NUMBER_OF_INPUT = 3
    BASE = 10
    BIT_SIZE = 64
  )

  // Declare variables.
  var arr [NUMBER_OF_INPUT]string // This is an array for input.
  var i2 uint64
  var d2 float64
  var s2 string
  var _ error

  // Read values for an integer, double, and String to your variables.
  for j := 0; j < len(arr); j++ {
    if scanner.Scan() {
      arr[j] = scanner.Text()
    }
  }

  // Convert to integer and double, respectively.
  i2, _ = strconv.ParseUint(arr[0], BASE, BIT_SIZE)
  d2, _ = strconv.ParseFloat(arr[1], BIT_SIZE)
  s2 = arr[2]

  // Print the sum of both integer variables on a new line.
  fmt.Printf("%d\n", i + i2)
  
  // Print the sum of the double variables on a new line.
  // The result is printed to 1 decimal place.
  fmt.Printf("%.1f\n", d + d2)

  // Concatenate and print the String variables on a new line
  // The 's' variable above should be printed first.
  fmt.Printf("%s%s\n", s, s2)
}
