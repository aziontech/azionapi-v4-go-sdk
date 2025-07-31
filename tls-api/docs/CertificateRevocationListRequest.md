# CertificateRevocationListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** | Indicates if the certificate revocation list is active. This field cannot be set to false. | [optional] 
**Crl** | **string** |  | 

## Methods

### NewCertificateRevocationListRequest

`func NewCertificateRevocationListRequest(name string, crl string, ) *CertificateRevocationListRequest`

NewCertificateRevocationListRequest instantiates a new CertificateRevocationListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRevocationListRequestWithDefaults

`func NewCertificateRevocationListRequestWithDefaults() *CertificateRevocationListRequest`

NewCertificateRevocationListRequestWithDefaults instantiates a new CertificateRevocationListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateRevocationListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateRevocationListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateRevocationListRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *CertificateRevocationListRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CertificateRevocationListRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CertificateRevocationListRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CertificateRevocationListRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCrl

`func (o *CertificateRevocationListRequest) GetCrl() string`

GetCrl returns the Crl field if non-nil, zero value otherwise.

### GetCrlOk

`func (o *CertificateRevocationListRequest) GetCrlOk() (*string, bool)`

GetCrlOk returns a tuple with the Crl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrl

`func (o *CertificateRevocationListRequest) SetCrl(v string)`

SetCrl sets Crl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


