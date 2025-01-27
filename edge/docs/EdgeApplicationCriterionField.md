# EdgeApplicationCriterionField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argument** | Pointer to [**NullableEdgeApplicationCriterionPolymorphicArgument**](EdgeApplicationCriterionPolymorphicArgument.md) |  | [optional] [default to ]
**Variable** | [**EdgeApplicationCriterionFieldVariableEnum**](EdgeApplicationCriterionFieldVariableEnum.md) |  | 
**Conditional** | [**ConditionalEnum**](ConditionalEnum.md) |  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 

## Methods

### NewEdgeApplicationCriterionField

`func NewEdgeApplicationCriterionField(variable EdgeApplicationCriterionFieldVariableEnum, conditional ConditionalEnum, operator OperatorEnum, ) *EdgeApplicationCriterionField`

NewEdgeApplicationCriterionField instantiates a new EdgeApplicationCriterionField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationCriterionFieldWithDefaults

`func NewEdgeApplicationCriterionFieldWithDefaults() *EdgeApplicationCriterionField`

NewEdgeApplicationCriterionFieldWithDefaults instantiates a new EdgeApplicationCriterionField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgument

`func (o *EdgeApplicationCriterionField) GetArgument() EdgeApplicationCriterionPolymorphicArgument`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeApplicationCriterionField) GetArgumentOk() (*EdgeApplicationCriterionPolymorphicArgument, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeApplicationCriterionField) SetArgument(v EdgeApplicationCriterionPolymorphicArgument)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeApplicationCriterionField) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeApplicationCriterionField) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeApplicationCriterionField) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil
### GetVariable

`func (o *EdgeApplicationCriterionField) GetVariable() EdgeApplicationCriterionFieldVariableEnum`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *EdgeApplicationCriterionField) GetVariableOk() (*EdgeApplicationCriterionFieldVariableEnum, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *EdgeApplicationCriterionField) SetVariable(v EdgeApplicationCriterionFieldVariableEnum)`

SetVariable sets Variable field to given value.


### GetConditional

`func (o *EdgeApplicationCriterionField) GetConditional() ConditionalEnum`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *EdgeApplicationCriterionField) GetConditionalOk() (*ConditionalEnum, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *EdgeApplicationCriterionField) SetConditional(v ConditionalEnum)`

SetConditional sets Conditional field to given value.


### GetOperator

`func (o *EdgeApplicationCriterionField) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EdgeApplicationCriterionField) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EdgeApplicationCriterionField) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


