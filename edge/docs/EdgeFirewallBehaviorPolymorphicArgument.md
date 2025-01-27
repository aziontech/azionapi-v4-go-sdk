# EdgeFirewallBehaviorPolymorphicArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SetRateLimitArgumentTypeEnum**](SetRateLimitArgumentTypeEnum.md) |  | [optional] [default to second]
**LimitBy** | [**LimitByEnum**](LimitByEnum.md) |  | 
**AverageRateLimit** | **int32** |  | 
**MaximumBurstSize** | Pointer to **int32** |  | [optional] 
**StatusCode** | **int32** |  | 
**ContentType** | Pointer to **string** |  | [optional] [default to ""]
**ContentBody** | Pointer to **string** |  | [optional] [default to ""]
**Id** | **int64** |  | 
**Mode** | [**ModeEnum**](ModeEnum.md) |  | 

## Methods

### NewEdgeFirewallBehaviorPolymorphicArgument

`func NewEdgeFirewallBehaviorPolymorphicArgument(limitBy LimitByEnum, averageRateLimit int32, statusCode int32, id int64, mode ModeEnum, ) *EdgeFirewallBehaviorPolymorphicArgument`

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

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetType() SetRateLimitArgumentTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetTypeOk() (*SetRateLimitArgumentTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetType(v SetRateLimitArgumentTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *EdgeFirewallBehaviorPolymorphicArgument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetLimitBy() LimitByEnum`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetLimitByOk() (*LimitByEnum, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetLimitBy(v LimitByEnum)`

SetLimitBy sets LimitBy field to given value.


### GetAverageRateLimit

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetAverageRateLimit() int32`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetAverageRateLimitOk() (*int32, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetAverageRateLimit(v int32)`

SetAverageRateLimit sets AverageRateLimit field to given value.


### GetMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetMaximumBurstSize() int32`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetMaximumBurstSizeOk() (*int32, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetMaximumBurstSize(v int32)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgument) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.

### GetStatusCode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetStatusCode(v int32)`

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

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgument) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EdgeFirewallBehaviorPolymorphicArgument) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


