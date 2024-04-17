package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

func startRepl(){
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
      command.callback()
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
  }
}

func normalizeInput(input string) []string{
  var lowerInput string = strings.ToLower(input)
  var fields []string = strings.Split(lowerInput, " ")

  return fields
}
