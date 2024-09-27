package chushi

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/jsonapi"
	"time"
)

type Runs struct {
	sdk *Sdk
}

type OidcTokenResponse struct {
	Token string `json:"token"`
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

type RunToken struct {
	ID          string    `jsonapi:"primary,authentication-tokens"`
	CreatedAt   time.Time `jsonapi:"attr,created-at,iso8601"`
	Description string    `jsonapi:"attr,description"`
	LastUsedAt  time.Time `jsonapi:"attr,last-used-at,iso8601"`
	Token       string    `jsonapi:"attr,token"`
	ExpiredAt   time.Time `jsonapi:"attr,expired-at,iso8601"`
}

func (p *Runs) Token(runId string) (string, error) {
	response := new(RunToken)
	resp, err := p.sdk.Client.
		R().
		//SetResult(response).
		Post(fmt.Sprintf("/api/v2/runs/%s/authentication-token", runId))
	if err != nil {
		return "", err
	}

	if err := jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response); err != nil {
		return "", err
	}
	return response.Token, nil
}
