package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
  "encoding/json"
  "errors"
)

type LocationAreas struct {
  Count int `json:"count"`
  Next *string `json:"next"`
  Previous *string `json:"previous"`
  Results []struct {
    Name string `json:"name"`
    URL string `json:"url"`
  } `json:"results"`
}

func commandHelp(config *Config) error {
  fmt.Println("\nWelcome to the Gokedex!")
  fmt.Println("Available commands and usage:\n\n")

  commandMap := getCommands()

  for _, val := range commandMap {
    fmt.Printf("%s: %s\n", val.name, val.description)
  }

  fmt.Println()

  return nil
}

func commandExit(config *Config) error {
  os.Exit(0)
  return nil
}

// Fetch and Parse locations from PokeAPI
func getLocations(URL string) (final LocationAreas, err error) {
  var result LocationAreas
  var baseRequestURL string = "https://pokeapi.co/api/v2/location-area/?limit=20"

  if URL == "" {
    URL = baseRequestURL
  }

  res, err := http.Get(URL)
  if err != nil {
    fmt.Printf("Error fetching locations: %s\n", err)
  }

  body, err := io.ReadAll(res.Body)
  res.Body.Close()

  if res.StatusCode > 299 {
    fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
    return LocationAreas{}, nil
  }
  if err != nil {
    fmt.Printf("Error during response: %s\n", err)
    return LocationAreas{}, nil
  }

  json.Unmarshal(body, &result)

  return result, nil
}

func baseMap(locationsData LocationAreas, config *Config) error {
  if locationsData.Next != nil {
    config.Next = *locationsData.Next
  } else {
    config.Next = "nil"
  }
  
  if locationsData.Previous != nil {
    config.Previous = *locationsData.Previous
  } else {
    config.Previous = ""
  }

  fmt.Println()
  for _, area := range locationsData.Results {
    fmt.Println(area.Name)
  }
  fmt.Println()
  return nil
}

func commandMap(config *Config) error {

  if config.Next == "nil" {
    return errors.New("There are no more locations.")
  }

  locationsData, err := getLocations(config.Next)

  if err != nil {
    return err
  }

  return baseMap(locationsData, config)
}

func commandMapb(config *Config) error {
  if config.Previous == ""{
    return errors.New("There are no previous locations.")
  }

  locationsData, err := getLocations(config.Previous)

  if err != nil {
    return err
  }

  return baseMap(locationsData, config)
}
