# PkiIssuersGenerateIntermediateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddBasicConstraints** | Pointer to **bool** | Whether to add a Basic Constraints extension with CA: true. Only needed as a workaround in some compatibility scenarios with Active Directory Certificate Services. | [optional] 
**AltNames** | Pointer to **string** | The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses. | [optional] 
**CommonName** | Pointer to **string** | The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans. | [optional] 
**Country** | Pointer to **[]string** | If set, Country will be set to this value. | [optional] 
**ExcludeCnFromSans** | Pointer to **bool** | If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included). | [optional] [default to false]
**Format** | Pointer to **string** | Format for returned data. Can be \&quot;pem\&quot;, \&quot;der\&quot;, or \&quot;pem_bundle\&quot;. If \&quot;pem_bundle\&quot;, any private key and issuing cert will be appended to the certificate pem. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**IpSans** | Pointer to **[]string** | The requested IP SANs, if any, in a comma-delimited list | [optional] 
**KeyBits** | Pointer to **int32** | The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519. | [optional] [default to 0]
**KeyName** | Pointer to **string** | Provide a name to the generated or existing key, the name must be unique across all keys and not be the reserved value &#39;default&#39; | [optional] 
**KeyRef** | Pointer to **string** | Reference to a existing key; either \&quot;default\&quot; for the configured default key, an identifier or the name assigned to the key. | [optional] [default to "default"]
**KeyType** | Pointer to **string** | The type of key to use; defaults to RSA. \&quot;rsa\&quot; \&quot;ec\&quot; and \&quot;ed25519\&quot; are the only valid values. | [optional] [default to "rsa"]
**Locality** | Pointer to **[]string** | If set, Locality will be set to this value. | [optional] 
**ManagedKeyId** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types. | [optional] 
**ManagedKeyName** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types. | [optional] 
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ | [optional] 
**NotBeforeDuration** | Pointer to **int32** | The duration before now which the certificate needs to be backdated by. | [optional] [default to 30]
**Organization** | Pointer to **[]string** | If set, O (Organization) will be set to this value. | [optional] 
**OtherSans** | Pointer to **[]string** | Requested other SANs, in an array with the format &lt;oid&gt;;UTF8:&lt;utf8 string value&gt; for each entry. | [optional] 
**Ou** | Pointer to **[]string** | If set, OU (OrganizationalUnit) will be set to this value. | [optional] 
**PostalCode** | Pointer to **[]string** | If set, Postal Code will be set to this value. | [optional] 
**PrivateKeyFormat** | Pointer to **string** | Format for the returned private key. Generally the default will be controlled by the \&quot;format\&quot; parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \&quot;pkcs8\&quot; to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \&quot;der\&quot;. | [optional] [default to "der"]
**Province** | Pointer to **[]string** | If set, Province will be set to this value. | [optional] 
**SerialNumber** | Pointer to **string** | The Subject&#39;s requested serial number, if any. See RFC 4519 Section 2.31 &#39;serialNumber&#39; for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate&#39;s Serial Number field. | [optional] 
**StreetAddress** | Pointer to **[]string** | If set, Street Address will be set to this value. | [optional] 
**Ttl** | Pointer to **int32** | The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA. | [optional] 
**UriSans** | Pointer to **[]string** | The requested URI SANs, if any, in a comma-delimited list. | [optional] 

## Methods

### NewPkiIssuersGenerateIntermediateRequest

`func NewPkiIssuersGenerateIntermediateRequest() *PkiIssuersGenerateIntermediateRequest`

NewPkiIssuersGenerateIntermediateRequest instantiates a new PkiIssuersGenerateIntermediateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiIssuersGenerateIntermediateRequestWithDefaults

`func NewPkiIssuersGenerateIntermediateRequestWithDefaults() *PkiIssuersGenerateIntermediateRequest`

NewPkiIssuersGenerateIntermediateRequestWithDefaults instantiates a new PkiIssuersGenerateIntermediateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddBasicConstraints

`func (o *PkiIssuersGenerateIntermediateRequest) GetAddBasicConstraints() bool`

GetAddBasicConstraints returns the AddBasicConstraints field if non-nil, zero value otherwise.

### GetAddBasicConstraintsOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetAddBasicConstraintsOk() (*bool, bool)`

