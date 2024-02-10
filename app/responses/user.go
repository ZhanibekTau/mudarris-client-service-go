package responses

type ClientAndUstaz struct {
	Ustaz  UstazStruct  `json:"ustaz"`
	Client ClientStruct `json:"client"`
}

type UstazStruct struct {
	Name           string  `json:"name"`
	Lastname       string  `json:"lastname"`
	Email          string  `json:"email"`
	Degree         string  `json:"degree"`
	University     string  `json:"university"`
	AdditionalInfo string  `json:"additional_info"`
	Experience     float32 `json:"experience"`
	Phone          string  `json:"phone"`
}

type ClientStruct struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
