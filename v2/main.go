package main

import (
	"fmt"
	"reflect"
)

type Config struct {
	Name string `env:"NAME,okay" defaultvalue:"Nitin"`
}

func Parse(cfg interface{}) error {
	s := reflect.ValueOf(cfg).Elem()

	fields := fields(s)
	for _, val := range fields {
		fmt.Printf("%v/ %v/ %v", val.typeOfVal, val.name, val.tags)
	}
	return nil
}

func main() {
	cfg := Config{}
	errors := Parse(&cfg)

	if errors != nil {
		fmt.Println(errors)
	}
}
