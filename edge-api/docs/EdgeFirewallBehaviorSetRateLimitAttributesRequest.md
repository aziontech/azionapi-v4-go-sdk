# EdgeFirewallBehaviorSetRateLimitAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | * &#x60;second&#x60; - second * &#x60;minute&#x60; - minute | [optional] 
**LimitBy** | **string** | * &#x60;client_ip&#x60; - client_ip * &#x60;global&#x60; - global | 
**AverageRateLimit** | **int64** |  | 
**MaximumBurstSize** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewEdgeFirewallBehaviorSetRateLimitAttributesRequest

`func NewEdgeFirewallBehaviorSetRateLimitAttributesRequest(limitBy string, averageRateLimit int64, ) *EdgeFirewallBehaviorSetRateLimitAttributesRequest`

NewEdgeFirewallBehaviorSetRateLimitAttributesRequest instantiates a new EdgeFirewallBehaviorSetRateLimitAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorSetRateLimitAttributesRequestWithDefaults

`func NewEdgeFirewallBehaviorSetRateLimitAttributesRequestWithDefaults() *EdgeFirewallBehaviorSetRateLimitAttributesRequest`

NewEdgeFirewallBehaviorSetRateLimitAttributesRequestWithDefaults instantiates a new EdgeFirewallBehaviorSetRateLimitAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetLimitBy() string`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetLimitByOk() (*string, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) SetLimitBy(v string)`

SetLimitBy sets LimitBy field to given value.


### GetAverageRateLimit

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetAverageRateLimit() int64`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetAverageRateLimitOk() (*int64, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) SetAverageRateLimit(v int64)`

SetAverageRateLimit sets AverageRateLimit field to given value.


### GetMaximumBurstSize

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetMaximumBurstSize() int64`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) GetMaximumBurstSizeOk() (*int64, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) SetMaximumBurstSize(v int64)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.

### SetMaximumBurstSizeNil

`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) SetMaximumBurstSizeNil(b bool)`

 SetMaximumBurstSizeNil sets the value for MaximumBurstSize to be an explicit nil

### UnsetMaximumBurstSize
`func (o *EdgeFirewallBehaviorSetRateLimitAttributesRequest) UnsetMaximumBurstSize()`

UnsetMaximumBurstSize ensures that no value is present for MaximumBurstSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


