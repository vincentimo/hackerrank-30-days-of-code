// https://www.hackerrank.com/challenges/30-hello-world

package main

import (  
  "bufio"
  "os"
  "fmt"
)

func main() {
  // Enter your code here. Read input from STDIN. Print output to STDOUT

  // := notation is used to both instantiate and assign a variable.
  // This scanner with bufio package is used because the
  // fmt.Scanln() does not accept space as a part of input.
  scanner := bufio.NewScanner(os.Stdin)

  // This notation is used to instantiate a variable.
  var in string

  // This performs the actual scan.
  if scanner.Scan() {
    in = scanner.Text()
  }

  // fmt.Println is used because it automatically adds a newline.
  fmt.Println("Hello, World.")

  // fmt.Printf is used because it formats the output to exclude the "%s".
  // Using fmt.Println results in the "%s" being literally printed because
  // it does not format the output, hence it does not remove the "%s".
  fmt.Printf("%s\n", in)
}