GetAddBasicConstraintsOk returns a tuple with the AddBasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddBasicConstraints

`func (o *PkiIssuersGenerateIntermediateRequest) SetAddBasicConstraints(v bool)`

SetAddBasicConstraints sets AddBasicConstraints field to given value.

### HasAddBasicConstraints

`func (o *PkiIssuersGenerateIntermediateRequest) HasAddBasicConstraints() bool`

HasAddBasicConstraints returns a boolean if a field has been set.

### GetAltNames

`func (o *PkiIssuersGenerateIntermediateRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *PkiIssuersGenerateIntermediateRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *PkiIssuersGenerateIntermediateRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetCommonName

`func (o *PkiIssuersGenerateIntermediateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkiIssuersGenerateIntermediateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *PkiIssuersGenerateIntermediateRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *PkiIssuersGenerateIntermediateRequest) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkiIssuersGenerateIntermediateRequest) SetCountry(v []string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PkiIssuersGenerateIntermediateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExcludeCnFromSans

`func (o *PkiIssuersGenerateIntermediateRequest) GetExcludeCnFromSans() bool`

GetExcludeCnFromSans returns the ExcludeCnFromSans field if non-nil, zero value otherwise.

### GetExcludeCnFromSansOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetExcludeCnFromSansOk() (*bool, bool)`

GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCnFromSans

`func (o *PkiIssuersGenerateIntermediateRequest) SetExcludeCnFromSans(v bool)`

SetExcludeCnFromSans sets ExcludeCnFromSans field to given value.

### HasExcludeCnFromSans

`func (o *PkiIssuersGenerateIntermediateRequest) HasExcludeCnFromSans() bool`

HasExcludeCnFromSans returns a boolean if a field has been set.

### GetFormat

`func (o *PkiIssuersGenerateIntermediateRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PkiIssuersGenerateIntermediateRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PkiIssuersGenerateIntermediateRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIpSans

`func (o *PkiIssuersGenerateIntermediateRequest) GetIpSans() []string`

GetIpSans returns the IpSans field if non-nil, zero value otherwise.

### GetIpSansOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetIpSansOk() (*[]string, bool)`

GetIpSansOk returns a tuple with the IpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSans

`func (o *PkiIssuersGenerateIntermediateRequest) SetIpSans(v []string)`

SetIpSans sets IpSans field to given value.

### HasIpSans

`func (o *PkiIssuersGenerateIntermediateRequest) HasIpSans() bool`

HasIpSans returns a boolean if a field has been set.

### GetKeyBits

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PkiIssuersGenerateIntermediateRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.

### HasKeyBits

`func (o *PkiIssuersGenerateIntermediateRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.

### GetKeyName

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PkiIssuersGenerateIntermediateRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *PkiIssuersGenerateIntermediateRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetKeyRef

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyRef() string`

GetKeyRef returns the KeyRef field if non-nil, zero value otherwise.

### GetKeyRefOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyRefOk() (*string, bool)`

GetKeyRefOk returns a tuple with the KeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRef

`func (o *PkiIssuersGenerateIntermediateRequest) SetKeyRef(v string)`

SetKeyRef sets KeyRef field to given value.

### HasKeyRef

`func (o *PkiIssuersGenerateIntermediateRequest) HasKeyRef() bool`

HasKeyRef returns a boolean if a field has been set.

### GetKeyType

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiIssuersGenerateIntermediateRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PkiIssuersGenerateIntermediateRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetLocality

`func (o *PkiIssuersGenerateIntermediateRequest) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkiIssuersGenerateIntermediateRequest) SetLocality(v []string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *PkiIssuersGenerateIntermediateRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetManagedKeyId

`func (o *PkiIssuersGenerateIntermediateRequest) GetManagedKeyId() string`

GetManagedKeyId returns the ManagedKeyId field if non-nil, zero value otherwise.

### GetManagedKeyIdOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetManagedKeyIdOk() (*string, bool)`

GetManagedKeyIdOk returns a tuple with the ManagedKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyId

`func (o *PkiIssuersGenerateIntermediateRequest) SetManagedKeyId(v string)`

SetManagedKeyId sets ManagedKeyId field to given value.

### HasManagedKeyId

`func (o *PkiIssuersGenerateIntermediateRequest) HasManagedKeyId() bool`

HasManagedKeyId returns a boolean if a field has been set.

### GetManagedKeyName

`func (o *PkiIssuersGenerateIntermediateRequest) GetManagedKeyName() string`

GetManagedKeyName returns the ManagedKeyName field if non-nil, zero value otherwise.

### GetManagedKeyNameOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetManagedKeyNameOk() (*string, bool)`

GetManagedKeyNameOk returns a tuple with the ManagedKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyName

`func (o *PkiIssuersGenerateIntermediateRequest) SetManagedKeyName(v string)`

SetManagedKeyName sets ManagedKeyName field to given value.

### HasManagedKeyName

`func (o *PkiIssuersGenerateIntermediateRequest) HasManagedKeyName() bool`

HasManagedKeyName returns a boolean if a field has been set.

### GetNotAfter

`func (o *PkiIssuersGenerateIntermediateRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PkiIssuersGenerateIntermediateRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *PkiIssuersGenerateIntermediateRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBeforeDuration

`func (o *PkiIssuersGenerateIntermediateRequest) GetNotBeforeDuration() int32`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetNotBeforeDurationOk() (*int32, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PkiIssuersGenerateIntermediateRequest) SetNotBeforeDuration(v int32)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.

### HasNotBeforeDuration

`func (o *PkiIssuersGenerateIntermediateRequest) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.

### GetOrganization

`func (o *PkiIssuersGenerateIntermediateRequest) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkiIssuersGenerateIntermediateRequest) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PkiIssuersGenerateIntermediateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOtherSans

`func (o *PkiIssuersGenerateIntermediateRequest) GetOtherSans() []string`

GetOtherSans returns the OtherSans field if non-nil, zero value otherwise.

### GetOtherSansOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetOtherSansOk() (*[]string, bool)`

GetOtherSansOk returns a tuple with the OtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSans

`func (o *PkiIssuersGenerateIntermediateRequest) SetOtherSans(v []string)`

SetOtherSans sets OtherSans field to given value.

### HasOtherSans

`func (o *PkiIssuersGenerateIntermediateRequest) HasOtherSans() bool`

HasOtherSans returns a boolean if a field has been set.

### GetOu

`func (o *PkiIssuersGenerateIntermediateRequest) GetOu() []string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetOuOk() (*[]string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *PkiIssuersGenerateIntermediateRequest) SetOu(v []string)`

SetOu sets Ou field to given value.

### HasOu

`func (o *PkiIssuersGenerateIntermediateRequest) HasOu() bool`

HasOu returns a boolean if a field has been set.

### GetPostalCode

`func (o *PkiIssuersGenerateIntermediateRequest) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PkiIssuersGenerateIntermediateRequest) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PkiIssuersGenerateIntermediateRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPrivateKeyFormat

`func (o *PkiIssuersGenerateIntermediateRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *PkiIssuersGenerateIntermediateRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.

### HasPrivateKeyFormat

`func (o *PkiIssuersGenerateIntermediateRequest) HasPrivateKeyFormat() bool`

HasPrivateKeyFormat returns a boolean if a field has been set.

### GetProvince

`func (o *PkiIssuersGenerateIntermediateRequest) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PkiIssuersGenerateIntermediateRequest) SetProvince(v []string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *PkiIssuersGenerateIntermediateRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PkiIssuersGenerateIntermediateRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiIssuersGenerateIntermediateRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PkiIssuersGenerateIntermediateRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *PkiIssuersGenerateIntermediateRequest) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PkiIssuersGenerateIntermediateRequest) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *PkiIssuersGenerateIntermediateRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetTtl

`func (o *PkiIssuersGenerateIntermediateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PkiIssuersGenerateIntermediateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PkiIssuersGenerateIntermediateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUriSans

`func (o *PkiIssuersGenerateIntermediateRequest) GetUriSans() []string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *PkiIssuersGenerateIntermediateRequest) GetUriSansOk() (*[]string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *PkiIssuersGenerateIntermediateRequest) SetUriSans(v []string)`

SetUriSans sets UriSans field to given value.

### HasUriSans

`func (o *PkiIssuersGenerateIntermediateRequest) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

