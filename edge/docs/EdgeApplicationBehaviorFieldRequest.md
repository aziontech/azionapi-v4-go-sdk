# EdgeApplicationBehaviorFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**EdgeApplicationBehaviorNameField**](EdgeApplicationBehaviorNameField.md) |  | 
**Argument** | Pointer to [**NullableEdgeApplicationBehaviorPolymorphicArgumentRequest**](EdgeApplicationBehaviorPolymorphicArgumentRequest.md) |  | [optional] 

## Methods

### NewEdgeApplicationBehaviorFieldRequest

`func NewEdgeApplicationBehaviorFieldRequest(name EdgeApplicationBehaviorNameField, ) *EdgeApplicationBehaviorFieldRequest`

NewEdgeApplicationBehaviorFieldRequest instantiates a new EdgeApplicationBehaviorFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationBehaviorFieldRequestWithDefaults

`func NewEdgeApplicationBehaviorFieldRequestWithDefaults() *EdgeApplicationBehaviorFieldRequest`

NewEdgeApplicationBehaviorFieldRequestWithDefaults instantiates a new EdgeApplicationBehaviorFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeApplicationBehaviorFieldRequest) GetName() EdgeApplicationBehaviorNameField`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationBehaviorFieldRequest) GetNameOk() (*EdgeApplicationBehaviorNameField, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationBehaviorFieldRequest) SetName(v EdgeApplicationBehaviorNameField)`

SetName sets Name field to given value.


### GetArgument

`func (o *EdgeApplicationBehaviorFieldRequest) GetArgument() EdgeApplicationBehaviorPolymorphicArgumentRequest`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeApplicationBehaviorFieldRequest) GetArgumentOk() (*EdgeApplicationBehaviorPolymorphicArgumentRequest, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeApplicationBehaviorFieldRequest) SetArgument(v EdgeApplicationBehaviorPolymorphicArgumentRequest)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeApplicationBehaviorFieldRequest) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeApplicationBehaviorFieldRequest) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeApplicationBehaviorFieldRequest) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


