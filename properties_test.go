package properties_test

import (
	"os"
	"path/filepath"
	"testing"

	prop "bitbucket.org/fabio_scotto_di_santolo/properties"
)

const FILENAME = "ini_test.properties"

func TestNew(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if p.FileName() != FILENAME {
		t.Failed()
	}
	if p.Path() != path {
		t.Failed()
	}
	if p.Length() != 0 {
		t.Failed()
	}
}

func TestFileName(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if p.FileName() != FILENAME {
		t.Failed()
	}
}

func TestSetFileName(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if p.FileName() != FILENAME {
		t.Failed()
	}
	if p.SetFileName("TEST"); p.FileName() == FILENAME {
		t.Failed()
	}

	if p.FileName() != "TEST" {
		t.Failed()
	}
}

func TestPath(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if p.Path() != path {
		t.Failed()
	}
}

func TestSetPath(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if p.Path() != path {
		t.Failed()
	}
	if p.SetPath("TEST"); p.Path() == path {
		t.Failed()
	}

	if p.Path() != "TEST" {
		t.Failed()
	}
}

func TestPutErrorKeyNoValue(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if err := p.Put("", "TEST"); err == nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] %s\n", err.Error())
	}
}

func TestPutErrorKeyValueWithSpaces(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if err := p.Put("    ", "TEST"); err == nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] %s\n", err.Error())
	}
}

func TestPutErrorPropertyNil(t *testing.T) {
	var p prop.Properties
	if err := p.Put("KEY", "VALUE"); err == nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] %s\n", err.Error())
	}
}

func TestPut(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if err := p.Put("KEY", "VALUE"); err != nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] Length of property: %d\n", p.Length())
	}
}

func TestGetErrorKeyNoValue(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if _, err := p.Get(""); err == nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] %s\n", err.Error())
	}
}

func TestGetErrorKeyValueWithSpaces(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if _, err := p.Get("     "); err == nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] %s\n", err.Error())
	}
}

func TestGetErrorKeyNotFound(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	if _, err := p.Get("TEST"); err == nil {
		t.Failed()
	} else {
		t.Logf("\n[DEBUG] %s\n", err.Error())
	}
}

func TestGet(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	p.Put("key1", "value1")
	p.Put("key2", "value2")
	p.Put("key3", "value3")
	p.Put("key4", "value4")
	p.Put("key5", "value5")
	if value, err := p.Get("key4"); err != nil {
		t.Failed()
	} else if value != "value4" {
		t.Failed()
	}
}

func TestGetProperties(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	p.Put("key1", "value1")
	p.Put("key2", "value2")
	p.Put("key3", "value3")
	p.Put("key4", "value4")
	p.Put("key5", "value5")
	keys := p.GetProperties()
	length := len(keys)
	if length != p.Length() {
		t.Failed()
	}
}

func TestRemove(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	p.Put("key1", "value1")
	p.Put("key2", "value2")
	p.Put("key3", "value3")
	p.Put("key4", "value4")
	p.Put("key5", "value5")
	lengthBefore := len(p.Values())
	p.Remove("key4")
	if p.Length() >= lengthBefore {
		t.Logf("\n[!!] Failed error %s\n", "Property length is major")
		t.Failed()
	}
}

func TestRemoveWithKeyNotFound(t *testing.T) {
	path, _ := filepath.Abs(FILENAME)
	p := prop.New(path, FILENAME)
	p.Put("key1", "value1")
	p.Put("key2", "value2")
	p.Put("key3", "value3")
	p.Put("key4", "value4")
	p.Put("key5", "value5")
	lengthBefore := len(p.Values())
	if _, err := p.Remove("KEY_NOT_FOUND"); err != nil {
		t.Logf("\n[!!] Failed error %s\n", err.Error())
		t.Failed()
	}
	if p.Length() != lengthBefore {
		t.Logf("\n[!!] Failed error %s\n", "Property length is different")
		t.Failed()
	}
}

func TestDefaultLoad(t *testing.T) {
	path, _ := os.Getwd()
	p := prop.New(path, FILENAME)
	if rowNumber, err := p.DefaultLoad(); rowNumber == 3 {
		if p.Length() != rowNumber {
			t.Failed()
		}
		for key, value := range p.Values() {
			t.Logf("%s = %s\n", key, value)
		}
	} else if err != nil {
		t.Logf("\n[!!] Failed error %s\n", err.Error())
		t.Failed()
	}
}

func TestDefaultStoreModifyOneValue(t *testing.T) {
	path, _ := os.Getwd()
	p := prop.New(path, FILENAME)
	if _, err := p.DefaultLoad(); err != nil {
		t.Logf("\n[!!] Failed error %s\n", err.Error())
		t.FailNow()
	}
	keys := p.GetProperties()
	if len(keys) > 0 {
		key := keys[0]
		p.Put(key, "MODIFY_TEST")

		file, err := p.DefaultStore()
		if err != nil {
			t.Logf("\n[!!] Failed error %s\n", err.Error())
			t.Failed()
		}

		if file == nil {
			t.Logf("\n[!!] Failed error %s\n", "File is nil")
			t.Failed()
		}
	}
}

func TestDefaultStoreAddNewValue(t *testing.T) {
	path, _ := os.Getwd()
	p := prop.New(path, FILENAME)
	if _, err := p.DefaultLoad(); err != nil {
		t.Logf("\n[!!] Failed error %s\n", err.Error())
		t.FailNow()
	}

	p.Put("NEW_KEY", "NEW_VALUE")
	file, err := p.DefaultStore()
	if err != nil {
		t.Logf("\n[!!] Failed error %s\n", err.Error())
		t.Failed()
	}

	if file == nil {
		t.Logf("\n[!!] Failed error %s\n", "File is nil")
		t.Failed()
	}
}

func TestDefaultStoreRemoveOneValue(t *testing.T) {
	path, _ := os.Getwd()
	p := prop.New(path, "testfile.properties")

	// Add values in Properties object
	p.Put("key1", "value1")
	p.Put("key2", "value2")
	p.Put("key3", "value3")
	p.Put("key4", "value4")
	p.Put("key5", "value5")

	if p.Length() != 5 {
		t.Logf("\n[!!] Failed error %s\n", "File is nil")
		t.Failed()
	}

	p.Remove("key3")

	file, err := p.DefaultStore()
	if err != nil {
		t.Logf("\n[!!] Failed error %s\n", err.Error())
		t.Failed()
	}

	if file == nil {
		t.Logf("\n[!!] Failed error %s\n", "File is nil")
		t.Failed()
	}
}
