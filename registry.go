package chushi

import (
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/jsonapi"
	"reflect"
)

type Registry struct {
	client *resty.Client
}

type Provider struct {
	ID        string `jsonapi:"primary,providers" json:"id"`
	Namespace string `jsonapi:"attr,namespace" json:"namespace"`
	Type      string `jsonapi:"attr,type" json:"type"`
}

type ProviderVersion struct {
	ID        string   `jsonapi:"primary,provider-versions" json:"id"`
	Version   string   `jsonapi:"attr,version" json:"version"`
	Protocols []string `jsonapi:"attr,protocols" json:"protocols"`
	KeyId     string   `jsonapi:"attr,key-id" json:"key_id"`
}

type ListProvidersInput struct {
	Namespace string
	Type      string
}

type ListProvidersOutput struct {
	Providers []*Provider `json:"providers"`
}

func (r *Registry) ListProviders(params *ListProvidersInput) (*ListProvidersOutput, error) {
	res, err := r.client.R().Get("registry/v1/providers")
	if err != nil {
		return nil, err
	}

	var providers []*Provider
	records, err := jsonapi.UnmarshalManyPayload(res.RawBody(), reflect.TypeOf(new(Provider)))
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		providers = append(providers, record.(*Provider))
	}

	return &ListProvidersOutput{Providers: providers}, nil
}

type GetProviderInput struct {
	Id string
}

type GetProviderOutput struct {
	Provider *Provider `json:"provider"`
}

func (r *Registry) GetProvider(params *GetProviderInput) (*GetProviderOutput, error) {
	res, err := r.client.R().SetPathParams(map[string]string{
		"providerId": params.Id,
	}).Get("registry/v1/providers/{providerId}")
	if err != nil {
		return nil, err
	}

	provider := new(Provider)
	if err = jsonapi.UnmarshalPayload(res.RawBody(), provider); err != nil {
		return nil, err
	}

	return &GetProviderOutput{Provider: provider}, nil
}

type ListProviderVersionsInput struct {
	Id string
}

type ListProviderVersionsOutput struct {
	ProviderVersions []*ProviderVersion `json:"versions"`
}

func (r *Registry) ListProviderVersions(params *ListProviderVersionsInput) (*ListProviderVersionsOutput, error) {
	return nil, nil
}

type CreateProviderVersionInput struct {
	ProviderId string
	Version    string
	Protocols  []string
	KeyId      string
}

type CreateProviderVersionOutput struct {
	ProviderVersion *ProviderVersion `json:"version"`
}

func (r *Registry) CreateProviderVersion(params *CreateProviderVersionInput) (*CreateProviderVersionOutput, error) {
	res, err := r.client.R().SetPathParams(map[string]string{
		"providerId": params.ProviderId,
	}).SetBody(&ProviderVersion{
		Version:   params.Version,
		Protocols: params.Protocols,
		KeyId:     params.KeyId,
	}).Post("registry/v1/providers/{providerId}/versions")
	if err != nil {
		return nil, err
	}

	version := new(ProviderVersion)
	if err = jsonapi.UnmarshalPayload(res.RawBody(), version); err != nil {
		return nil, err
	}

	return &CreateProviderVersionOutput{ProviderVersion: version}, nil
}
