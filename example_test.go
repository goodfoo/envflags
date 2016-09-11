package envflags_test

import (
	"fmt"
	"os"

	"github.com/goodfoo/envflags"
)

/*
	Environment variabable IVAL and SVAL will have highest precedence
	command line params -ival and -sval will have next precedence
	and finally defaults
*/

func ExampleParse() {
	// compatible with std lib flags
	// optional conversion from flag name to env
	// default is strings.ToUpper
	flags := envflags.New() //.Transform(strings.ToLower)

	// define flags as usual
	i := flags.Int("ival", 1, "provide and ival as a parameter or env var")
	s := flags.String("sval", "flags", "provide a sval as a parameter or env var")
	// this happens if user runs: go test -v
	flags.Bool("test.v", false, "verbosity")

	// simulate environment settings
	os.Setenv("SVAL", "awesome flags!")

	// Parse as usual
	// ENV is evaluated and injected where matched
	flags.Parse()

	fmt.Printf("i = %d\ns = %s", *i, *s)
	// Output:
	// i = 1
	// s = awesome flags!
}
