package env

import "flag"

var envs []envVar
var help = flag.Bool("help", false, "--help to show env help")

func init() {
	envs = make([]envVar, 0)
}

type envVar struct {
	value        interface{}
	name         string
	varType      string
	defaultValue interface{}
	setValue     func(interface{}, string) error
	setDefault   func(interface{}, interface{})
	envValue     *string
}

func String(name string, defaultValue string) *string {
	v := new(string)

	envs = append(envs, envVar{
		value:        v,
		name:         name,
		defaultValue: defaultValue,
		varType:      "string",
		setValue: func(i interface{}, s string) error {
			return nil
		},
		setDefault: func(i1, i2 interface{}) {},
		envValue:   new(string),
	})

	return v
}

func Int(name string, defaultValue int) *int {
	v := new(int)

	envs = append(envs, envVar{
		value:        v,
		name:         name,
		defaultValue: defaultValue,
		varType:      "int",
		setValue: func(i interface{}, s string) error {
			return nil
		},
		setDefault: func(i1, i2 interface{}) {},
	})

	return v
}

func Bool(name string, defaultValue bool) *bool {
	v := new(bool)

	envs = append(envs, envVar{
		value:        v,
		name:         name,
		defaultValue: defaultValue,
		varType:      "bool",
		setValue: func(i interface{}, s string) error {
			return nil
		},
		setDefault: func(i1, i2 interface{}) {},
	})

	return v
}

func Float64(name string, defaultValue float64) *float64 {
	v := new(float64)

	envs = append(envs, envVar{
		value:        v,
		name:         name,
		defaultValue: defaultValue,
		varType:      "bool",
		setValue: func(i interface{}, s string) error {
			return nil
		},
		setDefault: func(i1, i2 interface{}) {},
	})

	return v
}
