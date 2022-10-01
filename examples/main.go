package main

import (
	"fmt"
	"os"

	"github.com/nitintf/env/v1"
)

var BIND_ADDRESS = env.String("BIND_ADDRESS", "locahost")
var PORT = env.Int("PORT", 3000)

func main() {
	err := env.Parse()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	fmt.Println("BIND_ADDRESS", *BIND_ADDRESS)
	fmt.Println("PORT", *PORT)
}
