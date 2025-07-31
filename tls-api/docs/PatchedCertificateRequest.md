# PatchedCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** | The value can&#39;t be changed after the certificate creation.  * &#x60;edge_certificate&#x60; - Edge Certificate * &#x60;trusted_ca_certificate&#x60; - Trusted CA Certificate | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedCertificateRequest

`func NewPatchedCertificateRequest() *PatchedCertificateRequest`

NewPatchedCertificateRequest instantiates a new PatchedCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCertificateRequestWithDefaults

`func NewPatchedCertificateRequestWithDefaults() *PatchedCertificateRequest`

NewPatchedCertificateRequestWithDefaults instantiates a new PatchedCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCertificateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCertificateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCertificateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCertificateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificate

`func (o *PatchedCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PatchedCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PatchedCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PatchedCertificateRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *PatchedCertificateRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *PatchedCertificateRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetPrivateKey

`func (o *PatchedCertificateRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PatchedCertificateRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PatchedCertificateRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PatchedCertificateRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *PatchedCertificateRequest) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *PatchedCertificateRequest) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetType

`func (o *PatchedCertificateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedCertificateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedCertificateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedCertificateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActive

`func (o *PatchedCertificateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedCertificateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedCertificateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedCertificateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


