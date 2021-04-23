//
// filename:  int32.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetInt32(key string) (int32, error) {
	v := Get(key)
	ret, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(ret), nil
}

func GetDefaultInt32(key string, defValue int32) int32 {
	if !Has(key) {
		return defValue
	}
	v, err := GetInt32(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetInt32(key string) int32 {
	v := MustGet(key)
	ret, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(ret)
}

func GetUint32(key string) (uint32, error) {
	v := Get(key)
	ret, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(ret), nil
}

func GetDefaultUint32(key string, defValue uint32) uint32 {
	if !Has(key) {
		return defValue
	}
	v, err := GetUint32(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetUint32(key string) uint32 {
	v := MustGet(key)
	ret, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(ret)
}
