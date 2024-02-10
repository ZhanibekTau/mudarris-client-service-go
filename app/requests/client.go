package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientAndUstazId struct {
	ClientId int `json:"client_id"`
	UstazId  int `json:"ustaz_id"`
}
