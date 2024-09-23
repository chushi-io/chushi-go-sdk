package chushi

import (
	"fmt"
	"time"
)

type Runs struct {
	sdk *Sdk
}

func (p *Runs) OidcToken(runId string) (string, error) {
	var response OidcTokenResponse
	_, err := p.sdk.Client.
		R().
		SetResult(&response).
		Get(fmt.Sprintf("/api/v2/runs/%s/oidc-token", runId))
	if err != nil {
		return "", err
	}
	return response.Token, nil
}

type OidcTokenResponse struct {
	Token string `json:"token"`
}

type RunToken struct {
	ID          string    `jsonapi:"primary,authentication-tokens"`
	CreatedAt   time.Time `jsonapi:"attr,created-at,iso8601"`
	Description string    `jsonapi:"attr,description"`
	LastUsedAt  time.Time `jsonapi:"attr,last-used-at,iso8601"`
	Token       string    `jsonapi:"attr,token"`
	ExpiredAt   time.Time `jsonapi:"attr,expired-at,iso8601"`
}

func (p *Runs) Token(runId string) (string, error) {
	var response RunToken
	_, err := p.sdk.Client.
		R().
		SetResult(&response).
		Get(fmt.Sprintf("/api/v2/runs/%s/authentication-token", runId))
	if err != nil {
		return "", err
	}
	return response.Token, nil
}
