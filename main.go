package main

import (
	_ "time/tzdata" // embed IANA timezone database for containers without tzdata

	"github.com/edyoCampos/base365/cmd"
)

func main() {
	cmd.Execute()
}
