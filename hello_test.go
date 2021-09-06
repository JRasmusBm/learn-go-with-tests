package main

import (
  "testing"
)

func TestCanGreetWorld(t *testing.T) {
  result := Hello("world")
  expected := "Hello, world!"

  if result != expected {
    t.Errorf("result %s expected %s", result, expected)
  }
}

func TestCanGreetName(t *testing.T) {
  result := Hello("Rasmus")
  expected := "Hello, Rasmus!"

  if result != expected {
    t.Errorf("result %s got %s", result, expected)
  }
}
