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

// PKIRevokeRequest struct for PKIRevokeRequest
type PKIRevokeRequest struct {
	// Certificate to revoke in PEM format; must be signed by an issuer in this mount.
	Certificate string `json:"certificate"`
	// Certificate serial number, in colon- or hyphen-separated octal
	SerialNumber string `json:"serial_number"`
}

// NewPKIRevokeRequestWithDefaults instantiates a new PKIRevokeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIRevokeRequestWithDefaults() *PKIRevokeRequest {
	var this PKIRevokeRequest

	return &this
}

func (o PKIRevokeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["certificate"] = o.Certificate
	toSerialize["serial_number"] = o.SerialNumber

	return json.Marshal(toSerialize)
}