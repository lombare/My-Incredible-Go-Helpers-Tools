package iem

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
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
			if idx >= len(bytes) {
				Set(bytes[:idx], "")
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

func LoadJsonEnv(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return LoadJsonReader(f)
}

func LoadJsonReader(reader io.Reader) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		return err
	}
	for k, v := range mp {
		Set(k, fmt.Sprint(v))
	}
	return nil
}

func MustLoadJson(file string) {
	if err := LoadJsonEnv(file); err != nil {
		panic(err)
	}
}

func MustLoadJsonReader(reader io.Reader) {
	if err := LoadJsonReader(reader); err != nil {
		panic(err)
	}
}

func Load(name string) error {
	if st, err := os.Stat(name + ".env"); err != nil && !os.IsNotExist(err) {
		return err
	} else if st != nil && st.Mode().IsRegular() {
		return LoadDotenv(name + ".env")
	} else if st, err = os.Stat(name + ".env.json"); err != nil {
		return err
	} else if st.Mode().IsRegular() {
		return LoadJsonEnv(name + ".env.json")
	}
	return fmt.Errorf("can't find or stat %v.env or %v.env.json", name, name)
}

func MustLoad(name string) {
	if err := Load(name); err != nil {
		panic(err)
	}
}
