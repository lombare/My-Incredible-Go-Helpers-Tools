//
// filename:  float64.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetFloat64(key string) (float64, error) {
	v := Get(key)
	ret, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetDefaultFloat64(key string, defValue float64) float64 {
	if !Has(key) {
		return defValue
	}
	v, err := GetFloat64(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetFloat64(key string) float64 {
	v := MustGet(key)
	ret, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return ret
}
