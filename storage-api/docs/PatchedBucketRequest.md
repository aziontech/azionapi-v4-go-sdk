# PatchedBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeAccess** | Pointer to **string** | * &#x60;read_only&#x60; - read_only * &#x60;read_write&#x60; - read_write * &#x60;restricted&#x60; - restricted | [optional] 

## Methods

### NewPatchedBucketRequest

`func NewPatchedBucketRequest() *PatchedBucketRequest`

NewPatchedBucketRequest instantiates a new PatchedBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBucketRequestWithDefaults

`func NewPatchedBucketRequestWithDefaults() *PatchedBucketRequest`

NewPatchedBucketRequestWithDefaults instantiates a new PatchedBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeAccess

`func (o *PatchedBucketRequest) GetEdgeAccess() string`

GetEdgeAccess returns the EdgeAccess field if non-nil, zero value otherwise.

### GetEdgeAccessOk

`func (o *PatchedBucketRequest) GetEdgeAccessOk() (*string, bool)`

GetEdgeAccessOk returns a tuple with the EdgeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccess

`func (o *PatchedBucketRequest) SetEdgeAccess(v string)`

SetEdgeAccess sets EdgeAccess field to given value.

### HasEdgeAccess

`func (o *PatchedBucketRequest) HasEdgeAccess() bool`

HasEdgeAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


