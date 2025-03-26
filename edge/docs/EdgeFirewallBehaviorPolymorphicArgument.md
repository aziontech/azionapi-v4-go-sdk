# EdgeFirewallBehaviorPolymorphicArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | * &#x60;second&#x60; - second * &#x60;minute&#x60; - minute | [optional] 
**LimitBy** | **string** | * &#x60;client_ip&#x60; - client_ip * &#x60;global&#x60; - global | 
**AverageRateLimit** | **int64** |  | 
**MaximumBurstSize** | Pointer to **int64** |  | [optional] 
**StatusCode** | **int64** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**ContentBody** | Pointer to **string** |  | [optional] 
**Id** | **int64** |  | 
**Mode** | **string** | * &#x60;learning&#x60; - learning * &#x60;blocking&#x60; - blocking | 

## Methods

### NewEdgeFirewallBehaviorPolymorphicArgument

`func NewEdgeFirewallBehaviorPolymorphicArgument(limitBy string, averageRateLimit int64, statusCode int64, id int64, mode string, ) *EdgeFirewallBehaviorPolymorphicArgument`

NewEdgeFirewallBehaviorPolymorphicArgument instantiates a new EdgeFirewallBehaviorPolymorphicArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorPolymorphicArgumentWithDefaults

`func NewEdgeFirewallBehaviorPolymorphicArgumentWithDefaults() *EdgeFirewallBehaviorPolymorphicArgument`

NewEdgeFirewallBehaviorPolymorphicArgumentWithDefaults instantiates a new EdgeFirewallBehaviorPolymorphicArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetLimitBy() string`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetLimitByOk() (*string, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetLimitBy(v string)`

SetLimitBy sets LimitBy field to given value.


### GetAverageRateLimit

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetAverageRateLimit() int64`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetAverageRateLimitOk() (*int64, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetAverageRateLimit(v int64)`

SetAverageRateLimit sets AverageRateLimit field to given value.


### GetMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetMaximumBurstSize() int64`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetMaximumBurstSizeOk() (*int64, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetMaximumBurstSize(v int64)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgument) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.

### GetStatusCode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.


### GetContentType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContentBody

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetContentBody() string`

GetContentBody returns the ContentBody field if non-nil, zero value otherwise.

### GetContentBodyOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetContentBodyOk() (*string, bool)`

GetContentBodyOk returns a tuple with the ContentBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBody

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetContentBody(v string)`

SetContentBody sets ContentBody field to given value.

### HasContentBody

`func (o *EdgeFirewallBehaviorPolymorphicArgument) HasContentBody() bool`

HasContentBody returns a boolean if a field has been set.

### GetId

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetId(v int64)`

SetId sets Id field to given value.


### GetMode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


