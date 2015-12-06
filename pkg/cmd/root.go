package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

const (
	bashCompletionFunc = `
__kubernetes-manifest-scanner_itemlister()
{
    local itemlist_out

    if [ ! -z "${flaghash["-f"]}" ] ; then
        itemlist_out=$(kubernetes-manifest-scanner itemlist -f ${flaghash["-f"]} 2>/dev/null)
    elif [ ! -z "${flaghash["--filename"]}" ] ; then
        __debug "debug -> ${flaghash["--filename"]}"
        itemlist_out=$(kubernetes-manifest-scanner itemlist -f ${flaghash["--filename"]} 2>/dev/null)
    else
        itemlist_out=$(kubernetes-manifest-scanner itemlist 2>/dev/null)
    fi

    if [ ! -z "${itemlist_out}+x" ] ; then
        COMPREPLY=( $( compgen -W "${itemlist_out[*]}" -- "$cur" ) )
    fi
}

__custom_func() {
    case ${last_command} in

        kubernetes-manifest-scanner_sample)
           __kubernetes-manifest-scanner_itemlister $1
           return
            ;;
        *)
            ;;
    esac
}
`
)

func KmsNew(out io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "kubernetes-manifest-scanner",
		Short: "Refer to kubernetes or OpenShift v3 manifest configuration",
		Long:  "Refer to kubernetes or OpenShift v3 manifest configuration",
		Run:   runHelp,
		BashCompletionFunction: bashCompletionFunc,
	}

	cmds.AddCommand(NewCmdSample(out))
	cmds.AddCommand(NewCmdSnippet(out))
	cmds.AddCommand(NewCmdItemList(out))
	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
