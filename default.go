package orchid

import l "github.com/vanillaos/orchid/log"

// init sets defaults for all orchid libraries
// that can be customized with InitXXX functions.
// This helps to ensure consistency across VanillaOS
// applications.
func init() {
	// log defaults
	l.Prefix(l.DefaultPrefix)
	l.Flags(l.DefaultFlags)

	// other future defaults
}

// InitLog initializes the std logging package
// with the provided prefix and flags.
func InitLog(prefix string, flags int) {
	l.Prefix(prefix)
	l.Flags(flags)
}
