package command

import (
	"io"

	"github.com/whalecold/kuMana/pkg/command/clone"

	"github.com/spf13/cobra"
)

func New(_ io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "kubeMana",
		Short: "manager all kinds kubernetes cluster config.",
		Long:  "manager all kinds kubernetes cluster configs.",
	}

	cmds.AddCommand(clone.New())
	return cmds
}
