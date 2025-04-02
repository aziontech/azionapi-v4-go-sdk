# ResponseBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**Bucket**](Bucket.md) |  | 

## Methods

### NewResponseBucket

`func NewResponseBucket(state string, data Bucket, ) *ResponseBucket`

NewResponseBucket instantiates a new ResponseBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBucketWithDefaults

`func NewResponseBucketWithDefaults() *ResponseBucket`

NewResponseBucketWithDefaults instantiates a new ResponseBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseBucket) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseBucket) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseBucket) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseBucket) GetData() Bucket`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseBucket) GetDataOk() (*Bucket, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseBucket) SetData(v Bucket)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


