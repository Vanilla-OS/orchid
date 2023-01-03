package cmdr

type App struct {
	Name        string
	RootCommand *Command
	Out         *Out
}

// NewApp creates a new command line application
func NewApp(name string) *App {

	return &App{
		Name: name,
		Out:  newOut(),
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
