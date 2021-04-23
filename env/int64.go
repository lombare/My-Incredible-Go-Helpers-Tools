//
// filename:  int64.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"strconv"
)

func GetInt64(key string) (int64, error) {
	v := Get(key)
	ret, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetDefaultInt64(key string, defValue int64) int64 {
	if !Has(key) {
		return defValue
	}
	v, err := GetInt64(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetInt64(key string) int64 {
	v := MustGet(key)
	ret, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func GetUint64(key string) (uint64, error) {
	v := Get(key)
	ret, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetDefaultUint64(key string, defValue uint64) uint64 {
	if !Has(key) {
		return defValue
	}
	v, err := GetUint64(key)
	if err != nil {
		return defValue
	}
	return v
}

func MustGetUint64(key string) uint64 {
	v := MustGet(key)
	ret, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}
