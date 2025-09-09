# FirewallBehaviorTagEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;tag_event&#x60; - tag_event | 
**Attributes** | [**FirewallBehaviorTagEventAttributes**](FirewallBehaviorTagEventAttributes.md) |  | 

## Methods

### NewFirewallBehaviorTagEvent

`func NewFirewallBehaviorTagEvent(type_ string, attributes FirewallBehaviorTagEventAttributes, ) *FirewallBehaviorTagEvent`

NewFirewallBehaviorTagEvent instantiates a new FirewallBehaviorTagEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorTagEventWithDefaults

`func NewFirewallBehaviorTagEventWithDefaults() *FirewallBehaviorTagEvent`

NewFirewallBehaviorTagEventWithDefaults instantiates a new FirewallBehaviorTagEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorTagEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorTagEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorTagEvent) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorTagEvent) GetAttributes() FirewallBehaviorTagEventAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorTagEvent) GetAttributesOk() (*FirewallBehaviorTagEventAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorTagEvent) SetAttributes(v FirewallBehaviorTagEventAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


