package views

type TokensView struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type LoginView struct {
	Client ClientView `json:"client,omitempty"`
	Tokens TokensView `json:"tokens,omitempty"`
}
