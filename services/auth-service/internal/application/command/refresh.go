package command

type RefreshCommand struct {
	RefreshToken string
}

// RefreshResult same as LoginResult
type RefreshResult struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	TokenType    string `json:"tokenType"`
	Username     string `json:"username"`
}
