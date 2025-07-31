# EdgeFirewallBehaviorSetWafAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WafId** | **int64** |  | 
**Mode** | **string** | * &#x60;logging&#x60; - logging * &#x60;blocking&#x60; - blocking | 

## Methods

### NewEdgeFirewallBehaviorSetWafAttributesRequest

`func NewEdgeFirewallBehaviorSetWafAttributesRequest(wafId int64, mode string, ) *EdgeFirewallBehaviorSetWafAttributesRequest`

NewEdgeFirewallBehaviorSetWafAttributesRequest instantiates a new EdgeFirewallBehaviorSetWafAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorSetWafAttributesRequestWithDefaults

`func NewEdgeFirewallBehaviorSetWafAttributesRequestWithDefaults() *EdgeFirewallBehaviorSetWafAttributesRequest`

NewEdgeFirewallBehaviorSetWafAttributesRequestWithDefaults instantiates a new EdgeFirewallBehaviorSetWafAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWafId

`func (o *EdgeFirewallBehaviorSetWafAttributesRequest) GetWafId() int64`

GetWafId returns the WafId field if non-nil, zero value otherwise.

### GetWafIdOk

`func (o *EdgeFirewallBehaviorSetWafAttributesRequest) GetWafIdOk() (*int64, bool)`

GetWafIdOk returns a tuple with the WafId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafId

`func (o *EdgeFirewallBehaviorSetWafAttributesRequest) SetWafId(v int64)`

SetWafId sets WafId field to given value.


### GetMode

`func (o *EdgeFirewallBehaviorSetWafAttributesRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EdgeFirewallBehaviorSetWafAttributesRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EdgeFirewallBehaviorSetWafAttributesRequest) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


