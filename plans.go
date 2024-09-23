package chushi

import (
	"github.com/hashicorp/go-tfe"
)

type Plans struct {
	sdk *Sdk
}

type UpdatePlanRequest struct {
}

func (p *Plans) Update(planId string, req *UpdatePlanRequest) (*tfe.Plan, error) {
	return nil, nil
}
