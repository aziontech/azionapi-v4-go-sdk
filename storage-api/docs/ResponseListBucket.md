# ResponseListBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**EdgeAccess** | **string** | * &#x60;read_only&#x60; - read_only * &#x60;read_write&#x60; - read_write * &#x60;restricted&#x60; - restricted | 

## Methods

### NewResponseListBucket

`func NewResponseListBucket(name string, edgeAccess string, ) *ResponseListBucket`

NewResponseListBucket instantiates a new ResponseListBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListBucketWithDefaults

`func NewResponseListBucketWithDefaults() *ResponseListBucket`

NewResponseListBucketWithDefaults instantiates a new ResponseListBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseListBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListBucket) SetName(v string)`

SetName sets Name field to given value.


### GetEdgeAccess

`func (o *ResponseListBucket) GetEdgeAccess() string`

GetEdgeAccess returns the EdgeAccess field if non-nil, zero value otherwise.

### GetEdgeAccessOk

`func (o *ResponseListBucket) GetEdgeAccessOk() (*string, bool)`

GetEdgeAccessOk returns a tuple with the EdgeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccess

`func (o *ResponseListBucket) SetEdgeAccess(v string)`

SetEdgeAccess sets EdgeAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


