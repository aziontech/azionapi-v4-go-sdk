# TLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **NullableInt64** |  | [optional] 
**Ciphers** | Pointer to [**NullableTLSCiphers**](TLSCiphers.md) |  | [optional] 
**MinimumVersion** | Pointer to [**NullableTLSMinimumVersion**](TLSMinimumVersion.md) |  | [optional] [default to tls_1_2]

## Methods

### NewTLS

`func NewTLS() *TLS`

NewTLS instantiates a new TLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSWithDefaults

`func NewTLSWithDefaults() *TLS`

NewTLSWithDefaults instantiates a new TLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TLS) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TLS) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TLS) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TLS) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *TLS) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *TLS) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCiphers

`func (o *TLS) GetCiphers() TLSCiphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TLS) GetCiphersOk() (*TLSCiphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TLS) SetCiphers(v TLSCiphers)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TLS) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### SetCiphersNil

`func (o *TLS) SetCiphersNil(b bool)`

 SetCiphersNil sets the value for Ciphers to be an explicit nil

### UnsetCiphers
`func (o *TLS) UnsetCiphers()`

UnsetCiphers ensures that no value is present for Ciphers, not even an explicit nil
### GetMinimumVersion

`func (o *TLS) GetMinimumVersion() TLSMinimumVersion`

GetMinimumVersion returns the MinimumVersion field if non-nil, zero value otherwise.

### GetMinimumVersionOk

`func (o *TLS) GetMinimumVersionOk() (*TLSMinimumVersion, bool)`

GetMinimumVersionOk returns a tuple with the MinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersion

`func (o *TLS) SetMinimumVersion(v TLSMinimumVersion)`

SetMinimumVersion sets MinimumVersion field to given value.

### HasMinimumVersion

`func (o *TLS) HasMinimumVersion() bool`

HasMinimumVersion returns a boolean if a field has been set.

### SetMinimumVersionNil

`func (o *TLS) SetMinimumVersionNil(b bool)`

 SetMinimumVersionNil sets the value for MinimumVersion to be an explicit nil

### UnsetMinimumVersion
`func (o *TLS) UnsetMinimumVersion()`

UnsetMinimumVersion ensures that no value is present for MinimumVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


