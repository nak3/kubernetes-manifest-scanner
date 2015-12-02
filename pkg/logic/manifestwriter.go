package logic

import (
	jvmap "github.com/nak3/jvmap"
)

func ManifestScanner(jsondata map[string]interface{}, manifestedParameters map[string]interface{}, depth int) map[string]interface{} {
	for k, v := range manifestedParameters {
		switch v.(type) {
		case map[string]interface{}:
			for kk, vv := range v.(map[string]interface{}) {
				var searchKey2 string
				if vv == "string" {
					manifestedParameters[k] = "string"
				} else if vv == "integer" {
					manifestedParameters[k] = "integer"
				} else if vv == "boolean" {
					manifestedParameters[k] = "boolean"
				} else if vv == "array" { // e.g. args in kubernetes/test/images/network-tester/rc.json
					var arrayOutput interface{}
					arrayOutput, depth = arrayScanner(jsondata, manifestedParameters[k].(map[string]interface{}), depth)
					if arrayOutput != nil {
						manifestedParameters[k] = []interface{}{arrayOutput}
					} else {
						manifestedParameters[k] = []interface{}{"array string"}
					}
				} else if vv == "any" { // e.g. labels in kubernetes/test/images/network-tester/rc.json
					manifestedParameters[k] = map[string]interface{}{"any": "e.g. label"} // possible array
				} else if kk == "$ref" {
					if depth <= 0 {
						manifestedParameters[k] = "$ref"
					} else {
						searchKey2 = vv.(string)
						manifestedParameters[k] = getRef(jsondata, searchKey2, depth)
						depth = depth - 1
					}
				}
			}
		default:
			continue
		}

	}
	return manifestedParameters
}

func arrayScanner(jsondata map[string]interface{}, valuesInArray map[string]interface{}, depth int) (interface{}, int) {
	for _, v := range valuesInArray {
		switch v.(type) {
		case map[string]interface{}:
			for kk, vv := range v.(map[string]interface{}) {
				var searchKey string
				if kk == "string" {
					return nil, depth
				}
				if kk == "$ref" {
					searchKey = vv.(string)
					// TODO need accurate depth
					if depth <= 0 {
						v.(map[string]interface{})[kk] = "$ref"
					} else {
						valuesInArray = getRef(jsondata, searchKey, depth)
					}
					// TODO one more for $ref .. supplementalGroups and capabilities add/drop their properties => {}. I will handle them?
					return valuesInArray, depth
				}
			}
		}
	}
	return nil, depth
}

func getRef(jsondata map[string]interface{}, searchKey string, depth int) map[string]interface{} {
	const rootKey = "properties"

	outputresult := jvmap.JsonValueMap(jsondata, rootKey, searchKey)
	manifestedParameters := outputresult[0][0][rootKey].(map[string]interface{})
	delete(manifestedParameters, "status") //TODO: smart way?
	manifestedParameters = ManifestScanner(jsondata, manifestedParameters, depth)

	return manifestedParameters
}
