package main

import (
	"github.com/jfilipczyk/nbp-cli/internal/cmd"
)

var (
	version   = "0.0.0" // set by goreleaser during build
	commit    = ""      // set by goreleaser during build
	buildDate = ""      // set by goreleaser during build
)

func main() {
	cfg := &cmd.Config{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	}
	cmd.Execute(cfg)
}
