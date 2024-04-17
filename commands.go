package main

import (
  "fmt"
  "os"
)

func commandHelp() error {
  fmt.Println("\nWelcome to the Gokedex!")
  fmt.Println("Available commands and usage:\n\n")

  commandMap := getCommands()

  for _, val := range commandMap {
    fmt.Printf("%s: %s", val.name, val.description)
    fmt.Println()
  }

  fmt.Println("\n")

  return nil
}

func commandExit() error {
  os.Exit(0)
  return nil
}
