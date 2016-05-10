/*
	Package envflags exposes an enhanced flag.Flagset which
	is a normal Flagset except it also injects environment variables
	at the highest precedence.
*/

package envflags

import (
	"flag"
	"os"
	"strings"
)

// FlagSet has-a flag.Flagset
type FlagSet struct {
	flag.FlagSet
	uppercase bool
}

// New return a new *FlagSet
func New() *FlagSet {
	return &FlagSet{}
}

// Upper builder style uppercase transform
func (f *FlagSet) Upper() *FlagSet {
	f.uppercase = true
	return f
}

// Parse inject the environment at highest precedence
func (f *FlagSet) Parse() {
	// get command line stuff and defaults
	f.FlagSet.Parse(os.Args[1:])

	transformer := func(s string) string { return s }
	if f.uppercase {
		transformer = strings.ToUpper
	}
	// inject environment
	f.VisitAll(func(flag *flag.Flag) {
		if value, OK := os.LookupEnv(transformer(flag.Name)); OK {
			flag.Value.Set(value)
		}
	})
}
