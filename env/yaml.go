//
// filename:  yaml.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYaml(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return LoadYamlReader(f)
}

func LoadYamlReader(reader io.Reader) error {
	d := yaml.NewDecoder(reader)
	mp := make(map[string]interface{})
	if err := d.Decode(mp); err != nil {
		return err
	}
	for k, v := range mp {
		Set(k, fmt.Sprint(v))
	}
	return nil
}

func MustLoadYaml(file string) {
	if err := LoadYaml(file); err != nil {
		panic(err)
	}
}

func MustLoadYamlReader(reader io.Reader) {
	if err := LoadYamlReader(reader); err != nil {
		panic(err)
	}
}
