package main

import (
	"fmt"
	"os"

	"github.com/whalecold/kuMana/pkg/command"
)

func main() {

	command := command.New(os.Stdin, os.Stdout, os.Stderr)
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
