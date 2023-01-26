/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AppRoleReadTokenTTLResponse struct for AppRoleReadTokenTTLResponse
type AppRoleReadTokenTTLResponse struct {
	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`
}

// NewAppRoleReadTokenTTLResponseWithDefaults instantiates a new AppRoleReadTokenTTLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleReadTokenTTLResponseWithDefaults() *AppRoleReadTokenTTLResponse {
	var this AppRoleReadTokenTTLResponse

	return &this
}

func (o AppRoleReadTokenTTLResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["token_ttl"] = o.TokenTtl

	return json.Marshal(toSerialize)
}