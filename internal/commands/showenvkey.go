package commands

import (
	"encoding/hex"
	"fmt"
	"os"

	"pkeys/internal/pkcrypto"
)

type ShowEnvKey struct {
	Print   bool   `help:"Print entire key to stdout" short:"p"`
	Output bool   `help:"Write key to a file instead of stdout" short:"o"`
	Path   string `arg:"" optional:"" default:"." help:"Path to write file to"`
}

// Retrieve system key - defaults to peeking the key
func (c *ShowEnvKey) Run() error {
	// Get key
	rawKey, err := pkcrypto.GetEnvKey()
	if err != nil {
		return err
	}
	key := hex.EncodeToString(rawKey)

	// Check flags
	if c.Output {
		// if c.Path is a dir, generate a default filename
		path, err := resolvePath(c.Path, "env_key", ".txt")
		if err != nil {
			return err
		}

		// Write file
		err = os.WriteFile(path, []byte(key), 0644)
		if err != nil {
			return err
		}

		fmt.Printf("File output written to %s\n", path)
		return nil
	}

	// Print full key to stdout
	if c.Print {
		fmt.Printf("%s\n", key)
		return nil
	}

	// Peek: print first and last 4 chars
	fmt.Printf(
		"%s...%s\n",
		key[0:4],
		key[len(key)-4:len(key)-1],
	)

	return nil
}
