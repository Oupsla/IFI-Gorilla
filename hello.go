package main

import "fmt"


const hello string = "Hello world\n"

func main() {

  for i := 1; i <= 9; i++ {
    if i == 5 {
        fmt.Printf(hello)
    }
  }

}

