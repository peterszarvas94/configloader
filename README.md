# Config loader

Checks your environment variables at build time, based on a struct. Loads the
vars into the struct, and errors if one is missing.

Enables you to have an autocomplete on your env vars, and make sure they
exists.

## Install

```shell
go get github.com/peterszarvas94/configloader
```

## Usage

```go
package main

import (
  "fmt"
  "os"

  "github.com/peterszarvas94/configloader"

  // E.g. you can use godotenv to load vars from .env file
  // It's not a dependency, you can define env vars as you wish
  "github.com/joho/godotenv/autoload/"
)

type Config struct {
  DB_NAME string
}
var config Config

func main() {
  err := configloader.Load(&config)
  if err != nil {
    // One env var is missing
    fmt.Println(err)
    os.Exit(1)
  }

  // Type-safe and autocompleted
  fmt.Println(config.DB_NAME)
}
```
