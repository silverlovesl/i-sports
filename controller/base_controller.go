package controller

import "reflect"

// Element Array elements as object
type Element interface{}

// WrapArray Array elements to object
func WrapArray(arrayData interface{}) map[string]Element {
	jsonData := make(map[string]Element)
	dataSize := getArraySize(arrayData)
	if arrayData == nil || dataSize <= 0 {
		jsonData["data"] = make([]Element, 0)
	} else {
		jsonData["data"] = arrayData
	}

	return jsonData
}

func getArraySize(data interface{}) int {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		return reflect.ValueOf(data).Len()
	default:
		return -1
	}
}
