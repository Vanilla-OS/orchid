package main

import (
	"embed"
	"os"

	"github.com/fitv/go-i18n"
	"github.com/vanilla-os/orchid"
	"github.com/vanilla-os/orchid/cmdr"
)

//go:embed locales/*.yml
var fs embed.FS

func main() {
	i18n, err := i18n.New(fs, "locales")
	if err != nil {
		cmdr.Error.Println(err)
		os.Exit(1)
	}
	i18n.SetDefaultLocale(orchid.Locale())

	oapp := cmdr.NewApp("orchid")
	// this is output to the user

	// root command
	root := cmdr.NewCommand(i18n.Trans("orchid.usage"), i18n.Trans("orchid.long"), i18n.Trans("orchid.short"), nil)
	oapp.CreateRootCommand(root)

	// run the app
	err = oapp.Run()
	if err != nil {
		cmdr.Error.Println(err)

	}

}
