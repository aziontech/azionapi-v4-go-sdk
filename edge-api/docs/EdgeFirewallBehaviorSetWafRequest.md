# EdgeFirewallBehaviorSetWafRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_waf&#x60; - set_waf | 
**Attributes** | [**EdgeFirewallBehaviorSetWafAttributesRequest**](EdgeFirewallBehaviorSetWafAttributesRequest.md) |  | 

## Methods

### NewEdgeFirewallBehaviorSetWafRequest

`func NewEdgeFirewallBehaviorSetWafRequest(type_ string, attributes EdgeFirewallBehaviorSetWafAttributesRequest, ) *EdgeFirewallBehaviorSetWafRequest`

NewEdgeFirewallBehaviorSetWafRequest instantiates a new EdgeFirewallBehaviorSetWafRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorSetWafRequestWithDefaults

`func NewEdgeFirewallBehaviorSetWafRequestWithDefaults() *EdgeFirewallBehaviorSetWafRequest`

NewEdgeFirewallBehaviorSetWafRequestWithDefaults instantiates a new EdgeFirewallBehaviorSetWafRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorSetWafRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorSetWafRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorSetWafRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeFirewallBehaviorSetWafRequest) GetAttributes() EdgeFirewallBehaviorSetWafAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeFirewallBehaviorSetWafRequest) GetAttributesOk() (*EdgeFirewallBehaviorSetWafAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeFirewallBehaviorSetWafRequest) SetAttributes(v EdgeFirewallBehaviorSetWafAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


