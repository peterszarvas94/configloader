# Simple env loader

It checks your environment variables at build time, based on a predefined struct

## Install

```shell
go get github.com/peterszarvas94/envloader@v1.0.5
```

## Usage

```go
package main

import (
  "fmt"
  "os"

  "github.com/peterszarvas94/envloader"
)


func main() {
  type Config struct {
    DB_NAME string
  }
  var config Config

  // optional
  envloader.File(".env")

  err := envloader.Loader(&config)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(config.DB_NAME)
}
```
