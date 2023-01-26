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

// TokenRenewSelfRequest struct for TokenRenewSelfRequest
type TokenRenewSelfRequest struct {
	// The desired increment in seconds to the token expiration
	Increment int32 `json:"increment"`
	// Token to renew (unused, does not need to be set)
	Token string `json:"token"`
}

// NewTokenRenewSelfRequestWithDefaults instantiates a new TokenRenewSelfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRenewSelfRequestWithDefaults() *TokenRenewSelfRequest {
	var this TokenRenewSelfRequest

	this.Increment = 0

	return &this
}

func (o TokenRenewSelfRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["increment"] = o.Increment
	toSerialize["token"] = o.Token

	return json.Marshal(toSerialize)
}