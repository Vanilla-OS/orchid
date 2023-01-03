package cmdr

import "github.com/vanillaos/orchid"

type App struct {
	Name        string
	RootCommand *Command
	Out         *Out
}

// NewApp creates a new command line application
func NewApp(name string) *App {
	// for application logs
	orchid.InitLog(name, 0)
	return &App{
		Name: name,
		// for console output
		Out: newOut(),
	}
}

func (a *App) CreateRootCommand(c *Command) {
	a.RootCommand = c
}

/*

	apx := cmdr.NewApp("apx")
	apx.CreateRootCommand(rootCmd) // defined elsewhere
	apx.Out.Info.Println("This is information")

*/
