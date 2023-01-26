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

// RadiusLoginWithUsernameRequest struct for RadiusLoginWithUsernameRequest
type RadiusLoginWithUsernameRequest struct {
	// Password for this user.
	Password string `json:"password"`
	// Username to be used for login. (POST request body)
	Username string `json:"username"`
}

// NewRadiusLoginWithUsernameRequestWithDefaults instantiates a new RadiusLoginWithUsernameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusLoginWithUsernameRequestWithDefaults() *RadiusLoginWithUsernameRequest {
	var this RadiusLoginWithUsernameRequest

	return &this
}

func (o RadiusLoginWithUsernameRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username

	return json.Marshal(toSerialize)
}