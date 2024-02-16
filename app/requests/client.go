package requests

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientAndUstazId struct {
	ClientId string `json:"client_id"`
	UstazId  string `json:"ustaz_id"`
}
type ClientIds struct {
	Ids []int `json:"client_ids"`
}
