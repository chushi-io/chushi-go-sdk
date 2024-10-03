package chushi

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-tfe"
	"github.com/hashicorp/jsonapi"
)

type Plans struct {
	sdk *Sdk
}

type UpdatePlanRequest struct {
	ID     string         `jsonapi:"primary,plans"`
	Status tfe.PlanStatus `jsonapi:"attr,status"`
}

func (p *Plans) Update(planId string, req *UpdatePlanRequest) (*tfe.Plan, error) {
	response := new(tfe.Plan)
	resp, err := p.sdk.Client.
		R().
		SetBody(map[string]interface{}{
			"data": map[string]interface{}{
				"id":   planId,
				"type": "plans",
				"attributes": map[string]interface{}{
					"status": req.Status,
				},
			},
		}).
		Patch(fmt.Sprintf("/api/v2/plans/%s", planId))
	if err != nil {
		return nil, err
	}

	if err := jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response); err != nil {
		return nil, err
	}
	return response, nil
}
