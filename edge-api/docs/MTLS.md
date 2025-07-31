# MTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verification** | Pointer to [**NullableMTLSVerification**](MTLSVerification.md) |  | [optional] 
**Certificate** | Pointer to **NullableInt64** |  | [optional] 
**Crl** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewMTLS

`func NewMTLS() *MTLS`

NewMTLS instantiates a new MTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTLSWithDefaults

`func NewMTLSWithDefaults() *MTLS`

NewMTLSWithDefaults instantiates a new MTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerification

`func (o *MTLS) GetVerification() MTLSVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *MTLS) GetVerificationOk() (*MTLSVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *MTLS) SetVerification(v MTLSVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *MTLS) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### SetVerificationNil

`func (o *MTLS) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *MTLS) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil
### GetCertificate

`func (o *MTLS) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *MTLS) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *MTLS) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *MTLS) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *MTLS) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *MTLS) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCrl

`func (o *MTLS) GetCrl() []int64`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *MTLS) GetCrlOk() (*[]int64, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *MTLS) SetCrl(v []int64)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *MTLS) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### SetCrlNil

`func (o *MTLS) SetCrlNil(b bool)`

 SetCrlNil sets the value for Crl to be an explicit nil

### UnsetCrl
`func (o *MTLS) UnsetCrl()`

UnsetCrl ensures that no value is present for Crl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


