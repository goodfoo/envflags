package main

import (
	"fmt"

	"github.com/goodfoo/envflags"
)

/*
	Environment varaibable ival and sval will have highest precedence
	command line params -ival and -sval will have next precedence
	and finally defaults
*/

func main() {
	flags := envflags.New()

	i := flags.Int("ival", 1, "provide and ival as a parameter or env var")
	s := flags.String("sval", "flags!", "provide a sval as a parameter or env var")

	flags.Parse()

	fmt.Printf("i = %d\ns = %s\n", *i, *s)
}
