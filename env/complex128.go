//
// filename:  complex128.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetComplex128(key string) (complex128, error) {
	v := Get(key)
	ret, err := strconv.ParseComplex(v, 128)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetDefaultComplex128(key string, defValue complex128) complex128 {
	if !Has(key) {
		return defValue
	}
	v, err := GetComplex128(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetComplex128(key string) complex128 {
	v := MustGet(key)
	ret, err := strconv.ParseComplex(v, 128)
	if err != nil {
		panic(err)
	}
	return ret
}
