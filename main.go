package main

import (
	"os"

	"pkeys/internal/commands"

	"github.com/alecthomas/kong"
)

const appName = "pkeys"

var appVersion = "dev"

// Root
type CLI struct {
	Generate commands.GenerateCmd `cmd:"" help:"Generate a new unique key"`
	Encrypt  commands.EncryptCmd  `cmd:"" help:"Encrpy content"`
	Decrypt  commands.DecryptCmd  `cmd:"" help:"Decrypt content"`
	Show     commands.ShowEnvKey   `cmd:"" help:"Show the currently set key"`
	Version  kong.VersionFlag      `help:"Show version"`
}
func main() {
	// Print help if no args passed
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "--help")
	}

	var cli CLI
	// Parse and execute
	ctx := kong.Parse(&cli,
		kong.Name(appName),
		kong.Description("A CLI encryption tool for plaintext and files\n\nSet the system 'PKEYS_KEY' environment variable to the generated key"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.DefaultEnvars(appName),
		kong.Vars{"version": appVersion},
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
