# FirewallBehaviorSetCustomResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_custom_response&#x60; - set_custom_response | 
**Attributes** | [**FirewallBehaviorSetCustomResponseAttributesRequest**](FirewallBehaviorSetCustomResponseAttributesRequest.md) |  | 

## Methods

### NewFirewallBehaviorSetCustomResponseRequest

`func NewFirewallBehaviorSetCustomResponseRequest(type_ string, attributes FirewallBehaviorSetCustomResponseAttributesRequest, ) *FirewallBehaviorSetCustomResponseRequest`

NewFirewallBehaviorSetCustomResponseRequest instantiates a new FirewallBehaviorSetCustomResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorSetCustomResponseRequestWithDefaults

`func NewFirewallBehaviorSetCustomResponseRequestWithDefaults() *FirewallBehaviorSetCustomResponseRequest`

NewFirewallBehaviorSetCustomResponseRequestWithDefaults instantiates a new FirewallBehaviorSetCustomResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorSetCustomResponseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorSetCustomResponseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorSetCustomResponseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorSetCustomResponseRequest) GetAttributes() FirewallBehaviorSetCustomResponseAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorSetCustomResponseRequest) GetAttributesOk() (*FirewallBehaviorSetCustomResponseAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorSetCustomResponseRequest) SetAttributes(v FirewallBehaviorSetCustomResponseAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


