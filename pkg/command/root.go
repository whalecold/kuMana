package command

import (
	"io"

	"github.com/whalecold/kuMana/pkg/command/clone"

	"github.com/spf13/cobra"
)

func New(_ io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "kuMana",
		Short: "manag all kinds kubernetes cluster config.",
		Long:  "manag all kinds kubernetes cluster configs.",
	}

	cmds.AddCommand(clone.New())
	return cmds
}
