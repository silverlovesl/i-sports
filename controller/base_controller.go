package controller

type Element interface{}

func WrapArray(data interface{}) map[string]Element {
	jsonData := make(map[string]Element)
	jsonData["data"] = data
	return jsonData
}
