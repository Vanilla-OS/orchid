package main

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vanillaos/orchid/cmdr"
)

const (
	reallyFlag string = "really"
	doitFlag   string = "doit"
	nameFlag   string = "name"
)

func main() {
	bean := cmdr.NewApp("bean")
	// this is output to the user
	cmdr.Info.Println("I'm a bean")

	// root command
	root := cmdr.NewCommand("bean <options>", `bean has a long description 
that spans several lines
and has some line breaks

and other interesting things.`, "Bean is a test harness for orchid", nil)
	bean.CreateRootCommand(root)
	root.AddPersistentBoolFlag(
		cmdr.NewBoolFlag(
			doitFlag,
			"d",
			"do the thing",
			false))

	// first child command
	child := cmdr.NewCommand("do", `do has a longer description 
	that spans several lines
	and has some line breaks
	
	and other interesting things.`, "Do things with beans", doBean)
	child.AddBoolFlag(
		cmdr.NewBoolFlag(
			reallyFlag,
			"r",
			"really do it",
			false))

	root.AddCommand(child)

	roast := cmdr.NewCommand("roast <color>", "long description", "Roast warms up your coffee", roast)
	roast.AddStringFlag(
		cmdr.NewStringFlag(
			nameFlag,
			"n",
			"name of the bean",
			"",
		),
	)
	roast.Args = cobra.ExactArgs(1)
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
	if !really {
		log.Println("bad")

		return errors.New("don't do it")
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
