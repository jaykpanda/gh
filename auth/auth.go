package auth

import (
	"net/http"
)

type BasicAuthRequest struct {
	Username string
	Password string
}

type BearerAuthRequest struct {
	Token string
}

type OAuth2AuthorizationRequest struct {
	ResponseType string `json:"response_type"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	Scope        string `json:"scope,omitempty"`
	State        string `json:"state,omitempty"`
}

type OAuth2AuthorizationResponse struct {
	Code  string `json:"code"`
	State string `json:"state,omitempty"`
}

type OAuth2TokenRequest struct {
	GrantType         string `json:"grant_type"`
	Code              string `json:"code,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	ClientCredentials string `json:"client_credentials,omitempty"`
	RefreshToken      string `json:"refresh_token,omitempty"`
	RedirectURI       string `json:"redirect_uri,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
	Scope             string `json:"scope,omitempty"`
}

type OAuth2TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	State        string `json:"state,omitempty"`
}

func (bh BasicAuthRequest) SetAuthorizationHeader(req *http.Request) *http.Request {
	req.SetBasicAuth(bh.Username, bh.Password)
	return req
}

func (bh BearerAuthRequest) SetAuthorizationHeader(req *http.Request) *http.Request {
	bearerToken := "Bearer" + bh.Token
	req.Header.Add("Authorization", bearerToken)
	return req
}

func (oareq OAuth2AuthorizationRequest) DoAuthorizationGrantRequest(req *http.Request) *OAuth2AuthorizationResponse {
	return &OAuth2AuthorizationResponse{}
}

func (oareq OAuth2AuthorizationRequest) DoImplicitGrantRequest(req *http.Request) *OAuth2TokenResponse {
	return &OAuth2TokenResponse{}
}

func (otreq OAuth2TokenRequest) DoPasswordCredentialsGrantRequest(req *http.Request) *OAuth2TokenResponse {
	return &OAuth2TokenResponse{}
}

func (otreq OAuth2TokenRequest) DoClientCredentialsGrantRequest(req *http.Request) *OAuth2TokenResponse {
	return &OAuth2TokenResponse{}
}
