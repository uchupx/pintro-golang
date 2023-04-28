package payload

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenReponse struct {
	Token string `json:"token"`
}

type UserPostRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
