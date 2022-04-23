package auth

import (
	"net/http"
)

type BasicAuthParameter struct {
	Username string
	Password string
}

type BearerAuthParameter struct {
	Token string
}

type OAuth2AuthorizationRequest struct {
	ResponseType string `json:"response_type"`
}

type OAuth2TokenRequestParameters struct {
	GrantType         string `json:"grant_type"`
	Code              string `json:"code,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	ClientCredentials string `json:"client_credentials,omitempty"`
	RedirectURI       string `json:"redirect_uri,omitempty"`
	ClientID          string `json:"client_id,omitempty"`
}

func (bh BasicAuthParameter) SetAuthorizationHeader(req *http.Request) *http.Request {
	req.SetBasicAuth(bh.Username, bh.Password)
	return req
}

func (bh BearerAuthParameter) SetAuthorizationHeader(req *http.Request) *http.Request {
	bearerToken := "Bearer" + bh.Token
	req.Header.Add("Authorization", bearerToken)
	return req
}
