package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var envs []envVar

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
			*i.(*string) = s
			return nil
		},
		setDefault: func(i1, i2 interface{}) {
			*i1.(*string) = i2.(string)
		},
		envValue: new(string),
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
			v, err := strconv.Atoi(s)

			if err != nil {
				i = nil
				return err
			}

			*i.(*int) = int(v)

			return nil
		},
		setDefault: func(i1, i2 interface{}) {
			*i1.(*int) = i2.(int)
		},
		envValue: new(string),
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
			v, err := strconv.ParseBool(s)

			if err != nil {
				i = nil
				return err
			}

			*i.(*bool) = bool(v)

			return nil
		},
		setDefault: func(i1, i2 interface{}) {
			*i1.(*bool) = i2.(bool)
		},
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
			v, err := strconv.ParseFloat(s, 64)

			if err != nil {
				i = nil
				return err
			}

			*i.(*float64) = float64(v)
			return nil
		},
		setDefault: func(i1, i2 interface{}) {
			*i1.(*float64) = i2.(float64)
		},
	})

	return v
}

func Parse() error {
	errors := make([]string, 0)

	for _, e := range envs {
		err := processEnvVariable(e)

		if err != nil {
			errors = append(errors, fmt.Sprintf("expected: %s type: %s got: %s", e.name, e.varType, *e.envValue))
		}
	}

	if len(errors) > 0 {
		errString := strings.Join(errors, "\n")
		return fmt.Errorf(errString)
	}

	return nil
}

func processEnvVariable(e envVar) error {
	*e.envValue = os.Getenv(e.name)

	if *e.envValue == "" {
		e.setDefault(e.value, e.defaultValue)
		return nil
	}

	err := e.setValue(e.value, *e.envValue)

	if err != nil {
		return err
	}

	return nil
}
