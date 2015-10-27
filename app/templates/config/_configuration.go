package config

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "path"
)

type Environment struct {
  Name     string
  Logging  LoggingConfig
}

type LoggingConfig struct {
  LogToConsole bool
  LogToFile    bool
  LogFilePath  string
  LogLevel     string
}

type DatabaseConfig struct {
  Username    string
  Password    string
  Host        string
  Port        int
}

func LoadEnvironment(env string) *Environment {
  envFilePath := path.Join(".", "config", "environments", env + ".json")
  envFileContents, err := ioutil.ReadFile(envFilePath)
  if err != nil {
    panic(fmt.Sprintf("Could not load %s config file: %s", envFilePath, err.Error()))
  }

  environment := &Environment{Name: env}
  err = json.Unmarshal(envFileContents, &environment)
  if err != nil {
    panic(fmt.Sprintf("Found config but could not load JSON: %s", err.Error()))
  }

  return environment
}
