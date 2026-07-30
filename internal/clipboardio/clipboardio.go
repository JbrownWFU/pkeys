package clipboardio

import (
	"fmt"

	"golang.design/x/clipboard"
)

// Read returns the current text contents of the system clipboard
func Read() ([]byte, error) {
	if err := clipboard.Init(); err != nil {
		return nil, fmt.Errorf("could not initialize clipboard: %w", err)
	}

	data := clipboard.Read(clipboard.FmtText)
	if len(data) == 0 {
		return nil, fmt.Errorf("clipboard is empty")
	}

	return data, nil
}

// Write overwrites the system clipboard with the given text contents
func Write(data []byte) error {
	if err := clipboard.Init(); err != nil {
		return fmt.Errorf("could not initialize clipboard: %w", err)
	}

	clipboard.Write(clipboard.FmtText, data)
	return nil
}
