# ResponseListCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AccessKey** | **string** |  | [readonly] 
**SecretKey** | **string** |  | [readonly] 
**Capabilities** | **[]string** |  | 
**Bucket** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewResponseListCredential

`func NewResponseListCredential(name string, accessKey string, secretKey string, capabilities []string, createdAt time.Time, ) *ResponseListCredential`

NewResponseListCredential instantiates a new ResponseListCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListCredentialWithDefaults

`func NewResponseListCredentialWithDefaults() *ResponseListCredential`

NewResponseListCredentialWithDefaults instantiates a new ResponseListCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseListCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListCredential) SetName(v string)`

SetName sets Name field to given value.


### GetAccessKey

`func (o *ResponseListCredential) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ResponseListCredential) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ResponseListCredential) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *ResponseListCredential) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ResponseListCredential) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ResponseListCredential) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetCapabilities

`func (o *ResponseListCredential) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ResponseListCredential) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ResponseListCredential) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetBucket

`func (o *ResponseListCredential) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ResponseListCredential) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ResponseListCredential) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ResponseListCredential) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ResponseListCredential) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ResponseListCredential) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ResponseListCredential) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ResponseListCredential) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponseListCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponseListCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponseListCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


