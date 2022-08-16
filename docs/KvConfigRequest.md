# KvConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** | If true, the backend will require the cas parameter to be set for each write | [optional] 
**DeleteVersionAfter** | Pointer to **int32** | If set, the length of time before a version is deleted. A negative duration disables the use of delete_version_after on all keys. A zero duration clears the current setting. Accepts a Go duration format string. | [optional] 
**MaxVersions** | Pointer to **int32** | The number of versions to keep for each key. Defaults to 10 | [optional] 

## Methods

### NewKvConfigRequest

`func NewKvConfigRequest() *KvConfigRequest`

NewKvConfigRequest instantiates a new KvConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvConfigRequestWithDefaults

`func NewKvConfigRequestWithDefaults() *KvConfigRequest`

NewKvConfigRequestWithDefaults instantiates a new KvConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCasRequired

`func (o *KvConfigRequest) GetCasRequired() bool`

GetCasRequired returns the CasRequired field if non-nil, zero value otherwise.

### GetCasRequiredOk

`func (o *KvConfigRequest) GetCasRequiredOk() (*bool, bool)`

GetCasRequiredOk returns a tuple with the CasRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasRequired

`func (o *KvConfigRequest) SetCasRequired(v bool)`

SetCasRequired sets CasRequired field to given value.

### HasCasRequired

`func (o *KvConfigRequest) HasCasRequired() bool`

HasCasRequired returns a boolean if a field has been set.

### GetDeleteVersionAfter

`func (o *KvConfigRequest) GetDeleteVersionAfter() int32`

GetDeleteVersionAfter returns the DeleteVersionAfter field if non-nil, zero value otherwise.

### GetDeleteVersionAfterOk

`func (o *KvConfigRequest) GetDeleteVersionAfterOk() (*int32, bool)`

GetDeleteVersionAfterOk returns a tuple with the DeleteVersionAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVersionAfter

`func (o *KvConfigRequest) SetDeleteVersionAfter(v int32)`

SetDeleteVersionAfter sets DeleteVersionAfter field to given value.

### HasDeleteVersionAfter

`func (o *KvConfigRequest) HasDeleteVersionAfter() bool`

HasDeleteVersionAfter returns a boolean if a field has been set.

### GetMaxVersions

`func (o *KvConfigRequest) GetMaxVersions() int32`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *KvConfigRequest) GetMaxVersionsOk() (*int32, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *KvConfigRequest) SetMaxVersions(v int32)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *KvConfigRequest) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

