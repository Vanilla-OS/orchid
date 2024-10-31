package orchid

import (
	"os"
	"strings"

	l "github.com/vanilla-os/orchid/log"
)

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

// Locale returns the two digit locale code
// from the LANG environment variable, or "en"
// if unset.
func Locale() string {
	var locale string
	for _, env := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		locale = os.Getenv(env)
		if strings.TrimSpace(locale) != "" {
			break
		}
	}
	if strings.TrimSpace(locale) == "C" || strings.TrimSpace(locale) == "POSIX" {
		return strings.TrimSpace(locale)
	}

	langs := strings.Split(os.Getenv("LANGUAGE"), ":")
	if len(langs) > 0 {
		return strings.Split(langs[0], "_")[0]
	}
	return locale
}
