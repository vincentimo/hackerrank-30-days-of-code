// https://www.hackerrank.com/challenges/30-binary-numbers

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)



func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  n := int64(nTemp)
  
  sN := strconv.FormatInt(n, 2)
  
  fmt.Printf("%d\n", calculate_maximum_consecutive_characters(sN))
}

func readLine(reader *bufio.Reader) string {
  str, _, err := reader.ReadLine()
  if err == io.EOF {
    return ""
  }

  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
  if err != nil {
    panic(err)
  }
}

func calculate_maximum_consecutive_characters(s string) int {
  
  // Counter for the maximum count
  var iMaximumCount int
  
  // If len(s) == 0, the maximum count will be 0
  if len(s) == 0 {
    iMaximumCount = 0
  } else {
  
    // Counter for the current count
    var iCurrentCount int
    
    // Special condition for the first character
    if s[0] == '1' {
      iCurrentCount = 1
    } else {
      iCurrentCount = 0
    }
    iMaximumCount = iCurrentCount
    
    // If len(s) > 1, begin the iteration
    if len(s) > 1 {
      for i := 1; i < len(s); i++ {
      
        // Evaluate the current count
        if s[i] == '1' {
          iCurrentCount++
        } else {
          iCurrentCount = 0
        }
        
        // Evaluate the maximum count
        if iCurrentCount > iMaximumCount {
          iMaximumCount = iCurrentCount
        }
      }
    }
  }
  return iMaximumCount
}
