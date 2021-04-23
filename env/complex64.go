//
// filename:  complex64.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetComplex64(key string) (complex64, error) {
	v := Get(key)
	ret, err := strconv.ParseComplex(v, 64)
	if err != nil {
		return 0, err
	}
	return complex64(ret), nil
}

func GetDefaultComplex64(key string, defValue complex64) complex64 {
	if !Has(key) {
		return defValue
	}
	v, err := GetComplex64(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetComplex64(key string) complex64 {
	v := MustGet(key)
	ret, err := strconv.ParseComplex(v, 64)
	if err != nil {
		panic(err)
	}
	return complex64(ret)
}
