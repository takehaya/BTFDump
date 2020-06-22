package internal

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/takehaya/btfdump/pkg/utils"
	"github.com/takehaya/btfdump/pkg/version"
	"github.com/urfave/cli"
)

// NewApp is Application Prelude
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "btfdump"
	app.Version = version.Version

	app.Usage = "Dump and analyze BPF Type Format"

	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "elffile",
			Value: "./a.out",
			Usage: "eBPF ELF File",
		},
		cli.StringFlag{
			Name:  "configfile",
			Value: "./scripts/config.yaml",
			Usage: "btfdump config file",
		},
	}
	app.Action = run
	return app
}

func run(ctx *cli.Context) error {
	elffile := ctx.String("elffile")
	// configfile := ctx.String("configfile")

	if !utils.FileExists(elffile) {
		return errors.WithStack(fmt.Errorf("elf file not found: %s\nHave you run 'make'?", elffile))
	}

	return nil
}
