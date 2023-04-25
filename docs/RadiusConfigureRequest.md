# RadiusConfigureRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DialTimeout** | Pointer to **int32** | Number of seconds before connect times out (default: 10) | [optional] [default to 10]
**Host** | Pointer to **string** | RADIUS server host | [optional] 
**NasIdentifier** | Pointer to **string** | RADIUS NAS Identifier field (optional) | [optional] [default to ""]
**NasPort** | Pointer to **int32** | RADIUS NAS port field (default: 10) | [optional] [default to 10]
**Port** | Pointer to **int32** | RADIUS server port (default: 1812) | [optional] [default to 1812]
**ReadTimeout** | Pointer to **int32** | Number of seconds before response times out (default: 10) | [optional] [default to 10]
**Secret** | Pointer to **string** | Secret shared with the RADIUS server | [optional] 
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#x27;default&#x27; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies. This will apply to all tokens generated by this auth method, in addition to any configured for specific users. | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**UnregisteredUserPolicies** | Pointer to **string** | Comma-separated list of policies to grant upon successful RADIUS authentication of an unregistered user (default: empty) | [optional] [default to ""]



## Methods


### NewRadiusConfigureRequest

`func NewRadiusConfigureRequest() *RadiusConfigureRequest`

NewRadiusConfigureRequest instantiates a new RadiusConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusConfigureRequestWithDefaults

`func NewRadiusConfigureRequestWithDefaults() *RadiusConfigureRequest`

NewRadiusConfigureRequestWithDefaults instantiates a new RadiusConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetDialTimeout

`func (o *RadiusConfigureRequest) GetDialTimeout() int32`

GetDialTimeout returns the DialTimeout field if non-nil, zero value otherwise.

### GetDialTimeoutOk

`func (o *RadiusConfigureRequest) GetDialTimeoutOk() (*int32, bool)`

GetDialTimeoutOk returns a tuple with the DialTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialTimeout

`func (o *RadiusConfigureRequest) SetDialTimeout(v int32)`

SetDialTimeout sets DialTimeout field to given value.


### HasDialTimeout

`func (o *RadiusConfigureRequest) HasDialTimeout() bool`

HasDialTimeout returns a boolean if a field has been set.




### GetHost

`func (o *RadiusConfigureRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RadiusConfigureRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RadiusConfigureRequest) SetHost(v string)`

SetHost sets Host field to given value.


### HasHost

`func (o *RadiusConfigureRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.




### GetNasIdentifier

`func (o *RadiusConfigureRequest) GetNasIdentifier() string`

GetNasIdentifier returns the NasIdentifier field if non-nil, zero value otherwise.

### GetNasIdentifierOk

`func (o *RadiusConfigureRequest) GetNasIdentifierOk() (*string, bool)`

GetNasIdentifierOk returns a tuple with the NasIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasIdentifier

`func (o *RadiusConfigureRequest) SetNasIdentifier(v string)`

SetNasIdentifier sets NasIdentifier field to given value.


### HasNasIdentifier

`func (o *RadiusConfigureRequest) HasNasIdentifier() bool`

HasNasIdentifier returns a boolean if a field has been set.




### GetNasPort

`func (o *RadiusConfigureRequest) GetNasPort() int32`

GetNasPort returns the NasPort field if non-nil, zero value otherwise.

### GetNasPortOk

`func (o *RadiusConfigureRequest) GetNasPortOk() (*int32, bool)`

GetNasPortOk returns a tuple with the NasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasPort

`func (o *RadiusConfigureRequest) SetNasPort(v int32)`

SetNasPort sets NasPort field to given value.


### HasNasPort

`func (o *RadiusConfigureRequest) HasNasPort() bool`

HasNasPort returns a boolean if a field has been set.




### GetPort

`func (o *RadiusConfigureRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RadiusConfigureRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RadiusConfigureRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### HasPort

`func (o *RadiusConfigureRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.




### GetReadTimeout

`func (o *RadiusConfigureRequest) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *RadiusConfigureRequest) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *RadiusConfigureRequest) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.


### HasReadTimeout

`func (o *RadiusConfigureRequest) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.




### GetSecret

`func (o *RadiusConfigureRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RadiusConfigureRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RadiusConfigureRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### HasSecret

`func (o *RadiusConfigureRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.




### GetTokenBoundCidrs

`func (o *RadiusConfigureRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *RadiusConfigureRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *RadiusConfigureRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.


### HasTokenBoundCidrs

`func (o *RadiusConfigureRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.




### GetTokenExplicitMaxTtl

`func (o *RadiusConfigureRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *RadiusConfigureRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *RadiusConfigureRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.


### HasTokenExplicitMaxTtl

`func (o *RadiusConfigureRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.




### GetTokenMaxTtl

`func (o *RadiusConfigureRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *RadiusConfigureRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *RadiusConfigureRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.


### HasTokenMaxTtl

`func (o *RadiusConfigureRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.




### GetTokenNoDefaultPolicy

`func (o *RadiusConfigureRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *RadiusConfigureRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *RadiusConfigureRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.


### HasTokenNoDefaultPolicy

`func (o *RadiusConfigureRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.




### GetTokenNumUses

`func (o *RadiusConfigureRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *RadiusConfigureRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *RadiusConfigureRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.


### HasTokenNumUses

`func (o *RadiusConfigureRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.




### GetTokenPeriod

`func (o *RadiusConfigureRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *RadiusConfigureRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *RadiusConfigureRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.


### HasTokenPeriod

`func (o *RadiusConfigureRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.




### GetTokenPolicies

`func (o *RadiusConfigureRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *RadiusConfigureRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *RadiusConfigureRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.


### HasTokenPolicies

`func (o *RadiusConfigureRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.




### GetTokenTtl

`func (o *RadiusConfigureRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *RadiusConfigureRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *RadiusConfigureRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.


### HasTokenTtl

`func (o *RadiusConfigureRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.




### GetTokenType

`func (o *RadiusConfigureRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *RadiusConfigureRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *RadiusConfigureRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### HasTokenType

`func (o *RadiusConfigureRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.




### GetUnregisteredUserPolicies

`func (o *RadiusConfigureRequest) GetUnregisteredUserPolicies() string`

GetUnregisteredUserPolicies returns the UnregisteredUserPolicies field if non-nil, zero value otherwise.

### GetUnregisteredUserPoliciesOk

`func (o *RadiusConfigureRequest) GetUnregisteredUserPoliciesOk() (*string, bool)`

GetUnregisteredUserPoliciesOk returns a tuple with the UnregisteredUserPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnregisteredUserPolicies

`func (o *RadiusConfigureRequest) SetUnregisteredUserPolicies(v string)`

SetUnregisteredUserPolicies sets UnregisteredUserPolicies field to given value.


### HasUnregisteredUserPolicies

`func (o *RadiusConfigureRequest) HasUnregisteredUserPolicies() bool`

HasUnregisteredUserPolicies returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

