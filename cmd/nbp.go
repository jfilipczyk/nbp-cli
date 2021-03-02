package main

import (
	"os"

	"github.com/jfilipczyk/nbp/internal/pkg/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
