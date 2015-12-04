package logic

import (
	"reflect"
)

func jsonValueParentChain(jsonDecodedData map[string]interface{}, searchKeyData map[string]interface{}, resultCollection []map[string]interface{}, initialFlag bool) (map[string]interface{}, []map[string]interface{}, bool, bool) {
	isFound := false
	for k, v := range jsonDecodedData {
		resultData := map[string]interface{}{}
		for _, vv := range searchKeyData {
			if reflect.DeepEqual(v, vv) {
				resultData = map[string]interface{}{k: vv.(map[string]interface{})}
				isFound = true
				// initialFlag = false
				return resultData, resultCollection, isFound, initialFlag
			} else {
				var returnData map[string]interface{}
				switch v.(type) {
				case map[string]interface{}:
					returnData, resultCollection, isFound, initialFlag = jsonValueParentChain(v.(map[string]interface{}), searchKeyData, resultCollection, initialFlag)

					if initialFlag == false && isFound == true {
						resultData = map[string]interface{}{k: returnData}
						resultCollection = append(resultCollection, resultData)
						initialFlag = true
						isFound = false
					} else if isFound == true {
						initialFlag = false
						resultData = map[string]interface{}{k: returnData}
						return resultData, resultCollection, isFound, initialFlag
					}
				default:
					continue
				}
			}
		}
	}
	return jsonDecodedData, resultCollection, isFound, initialFlag // not found
}

func JsonValueParentChain(jsonDecodedData map[string]interface{}, searchData map[string]interface{}, keys ...string) map[string]interface{} {
	var resultCollection []map[string]interface{}
	initialFlag := true

	resultData, resultCollection, _, _ := jsonValueParentChain(jsonDecodedData, searchData, resultCollection, initialFlag)

	// TODO
	if resultCollection == nil {
		resultCollection = append(resultCollection, resultData)
	}

	return resultCollection[0]
}
