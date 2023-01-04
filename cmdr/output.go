package cmdr

import "github.com/pterm/pterm"

var (
	Info, Warning, Success, Fatal, Debug, Description, Error pterm.PrefixPrinter
)

func init() {
	Error = pterm.Error
	Info = pterm.Info
	Warning = pterm.Warning
	Success = pterm.Success
	Fatal = pterm.Fatal
	Debug = pterm.Debug
	Description = pterm.Description
}
