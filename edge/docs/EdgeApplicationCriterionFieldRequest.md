# EdgeApplicationCriterionFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argument** | Pointer to [**NullableEdgeApplicationCriterionPolymorphicArgumentRequest**](EdgeApplicationCriterionPolymorphicArgumentRequest.md) |  | [optional] [default to ]
**Variable** | [**EdgeApplicationCriterionFieldVariableEnum**](EdgeApplicationCriterionFieldVariableEnum.md) |  | 
**Conditional** | [**ConditionalEnum**](ConditionalEnum.md) |  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 

## Methods

### NewEdgeApplicationCriterionFieldRequest

`func NewEdgeApplicationCriterionFieldRequest(variable EdgeApplicationCriterionFieldVariableEnum, conditional ConditionalEnum, operator OperatorEnum, ) *EdgeApplicationCriterionFieldRequest`

NewEdgeApplicationCriterionFieldRequest instantiates a new EdgeApplicationCriterionFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationCriterionFieldRequestWithDefaults

`func NewEdgeApplicationCriterionFieldRequestWithDefaults() *EdgeApplicationCriterionFieldRequest`

NewEdgeApplicationCriterionFieldRequestWithDefaults instantiates a new EdgeApplicationCriterionFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgument

`func (o *EdgeApplicationCriterionFieldRequest) GetArgument() EdgeApplicationCriterionPolymorphicArgumentRequest`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeApplicationCriterionFieldRequest) GetArgumentOk() (*EdgeApplicationCriterionPolymorphicArgumentRequest, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeApplicationCriterionFieldRequest) SetArgument(v EdgeApplicationCriterionPolymorphicArgumentRequest)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeApplicationCriterionFieldRequest) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeApplicationCriterionFieldRequest) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeApplicationCriterionFieldRequest) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil
### GetVariable

`func (o *EdgeApplicationCriterionFieldRequest) GetVariable() EdgeApplicationCriterionFieldVariableEnum`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *EdgeApplicationCriterionFieldRequest) GetVariableOk() (*EdgeApplicationCriterionFieldVariableEnum, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *EdgeApplicationCriterionFieldRequest) SetVariable(v EdgeApplicationCriterionFieldVariableEnum)`

SetVariable sets Variable field to given value.


### GetConditional

`func (o *EdgeApplicationCriterionFieldRequest) GetConditional() ConditionalEnum`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *EdgeApplicationCriterionFieldRequest) GetConditionalOk() (*ConditionalEnum, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *EdgeApplicationCriterionFieldRequest) SetConditional(v ConditionalEnum)`

SetConditional sets Conditional field to given value.


### GetOperator

`func (o *EdgeApplicationCriterionFieldRequest) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EdgeApplicationCriterionFieldRequest) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EdgeApplicationCriterionFieldRequest) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


