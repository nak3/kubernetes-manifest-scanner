package main

import (
	"io"
	"os"
	"runtime"

	"github.com/nak3/kubernetes-manifest-scanner/pkg/cmd"
	"github.com/spf13/cobra"
)

func main() {

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	cmd := cmdNew(os.Stdin, os.Stdout, os.Stderr)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func cmdNew(in io.Reader, out, err io.Writer) *cobra.Command {
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
