# MFAMethodWriteDuoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiHostname** | Pointer to **string** | API host name for Duo. | [optional] 
**IntegrationKey** | Pointer to **string** | Integration key for Duo. | [optional] 
**MethodId** | Pointer to **string** | The unique identifier for this MFA method. | [optional] 
**PushInfo** | Pointer to **string** | Push information for Duo. | [optional] 
**SecretKey** | Pointer to **string** | Secret key for Duo. | [optional] 
**UsePasscode** | Pointer to **bool** | If true, the user is reminded to use the passcode upon MFA validation. This option does not enforce using the passcode. Defaults to false. | [optional] 
**UsernameFormat** | Pointer to **string** | A template string for mapping Identity names to MFA method names. Values to subtitute should be placed in {{}}. For example, \&quot;{{alias.name}}@example.com\&quot;. Currently-supported mappings: alias.name: The name returned by the mount configured via the mount_accessor parameter If blank, the Alias&#39;s name field will be used as-is. | [optional] 

## Methods

### NewMFAMethodWriteDuoRequest

`func NewMFAMethodWriteDuoRequest() *MFAMethodWriteDuoRequest`

NewMFAMethodWriteDuoRequest instantiates a new MFAMethodWriteDuoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAMethodWriteDuoRequestWithDefaults

`func NewMFAMethodWriteDuoRequestWithDefaults() *MFAMethodWriteDuoRequest`

NewMFAMethodWriteDuoRequestWithDefaults instantiates a new MFAMethodWriteDuoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiHostname

`func (o *MFAMethodWriteDuoRequest) GetApiHostname() string`

GetApiHostname returns the ApiHostname field if non-nil, zero value otherwise.

### GetApiHostnameOk

`func (o *MFAMethodWriteDuoRequest) GetApiHostnameOk() (*string, bool)`

GetApiHostnameOk returns a tuple with the ApiHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHostname

`func (o *MFAMethodWriteDuoRequest) SetApiHostname(v string)`

SetApiHostname sets ApiHostname field to given value.

### HasApiHostname

`func (o *MFAMethodWriteDuoRequest) HasApiHostname() bool`

HasApiHostname returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *MFAMethodWriteDuoRequest) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *MFAMethodWriteDuoRequest) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *MFAMethodWriteDuoRequest) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *MFAMethodWriteDuoRequest) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetMethodId

`func (o *MFAMethodWriteDuoRequest) GetMethodId() string`

GetMethodId returns the MethodId field if non-nil, zero value otherwise.

### GetMethodIdOk

`func (o *MFAMethodWriteDuoRequest) GetMethodIdOk() (*string, bool)`

GetMethodIdOk returns a tuple with the MethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodId

`func (o *MFAMethodWriteDuoRequest) SetMethodId(v string)`

SetMethodId sets MethodId field to given value.

### HasMethodId

`func (o *MFAMethodWriteDuoRequest) HasMethodId() bool`

HasMethodId returns a boolean if a field has been set.

### GetPushInfo

`func (o *MFAMethodWriteDuoRequest) GetPushInfo() string`

GetPushInfo returns the PushInfo field if non-nil, zero value otherwise.

### GetPushInfoOk

`func (o *MFAMethodWriteDuoRequest) GetPushInfoOk() (*string, bool)`

GetPushInfoOk returns a tuple with the PushInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushInfo

`func (o *MFAMethodWriteDuoRequest) SetPushInfo(v string)`

SetPushInfo sets PushInfo field to given value.

### HasPushInfo

`func (o *MFAMethodWriteDuoRequest) HasPushInfo() bool`

HasPushInfo returns a boolean if a field has been set.

### GetSecretKey

`func (o *MFAMethodWriteDuoRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *MFAMethodWriteDuoRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *MFAMethodWriteDuoRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *MFAMethodWriteDuoRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUsePasscode

`func (o *MFAMethodWriteDuoRequest) GetUsePasscode() bool`

GetUsePasscode returns the UsePasscode field if non-nil, zero value otherwise.

### GetUsePasscodeOk

`func (o *MFAMethodWriteDuoRequest) GetUsePasscodeOk() (*bool, bool)`

GetUsePasscodeOk returns a tuple with the UsePasscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasscode

`func (o *MFAMethodWriteDuoRequest) SetUsePasscode(v bool)`

SetUsePasscode sets UsePasscode field to given value.

### HasUsePasscode

`func (o *MFAMethodWriteDuoRequest) HasUsePasscode() bool`

HasUsePasscode returns a boolean if a field has been set.

### GetUsernameFormat

`func (o *MFAMethodWriteDuoRequest) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *MFAMethodWriteDuoRequest) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *MFAMethodWriteDuoRequest) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.

### HasUsernameFormat

`func (o *MFAMethodWriteDuoRequest) HasUsernameFormat() bool`

HasUsernameFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

