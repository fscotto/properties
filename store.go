package properties

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

// StoringFunction -- type
type StoringFunction func(Properties) (*os.File, error)

func defaultStore(p Properties) (*os.File, error) {
	absolutePathFile, err := filepath.Abs(filepath.Join(p.Path(), p.FileName()))
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(absolutePathFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	// Write this file
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, pair := range p.values {
		line := fmt.Sprintf("%s=%s\n", escape(pair.First, true), escape(pair.Second, false))
		log.Print(line)
		if _, err := writer.Write([]byte(line)); err != nil {
			return nil, err
		}
	}

	if err := writer.Flush(); err != nil {
		return nil, err
	}
	return file, nil
}

// escape returns a string that is safe to use as either a key or value in a
// property file. Whitespace characters, key separators, and comment markers
// should always be escaped.
func escape(s string, key bool) string {

	leading := true
	var buf bytes.Buffer
	for _, ch := range s {
		wasSpace := false
		if ch == '\t' {
			buf.WriteString(`\t`)
		} else if ch == '\n' {
			buf.WriteString(`\n`)
		} else if ch == '\r' {
			buf.WriteString(`\r`)
		} else if ch == '\f' {
			buf.WriteString(`\f`)
		} else if ch == ' ' {
			if key || leading {
				buf.WriteString(`\ `)
				wasSpace = true
			} else {
				buf.WriteRune(ch)
			}
		} else if ch == ':' {
			buf.WriteString(`\:`)
		} else if ch == '=' {
			buf.WriteString(`\=`)
		} else if ch == '#' {
			buf.WriteString(`\#`)
		} else if ch == '!' {
			buf.WriteString(`\!`)
		} else if !unicode.IsPrint(ch) || ch > 126 {
			buf.WriteString(fmt.Sprintf(`\u%04x`, ch))
		} else {
			buf.WriteRune(ch)
		}

		if !wasSpace {
			leading = false
		}
	}
	return buf.String()
}
