package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/nak3/jvmap"
	"github.com/nak3/kubernetes-manifest-scanner/pkg/logic"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

func NewCmdSnippet(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "snippet -f FILENAME RESOURCE",
		Short: "Refer a configuration parameter as snippet"
		Long:  "Refer a configuration parameter as snippet"
		//		Example: get_example,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(validateArgsSnippet(cmd, args))
			// cmdutil.CheckErr(cmdutil.ValidateOutputArgs(cmd))
			cmdutil.CheckErr(RunSnippet(cmd, args[0]))
		},
	}
	cmd.MarkFlagRequired("filename")
	cmd.PersistentFlags().StringP("filename", "f", "https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/swagger-spec/v1.json", "Path to swagger API json")
	cmd.PersistentFlags().BoolP("insecure", "k", false, "Allow insecure SSL connections to swagger JSON file")

	return cmd
}

func validateArgsSnippet(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return cmdutil.UsageError(cmd, "Unexpected args: %v", args)
	} else if len(args) < 1 {
		return cmdutil.UsageError(cmd, "You need specify resource name. e.g. v1.Pod")
	}

	return nil
}

func refPart(jsondata map[string]interface{}) error {
	if err := resultOutput(jsondata); err != nil {
		return err
	}
	return nil
}

func RunSnippet(cmd *cobra.Command, searchKey string) error {
	filelocation := cmdutil.GetFlagString(cmd, "filename")
	insecure := cmdutil.GetFlagBool(cmd, "insecure")

	jsondataRaw, err := ReadConfigDataFromLocation(filelocation, insecure)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	jsondata := map[string]interface{}{}
	err = json.Unmarshal(jsondataRaw, &jsondata)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	descripitonresult := jvmap.JsonValueMap(jsondata, searchKey)

	var foundSnippetList []map[string]interface{}
	for k, _ := range descripitonresult {
		foundSnippetList = append(foundSnippetList, logic.JsonValueParentChain(jsondata, descripitonresult[k][0], searchKey))
	}

	if foundSnippetList == nil {
		return fmt.Errorf("Not match parameter %s in %s", searchKey, filelocation)
	} else if n := len(foundSnippetList); n > 1 {
		fmt.Printf("\"%s\" found at %d locations\n", searchKey, n)
		for k, _ := range descripitonresult {
			fmt.Printf("\n")
			if err = refPart(foundSnippetList[k]); err != nil {
				return fmt.Errorf("%s", err)
			}

		}
	} else {
		if err = refPart(descripitonresult[0][0]); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}
