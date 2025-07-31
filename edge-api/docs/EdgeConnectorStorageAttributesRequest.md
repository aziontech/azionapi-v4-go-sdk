# EdgeConnectorStorageAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewEdgeConnectorStorageAttributesRequest

`func NewEdgeConnectorStorageAttributesRequest(bucket string, ) *EdgeConnectorStorageAttributesRequest`

NewEdgeConnectorStorageAttributesRequest instantiates a new EdgeConnectorStorageAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorStorageAttributesRequestWithDefaults

`func NewEdgeConnectorStorageAttributesRequestWithDefaults() *EdgeConnectorStorageAttributesRequest`

NewEdgeConnectorStorageAttributesRequestWithDefaults instantiates a new EdgeConnectorStorageAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *EdgeConnectorStorageAttributesRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *EdgeConnectorStorageAttributesRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *EdgeConnectorStorageAttributesRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPrefix

`func (o *EdgeConnectorStorageAttributesRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EdgeConnectorStorageAttributesRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EdgeConnectorStorageAttributesRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *EdgeConnectorStorageAttributesRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


