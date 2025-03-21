# EdgeFirewallBehaviorField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | * &#x60;deny&#x60; - deny * &#x60;run_function&#x60; - run_function * &#x60;drop&#x60; - drop * &#x60;set_rate_limit&#x60; - set_rate_limit * &#x60;tag_event&#x60; - tag_event * &#x60;set_custom_response&#x60; - set_custom_response * &#x60;set_waf_ruleset&#x60; - set_waf_ruleset | 
**Argument** | Pointer to [**NullableEdgeFirewallBehaviorPolymorphicArgument**](EdgeFirewallBehaviorPolymorphicArgument.md) |  | [optional] 

## Methods

### NewEdgeFirewallBehaviorField

`func NewEdgeFirewallBehaviorField(name string, ) *EdgeFirewallBehaviorField`

NewEdgeFirewallBehaviorField instantiates a new EdgeFirewallBehaviorField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorFieldWithDefaults

`func NewEdgeFirewallBehaviorFieldWithDefaults() *EdgeFirewallBehaviorField`

NewEdgeFirewallBehaviorFieldWithDefaults instantiates a new EdgeFirewallBehaviorField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFirewallBehaviorField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewallBehaviorField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewallBehaviorField) SetName(v string)`

SetName sets Name field to given value.


### GetArgument

`func (o *EdgeFirewallBehaviorField) GetArgument() EdgeFirewallBehaviorPolymorphicArgument`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeFirewallBehaviorField) GetArgumentOk() (*EdgeFirewallBehaviorPolymorphicArgument, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeFirewallBehaviorField) SetArgument(v EdgeFirewallBehaviorPolymorphicArgument)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeFirewallBehaviorField) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeFirewallBehaviorField) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeFirewallBehaviorField) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


