package main

import (
  "fmt"
  "bufio"
  "os"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

func main(){
  availableCommands := getCommands()

  var prompt string = "pokedex > "
  scanner := bufio.NewScanner(os.Stdin)

  for true{
    fmt.Print(prompt)
    scanner.Scan()
    command := scanner.Text()

    if _, ok := availableCommands[command]; ok {

      if command == "exit"{
        return
      }

      availableCommands[command].callback()
    }

  }
}

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
  return nil
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand {
    "help": {
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
    "exit": {
      name: "exit",
      description: "Exit the program",
      callback: commandExit,
    },
  }
}
