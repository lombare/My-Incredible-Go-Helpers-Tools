package iem

import (
	"fmt"
	"os"
)

func Get(key string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return ""
}

func GetDefault(key, defValue string) string {
	if !Has(key) {
		return defValue
	}
	return Get(key)
}

func MustGet(key string) string {
	if !Has(key) {
		panic(fmt.Errorf("you must provide the variable %v in your environment"))
	}
	return Get(key)
}
