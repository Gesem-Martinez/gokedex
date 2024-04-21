package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Config struct {
  Next string
  Previous string
}

type cliCommand struct {
  name string
  description string
  callback func(config *Config) error
}

func startRepl(){
  config := Config{}

  config.Next = ""
  config.Previous = ""

  availableCommands := getCommands()

  var prompt string = "Gokedex > "
  scanner := bufio.NewScanner(os.Stdin)

  for true{
    fmt.Print(prompt)
    scanner.Scan()

    commandFields := normalizeInput(scanner.Text())

    if len(commandFields) == 0 {
      continue
    }

    if command, ok := availableCommands[commandFields[0]]; ok {
      err := command.callback(&config)
      if err != nil {
        fmt.Println(err)
      }
    }

  }
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
    "map": {
      name: "map",
      description: "Display the names of the next 20 location areas in the Pokemon world.",
      callback: commandMap,
    },
    "mapb": {
      name: "mapb",
      description: "Display the names of the previous 20 location areas in the Pokemon world.",
      callback: commandMapb,
    },
  }
}

func normalizeInput(input string) []string{
  var lowerInput string = strings.ToLower(input)
  var fields []string = strings.Split(lowerInput, " ")

  return fields
}
