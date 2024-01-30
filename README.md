# Simple env loader

It checks your environment variables at build time, based on a predefined struct

## Install

```shell
go get github.com/peterszarvas94/envloader@v1.0.8
```

## Usage

```go
package main

import (
  "fmt"
  "os"

  "github.com/peterszarvas94/envloader"
)

type Config struct {
  DB_NAME string
}
var config Config

func main() {
  // optionally load env from file
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()
  envloader.File(file)

  // check fields and load values to config
  err := envloader.Load(&config)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(config.DB_NAME)
}
```
