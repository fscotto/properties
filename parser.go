package properties

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ParseFunction -- type
type ParseFunction func(string, string) (map[string]string, error)

// Default parse method for parsing key - value file
func defaultParse(path, fileName string) (m map[string]string, err error) {
	absolutePathFile, err := filepath.Abs(filepath.Join(path, fileName))
	if err != nil {
		return nil, err
	}
	file, err := os.Open(absolutePathFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	m = make(map[string]string)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		// check if the line has = sign
		// and process the line. Ignore the rest.
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(strings.Replace(line[equal+1:], "\"", "", -1))
				}
				// assign the values map
				m[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
