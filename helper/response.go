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

type LoginResponse struct {
	Message any `json:"message"`
	User    any `json:"user"`
	Token   any `json:"token"`
	Cart    any `json:"cart"`
	Order   any `json:"order"`
}

func ResponseFormatLogin(message any, user any, token any, cart any, order any) LoginResponse {
	return LoginResponse{Message: message, User: user, Token: token, Cart: cart, Order: order}
}

func ResponseGetAllProducts(code int, message any, totalPages any, data any) map[string]any {
	var result = map[string]any{}
	result["code"] = code
	result["message"] = message
	result["total_pages"] = totalPages
	if data != nil {
		result["data"] = data
	}
	return result
}

func ResponseGetOrder(code int, message any, totalPages any, data any, items any) map[string]any {
	var result = map[string]any{}
	result["code"] = code
	result["message"] = message
	result["items"] = items
	if data != nil {
		result["data"] = data
	}
	return result
}