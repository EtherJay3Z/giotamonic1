package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mcpride/giotamonic/iota_mnemonic"
	"github.com/tyler-smith/go-bip39"
)

// ToSeedCommand sets up a "to-seed" command to convert mnemonic to iota seed
func NewCommand() cli.Command {

	return cli.Command{
		Name:        "new",
		Usage:       "Create new mnemonic and iota seed",
		Description: "Creates new BIP0039 mnemonic word list and corresponding - optional passphrase encrypted - iota seed.",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "passphrase", Usage: "Additional passphrase", Value: ""},
			cli.IntFlag{Name: "bitsize", Usage: "Bit length of the mnemonic entropy. Must be [128, 256] and a multiple of 32", Value: 256},
		},
		Action: newAction,
	}
}

// ToSeedCommand sets up a "to-seed" command to convert mnemonic to iota seed
func ToSeedCommand() cli.Command {

	return cli.Command{
		Name:        "to-seed",
		Usage:       "Convert mnemonic to iota seed",
		Description: "Converts mnemonic words to iota seed.",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "passphrase", Usage: "Additional passphrase", Value: ""},
		},
		Action: toSeedAction,
	}
}

func toSeedAction(c *cli.Context) {

	mnemonic := ""

	if len(c.Args()) != 1 {
		reader := bufio.NewReader(os.Stdin)
		mnemonic, _ = reader.ReadString('\n')
	} else {
		mnemonic = c.Args()[0]
	}

	passphrase := ""
	if c.IsSet("passphrase") {
		passphrase = c.String("passphrase")
	}
	if mnemonic != "" {
		seed, err := iota_mnemonic.ToSeed(mnemonic, passphrase)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Fprint(os.Stdout, seed)
	} else {
		fmt.Fprintln(os.Stderr, "Mnemonic must be provided.")
		os.Exit(1)
	}
}

func newAction(c *cli.Context) {

	passphrase := ""
	if c.IsSet("passphrase") {
		passphrase = c.String("passphrase")
	}

	bitSize := 256
	if c.IsSet("bitsize") {
		bitSize = c.Int("bitsize")
	}

	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	seed, err := iota_mnemonic.ToSeed(mnemonic, passphrase)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, "Mnemonic: "+mnemonic+"\n")
	fmt.Fprint(os.Stdout, "Seed: "+seed)
}

func main() {
	app := cli.NewApp()
	app.Name = "giotamonic"
	app.Version = "0.1.0"
	app.Usage = "Generate, restore IOTA seed from Bitcoin BIP39 mnemonic word list."
	app.Flags = []cli.Flag{}
	app.Author = "M. Stolze (mcpride)"
	app.Email = ""
	app.Commands = []cli.Command{
		NewCommand(),
		ToSeedCommand(),
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
