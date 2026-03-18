# BehaviorRunFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**BehaviorRunFunctionAttributesRequest**](BehaviorRunFunctionAttributesRequest.md) |  | 

## Methods

### NewBehaviorRunFunctionRequest

`func NewBehaviorRunFunctionRequest(type_ string, attributes BehaviorRunFunctionAttributesRequest, ) *BehaviorRunFunctionRequest`

NewBehaviorRunFunctionRequest instantiates a new BehaviorRunFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRunFunctionRequestWithDefaults

`func NewBehaviorRunFunctionRequestWithDefaults() *BehaviorRunFunctionRequest`

NewBehaviorRunFunctionRequestWithDefaults instantiates a new BehaviorRunFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorRunFunctionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorRunFunctionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorRunFunctionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorRunFunctionRequest) GetAttributes() BehaviorRunFunctionAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorRunFunctionRequest) GetAttributesOk() (*BehaviorRunFunctionAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorRunFunctionRequest) SetAttributes(v BehaviorRunFunctionAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


