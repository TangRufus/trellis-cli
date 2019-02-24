package cmd

import (
	"strings"

	"github.com/mitchellh/cli"
	"github.com/roots/trellis-cli/trellis"
)

type GalaxyCommand struct {
	UI      cli.Ui
	Trellis *trellis.Trellis
}

func (c *GalaxyCommand) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *GalaxyCommand) Synopsis() string {
	return "Commands for Ansible Galaxy"
}

func (c *GalaxyCommand) Help() string {
	helpText := `
Usage: trellis galaxy <subcommand> [<args>]
`

	return strings.TrimSpace(helpText)
}
