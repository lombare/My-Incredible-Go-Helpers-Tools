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
	if err := os.Setenv(key, value); err != nil {
		panic(err)
	}
}
