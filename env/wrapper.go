//
// filename:  wrapper.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"os"
)

func ClearEnv() {
	os.Clearenv()
}

func Has(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

func Set(key, value string) {
	if !Has(key) {
		if err := os.Setenv(key, value); err != nil {
			panic(err)
		}
	}
}
