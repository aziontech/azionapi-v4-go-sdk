# ResponseBadRequestCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **[]string** |  | [optional] 
**ExpirationDate** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Capabilities** | Pointer to [**ResponseBadRequestCredentialCapabilities**](ResponseBadRequestCredentialCapabilities.md) |  | [optional] 
**Bucket** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestCredential

`func NewResponseBadRequestCredential() *ResponseBadRequestCredential`

NewResponseBadRequestCredential instantiates a new ResponseBadRequestCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestCredentialWithDefaults

`func NewResponseBadRequestCredentialWithDefaults() *ResponseBadRequestCredential`

NewResponseBadRequestCredentialWithDefaults instantiates a new ResponseBadRequestCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseBadRequestCredential) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestCredential) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestCredential) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ResponseBadRequestCredential) GetExpirationDate() []string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResponseBadRequestCredential) GetExpirationDateOk() (*[]string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResponseBadRequestCredential) SetExpirationDate(v []string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ResponseBadRequestCredential) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestCredential) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestCredential) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestCredential) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestCredential) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetCapabilities

`func (o *ResponseBadRequestCredential) GetCapabilities() ResponseBadRequestCredentialCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ResponseBadRequestCredential) GetCapabilitiesOk() (*ResponseBadRequestCredentialCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ResponseBadRequestCredential) SetCapabilities(v ResponseBadRequestCredentialCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ResponseBadRequestCredential) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetBucket

`func (o *ResponseBadRequestCredential) GetBucket() []string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ResponseBadRequestCredential) GetBucketOk() (*[]string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ResponseBadRequestCredential) SetBucket(v []string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ResponseBadRequestCredential) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestCredential) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestCredential) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestCredential) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestCredential) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


