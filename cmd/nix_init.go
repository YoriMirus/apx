package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2023
	Description: Apx is a wrapper around multiple package managers to install packages and run commands inside a managed container.
*/

import (
	"github.com/spf13/cobra"
	"github.com/vanilla-os/apx/core"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewNixInitCommand() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"init",
		apx.Trans("nixinit.long"),
		apx.Trans("nixinit.short"),
		initNix,
	)
	cmd.Example = "apx nix init"
	return cmd
}
func initNix(cmd *cobra.Command, args []string) error {
	// prompt for confirmation

	b, err := cmdr.Confirm.Show(apx.Trans("nixinit.confirm"))

	if err != nil {
		return err
	}

	if !b {
		cmdr.Info.Println(apx.Trans("apx.cxl"))
		return nil
	}
	err = core.NixInit()
	if err != nil {
		return err
	}
	cmdr.Success.Println(apx.Trans("nixinit.success"))
	return nil

}