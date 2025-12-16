# InputPolymorphicInputDataSourceAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type identifier for this endpoint (raw_logs) | 
**Attributes** | [**InputDataSourceRequest**](InputDataSourceRequest.md) |  | 

## Methods

### NewInputPolymorphicInputDataSourceAttributesRequest

`func NewInputPolymorphicInputDataSourceAttributesRequest(type_ string, attributes InputDataSourceRequest, ) *InputPolymorphicInputDataSourceAttributesRequest`

NewInputPolymorphicInputDataSourceAttributesRequest instantiates a new InputPolymorphicInputDataSourceAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputPolymorphicInputDataSourceAttributesRequestWithDefaults

`func NewInputPolymorphicInputDataSourceAttributesRequestWithDefaults() *InputPolymorphicInputDataSourceAttributesRequest`

NewInputPolymorphicInputDataSourceAttributesRequestWithDefaults instantiates a new InputPolymorphicInputDataSourceAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputPolymorphicInputDataSourceAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputPolymorphicInputDataSourceAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputPolymorphicInputDataSourceAttributesRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InputPolymorphicInputDataSourceAttributesRequest) GetAttributes() InputDataSourceRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InputPolymorphicInputDataSourceAttributesRequest) GetAttributesOk() (*InputDataSourceRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InputPolymorphicInputDataSourceAttributesRequest) SetAttributes(v InputDataSourceRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


