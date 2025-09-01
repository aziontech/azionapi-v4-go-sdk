# FirewallBehaviorsFirewallBehaviorSetWafRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_waf&#x60; - set_waf | 
**Attributes** | [**FirewallBehaviorSetWafAttributesRequest**](FirewallBehaviorSetWafAttributesRequest.md) |  | 

## Methods

### NewFirewallBehaviorsFirewallBehaviorSetWafRequest

`func NewFirewallBehaviorsFirewallBehaviorSetWafRequest(type_ string, attributes FirewallBehaviorSetWafAttributesRequest, ) *FirewallBehaviorsFirewallBehaviorSetWafRequest`

NewFirewallBehaviorsFirewallBehaviorSetWafRequest instantiates a new FirewallBehaviorsFirewallBehaviorSetWafRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorsFirewallBehaviorSetWafRequestWithDefaults

`func NewFirewallBehaviorsFirewallBehaviorSetWafRequestWithDefaults() *FirewallBehaviorsFirewallBehaviorSetWafRequest`

NewFirewallBehaviorsFirewallBehaviorSetWafRequestWithDefaults instantiates a new FirewallBehaviorsFirewallBehaviorSetWafRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorsFirewallBehaviorSetWafRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorsFirewallBehaviorSetWafRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorsFirewallBehaviorSetWafRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorsFirewallBehaviorSetWafRequest) GetAttributes() FirewallBehaviorSetWafAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorsFirewallBehaviorSetWafRequest) GetAttributesOk() (*FirewallBehaviorSetWafAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorsFirewallBehaviorSetWafRequest) SetAttributes(v FirewallBehaviorSetWafAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


