//
// filename:  ints.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetInt(key string) (int, error) {
	v := Get(key)
	ret, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(ret), nil
}

func GetDefaultInt(key string, defValue int) int {
	if !Has(key) {
		return defValue
	}
	v, err := GetInt(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetInt(key string) int {
	v := MustGet(key)
	ret, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(ret)
}

func GetUint(key string) (uint, error) {
	v := Get(key)
	ret, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(ret), nil
}

func GetDefaultUint(key string, defValue uint) uint {
	if !Has(key) {
		return defValue
	}
	v, err := GetUint(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetUint(key string) uint {
	v := MustGet(key)
	ret, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(ret)
}
