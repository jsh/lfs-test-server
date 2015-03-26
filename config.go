package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Configuration struct {
	Listen string `config:"tcp://:8080"`
	Host   string `config:"localhost:8080"`
	Scheme string `config:"https"`
}

var Config = &Configuration{}

const keyPrefix = "HARBOUR"

func init() {
	te := reflect.TypeOf(Config).Elem()
	ve := reflect.ValueOf(Config).Elem()

	for i := 0; i < te.NumField(); i++ {
		sf := te.Field(i)
		name := sf.Name
		field := ve.FieldByName(name)

		envVar := strings.ToUpper(fmt.Sprintf("%s_%s", keyPrefix, name))
		env := os.Getenv(envVar)
		tag := sf.Tag.Get("config")

		if env == "" && tag != "" {
			env = tag
		}

		field.SetString(env)
	}

	if port := os.Getenv("PORT"); port != "" {
		// If $PORT is set, override HARBOUR_LISTEN. This is useful for deploying to Heroku.
		Config.Listen = "tcp://:" + port
	}
}
