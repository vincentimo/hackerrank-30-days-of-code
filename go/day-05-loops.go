// https://www.hackerrank.com/challenges/30-loops

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func print_first_10_multiples(n int) {
  for i := 1; i <= 10; i++ {
    fmt.Printf("%d x %d = %d\n", n, i, n * i)    
  }
}

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  n := int(nTemp)

  print_first_10_multiples(n)
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
