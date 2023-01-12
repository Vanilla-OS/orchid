package main

import (
	"embed"

	"github.com/vanilla-os/orchid/cmdr"
)

//go:embed locales/*.yml
var fs embed.FS

func main() {

	oapp := cmdr.NewApp("orchid", fs)

	// root command
	root := cmdr.NewCommand(oapp.Trans("orchid.usage"), oapp.Trans("orchid.long"), oapp.Trans("orchid.short"), nil)
	oapp.CreateRootCommand(root)

	// run the app
	err := oapp.Run()
	if err != nil {
		cmdr.Error.Println(err)

	}

}
