package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func KmsNew(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "kubernetes-manifest-scanner",
		Short: "Refer to kubernetes or OpenShift v3 manifest configuration",
		Long:  "Refer to kubernetes or OpenShift v3 manifest configuration",
		Run:   runHelp,
	}

	cmds.AddCommand(cmd.NewCmdSample(out))
	cmds.AddCommand(cmd.NewCmdSnippet(out))
	cmds.AddCommand(cmd.NewCmdItemList(out))
	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
