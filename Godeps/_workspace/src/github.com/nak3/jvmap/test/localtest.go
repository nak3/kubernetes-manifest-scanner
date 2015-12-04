package main

import (
	"../"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileReader(filename string) ([]byte, error) {

	inf, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	raw_data, err_ := ioutil.ReadAll(inf)
	if err_ != nil {
		return nil, err
	}
	return raw_data, nil
}

func main() {
	filename := "test1.json"
	json_rawdata, _ := fileReader(filename)

	jsondata := map[string]interface{}{}
	err := json.Unmarshal(json_rawdata, &jsondata)
	if err != nil {
		log.Fatal(err)
	}

	searchKey := "tokyo"
	fmt.Printf("\n 1. \n")
	fmt.Printf("%v\n", jvmap.JsonValueMap(jsondata, searchKey)) // [[map[tokyo:[123 456]] map[tokyo:[abc def]] map[tokyo:foo]]]

	rootKey := "T"
	searchKey = "tokyo"

	fmt.Printf("\n 2. \n")
	fmt.Printf("%v\n", jvmap.JsonValueMap(jsondata, searchKey, rootKey)) // [[map[tokyo:[abc def]]]]

	rootKey = "O"
	searchKey = "tokyo"

	fmt.Printf("\n 3. \n")
	fmt.Printf("%v\n", jvmap.JsonValueMap(jsondata, searchKey, rootKey)) // [[map[tokyo:[123 456]]]]

	rootKey = "models"
	searchKey = "tokyo"

	fmt.Printf("\n 4. \n")
	fmt.Printf("%v\n", jvmap.JsonValueMap(jsondata, searchKey, rootKey)) // [[map[tokyo:[123 456]] map[tokyo:[abc def]] map[tokyo:foo]]]

}
