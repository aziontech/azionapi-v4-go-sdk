# PatchedCertificateRevocationListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** | Indicates if the certificate revocation list is active. This field cannot be set to false. | [optional] 
**Crl** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedCertificateRevocationListRequest

`func NewPatchedCertificateRevocationListRequest() *PatchedCertificateRevocationListRequest`

NewPatchedCertificateRevocationListRequest instantiates a new PatchedCertificateRevocationListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCertificateRevocationListRequestWithDefaults

`func NewPatchedCertificateRevocationListRequestWithDefaults() *PatchedCertificateRevocationListRequest`

NewPatchedCertificateRevocationListRequestWithDefaults instantiates a new PatchedCertificateRevocationListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCertificateRevocationListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCertificateRevocationListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCertificateRevocationListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCertificateRevocationListRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedCertificateRevocationListRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedCertificateRevocationListRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedCertificateRevocationListRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedCertificateRevocationListRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCrl

`func (o *PatchedCertificateRevocationListRequest) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *PatchedCertificateRevocationListRequest) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *PatchedCertificateRevocationListRequest) SetCrl(v string)`

SetCrl sets Crl field to given value.

### HasCrl

`func (o *PatchedCertificateRevocationListRequest) HasCrl() bool`

HasCrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


