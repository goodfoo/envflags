/*
Package envflags exposes an enhanced flag.Flagset
which injects environment variables at the highest precedence.

see: https://golang.org/pkg/flag/
*/
package envflags

import (
	"flag"
	"os"
	"strings"
)

// FlagSet is a flag.Flagset and exposes all normal Flagset methods.
//
// see: https://golang.org/pkg/flag/
type FlagSet struct {
	flag.FlagSet
	transform func(string) string
}

// New return a new *FlagSet
func New() *FlagSet {
	return &FlagSet{
		flag.FlagSet{},
		strings.ToUpper,
	}
}

// Transform apply a transformation to the flag name when searching the ENVIRONMENT
// Any func string -> string is suitable.
//
// default is https://golang.org/pkg/strings/#ToUpper
func (f *FlagSet) Transform(transform func(string) string) *FlagSet {
	f.transform = transform
	return f
}

// Parse flag.Parse and inject the environment at highest precedence.
func (f *FlagSet) Parse() {
	// get command line stuff and defaults
	f.FlagSet.Parse(os.Args[1:])

	// inject environment
	f.VisitAll(func(flag *flag.Flag) {
		if value, OK := os.LookupEnv(f.transform(flag.Name)); OK {
			flag.Value.Set(value)
		}
	})
}
