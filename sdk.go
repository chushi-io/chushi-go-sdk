package chushi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/go-tfe"
)

type Sdk struct {
	Address string
	Token   string
	Client  *resty.Client
	Runs    *Runs
	Plans   *Plans
	Applies *Applies
}

func New(config *tfe.Config) (*Sdk, error) {
	sdk := &Sdk{
		Address: config.Address,
		Token:   config.Token,
	}

	client := resty.New()
	client.SetBaseURL(config.Address)
	client.SetHeader("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	sdk.Client = client

	sdk.Plans = &Plans{sdk}
	sdk.Applies = &Applies{sdk}
	sdk.Runs = &Runs{sdk}
	return sdk, nil
}
