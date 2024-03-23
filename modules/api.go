package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

const (
	firebaseEndpoint        = "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s"
	firebaseRefreshEndpoint = "https://securetoken.googleapis.com/v1/token?key=%s"
)

type (
	httpClient interface {
		post(path string, value io.Reader, header map[string]string) ([]byte, error)
	}

	ApiClient struct {
		apiKey       string
		refreshToken string
		httpClient
	}

	TokenInfo struct {
		Kind         string `json:"kind"`
		IdToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		LocalId      string `json:"localId"`
	}

	refreshAPIResponse struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    string `json:"expires_in"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		IdToken      string `json:"id_token"`
		UserId       string `json:"user_id"`
		ProjectId    string `json:"project_id"`
	}

	RefreshResponse struct {
		AccessToken  string `json:"accessToken"`
		ExpiresIn    int    `json:"expiresIn"`
		TokenType    string `json:"tokenType"`
		RefreshToken string `json:"refreshToken"`
		IdToken      string `json:"idToken"`
		UserId       string `json:"userId"`
		ProjectId    string `json:"projectId"`
	}
)

func NewApiClient(httpClient httpClient) *ApiClient {
	return &ApiClient{
		httpClient: httpClient,
	}
}

func (a *ApiClient) SetApiClient(apiKey, refreshToken string) {
	a.apiKey = apiKey
	a.refreshToken = refreshToken
}

func (a *ApiClient) Refresh() (RefreshResponse, error) {
	if a.apiKey == "" || a.refreshToken == "" {
		return RefreshResponse{}, fmt.Errorf("apiKey or refreshToken is not found")
	}
	body, err := a.httpClient.post(fmt.Sprintf(firebaseRefreshEndpoint, a.apiKey),
		bytes.NewBuffer([]byte(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", a.refreshToken))),
		map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return RefreshResponse{}, err
	}
	var resp refreshAPIResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return RefreshResponse{}, err
	}
	return resp.ToRefreshResponse(), nil
}

func (a *ApiClient) GetTokenInfo() (TokenInfo, error) {
	var info TokenInfo
	if a.apiKey == "" {
		return info, fmt.Errorf("apiKey is not found")
	}
	body, err := a.httpClient.post(fmt.Sprintf(firebaseEndpoint, a.apiKey), nil, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		return info, err
	}
	if err := json.Unmarshal(body, &info); err != nil {
		return info, err
	}
	a.refreshToken = info.RefreshToken
	return info, nil
}

func (rat refreshAPIResponse) ToRefreshResponse() RefreshResponse {
	expires, err := strconv.Atoi(rat.ExpiresIn)
	if err != nil {
		expires = 0
	}
	return RefreshResponse{
		AccessToken:  rat.AccessToken,
		ExpiresIn:    expires,
		TokenType:    rat.TokenType,
		RefreshToken: rat.RefreshToken,
		IdToken:      rat.IdToken,
		UserId:       rat.UserId,
		ProjectId:    rat.ProjectId,
	}
}
