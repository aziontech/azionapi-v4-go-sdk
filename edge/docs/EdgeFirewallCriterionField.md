# EdgeFirewallCriterionField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argument** | Pointer to [**NullableEdgeFirewallCriterionPolymorphicArgument**](EdgeFirewallCriterionPolymorphicArgument.md) |  | [optional] [default to ]
**Variable** | [**EdgeFirewallCriterionFieldVariableEnum**](EdgeFirewallCriterionFieldVariableEnum.md) |  | 
**Conditional** | [**ConditionalEnum**](ConditionalEnum.md) |  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) |  | 

## Methods

### NewEdgeFirewallCriterionField

`func NewEdgeFirewallCriterionField(variable EdgeFirewallCriterionFieldVariableEnum, conditional ConditionalEnum, operator OperatorEnum, ) *EdgeFirewallCriterionField`

NewEdgeFirewallCriterionField instantiates a new EdgeFirewallCriterionField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallCriterionFieldWithDefaults

`func NewEdgeFirewallCriterionFieldWithDefaults() *EdgeFirewallCriterionField`

NewEdgeFirewallCriterionFieldWithDefaults instantiates a new EdgeFirewallCriterionField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgument

`func (o *EdgeFirewallCriterionField) GetArgument() EdgeFirewallCriterionPolymorphicArgument`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeFirewallCriterionField) GetArgumentOk() (*EdgeFirewallCriterionPolymorphicArgument, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeFirewallCriterionField) SetArgument(v EdgeFirewallCriterionPolymorphicArgument)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeFirewallCriterionField) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeFirewallCriterionField) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeFirewallCriterionField) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil
### GetVariable

`func (o *EdgeFirewallCriterionField) GetVariable() EdgeFirewallCriterionFieldVariableEnum`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *EdgeFirewallCriterionField) GetVariableOk() (*EdgeFirewallCriterionFieldVariableEnum, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *EdgeFirewallCriterionField) SetVariable(v EdgeFirewallCriterionFieldVariableEnum)`

SetVariable sets Variable field to given value.


### GetConditional

`func (o *EdgeFirewallCriterionField) GetConditional() ConditionalEnum`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *EdgeFirewallCriterionField) GetConditionalOk() (*ConditionalEnum, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *EdgeFirewallCriterionField) SetConditional(v ConditionalEnum)`

SetConditional sets Conditional field to given value.


### GetOperator

`func (o *EdgeFirewallCriterionField) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EdgeFirewallCriterionField) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EdgeFirewallCriterionField) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


