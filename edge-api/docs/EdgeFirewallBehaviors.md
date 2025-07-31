# EdgeFirewallBehaviors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**EdgeFirewallBehaviorRunFunctionAttributes**](EdgeFirewallBehaviorRunFunctionAttributes.md) |  | 

## Methods

### NewEdgeFirewallBehaviors

`func NewEdgeFirewallBehaviors(type_ string, attributes EdgeFirewallBehaviorRunFunctionAttributes, ) *EdgeFirewallBehaviors`

NewEdgeFirewallBehaviors instantiates a new EdgeFirewallBehaviors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorsWithDefaults

`func NewEdgeFirewallBehaviorsWithDefaults() *EdgeFirewallBehaviors`

NewEdgeFirewallBehaviorsWithDefaults instantiates a new EdgeFirewallBehaviors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviors) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeFirewallBehaviors) GetAttributes() EdgeFirewallBehaviorRunFunctionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeFirewallBehaviors) GetAttributesOk() (*EdgeFirewallBehaviorRunFunctionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeFirewallBehaviors) SetAttributes(v EdgeFirewallBehaviorRunFunctionAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


