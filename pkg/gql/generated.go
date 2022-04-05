// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// integrationsIntegrationsIntegration_v1 includes the requested fields of the GraphQL type Integration_v1.
type integrationsIntegrationsIntegration_v1 struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Schemas     []string `json:"schemas"`
}

// GetName returns integrationsIntegrationsIntegration_v1.Name, and is useful for accessing the field via an interface.
func (v *integrationsIntegrationsIntegration_v1) GetName() string { return v.Name }

// GetDescription returns integrationsIntegrationsIntegration_v1.Description, and is useful for accessing the field via an interface.
func (v *integrationsIntegrationsIntegration_v1) GetDescription() string { return v.Description }

// GetSchemas returns integrationsIntegrationsIntegration_v1.Schemas, and is useful for accessing the field via an interface.
func (v *integrationsIntegrationsIntegration_v1) GetSchemas() []string { return v.Schemas }

// integrationsResponse is returned by integrations on success.
type integrationsResponse struct {
	Integrations []integrationsIntegrationsIntegration_v1 `json:"integrations"`
}

// GetIntegrations returns integrationsResponse.Integrations, and is useful for accessing the field via an interface.
func (v *integrationsResponse) GetIntegrations() []integrationsIntegrationsIntegration_v1 {
	return v.Integrations
}

func integrations(
	ctx context.Context,
	client graphql.Client,
) (*integrationsResponse, error) {
	req := &graphql.Request{
		OpName: "integrations",
		Query: `
query integrations {
	integrations: integrations_v1 {
		name
		description
		schemas
	}
}
`,
	}
	var err error

	var data integrationsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}