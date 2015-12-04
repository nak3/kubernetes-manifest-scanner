package jvmap

// jsonValueSearch returns valueList found under searchRoot
func jsonValueSearch(decoded map[string]interface{}, searchRoot string, valueList []map[string]interface{}) []map[string]interface{} {
	for k, v := range decoded {
		if k != searchRoot {
			switch v.(type) {
			case map[string]interface{}:
				valueList = jsonValueSearch(v.(map[string]interface{}), searchRoot, valueList)
			default:
				continue
			}
		} else {
			valueMap := make(map[string]interface{})
			valueMap[k] = v
			valueList = append(valueList, valueMap)
			switch v.(type) {
			case map[string]interface{}:
				valueList = jsonValueSearch(v.(map[string]interface{}), searchRoot, valueList)
			default:
				continue
			}
		}
	}
	return valueList
}

// JsonValueMap makes map(s) list from jsondata by searchKey under the values of rootKey
func JsonValueMap(jsonData map[string]interface{}, keys ...string) [][]map[string]interface{} {
	searchKey := keys[0]
	rootKey := searchKey
	if len(keys) >= 2 {
		rootKey = keys[1]
	}

	var scopedData []map[string]interface{}

	// NOTE: Should not skip rootKey == searchKey, since rootKey is not only one in JSON data.
	scopedData = jsonValueSearch(jsonData, rootKey, scopedData)

	var jsonValueMapList [][]map[string]interface{}
	for _, val := range scopedData {
		var foundValues []map[string]interface{}
		foundValues = jsonValueSearch(val, searchKey, foundValues)
		if foundValues != nil {
			jsonValueMapList = append(jsonValueMapList, foundValues)
		}
	}
	return jsonValueMapList
}
