package utility

import (
	"os"
	"path/filepath"
)

// BinaryPath return the path where the executable file are
func BinaryPath() (path string, err error) {
	e, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(e), nil
}
