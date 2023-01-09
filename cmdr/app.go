package cmdr

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/vanilla-os/orchid"
)

type App struct {
	Name        string
	RootCommand *Command
}

// NewApp creates a new command line application
func NewApp(name string) *App {
	// for application logs
	orchid.InitLog(name+" : 	", 0)

	viper.SetEnvPrefix(name)
	viper.AutomaticEnv()

	return &App{
		Name: name,
	}
}

func (a *App) CreateRootCommand(c *Command) {
	a.RootCommand = c
}

func (a *App) Run() error {
	err := a.ensureLogDir()
	if err != nil {
		return err
	}
	logDir, err := getLogDir(a.Name)
	if err != nil {
		return err
	}
	logFile := path.Join(logDir, a.Name+".log")
	//create your file with desired read/write permissions
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	//set output of logs to f
	log.SetOutput(f)

	if a.RootCommand != nil {

		return a.RootCommand.Execute()
	}

	return errors.New("no root command defined")
}

/*

	apx := cmdr.NewApp("apx")
	apx.CreateRootCommand(rootCmd) // defined elsewhere
	apx.Info.Println("This is information")

*/

func (a *App) ensureLogDir() error {
	logPath, err := getLogDir(a.Name)
	if err != nil {
		return err
	}
	return os.MkdirAll(logPath, 0755)
}

func getLogDir(app string) (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".local", "share", app), nil
}
