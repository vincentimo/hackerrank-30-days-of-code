// https://www.hackerrank.com/challenges/30-class-vs-instance

package main

import "fmt"

type person struct {
  age int
}

func (p person) NewPerson(initialAge int) person {
  //Add some more code to run some checks on initialAge
  
  p.age = initialAge

  if p.age < 0 {
    fmt.Printf("Age is not valid, setting age to 0.\n")
    p.age = 0
  }

  return p
}

func (p person) amIOld() {
  //Do some computation in here and print out the correct statement to the console
  
  var output string

  if p.age < 13 {
    output = "You are young."
  } else if p.age < 18 {
    output = "You are a teenager."
  } else {
    output = "You are old."
  }

  fmt.Printf("%s\n", output)
}

func (p person) yearPasses() person {
  //Increment the age of the person in here
  
  p.age++

  return p
}

func main() {
  var T, age int

  fmt.Scan(&T)

  for i := 0; i < T; i++ {
    fmt.Scan(&age)
    p := person{age: age}
    p = p.NewPerson(age)
    p.amIOld()
    for j := 0; j < 3; j++ {
      p = p.yearPasses()
    }
    p.amIOld()
    fmt.Println()
  }
}
