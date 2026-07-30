package commands

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"pkeys/internal/clipboardio"
	"pkeys/internal/pkcrypto"
)

type EncryptCmd struct {
	Content   string `arg:"" optional:"" help:"Text to encrypt"`
	File      string `short:"f" help:"Path to a file to encrypt"`
	Output    string `short:"o" help:"Write output to a file instead of stdout"`
	Clipboard bool   `short:"c" help:"Read input from / write result to the system clipboard"`
}

// Encrypt
func (c *EncryptCmd) Run() error {
	// TODO pull in key from env

	var data []byte

	// Read in filepath
	if c.File != "" {
		f, err := os.Open(c.File)
		if err != nil {
			return err
		}
		defer f.Close()

		data, err = io.ReadAll(f)
		if err != nil {
			return err
		}

	} else if c.Content != "" {
		// Read in passed text
		data = []byte(c.Content)
	} else {
		// Read from clipboard
		clipData, err := clipboardio.Read()
		if err != nil {
			return fmt.Errorf("could not read clipboard contents: %w", err)
		}
		data = clipData
	}

	// Get key
	key, err := pkcrypto.GetEnvKey()
	if err != nil {
		return err
	}

	// Encrypt
	enc, err := pkcrypto.Encrypt(key, data)
	if err != nil {
		return err
	}

	// Write encrypted file
	if c.Output != "" {
		path, err := resolvePath(c.Output, "encrypted", ".enc")
		if err != nil {
			return err
		}

		// Write file
		err = os.WriteFile(path, []byte(hex.EncodeToString(enc)), 0644)
		if err != nil {
			return err
		}

		fmt.Printf("File output written to %s\n", path)
		return nil
	}

	// Write encrypted result to clipboard
	if c.Clipboard {
		if err := clipboardio.Write([]byte(hex.EncodeToString(enc))); err != nil {
			return fmt.Errorf("could not write clipboard contents: %w", err)
		}

		fmt.Println("Encrypted result copied to clipboard")
		return nil
	}

	// Print to stdout
	fmt.Println(hex.EncodeToString(enc))
	return nil
}

// Flag validation for Kong
func (c *EncryptCmd) Validate() error {
	if c.Content != "" && c.File != "" {
		return fmt.Errorf("cannot pass both content and -f/--file")
	}
	if c.Content == "" && c.File == "" && !c.Clipboard {
		return fmt.Errorf("must pass either content, -f/--file, or -c/--clipboard")
	}
	return nil
}
