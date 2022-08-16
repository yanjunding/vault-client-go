# SystemRekeyInitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to **bool** | Specifies if using PGP-encrypted keys, whether Vault should also store a plaintext backup of the PGP-encrypted keys. | [optional] 
**PgpKeys** | Pointer to **[]string** | Specifies an array of PGP public keys used to encrypt the output unseal keys. Ordering is preserved. The keys must be base64-encoded from their original binary representation. The size of this array must be the same as secret_shares. | [optional] 
**RequireVerification** | Pointer to **bool** | Turns on verification functionality | [optional] 
**SecretShares** | Pointer to **int32** | Specifies the number of shares to split the unseal key into. | [optional] 
**SecretThreshold** | Pointer to **int32** | Specifies the number of shares required to reconstruct the unseal key. This must be less than or equal secret_shares. If using Vault HSM with auto-unsealing, this value must be the same as secret_shares. | [optional] 

## Methods

### NewSystemRekeyInitRequest

`func NewSystemRekeyInitRequest() *SystemRekeyInitRequest`

NewSystemRekeyInitRequest instantiates a new SystemRekeyInitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRekeyInitRequestWithDefaults

`func NewSystemRekeyInitRequestWithDefaults() *SystemRekeyInitRequest`

NewSystemRekeyInitRequestWithDefaults instantiates a new SystemRekeyInitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *SystemRekeyInitRequest) GetBackup() bool`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *SystemRekeyInitRequest) GetBackupOk() (*bool, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *SystemRekeyInitRequest) SetBackup(v bool)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *SystemRekeyInitRequest) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetPgpKeys

`func (o *SystemRekeyInitRequest) GetPgpKeys() []string`

GetPgpKeys returns the PgpKeys field if non-nil, zero value otherwise.

### GetPgpKeysOk

`func (o *SystemRekeyInitRequest) GetPgpKeysOk() (*[]string, bool)`

GetPgpKeysOk returns a tuple with the PgpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgpKeys

`func (o *SystemRekeyInitRequest) SetPgpKeys(v []string)`

SetPgpKeys sets PgpKeys field to given value.

### HasPgpKeys

`func (o *SystemRekeyInitRequest) HasPgpKeys() bool`

HasPgpKeys returns a boolean if a field has been set.

### GetRequireVerification

`func (o *SystemRekeyInitRequest) GetRequireVerification() bool`

GetRequireVerification returns the RequireVerification field if non-nil, zero value otherwise.

### GetRequireVerificationOk

`func (o *SystemRekeyInitRequest) GetRequireVerificationOk() (*bool, bool)`

GetRequireVerificationOk returns a tuple with the RequireVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireVerification

`func (o *SystemRekeyInitRequest) SetRequireVerification(v bool)`

SetRequireVerification sets RequireVerification field to given value.

### HasRequireVerification

`func (o *SystemRekeyInitRequest) HasRequireVerification() bool`

HasRequireVerification returns a boolean if a field has been set.

### GetSecretShares

`func (o *SystemRekeyInitRequest) GetSecretShares() int32`

GetSecretShares returns the SecretShares field if non-nil, zero value otherwise.

### GetSecretSharesOk

`func (o *SystemRekeyInitRequest) GetSecretSharesOk() (*int32, bool)`

GetSecretSharesOk returns a tuple with the SecretShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretShares

`func (o *SystemRekeyInitRequest) SetSecretShares(v int32)`

SetSecretShares sets SecretShares field to given value.

### HasSecretShares

`func (o *SystemRekeyInitRequest) HasSecretShares() bool`

HasSecretShares returns a boolean if a field has been set.

### GetSecretThreshold

`func (o *SystemRekeyInitRequest) GetSecretThreshold() int32`

GetSecretThreshold returns the SecretThreshold field if non-nil, zero value otherwise.

### GetSecretThresholdOk

`func (o *SystemRekeyInitRequest) GetSecretThresholdOk() (*int32, bool)`

GetSecretThresholdOk returns a tuple with the SecretThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretThreshold

`func (o *SystemRekeyInitRequest) SetSecretThreshold(v int32)`

SetSecretThreshold sets SecretThreshold field to given value.

### HasSecretThreshold

`func (o *SystemRekeyInitRequest) HasSecretThreshold() bool`

HasSecretThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

