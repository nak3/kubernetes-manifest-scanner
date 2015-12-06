package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"

	"github.com/nak3/jvmap"
	"github.com/spf13/cobra"
)

func NewCmdItemList(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "itemlist -f FILENAME",
		Short: "Output item list",
		Long:  "Output item list",
		//		Example: get_example,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(validateExtraArgs(cmd, args))
			// cmdutil.CheckErr(cmdutil.ValidateOutputArgs(cmd))
			cmdutil.CheckErr(RunItemList(cmd))
		},
	}

	cmd.MarkFlagRequired("filename")
	cmd.PersistentFlags().StringP("filename", "f", "https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/swagger-spec/v1.json", "Path to swagger API json")
	cmd.PersistentFlags().BoolP("insecure", "k", false, "Allow insecure SSL connections to swagger JSON file")

	return cmd
}

func validateExtraArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return cmdutil.UsageError(cmd, "Unexpected args: %v", args)
	}
	return nil
}

func topList(jsondata map[string]interface{}) {
	const searchKey = "models"

	//TODO
	outputresult := jvmap.JsonValueMap(jsondata, searchKey)
	for k, _ := range outputresult[0][0][searchKey].(map[string]interface{}) {
		if jvmap.JsonValueMap(jsondata, "kind", k) != nil {
			fmt.Printf("%v ", k)
		}
	}
}

func RunItemList(cmd *cobra.Command) error {
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

	topList(jsondata)

	return nil
}
