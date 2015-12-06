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
			cmdutil.CheckErr(ValidateArgs(cmd, args))
			// cmdutil.CheckErr(cmdutil.ValidateOutputArgs(cmd))
			cmdutil.CheckErr(RunItemList(cmd))
		},
	}

	cmd.MarkFlagRequired("filename")
	cmd.PersistentFlags().StringP("filename", "f", "https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/swagger-spec/v1.json", "Path to swagger API json")

	return cmd
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

	jsondataRaw, err := cmdutil.ReadConfigDataFromLocation(filelocation)
	if err != nil {
		return fmt.Errorf("Input file reading error")
	}

	jsondata := map[string]interface{}{}
	err = json.Unmarshal(jsondataRaw, &jsondata)
	if err != nil {
		return fmt.Errorf("Json unmarshal error. Probably json input was invalid.")
	}

	topList(jsondata)

	return nil
}
