package cmdr

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Command struct {
	*cobra.Command
	children []*Command
}

// AddCommand adds a command to the slice and to the underlying
// cobra command.
func (c *Command) AddCommand(commands ...*Command) {
	c.children = append(c.children, commands...)
	for _, cmd := range commands {
		c.Command.AddCommand(cmd.Command)
	}
}

// Children returns the children commands.
func (c *Command) Children() []*Command {
	return c.children
}

func (c *Command) AddBoolFlag(f BoolFlag) {
	c.Command.Flags().BoolP(f.Name, f.Shorthand, f.Value, f.Usage)
	viper.BindPFlag(f.Name, c.Command.Flags().Lookup(f.Name))
}
func (c *Command) AddPersistentBoolFlag(f BoolFlag) {
	c.Command.PersistentFlags().BoolP(f.Name, f.Shorthand, f.Value, f.Usage)
	viper.BindPFlag(f.Name, c.Command.PersistentFlags().Lookup(f.Name))
}

func (c *Command) AddStringFlag(f StringFlag) {
	c.Command.Flags().StringP(f.Name, f.Shorthand, f.Value, f.Usage)
	viper.BindPFlag(f.Name, c.Command.Flags().Lookup(f.Name))
}
func (c *Command) AddPersistentStringFlag(f BoolFlag) {
	c.Command.PersistentFlags().BoolP(f.Name, f.Shorthand, f.Value, f.Usage)
	viper.BindPFlag(f.Name, c.Command.PersistentFlags().Lookup(f.Name))
}

func NewCommand(use, long, short string, runE func(cmd *cobra.Command, args []string) error) *Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE:  runE,
	}
	return &Command{
		Command:  cmd,
		children: make([]*Command, 0),
	}
}

func NewCommandCustom(cmd *cobra.Command) *Command {
	return &Command{
		Command:  cmd,
		children: make([]*Command, 0),
	}
}
