//
// filename:  int16.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetInt8(key string) (int8, error) {
	v := Get(key)
	ret, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(ret), nil
}

func GetDefaultInt8(key string, defValue int8) int8 {
	if !Has(key) {
		return defValue
	}
	v, err := GetInt8(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetInt8(key string) int8 {
	v := MustGet(key)
	ret, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		panic(err)
	}
	return int8(ret)
}

func GetUint8(key string) (uint8, error) {
	v := Get(key)
	ret, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(ret), nil
}

func GetDefaultUint8(key string, defValue uint8) uint8 {
	if !Has(key) {
		return defValue
	}
	v, err := GetUint8(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetUint8(key string) uint8 {
	v := MustGet(key)
	ret, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(ret)
}
