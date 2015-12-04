package cmd

import (
	"encoding/json"
	"fmt"
)

func resultOutput(podProperties map[string]interface{}) error {
	testOutput, err := json.MarshalIndent(podProperties, "", "\t")
	if err != nil {
		return fmt.Errorf("JSON output failed")
	}
	fmt.Printf("%s\n", testOutput)

	return nil
}
