# ENV

env is a simple package to get enviorment variables from env files. you can also
provide default value as fallback we env is not found in env files

## Basic Usage

```go

package main

import (
	"fmt"
	"os"

	"github.com/nitintf/env/v1"
)

var PORT = env.Int("PORT", 3000)

func main() {
	err := env.Parse()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	fmt.Println("PORT", *PORT)
}

```
