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
