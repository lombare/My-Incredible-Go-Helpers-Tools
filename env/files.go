//
// filename:  files.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"io"
	"os"
)

func GetFile(key string) (io.ReadWriteCloser, error) {
	v := Get(key)
	f, err := os.Open(v)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetDefaultFile(key string, def io.ReadWriteCloser) io.ReadWriteCloser {
	if !Has(key) {
		return def
	}
	v, err := GetFile(key)
	if err != nil {
		return def
	}
	return v
}

func MustGetFile(key string) io.ReadWriteCloser {
	v := MustGet(key)
	ret, err := os.Open(v)
	if err != nil {
		panic(err)
	}
	return ret
}

func GetFileContent(key string) ([]byte, error) {
	v := Get(key)
	f, err := os.Open(v)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetDefaultFileContent(key string, def []byte) []byte {
	if !Has(key) {
		return def
	}
	v, err := GetFileContent(key)
	if err != nil {
		return def
	}
	return v
}

func MustGetFileContent(key string) []byte {
	v := MustGet(key)
	f, err := os.Open(v)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bytes
}
