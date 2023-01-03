package cmdr

import "github.com/pterm/pterm"

type Out struct {
	Info, Warning, Success, Fatal, Debug, Description, Error pterm.PrefixPrinter
}

func newOut() *Out {
	return &Out{
		Error:       pterm.Error,
		Info:        pterm.Info,
		Warning:     pterm.Warning,
		Success:     pterm.Success,
		Fatal:       pterm.Fatal,
		Debug:       pterm.Debug,
		Description: pterm.Description,
	}
}
