// https://www.hackerrank.com/challenges/30-2d-arrays

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

  var arr [][]int32
  for i := 0; i < 6; i++ {
    arrRowTemp := strings.Split(readLine(reader), " ")

    var arrRow []int32
    for _, arrRowItem := range arrRowTemp {
      arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
      checkError(err)
      arrItem := int32(arrItemTemp)
      arrRow = append(arrRow, arrItem)
    }

    if len(arrRow) != int(6) {
      panic("Bad input")
    }

    arr = append(arr, arrRow)
  }
  
  fmt.Printf("%d\n", calculate_maximum_hourglass_sum(arr))
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

func calculate_maximum_hourglass_sum(arr [][]int32) int32 {

  var iLenI, iLenJ int
  var iCurrentSum, iMaximumSum int32
  
  iLenI = len(arr)
  iLenJ = len(arr[0])
  
  for i := 0; i < iLenI - 2; i++ {
    for j := 0; j < iLenJ - 2; j++ {
      
      // Evaluate the current sum
      iCurrentSum = arr[i][j] + arr[i][j + 1] + arr[i][j + 2] + arr[i + 1][j + 1] + arr[i + 2][j] + arr[i + 2][j + 1] + arr[i + 2][j + 2]
      
      // For the first iteration, assign the maximum sum with the current sum
      if (i == 0) && (j == 0) {
        iMaximumSum = iCurrentSum
      }
      
      // Evaluate the maximum sum
      if iCurrentSum > iMaximumSum {
        iMaximumSum = iCurrentSum
      }
    }
  }
  
  return iMaximumSum
}
