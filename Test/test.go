package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func main() {
  p := person{name: "Chris", age: 19}
  fmt.Println(p.age)
}
