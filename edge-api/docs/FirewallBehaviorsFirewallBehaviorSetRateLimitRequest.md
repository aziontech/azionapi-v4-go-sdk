# FirewallBehaviorsFirewallBehaviorSetRateLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_rate_limit&#x60; - set_rate_limit | 
**Attributes** | [**FirewallBehaviorSetRateLimitAttributesRequest**](FirewallBehaviorSetRateLimitAttributesRequest.md) |  | 

## Methods

### NewFirewallBehaviorsFirewallBehaviorSetRateLimitRequest

`func NewFirewallBehaviorsFirewallBehaviorSetRateLimitRequest(type_ string, attributes FirewallBehaviorSetRateLimitAttributesRequest, ) *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest`

NewFirewallBehaviorsFirewallBehaviorSetRateLimitRequest instantiates a new FirewallBehaviorsFirewallBehaviorSetRateLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorsFirewallBehaviorSetRateLimitRequestWithDefaults

`func NewFirewallBehaviorsFirewallBehaviorSetRateLimitRequestWithDefaults() *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest`

NewFirewallBehaviorsFirewallBehaviorSetRateLimitRequestWithDefaults instantiates a new FirewallBehaviorsFirewallBehaviorSetRateLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest) GetAttributes() FirewallBehaviorSetRateLimitAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest) GetAttributesOk() (*FirewallBehaviorSetRateLimitAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorsFirewallBehaviorSetRateLimitRequest) SetAttributes(v FirewallBehaviorSetRateLimitAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


