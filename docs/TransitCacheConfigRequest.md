# TransitCacheConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | Size of cache, use 0 for an unlimited cache size, defaults to 0 | [optional] [default to 0]

## Methods

### NewTransitCacheConfigRequest

`func NewTransitCacheConfigRequest() *TransitCacheConfigRequest`

NewTransitCacheConfigRequest instantiates a new TransitCacheConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitCacheConfigRequestWithDefaults

`func NewTransitCacheConfigRequestWithDefaults() *TransitCacheConfigRequest`

NewTransitCacheConfigRequestWithDefaults instantiates a new TransitCacheConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *TransitCacheConfigRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TransitCacheConfigRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TransitCacheConfigRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TransitCacheConfigRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

