package cmdr

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2023
	Description: Orchid is a cli helper for VanillaOS projects
*/

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/vanilla-os/orchid"
)

func NewManCommand(title string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "man",
		Short:                 "Generates manpages",
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		Hidden:                true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			/*		manPage, err := mcoral.NewManPage(1, cmd.Root())
					if err != nil {
						return err
					}

					_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
					return err
			*/
			header := &doc.GenManHeader{
				Title:   title,
				Source:  "VanillaOS/orchid man page generator",
				Manual:  "Some Manual",
				Section: "1",
			}
			manpath := path.Join("manpages", orchid.Locale())
			err := os.MkdirAll(manpath, 0755)
			if err != nil {
				return err
			}
			return doc.GenManTree(cmd.Root(), header, manpath)

		},
	}
	return cmd
}
