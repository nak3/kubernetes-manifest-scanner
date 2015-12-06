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

func NewCmdSample(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "sample -f FILENAME RESOURCE",
		Short: "Output manifest with all parameteres",
		Long:  "Output manifest with all parameteres",
		//		Example: get_example,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(validateArgs(cmd, args))
			// cmdutil.CheckErr(cmdutil.ValidateOutputArgs(cmd))
			cmdutil.CheckErr(RunSample(cmd, args[0]))
		},
	}
	cmd.MarkFlagRequired("filename")
	cmd.PersistentFlags().StringP("filename", "f", "https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/swagger-spec/v1.json", "Path to swagger API json")
	cmd.PersistentFlags().IntP("depth", "d", 5, "Depth to expand $ref")

	return cmd
}

func allwriter(jsondata map[string]interface{}, properties map[string]interface{}, depth int) error {

	properties = logic.ManifestScanner(jsondata, properties, depth)
	if err := resultOutput(properties); err != nil {
		return err
	}
	return nil
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return cmdutil.UsageError(cmd, "Unexpected args: %v", args)
	} else if len(args) < 1 {
		return cmdutil.UsageError(cmd, "You need specify resource name. You can get the list by itemlist subcommand or bash completion")
	}

	return nil
}

func RunSample(cmd *cobra.Command, rootKey string) error {
	const searchKey = "properties"

	filelocation := cmdutil.GetFlagString(cmd, "filename")

	jsondataRaw, err := cmdutil.ReadConfigDataFromLocation(filelocation)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	jsondata := map[string]interface{}{}
	err = json.Unmarshal(jsondataRaw, &jsondata)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	if rootKey == "" {
		return fmt.Errorf("Need to set RESOURCE NAME as argmument")
	}

	outputresult := jvmap.JsonValueMap(jsondata, searchKey, rootKey)
	if outputresult == nil {
		return fmt.Errorf("Not match your item %s in %s", rootKey, filelocation)
	}

	depth := cmdutil.GetFlagInt(cmd, "depth")

	podProperties := outputresult[0][0][searchKey].(map[string]interface{})
	delete(podProperties, "status") //TODO: smart way?

	err = allwriter(jsondata, podProperties, depth)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
