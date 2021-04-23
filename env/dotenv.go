//
// filename:  dotenv.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package iem

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func LoadDotenv(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return LoadDotenvReader(f)
}

func LoadDotenvReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		bytes := string(scanner.Bytes())

		if idx := strings.Index(bytes, "="); idx > 0 {
			key := bytes[:idx]
			if idx >= len(bytes) {
				Set(key, "")
				continue
			}
			value := bytes[idx+1:]
			if value != "" && value[0] == '`' || value[0] == '\'' || value[0] == '"' {
				str, err := strconv.Unquote(bytes[idx+1:])
				if err != nil {
					return err
				}
				Set(bytes[:idx], str)
			} else {
				Set(bytes[:idx], bytes[idx+1:])
			}
		} else {
			Set(bytes[:idx], "")
		}
	}

	return nil
}

func MustLoadDotenv(file string) {
	if err := LoadDotenv(file); err != nil {
		panic(err)
	}
}

func MustLoadDotenvReader(reader io.Reader) {
	if err := LoadDotenvReader(reader); err != nil {
		panic(err)
	}
}
