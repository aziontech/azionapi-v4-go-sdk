# EdgeFirewallCriterionField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditional** | **string** | * &#x60;if&#x60; - if * &#x60;or&#x60; - or * &#x60;and&#x60; - and | 
**Variable** | **string** | * &#x60;${header_accept}&#x60; - ${header_accept} * &#x60;${header_accept_encoding}&#x60; - ${header_accept_encoding} * &#x60;${header_accept_language}&#x60; - ${header_accept_language} * &#x60;${header_cookie}&#x60; - ${header_cookie} * &#x60;${header_origin}&#x60; - ${header_origin} * &#x60;${header_referer}&#x60; - ${header_referer} * &#x60;${header_user_agent}&#x60; - ${header_user_agent} * &#x60;${host}&#x60; - ${host} * &#x60;${network}&#x60; - ${network} * &#x60;${request_args}&#x60; - ${request_args} * &#x60;${request_method}&#x60; - ${request_method} * &#x60;${request_uri}&#x60; - ${request_uri} * &#x60;${scheme}&#x60; - ${scheme} * &#x60;${ssl_verification_status}&#x60; - ${ssl_verification_status} * &#x60;${client_certificate_validation}&#x60; - ${client_certificate_validation} | 
**Operator** | **string** | * &#x60;does_not_exist&#x60; - does_not_exist * &#x60;does_not_match&#x60; - does_not_match * &#x60;does_not_start_with&#x60; - does_not_start_with * &#x60;exists&#x60; - exists * &#x60;is_equal&#x60; - is_equal * &#x60;is_in_list&#x60; - is_in_list * &#x60;is_not_equal&#x60; - is_not_equal * &#x60;is_not_in_list&#x60; - is_not_in_list * &#x60;matches&#x60; - matches * &#x60;starts_with&#x60; - starts_with | 
**Argument** | Pointer to [**NullableEdgeFirewallCriterionPolymorphicArgument**](EdgeFirewallCriterionPolymorphicArgument.md) |  | [optional] 

## Methods

### NewEdgeFirewallCriterionField

`func NewEdgeFirewallCriterionField(conditional string, variable string, operator string, ) *EdgeFirewallCriterionField`

NewEdgeFirewallCriterionField instantiates a new EdgeFirewallCriterionField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallCriterionFieldWithDefaults

`func NewEdgeFirewallCriterionFieldWithDefaults() *EdgeFirewallCriterionField`

NewEdgeFirewallCriterionFieldWithDefaults instantiates a new EdgeFirewallCriterionField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditional

`func (o *EdgeFirewallCriterionField) GetConditional() string`

GetConditional returns the Conditional field if non-nil, zero value otherwise.

### GetConditionalOk

`func (o *EdgeFirewallCriterionField) GetConditionalOk() (*string, bool)`

GetConditionalOk returns a tuple with the Conditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditional

`func (o *EdgeFirewallCriterionField) SetConditional(v string)`

SetConditional sets Conditional field to given value.


### GetVariable

`func (o *EdgeFirewallCriterionField) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *EdgeFirewallCriterionField) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *EdgeFirewallCriterionField) SetVariable(v string)`

SetVariable sets Variable field to given value.


### GetOperator

`func (o *EdgeFirewallCriterionField) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EdgeFirewallCriterionField) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EdgeFirewallCriterionField) SetOperator(v string)`

SetOperator sets Operator field to given value.


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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


