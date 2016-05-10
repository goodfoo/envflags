# envflags
Go Flags plus Environment overrides

This is a regular [flag.FlagSet](https://golang.org/pkg/flag/) which, after is resolves the flags, will then inject ENVIRONMENT varaiables if they are set.

This has the affect of promoting ENVIRONMENT vars to highest precedence.
