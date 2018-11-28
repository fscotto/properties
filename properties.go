package properties

import (
	e "errors"
	"os"
	"path/filepath"
	"strings"
)

// Pair --
type Pair struct {
	First  string
	Second string
}

// Properties -- type
type Properties struct {
	fileName string
	path     string
	values   map[int]Pair
	length   int
}

// New -- Make new Properties object
func New(path, fileName string) Properties {
	return Properties{
		fileName: fileName,
		path:     filepath.Clean(path),
		values:   make(map[int]Pair),
		length:   0,
	}
}

// FileName -- Getter for filename's property file
func (p Properties) FileName() string {
	return p.fileName
}

// SetFileName -- Setter filename's property file
func (p *Properties) SetFileName(fileName string) {
	p.fileName = fileName
}

// Path -- Getter for path of property file
func (p Properties) Path() string {
	return p.path
}

// SetPath -- Setter for path of property file
func (p *Properties) SetPath(path string) {
	p.path = path
}

// Length -- Getter length of property file
func (p Properties) Length() int {
	return p.length
}

// Values -- Getter values of property file
func (p Properties) Values() map[int]Pair {
	return p.values
}

// Put -- Put key - value in the Properties object
func (p *Properties) Put(key, value string) error {
	if key == "" || len(strings.TrimSpace(key)) == 0 {
		return e.New("Key value is nil")
	}
	if p.values != nil {
		p.values[p.length+1] = Pair{key, value}
		p.length++
	} else {
		return e.New("Property values is nil")
	}
	return nil
}

// Get -- Get value associated with key
func (p Properties) Get(key string) (string, error) {
	if key == "" || len(strings.TrimSpace(key)) == 0 {
		return "", e.New("Key value is nil")
	}
	index := p.index(key)
	if _, ok := p.values[index]; !ok {
		return "", e.New("Key not found")
	}
	return p.values[index].Second, nil
}

// Remove -- Remove property with key
func (p *Properties) Remove(key string) (string, error) {
	if key == "" || len(strings.TrimSpace(key)) == 0 {
		return "", e.New("Key value is nil")
	}
	lenghtBefore := len(p.values)
	index := p.index(key)
	delete(p.values, index)
	if len(p.values) != lenghtBefore {
		p.length--
	}
	return key, nil
}

// GetProperties -- Get all key value in Properties object
func (p Properties) GetProperties() (keys []string) {
	for i := range p.values {
		key := p.values[i].First
		keys = append(keys, key)
	}
	return keys
}

func (p Properties) index(key string) int {
	index := -1
	for i, item := range p.values {
		if item.First == key {
			index = i
			break
		}
	}
	return index
}

// DefaultLoad -- Load file in Properties object using default parse function
func (p *Properties) DefaultLoad() (int, error) {
	return p.Load(defaultParse)
}

// Load -- Load file in Properties object with specific parse function
func (p *Properties) Load(pf ParseFunction) (int, error) {
	m, err := pf(p.path, p.fileName)
	if err != nil {
		return 0, err
	}
	p.values = m
	p.length = len(m)
	return len(m), nil
}

// DefaultStore -- Create or modify property file
func (p Properties) DefaultStore() (*os.File, error) {
	return p.Store(defaultStore)
}

// Store -- Create or modify property file with specific StoringFunction
func (p Properties) Store(sf StoringFunction) (*os.File, error) {
	file, err := sf(p)
	if err != nil {
		return nil, err
	}
	return file, nil
}
