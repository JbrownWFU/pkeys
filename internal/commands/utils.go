package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"pkeys/internal/clipboardio"
)

// resolveInput determines the input bytes for encrypt/decrypt, in order of
// precedence: positional content, -f/--file, -c/--clipboard, piped stdin,
// falling back to the clipboard if stdin isn't piped (e.g. bare invocation
// in an interactive terminal).
func resolveInput(content string, file string, useClipboard bool) ([]byte, error) {
	switch {
	case file != "":
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		return io.ReadAll(f)

	case content != "":
		return []byte(content), nil

	case useClipboard:
		data, err := clipboardio.Read()
		if err != nil {
			return nil, fmt.Errorf("could not read clipboard contents: %w", err)
		}
		return data, nil

	default:
		if isStdinPiped() {
			return io.ReadAll(os.Stdin)
		}

		data, err := clipboardio.Read()
		if err != nil {
			return nil, fmt.Errorf("could not read clipboard contents: %w", err)
		}
		return data, nil
	}
}

// isStdinPiped reports whether stdin is connected to a pipe/redirect rather
// than an interactive terminal.
func isStdinPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

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
