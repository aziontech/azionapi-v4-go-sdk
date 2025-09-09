# FirewallBehaviorTagEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;tag_event&#x60; - tag_event | 
**Attributes** | [**FirewallBehaviorTagEventAttributesRequest**](FirewallBehaviorTagEventAttributesRequest.md) |  | 

## Methods

### NewFirewallBehaviorTagEventRequest

`func NewFirewallBehaviorTagEventRequest(type_ string, attributes FirewallBehaviorTagEventAttributesRequest, ) *FirewallBehaviorTagEventRequest`

NewFirewallBehaviorTagEventRequest instantiates a new FirewallBehaviorTagEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallBehaviorTagEventRequestWithDefaults

`func NewFirewallBehaviorTagEventRequestWithDefaults() *FirewallBehaviorTagEventRequest`

NewFirewallBehaviorTagEventRequestWithDefaults instantiates a new FirewallBehaviorTagEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FirewallBehaviorTagEventRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FirewallBehaviorTagEventRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FirewallBehaviorTagEventRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FirewallBehaviorTagEventRequest) GetAttributes() FirewallBehaviorTagEventAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FirewallBehaviorTagEventRequest) GetAttributesOk() (*FirewallBehaviorTagEventAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FirewallBehaviorTagEventRequest) SetAttributes(v FirewallBehaviorTagEventAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


