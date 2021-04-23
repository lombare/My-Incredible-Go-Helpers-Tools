//
// filename:  float32.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetFloat32(key string) (float32, error) {
	v := Get(key)
	ret, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, err
	}
	return float32(ret), nil
}

func GetDefaultFloat32(key string, defValue float32) float32 {
	if !Has(key) {
		return defValue
	}
	v, err := GetFloat32(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetFloat32(key string) float32 {
	v := MustGet(key)
	ret, err := strconv.ParseFloat(v, 32)
	if err != nil {
		panic(err)
	}
	return float32(ret)
}
