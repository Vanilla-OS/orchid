package main

import (
	"embed"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/orchid/cmdr"
)

const (
	reallyFlag string = "really"
	doitFlag   string = "doit"
	nameFlag   string = "name"
)

//go:embed locales/*.yml
var fs embed.FS
var bean *cmdr.App

func main() {

	bean = cmdr.NewApp("bean", "0.1.1", fs)
	// this is output to the user outside the scope of a command
	//cmdr.Info.Println("I'm a bean")
	// this is written to ~/.local/share/[appname]/[appname].log
	bean.Logger.Println("I'm written to the logs")

	// root command
	root := cmdr.NewCommand(
		bean.Trans("bean.use"),
		bean.Trans("bean.long"),
		bean.Trans("bean.short"),
		nil).
		WithPersistentBoolFlag(
			cmdr.NewBoolFlag(
				doitFlag,
				"d",
				bean.Trans("bean.doitFlag"),
				false))

	bean.CreateRootCommand(root)

	// first child command
	child := cmdr.NewCommand(
		bean.Trans("do.use"),
		bean.Trans("do.long"),
		bean.Trans("do.short"),
		doBean).
		WithBoolFlag(
			cmdr.NewBoolFlag(
				reallyFlag,
				"r",
				bean.Trans("do.reallyFlag"),
				false))
	child.Example = "bean do -r"

	root.AddCommand(child)

	roast := cmdr.NewCommand(
		bean.Trans("roast.use"),
		bean.Trans("roast.long"),
		bean.Trans("roast.short"),
		roast).
		WithStringFlag(
			cmdr.NewStringFlag(
				nameFlag,
				"n",
				bean.Trans("roast.nameFlag"),
				"defaultedBean",
			),
		)
	roast.Example = "bean roast -n arabica"
	root.AddCommand(roast)

	// run the app
	err := bean.Run()
	if err != nil {
		cmdr.Error.Println(err)

	}

}

func doBean(cmd *cobra.Command, args []string) error {
	cmdr.Warning.Println("We are considering doing it.")
	// local flag on do
	really := cmdr.FlagValBool(reallyFlag)
	// persistent flag on root
	doit := cmdr.FlagValBool(doitFlag)
	cmdr.Info.Println("Really?", really)
	cmdr.Info.Println("Do it?", doit)
	b, err := cmdr.Confirm.Show(bean.Trans("do.confirm"))
	if err != nil {
		return err
	}
	if !b {
		cmdr.Warning.Println("They didn't want to do it")
		return nil
	}

	if !really {
		log.Println("they didn't really want to do it")

		return errors.New("didn't do it, supply -r flag to make it happen")
	}

	pb, _ := cmdr.ProgressBar.WithTotal(3).WithTitle("Your Mom").Start()
	time.Sleep(2 * time.Second)
	pb.UpdateTitle("doing the first thing")
	pb.Increment()

	pb.UpdateTitle("doing the second thing")
	time.Sleep(2 * time.Second)
	pb.Increment()

	pb.UpdateTitle("doing the third thing")

	time.Sleep(2 * time.Second)
	pb.Increment()
	pb.Stop()

	cmdr.Success.Println("It is done.")
	return nil

}

func roast(cmd *cobra.Command, args []string) error {
	cmdr.Warning.Println("Warming up the roaster")
	// local flag on do
	var name string
	name = cmdr.FlagValString(nameFlag)
	if name == "" {
		name = "generic bean"
	}
	// persistent flag on root
	doit := cmdr.FlagValBool(doitFlag)
	cmdr.Info.Println("Do it?", doit)
	cmdr.Info.Printf("Roasting %s\n", name)
	var second = time.Second

	var phases = strings.Split("warming roasting cooling off-gassing", " ")

	spinner, _ := cmdr.Spinner.Start("Roasting your beans")
	time.Sleep(second)

	for i := 0; i < len(phases); i++ {
		spinner.UpdateText("Roaster " + phases[i])
		if phases[i] == "cooling" {
			cmdr.Warning.Println("temperatures above normal")
		} else {
			cmdr.Success.Println("Completed " + phases[i])
		}
		time.Sleep(second)
	}
	spinner.Success()
	return nil

}
