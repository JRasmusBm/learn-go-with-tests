package main

import (
	"fmt"
)

const englishPrefix = "Hello, "

func Hello(name string) string {
  if name == "" {
    name = "world"
  }

	return englishPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}
