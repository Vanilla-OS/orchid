package cmdr

import "github.com/spf13/viper"

type flagBase struct {
	Name      string
	Shorthand string
	Usage     string
}

type BoolFlag struct {
	flagBase
	Value bool
}

type StringFlag struct {
	flagBase
	Value string
}

func NewBoolFlag(name, shorthand, usage string, value bool) BoolFlag {
	return BoolFlag{
		flagBase: flagBase{
			Name:      name,
			Shorthand: shorthand,
			Usage:     usage,
		},
		Value: value,
	}
}

func NewStringFlag(name, shorthand, usage, value string) StringFlag {
	return StringFlag{
		flagBase: flagBase{
			Name:      name,
			Shorthand: shorthand,
			Usage:     usage,
		},
		Value: value,
	}
}

func FlagValBool(name string) bool {
	return viper.GetBool(name)
}
func FlagValString(name string) string {
	return viper.GetString(name)
}
