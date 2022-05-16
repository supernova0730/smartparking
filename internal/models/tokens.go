package models

import "smartparking/internal/views"

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func (t Tokens) ToView() views.TokensView {
	return views.TokensView{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
	}
}
