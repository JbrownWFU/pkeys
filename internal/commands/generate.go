package commands

import (
	"fmt"
	"runtime"
	"os"

	"pkeys/internal/pkcrypto"
)

type GenerateCmd struct {
	Output string `short:"o" help:"Write output to a file instead of stdout"`
	Shell  bool   `help:"Print with instructions to set for your system" short:"s"`
}

// Generate a new key
func (c *GenerateCmd) Run() error {
	key, err := pkcrypto.Generate()
	if err != nil {
		return err
	}

	// Global write out
	if c.Output != "" {
		// if c.Output is a dir, generate a default filename
		path, err := resolvePath(c.Output, "key", ".txt")
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

	// Print with help
	if c.Shell {
		switch runtime.GOOS {
		case "windows":
			fmt.Printf("Manually edit your system variables (Settings -> System Variables) and add:\n\nPKEYS_KEY = %s\n\nOr set through terminal: \n\nsetx PKEYS_KEY %s\n\nOpen a new terminal window to reload\n", key, key)
			return nil
		default:
			// Mac / Linux
			fmt.Printf("Add this to your ~/.zshrc (or ~/.bashrc):\n\nexport PKEYS_KEY=%s\n\nthen reload your terminal session:\n\nsource ~/.zshrc\n", key)
			return nil
		}
	}

	fmt.Printf("%s\n", key)
	return nil
}
