package chushi

import (
	"bytes"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/hashicorp/go-tfe"
	"reflect"
	"time"
)

type Job struct {
	ID          string    `jsonapi:"primary,jobs"`
	CreatedAt   time.Time `jsonapi:"attr,created-at,iso8601"`
	Description string    `jsonapi:"attr,description"`
	LastUsedAt  time.Time `jsonapi:"attr,last-used-at,iso8601"`
	Token       string    `jsonapi:"attr,token"`
	ExpiredAt   time.Time `jsonapi:"attr,expired-at,iso8601"`

	Run       *tfe.Run       `jsonapi:"relation,run"`
	Workspace *tfe.Workspace `jsonapi:"relation,workspace"`
	AgentPool *tfe.AgentPool `jsonapi:"relation,agent-pool"`
}

type Jobs struct {
	sdk *Sdk
}

func (j *Jobs) List(agentPoolId string) ([]Job, error) {
	var jobs []Job
	resp, err := j.sdk.Client.R().Get(fmt.Sprintf("/api/v2/agent-pools/%s/jobs", agentPoolId))
	if err != nil {
		return nil, err
	}
	items, err := jsonapi.UnmarshalManyPayload(bytes.NewReader(resp.Body()), reflect.TypeOf(new(Job)))
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		jobs = append(jobs, item.(Job))
	}
	return jobs, nil
}

func (j *Jobs) Read(jobId string) (*Job, error) {
	response := new(Job)
	resp, err := j.sdk.Client.R().Get(fmt.Sprintf("/api/v2/jobs/%s", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}

func (j *Jobs) Lock(jobId string) (*Job, error) {
	response := new(Job)
	resp, err := j.sdk.Client.R().Post(fmt.Sprintf("/api/v2/jobs/%s/lock", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}

func (j *Jobs) Unlock(jobId string) (*Job, error) {
	response := new(Job)
	resp, err := j.sdk.Client.R().Post(fmt.Sprintf("/api/v2/jobs/%s/unlock", jobId))
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
		SetBody(map[string]interface{}{"status": status}).
		Patch(fmt.Sprintf("/api/v2/jobs/%s", jobId))
	if err != nil {
		return nil, err
	}

	err = jsonapi.UnmarshalPayload(bytes.NewReader(resp.Body()), response)
	return response, err
}
