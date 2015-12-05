package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

func KmsNew(out io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "kubernetes-manifest-scanner",
		Short: "Refer to kubernetes or OpenShift v3 manifest configuration",
		Long:  "Refer to kubernetes or OpenShift v3 manifest configuration",
		Run:   runHelp,
	}

	cmds.AddCommand(NewCmdSample(out))
	cmds.AddCommand(NewCmdSnippet(out))
	cmds.AddCommand(NewCmdItemList(out))
	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
