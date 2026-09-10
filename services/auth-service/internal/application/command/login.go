package command

// LoginCommand request for login
type LoginCommand struct {
	Username string
	Password string
}

// LoginResult login result
type LoginResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
