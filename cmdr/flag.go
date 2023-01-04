package cmdr

type BoolFlag struct {
	Name      string
	Shorthand string
	Usage     string
	Value     bool
}

func NewBoolFlag(name, shorthand, usage string, value bool) BoolFlag {
	return BoolFlag{
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		Value:     value,
	}
}
