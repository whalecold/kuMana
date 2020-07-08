package main

import (
	"log"
	"os"

	"github.com/whalecold/kuMana/pkg/command"

	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := command.New(os.Stdin, os.Stdout, os.Stderr)
	err := doc.GenMarkdownTree(cmd, "./doc")
	if err != nil {
		log.Fatal(err)
	}
}
