//
// filename:  int8.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetInt16(key string) (int16, error) {
	v := Get(key)
	ret, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(ret), nil
}

func GetDefaultInt16(key string, defValue int16) int16 {
	if !Has(key) {
		return defValue
	}
	v, err := GetInt16(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetInt16(key string) int16 {
	v := MustGet(key)
	ret, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(ret)
}

func GetUint16(key string) (uint16, error) {
	v := Get(key)
	ret, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(ret), nil
}

func GetDefaultUint16(key string, defValue uint16) uint16 {
	if !Has(key) {
		return defValue
	}
	v, err := GetUint16(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetUint16(key string) uint16 {
	v := MustGet(key)
	ret, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(ret)
}
