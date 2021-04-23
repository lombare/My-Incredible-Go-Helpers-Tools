//
// filename:  json.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadJsonEnv(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return LoadJsonReader(f)
}

func LoadJsonReader(reader io.Reader) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		Set(k, fmt.Sprint(v))
	}
	return nil
}

func MustLoadJson(file string) {
	if err := LoadJsonEnv(file); err != nil {
		panic(err)
	}
}

func MustLoadJsonReader(reader io.Reader) {
	if err := LoadJsonReader(reader); err != nil {
		panic(err)
	}
}
