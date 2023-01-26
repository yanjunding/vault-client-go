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

// CloudFoundryWriteRoleRequest struct for CloudFoundryWriteRoleRequest
type CloudFoundryWriteRoleRequest struct {
	// Require that the client certificate presented has at least one of these app IDs.
	BoundApplicationIds []string `json:"bound_application_ids"`
	// Use \"token_bound_cidrs\" instead. If this and \"token_bound_cidrs\" are both specified, only \"token_bound_cidrs\" will be used.
	// Deprecated
	BoundCidrs []string `json:"bound_cidrs"`
	// Require that the client certificate presented has at least one of these instance IDs.
	BoundInstanceIds []string `json:"bound_instance_ids"`
	// Require that the client certificate presented has at least one of these org IDs.
	BoundOrganizationIds []string `json:"bound_organization_ids"`
	// Require that the client certificate presented has at least one of these space IDs.
	BoundSpaceIds []string `json:"bound_space_ids"`
	// If set to true, disables the default behavior that logging in must be performed from an acceptable IP address described by the certificate presented.
	DisableIpMatching bool `json:"disable_ip_matching"`
	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl int32 `json:"max_ttl"`
	// Use \"token_period\" instead. If this and \"token_period\" are both specified, only \"token_period\" will be used.
	// Deprecated
	Period int32 `json:"period"`
	// Use \"token_policies\" instead. If this and \"token_policies\" are both specified, only \"token_policies\" will be used.
	// Deprecated
	Policies []string `json:"policies"`
	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`
	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl int32 `json:"token_explicit_max_ttl"`
	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`
	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`
	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`
	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`
	// Comma-separated list of policies
	TokenPolicies []string `json:"token_policies"`
	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`
	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`
	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl int32 `json:"ttl"`
}

// NewCloudFoundryWriteRoleRequestWithDefaults instantiates a new CloudFoundryWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFoundryWriteRoleRequestWithDefaults() *CloudFoundryWriteRoleRequest {
	var this CloudFoundryWriteRoleRequest

	this.DisableIpMatching = false
	this.TokenType = "default-service"

	return &this
}

func (o CloudFoundryWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["bound_application_ids"] = o.BoundApplicationIds
	toSerialize["bound_cidrs"] = o.BoundCidrs
	toSerialize["bound_instance_ids"] = o.BoundInstanceIds
	toSerialize["bound_organization_ids"] = o.BoundOrganizationIds
	toSerialize["bound_space_ids"] = o.BoundSpaceIds
	toSerialize["disable_ip_matching"] = o.DisableIpMatching
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["period"] = o.Period
	toSerialize["policies"] = o.Policies
	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	toSerialize["token_max_ttl"] = o.TokenMaxTtl
	toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	toSerialize["token_num_uses"] = o.TokenNumUses
	toSerialize["token_period"] = o.TokenPeriod
	toSerialize["token_policies"] = o.TokenPolicies
	toSerialize["token_ttl"] = o.TokenTtl
	toSerialize["token_type"] = o.TokenType
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}