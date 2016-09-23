package annict

import (
	"net/http"
)

type AuthenticationService service

type TokenInfo struct {
	ResourceOwnerId  int64    `json:"resource_owner_id"`
	Scopes           []string `json:"scopes"`
	ExpiresInSeconds int64    `json:"expires_in_seconds"`
	Application      struct {
		Uid string `json:"uid"`
	} `json:"application"`
	CreatedAt int64 `json:"created_at"`
}

func (s *AuthenticationService) Info() (*TokenInfo, *http.Response, error) {
	u, err := addOptions("oauth/token/info", nil)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	info := new(*TokenInfo)
	resp, err := s.client.Do(req, info)
	if err != nil {
		return nil, resp, err
	}

	return *info, resp, err
}
