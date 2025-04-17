# TLSWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **NullableInt64** |  | [optional] 
**Ciphers** | Pointer to [**NullableTLSWorkloadCiphers**](TLSWorkloadCiphers.md) |  | [optional] 
**MinimumVersion** | Pointer to [**NullableTLSWorkloadMinimumVersion**](TLSWorkloadMinimumVersion.md) |  | [optional] 

## Methods

### NewTLSWorkload

`func NewTLSWorkload() *TLSWorkload`

NewTLSWorkload instantiates a new TLSWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSWorkloadWithDefaults

`func NewTLSWorkloadWithDefaults() *TLSWorkload`

NewTLSWorkloadWithDefaults instantiates a new TLSWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TLSWorkload) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TLSWorkload) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TLSWorkload) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TLSWorkload) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *TLSWorkload) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *TLSWorkload) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCiphers

`func (o *TLSWorkload) GetCiphers() TLSWorkloadCiphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TLSWorkload) GetCiphersOk() (*TLSWorkloadCiphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TLSWorkload) SetCiphers(v TLSWorkloadCiphers)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TLSWorkload) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### SetCiphersNil

`func (o *TLSWorkload) SetCiphersNil(b bool)`

 SetCiphersNil sets the value for Ciphers to be an explicit nil

### UnsetCiphers
`func (o *TLSWorkload) UnsetCiphers()`

UnsetCiphers ensures that no value is present for Ciphers, not even an explicit nil
### GetMinimumVersion

`func (o *TLSWorkload) GetMinimumVersion() TLSWorkloadMinimumVersion`

GetMinimumVersion returns the MinimumVersion field if non-nil, zero value otherwise.

### GetMinimumVersionOk

`func (o *TLSWorkload) GetMinimumVersionOk() (*TLSWorkloadMinimumVersion, bool)`

GetMinimumVersionOk returns a tuple with the MinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersion

`func (o *TLSWorkload) SetMinimumVersion(v TLSWorkloadMinimumVersion)`

SetMinimumVersion sets MinimumVersion field to given value.

### HasMinimumVersion

`func (o *TLSWorkload) HasMinimumVersion() bool`

HasMinimumVersion returns a boolean if a field has been set.

### SetMinimumVersionNil

`func (o *TLSWorkload) SetMinimumVersionNil(b bool)`

 SetMinimumVersionNil sets the value for MinimumVersion to be an explicit nil

### UnsetMinimumVersion
`func (o *TLSWorkload) UnsetMinimumVersion()`

UnsetMinimumVersion ensures that no value is present for MinimumVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


