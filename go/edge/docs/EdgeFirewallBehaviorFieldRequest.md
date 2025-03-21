# EdgeFirewallBehaviorFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | * &#x60;deny&#x60; - deny * &#x60;run_function&#x60; - run_function * &#x60;drop&#x60; - drop * &#x60;set_rate_limit&#x60; - set_rate_limit * &#x60;tag_event&#x60; - tag_event * &#x60;set_custom_response&#x60; - set_custom_response * &#x60;set_waf_ruleset&#x60; - set_waf_ruleset | 
**Argument** | Pointer to [**NullableEdgeFirewallBehaviorPolymorphicArgumentRequest**](EdgeFirewallBehaviorPolymorphicArgumentRequest.md) |  | [optional] 

## Methods

### NewEdgeFirewallBehaviorFieldRequest

`func NewEdgeFirewallBehaviorFieldRequest(name string, ) *EdgeFirewallBehaviorFieldRequest`

NewEdgeFirewallBehaviorFieldRequest instantiates a new EdgeFirewallBehaviorFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorFieldRequestWithDefaults

`func NewEdgeFirewallBehaviorFieldRequestWithDefaults() *EdgeFirewallBehaviorFieldRequest`

NewEdgeFirewallBehaviorFieldRequestWithDefaults instantiates a new EdgeFirewallBehaviorFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFirewallBehaviorFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewallBehaviorFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewallBehaviorFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgument

`func (o *EdgeFirewallBehaviorFieldRequest) GetArgument() EdgeFirewallBehaviorPolymorphicArgumentRequest`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *EdgeFirewallBehaviorFieldRequest) GetArgumentOk() (*EdgeFirewallBehaviorPolymorphicArgumentRequest, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *EdgeFirewallBehaviorFieldRequest) SetArgument(v EdgeFirewallBehaviorPolymorphicArgumentRequest)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *EdgeFirewallBehaviorFieldRequest) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### SetArgumentNil

`func (o *EdgeFirewallBehaviorFieldRequest) SetArgumentNil(b bool)`

 SetArgumentNil sets the value for Argument to be an explicit nil

### UnsetArgument
`func (o *EdgeFirewallBehaviorFieldRequest) UnsetArgument()`

UnsetArgument ensures that no value is present for Argument, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


