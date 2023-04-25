// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AwsGenerateStsCredentials2Request struct for AwsGenerateStsCredentials2Request
type AwsGenerateStsCredentials2Request struct {
	// ARN of role to assume when credential_type is assumed_role
	RoleArn string `json:"role_arn"`

	// Session name to use when assuming role. Max chars: 64
	RoleSessionName string `json:"role_session_name"`

	// Lifetime of the returned credentials in seconds
	Ttl int32 `json:"ttl"`
}

// NewAwsGenerateStsCredentials2RequestWithDefaults instantiates a new AwsGenerateStsCredentials2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsGenerateStsCredentials2RequestWithDefaults() *AwsGenerateStsCredentials2Request {
	var this AwsGenerateStsCredentials2Request

	this.Ttl = 3600

	return &this
}

func (o AwsGenerateStsCredentials2Request) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["role_arn"] = o.RoleArn
	toSerialize["role_session_name"] = o.RoleSessionName
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}