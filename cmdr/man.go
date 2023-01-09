package cmdr

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2023
	Description: Orchid is a cli helper for VanillaOS projects
*/

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
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
				Section: "1",
			}
			return doc.GenManTree(cmd.Root(), header, "manpages")

		},
	}
	return cmd
}
