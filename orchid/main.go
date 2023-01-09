package main

import (
	"github.com/vanilla-os/orchid/cmdr"
)

func main() {
	oapp := cmdr.NewApp("orchid")
	// this is output to the user

	// root command
	root := cmdr.NewCommand("orchid <options>", "orchid is a cli helper for VanillaOS projects", "orchid is a cli helper for VanillaOS projects", nil)
	oapp.CreateRootCommand(root)

	// run the app
	err := oapp.Run()
	if err != nil {
		cmdr.Error.Println(err)

	}

}
