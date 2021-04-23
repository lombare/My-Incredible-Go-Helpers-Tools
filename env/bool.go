//
// filename:  bool.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetBool(key string) (bool, error) {
	v := Get(key)
	ret, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func GetDefaultBool(key string, defValue bool) bool {
	if !Has(key) {
		return defValue
	}
	v, err := GetBool(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetBool(key string) bool {
	v := MustGet(key)
	ret, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return ret
}
