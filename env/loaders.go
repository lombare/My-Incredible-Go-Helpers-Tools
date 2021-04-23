//
// filename:  loaders.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"fmt"
	"os"
)

func Load(name string) error {
	if st, err := os.Stat(name + ".env"); err != nil && !os.IsNotExist(err) {
		return err
	} else if st != nil && st.Mode().IsRegular() {
		return LoadDotenv(name + ".env")
	} else if st, err = os.Stat(name + ".env.json"); err != nil {
		return err
	} else if st.Mode().IsRegular() {
		return LoadJsonEnv(name + ".env.json")
	}
	return fmt.Errorf("can't find or stat %v.env or %v.env.json", name, name)
}

func MustLoad(name string) {
	if err := Load(name); err != nil {
		panic(err)
	}
}
