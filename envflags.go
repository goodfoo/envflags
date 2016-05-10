/*
	Package envflags exposes an enhanced flag.Flagset which
	is a normal Flagset except it also injects environment variables
	at the highest precedence.
*/

package envflags

import (
	"flag"
	"os"
)

// FlagSet has-a flag.Flagset
type FlagSet struct {
	flag.FlagSet
}

// Parse inject the environment at highest precedence
func (f *FlagSet) Parse() {
	// get command line stuff and defaults
	f.FlagSet.Parse(os.Args[1:])
	// inject environment
	f.VisitAll(func(flag *flag.Flag) {
		if value, OK := os.LookupEnv(flag.Name); OK {
			flag.Value.Set(value)
		}
	})
}
