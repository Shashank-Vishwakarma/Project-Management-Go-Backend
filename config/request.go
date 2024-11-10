package config

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}