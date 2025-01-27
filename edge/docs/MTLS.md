# MTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verification** | Pointer to [**VerificationEnum**](VerificationEnum.md) |  | [optional] [default to enforce]
**Certificate** | Pointer to **NullableInt32** |  | [optional] 
**Crl** | Pointer to **[]int32** |  | [optional] 

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

`func (o *MTLS) GetVerification() VerificationEnum`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *MTLS) GetVerificationOk() (*VerificationEnum, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *MTLS) SetVerification(v VerificationEnum)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *MTLS) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetCertificate

`func (o *MTLS) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *MTLS) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *MTLS) SetCertificate(v int32)`

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

`func (o *MTLS) GetCrl() []int32`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *MTLS) GetCrlOk() (*[]int32, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *MTLS) SetCrl(v []int32)`

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


