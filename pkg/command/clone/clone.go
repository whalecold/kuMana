package clone

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cloneExample = `
# kubeMana clone -p password -u user -h host
`

type options struct {
	password string
	user     string
	host     string
}

func New() *cobra.Command {
	opts := options{}
	cmds := &cobra.Command{
		Use:     "clone",
		Short:   "clone the kubernetes cluster config from specify host",
		Long:    "clone the kubernetes cluster config from specify host",
		Example: cloneExample,
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.clone(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
		},
	}
	cmds.Flags().StringVar(&opts.host, "host", opts.host, "the host name, [required]")
	cmds.Flags().StringVar(&opts.password, "passwd", opts.password, "the password of host, [required]")
	cmds.Flags().StringVar(&opts.user, "user", "root", "the host user, default root, [optional]")
	return cmds
}

func (opts *options) validate() error {
	if opts.password == "" {
		return fmt.Errorf("passwd should't be empty")
	}
	if opts.host == "" {
		return fmt.Errorf("host should't be empty")
	}
	return nil
}

func (opts *options) clone() error {
	if err := opts.validate(); err != nil {
		return err
	}

	dstPath := fmt.Sprintf("%s@%s:/root/.kube/config", opts.user, opts.host)
	localPath := fmt.Sprintf("%s/.kube/config.%s", os.Getenv("HOME"), opts.host)

	_, err := os.Stat(localPath)
	if err == nil {
		// if exist, remove it.
		if err = os.Remove(localPath); err != nil {
			return err
		}
	}

	args := []string{"-p", opts.password, "scp", dstPath, localPath}
	exeCmd := exec.Command("sshpass", args...)

	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr

	if err := exeCmd.Run(); err != nil {
		return err
	}
	return nil
}
