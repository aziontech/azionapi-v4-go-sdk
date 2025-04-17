# TLSWorkloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **NullableInt64** |  | [optional] 
**Ciphers** | Pointer to [**NullableTLSWorkloadCiphers**](TLSWorkloadCiphers.md) |  | [optional] 
**MinimumVersion** | Pointer to [**NullableTLSWorkloadMinimumVersion**](TLSWorkloadMinimumVersion.md) |  | [optional] 

## Methods

### NewTLSWorkloadRequest

`func NewTLSWorkloadRequest() *TLSWorkloadRequest`

NewTLSWorkloadRequest instantiates a new TLSWorkloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSWorkloadRequestWithDefaults

`func NewTLSWorkloadRequestWithDefaults() *TLSWorkloadRequest`

NewTLSWorkloadRequestWithDefaults instantiates a new TLSWorkloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TLSWorkloadRequest) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TLSWorkloadRequest) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TLSWorkloadRequest) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TLSWorkloadRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *TLSWorkloadRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *TLSWorkloadRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCiphers

`func (o *TLSWorkloadRequest) GetCiphers() TLSWorkloadCiphers`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *TLSWorkloadRequest) GetCiphersOk() (*TLSWorkloadCiphers, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *TLSWorkloadRequest) SetCiphers(v TLSWorkloadCiphers)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *TLSWorkloadRequest) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### SetCiphersNil

`func (o *TLSWorkloadRequest) SetCiphersNil(b bool)`

 SetCiphersNil sets the value for Ciphers to be an explicit nil

### UnsetCiphers
`func (o *TLSWorkloadRequest) UnsetCiphers()`

UnsetCiphers ensures that no value is present for Ciphers, not even an explicit nil
### GetMinimumVersion

`func (o *TLSWorkloadRequest) GetMinimumVersion() TLSWorkloadMinimumVersion`

GetMinimumVersion returns the MinimumVersion field if non-nil, zero value otherwise.

### GetMinimumVersionOk

`func (o *TLSWorkloadRequest) GetMinimumVersionOk() (*TLSWorkloadMinimumVersion, bool)`

GetMinimumVersionOk returns a tuple with the MinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersion

`func (o *TLSWorkloadRequest) SetMinimumVersion(v TLSWorkloadMinimumVersion)`

SetMinimumVersion sets MinimumVersion field to given value.

### HasMinimumVersion

`func (o *TLSWorkloadRequest) HasMinimumVersion() bool`

HasMinimumVersion returns a boolean if a field has been set.

### SetMinimumVersionNil

`func (o *TLSWorkloadRequest) SetMinimumVersionNil(b bool)`

 SetMinimumVersionNil sets the value for MinimumVersion to be an explicit nil

### UnsetMinimumVersion
`func (o *TLSWorkloadRequest) UnsetMinimumVersion()`

UnsetMinimumVersion ensures that no value is present for MinimumVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


