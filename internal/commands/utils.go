package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Resolve user passed paths and return final path with default name if needed
func resolvePath(path string, name string, ext string) (string, error) {
	var filePath string

	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		// Unexpected error i.e. permissions error
		return "", err
	}

	if err == nil && info.IsDir() {
		now := time.Now()
		// Path is a dir

		// Create default filename with timestamp to prevent overwrites
		filePath = filepath.Join(path,
			fmt.Sprintf("%s_%s%s",
				name,
				now.Format("2006_01_02_150405"),
				ext,
			))

	} else {
		// Path is a filepath
		// Append default extension if the passed path doesn't already have one
		if filepath.Ext(path) == "" {
			filePath = path + ext
		} else {
			filePath = path
		}
	}

	return filePath, nil
}
