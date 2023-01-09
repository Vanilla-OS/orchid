package main

import (
	"github.com/vanilla-os/orchid/cmdr"
)

func main() {
	bean := cmdr.NewApp("orchid")
	// this is output to the user

	// root command
	root := cmdr.NewCommand("orchid <options>", "orchid is a cli helper for VanillaOS projects", "orchid is a cli helper for VanillaOS projects", nil)
	bean.CreateRootCommand(root)

	// run the app
	err := bean.Run()
	if err != nil {
		cmdr.Error.Println(err)

	}

}
