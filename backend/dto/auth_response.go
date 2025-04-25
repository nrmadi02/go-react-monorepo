package dto

import "github.com/nrmadi02/go_react_monorepo/backend/models"

type AuthResponse struct {
	AccessToken string      `json:"access_token"`
	User        models.User `json:"user"`
}

type MeResponse struct {
	User models.User `json:"user"`
}

func NewAuthResponse(account *models.Account) AuthResponse {
	return AuthResponse{
		AccessToken: account.AccessToken,
		User:        account.User,
	}
}

func NewMeResponse(account *models.Account) MeResponse {
	return MeResponse{
		User: account.User,
	}
}
