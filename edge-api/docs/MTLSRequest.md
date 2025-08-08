# MTLSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verification** | Pointer to [**NullableMTLSVerification**](MTLSVerification.md) |  | [optional] 
**Certificate** | Pointer to **NullableInt64** |  | [optional] 
**Crl** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewMTLSRequest

`func NewMTLSRequest() *MTLSRequest`

NewMTLSRequest instantiates a new MTLSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTLSRequestWithDefaults

`func NewMTLSRequestWithDefaults() *MTLSRequest`

NewMTLSRequestWithDefaults instantiates a new MTLSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerification

`func (o *MTLSRequest) GetVerification() MTLSVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *MTLSRequest) GetVerificationOk() (*MTLSVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *MTLSRequest) SetVerification(v MTLSVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *MTLSRequest) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### SetVerificationNil

`func (o *MTLSRequest) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *MTLSRequest) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil
### GetCertificate

`func (o *MTLSRequest) GetCertificate() int64`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *MTLSRequest) GetCertificateOk() (*int64, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *MTLSRequest) SetCertificate(v int64)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *MTLSRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *MTLSRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *MTLSRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCrl

`func (o *MTLSRequest) GetCrl() []int64`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *MTLSRequest) GetCrlOk() (*[]int64, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *MTLSRequest) SetCrl(v []int64)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *MTLSRequest) HasCrl() bool`

HasCrl returns a boolean if a field has been set.

### SetCrlNil

`func (o *MTLSRequest) SetCrlNil(b bool)`

 SetCrlNil sets the value for Crl to be an explicit nil

### UnsetCrl
`func (o *MTLSRequest) UnsetCrl()`

UnsetCrl ensures that no value is present for Crl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


