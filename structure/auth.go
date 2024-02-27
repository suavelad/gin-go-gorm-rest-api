package structure

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
	Id           int    `json:"id"`
}
