package helper

func ResponseFormat(code int, message any, data any) map[string]any {
	var result = map[string]any{}
	result["code"] = code
	result["message"] = message
	if data != nil {
		result["data"] = data
	}
	return result
}