package help

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadFile read file content
func ReadFile(filename string) string {
	bytes_, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(bytes_)
}

// WriteFile write file content
func WriteFile(filename, content string) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
