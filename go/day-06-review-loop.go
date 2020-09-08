// https://www.hackerrank.com/challenges/30-review-loop

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
  // Enter your code here. Read input from STDIN. Print output to STDOUT
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  // Get T (the number of test cases)
  TTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  T := int(TTemp)

  // Make arrays to store even- and odd-separated input string
  astrEven := make([]string, T)
  astrOdd := make([]string, T)

  for i := 0; i < T; i++ {
    S := readLine(reader)

    // Separate input string
    for j := 0; j < len(S); j++ {
      if isEven(j) {
        astrEven[i] += string(S[j]) // The output of S[j] is char; convert it to string
      } else {
        astrOdd[i] += string(S[j])
      }
    }
  }

  // Print the arrays
  for i := 0; i < T; i++ {
    fmt.Printf("%s %s\n", astrEven[i], astrOdd[i])
  }
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

func isEven(n int) bool {
  return n % 2 == 0
}
