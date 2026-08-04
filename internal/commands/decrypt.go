
package commands

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"pkeys/internal/clipboardio"
	"pkeys/internal/pkcrypto"
)

type DecryptCmd struct {
	Content   string `arg:"" optional:"" help:"Hex-encoded ciphertext to decrypt"`
	File      string `short:"f" help:"Path to a file containing hex-encoded ciphertext"`
	Output    string `short:"o" help:"Write output to a file instead of stdout"`
	Clipboard bool   `short:"c" help:"Read input from / write result to the system clipboard"`
}

// Decrypt
func (c *DecryptCmd) Run() error {
	hexData, err := resolveInput(c.Content, c.File, c.Clipboard)
	if err != nil {
		return err
	}

	// Decode hex ciphertext
	ciphertext, err := hex.DecodeString(strings.TrimSpace(string(hexData)))
	if err != nil {
		return fmt.Errorf("content is not valid hex: %w", err)
	}

	// Get key
	key, err := pkcrypto.GetEnvKey()
	if err != nil {
		return err
	}

	// Decrypt
	dec, err := pkcrypto.Decrypt(key, ciphertext)
	if err != nil {
		return err
	}


	// Write decrypted file
	if c.Output != "" {
		path, err := resolvePath(c.Output, "decrypted", ".dec")
		if err != nil {
			return err
		}

		// Write file
		err = os.WriteFile(path, dec, 0644)
		if err != nil {
			return err
		}

		fmt.Printf("File output written to %s\n", path)
		return nil
	}

	// Write decrypted result to clipboard
	if c.Clipboard {
		if err := clipboardio.Write(dec); err != nil {
			return fmt.Errorf("could not write clipboard contents: %w", err)
		}

		fmt.Println("Decrypted result copied to clipboard")
		return nil
	}

	// Print to stdout
	fmt.Println(string(dec))
	return nil
}

// Flag validation for Kong
func (c *DecryptCmd) Validate() error {
	if c.Content != "" && c.File != "" {
		return fmt.Errorf("cannot pass both content and -f/--file")
	}
	return nil
}
