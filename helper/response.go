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
	User  any `json:"user"`
	Token any `json:"token"`
	Cart  any `json:"cart"`
}

func ResponseFormatLogin(user any, token any, cart any) LoginResponse {
	return LoginResponse{User: user, Token: token, Cart: cart}
}