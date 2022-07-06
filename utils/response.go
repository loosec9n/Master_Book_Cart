package utils

func PrepareResponse(sucess bool, message string, result interface{}) map[string]interface{} {
	output := map[string]interface{}{
		"1 : sucess":  sucess,
		"2 : message": message,
		"3 : result":  result,
	}
	return output
}
