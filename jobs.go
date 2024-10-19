package chushi

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-tfe"
	"github.com/hashicorp/jsonapi"
	"reflect"
)

type Job struct {
	ID        string  `jsonapi:"primary,jobs"`
	LockedBy  *string `jsonapi:"attr,locked-by"`
	Locked    bool    `jsonapi:"attr,locked"`
	Operation string  `jsonapi:"attr,operation"`
	CreatedAt string  `jsonapi:"attr,created-at"`
	UpdatedAt string  `jsonapi:"attr,updated-at"`
	Status    string  `jsonapi:"attr,status"`

	// Relations
	Workspace    *tfe.Workspace    `jsonapi:"relation,workspace"`
	Run          *tfe.Run          `jsonapi:"relation,run"`
	AgentPool    *tfe.AgentPool    `jsonapi:"relation,agent-pool"`
	Organization *tfe.Organization `jsonapi:"relation,organization"`

	Links map[string]interface{} `jsonapi:"links,omitempty"`
}

type JobList struct {
	Items []*Job
}

type Jobs struct {
	sdk *Sdk
}

func (j *Jobs) List(agentPoolId string) ([]*Job, error) {
	var jobs []*Job
	resp, err := j.sdk.Client.R().Get(fmt.Sprintf("/api/v2/agent-pools/%s/jobs", agentPoolId))
	if err != nil {
		return nil, err
	}
	items, err := jsonapi.UnmarshalManyPayload(bytes.NewReader(resp.Body()), reflect.TypeOf(new(Job)))
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		jobs = append(jobs, item.(*Job))
	}
	return jobs, nil
}

func (j *Jobs) Read(jobId string) (*JobList, error) {
	response := new(JobList)
	resp, err := j.sdk.Client.R().Get(fmt.Sprintf("/api/v2/jobs/%s", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}

func (j *Jobs) Lock(jobId string, lockId string) (*Job, error) {
	response := new(Job)
	resp, err := j.sdk.Client.R().
		SetBody(map[string]interface{}{
			"data": map[string]interface{}{
				"id":   jobId,
				"type": "jobs",
				"attributes": map[string]interface{}{
					"status": lockId,
				},
			},
		}).
		Post(fmt.Sprintf("/api/v2/jobs/%s/lock", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}

func (j *Jobs) Unlock(jobId string, lockId string) (*Job, error) {
	response := new(Job)
	resp, err := j.sdk.Client.R().
		SetBody(map[string]interface{}{
			"data": map[string]interface{}{
				"id":   jobId,
				"type": "jobs",
				"attributes": map[string]interface{}{
					"locked_by": lockId,
				},
			},
		}).
		Post(fmt.Sprintf("/api/v2/jobs/%s/unlock", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}

func (j *Jobs) Update(jobId string, status string) (*Job, error) {
	response := new(Job)
	resp, err := j.sdk.Client.
		R().
		SetBody(map[string]interface{}{
			"data": map[string]interface{}{
				"id":   jobId,
				"type": "jobs",
				"attributes": map[string]interface{}{
					"status": status,
				},
			},
		}).
		Patch(fmt.Sprintf("/api/v2/jobs/%s", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}
