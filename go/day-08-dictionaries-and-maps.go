// https://www.hackerrank.com/challenges/30-dictionaries-and-maps
// This fails on Test Case #1
// Alternative: https://johnny2136.github.io/30DaysOfCodeIssue/
// But this fails on Test Case #2 and #3

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
  // Initialize a reader for getting user's inputs
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  // Get the number of entries in the phone book
  nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  n := int(nTemp)

  // Initialize a map
  mPhoneBook := make(map[string]string)

  // Get phone book entries in the form of "${name} ${phone number}"
  // Store them to the map; key = "${name}", value = "${phone number}"
  for i := 0; i < n; i++ {
    arrTemp := strings.Split(readLine(reader), " ")
    mPhoneBook[arrTemp[0]] = arrTemp[1]
  }

  // Initialize an array
  aQuery := make([]string, n)

  // Get query entries in the form of "${name}"
  // Store them to the array
  for i := 0; i < n; i++ {
    aQuery[i] = readLine(reader)
  }

  // Access the array and check if key exist
  for i := 0; i < n; i++ {
    sValue, bExist := mPhoneBook[aQuery[i]]
    if bExist {
      fmt.Printf("%s=%s\n", aQuery[i], sValue)
    } else {
      fmt.Println("Not found")
    }
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
