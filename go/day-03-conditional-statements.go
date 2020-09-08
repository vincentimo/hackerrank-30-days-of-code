// https://www.hackerrank.com/challenges/30-conditional-statements

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

  NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  N := int32(NTemp)

  var output string

  if N % 2 == 1 {
    output = "Weird"
  } else if (N >= 2) && (N <= 5) {
    output = "Not Weird"
  } else if (N >= 6) && (N <= 20) {
    output = "Weird"
  } else if (N > 20) {
    output = "Not Weird"
  }

  fmt.Printf("%s\n", output)
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
