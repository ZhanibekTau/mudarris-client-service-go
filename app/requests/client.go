package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
