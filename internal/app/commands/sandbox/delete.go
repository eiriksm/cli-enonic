package sandbox

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/enonic/xp-cli/internal/app/util"
	"os"
)

var Delete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a sandbox",
	Aliases: []string{"del", "rm"},
	Action: func(c *cli.Context) error {

		sandbox := ensureSandboxNameExists(c, "Select sandbox to delete:")

		if !acceptToDeleteSandbox(sandbox.Name) {
			os.Exit(1)
		}

		if boxesData := readSandboxesData(); boxesData.Running == sandbox.Name {
			fmt.Fprintf(os.Stderr, "Sandbox '%s' is currently running, stop it first!", sandbox.Name)
			os.Exit(1)
		}

		boxes := getSandboxesUsingDistro(sandbox.Distro)
		if len(boxes) == 1 && boxes[0].Name == sandbox.Name && acceptToDeleteDistro(sandbox.Distro) {
			deleteDistro(sandbox.Distro)
		}

		deleteSandbox(sandbox.Name)
		fmt.Fprintf(os.Stderr, "Sandbox '%s' deleted", sandbox.Name)

		return nil
	},
}

func acceptToDeleteSandbox(name string) bool {
	return util.YesNoPrompt(fmt.Sprintf("WARNING: This can not be undone ! Do you still want to delete sandbox '%s' ?", name))
}

func acceptToDeleteDistro(distro string) bool {
	return util.YesNoPrompt(fmt.Sprintf("Distro '%s' is not used any more. Do you want to delete it ?", distro))
}