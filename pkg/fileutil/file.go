package fileutil

import (
	"io/ioutil"
)

// Utility functions for file handling
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}
