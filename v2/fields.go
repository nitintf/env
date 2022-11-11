package main

import "reflect"

type field struct {
	value     reflect.Value
	name      string
	tags      reflect.StructTag
	typeOfVal reflect.Type
}

func fields(s reflect.Value) []field {
	fields := make([]field, 0, s.NumField())

	for i := 0; i < s.NumField(); i++ {
		tags := s.Type().Field(i).Tag
		name := s.Type().Field(i).Name
		typeOfVal := s.Type().Field(i).Type
		value := s.Field(i)
		fields = append(fields, field{value, name, tags, typeOfVal})
	}

	return fields
}
