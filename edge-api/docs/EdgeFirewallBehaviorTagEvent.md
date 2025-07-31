# EdgeFirewallBehaviorTagEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;tag_event&#x60; - tag_event | 
**Attributes** | [**EdgeFirewallBehaviorTagEventAttributes**](EdgeFirewallBehaviorTagEventAttributes.md) |  | 

## Methods

### NewEdgeFirewallBehaviorTagEvent

`func NewEdgeFirewallBehaviorTagEvent(type_ string, attributes EdgeFirewallBehaviorTagEventAttributes, ) *EdgeFirewallBehaviorTagEvent`

NewEdgeFirewallBehaviorTagEvent instantiates a new EdgeFirewallBehaviorTagEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorTagEventWithDefaults

`func NewEdgeFirewallBehaviorTagEventWithDefaults() *EdgeFirewallBehaviorTagEvent`

NewEdgeFirewallBehaviorTagEventWithDefaults instantiates a new EdgeFirewallBehaviorTagEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorTagEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorTagEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorTagEvent) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeFirewallBehaviorTagEvent) GetAttributes() EdgeFirewallBehaviorTagEventAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeFirewallBehaviorTagEvent) GetAttributesOk() (*EdgeFirewallBehaviorTagEventAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeFirewallBehaviorTagEvent) SetAttributes(v EdgeFirewallBehaviorTagEventAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


