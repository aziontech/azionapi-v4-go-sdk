# CredentialCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Capabilities** | **[]string** |  | 
**Bucket** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCredentialCreate

`func NewCredentialCreate(name string, capabilities []string, ) *CredentialCreate`

NewCredentialCreate instantiates a new CredentialCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialCreateWithDefaults

`func NewCredentialCreateWithDefaults() *CredentialCreate`

NewCredentialCreateWithDefaults instantiates a new CredentialCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CredentialCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialCreate) SetName(v string)`

SetName sets Name field to given value.


### GetCapabilities

`func (o *CredentialCreate) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CredentialCreate) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CredentialCreate) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetBucket

`func (o *CredentialCreate) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CredentialCreate) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CredentialCreate) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CredentialCreate) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CredentialCreate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CredentialCreate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CredentialCreate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CredentialCreate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


