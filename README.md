# simplerandom
Random string generator with go.

## Usage

1. go get `github.com/songjiayang/simplerandom`
2. use `Number`, `Literal`, `URL` seed to build your random string.

  ```golang
  package main

  import (
    "fmt"

    "github.com/songjiayang/simplerandom"
  )

  func main() {
    randomNumber, _ := simplerandom.Number.GenerateString(6)
    fmt.Println(randomNumber)

    randomLiteral, _ := simplerandom.Literal.GenerateString(8)
    fmt.Println(randomLiteral)

    randomURL, _ := simplerandom.URL.GenerateString(10)
    fmt.Println(randomURL)
  }

  ```
