package properties

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// StoringFunction -- type
type StoringFunction func(Properties) (*os.File, error)

// FIXME when you remove a property it not work fine
func defaultStore(p Properties) (*os.File, error) {
	absolutePathFile, err := filepath.Abs(filepath.Join(p.Path(), p.FileName()))
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(absolutePathFile, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	// Write this file
	defer file.Close()
	for _, key := range p.GetProperties() {
		value, err := p.Get(key)
		if err != nil {
			return nil, err
		}

		row, err := buildRow(key, value)
		if err != nil {
			return nil, err
		}

		if _, err := file.Write([]byte(row + "\n")); err != nil {
			return nil, err
		}
	}

	return file, nil
}

func buildRow(key, value string) (row string, err error) {
	if key == "" || value == "" {
		return "", errors.New("Key or Value param is a empty string")
	}

	newKey := strings.Replace(key, "\"", "", -1)
	newValue := strings.Join([]string{"\"", value, "\""}, "")
	row = strings.Join([]string{newKey, newValue}, "=")
	return strings.TrimSpace(row), nil
}
