# EdgeFirewallBehaviorPolymorphicArgumentRequest

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

### NewEdgeFirewallBehaviorPolymorphicArgumentRequest

`func NewEdgeFirewallBehaviorPolymorphicArgumentRequest(limitBy LimitByEnum, averageRateLimit int32, statusCode int32, id int64, mode ModeEnum, ) *EdgeFirewallBehaviorPolymorphicArgumentRequest`

NewEdgeFirewallBehaviorPolymorphicArgumentRequest instantiates a new EdgeFirewallBehaviorPolymorphicArgumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallBehaviorPolymorphicArgumentRequestWithDefaults

`func NewEdgeFirewallBehaviorPolymorphicArgumentRequestWithDefaults() *EdgeFirewallBehaviorPolymorphicArgumentRequest`

NewEdgeFirewallBehaviorPolymorphicArgumentRequestWithDefaults instantiates a new EdgeFirewallBehaviorPolymorphicArgumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetType() SetRateLimitArgumentTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetTypeOk() (*SetRateLimitArgumentTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetType(v SetRateLimitArgumentTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetLimitBy() LimitByEnum`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetLimitByOk() (*LimitByEnum, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetLimitBy(v LimitByEnum)`

SetLimitBy sets LimitBy field to given value.


### GetAverageRateLimit

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetAverageRateLimit() int32`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetAverageRateLimitOk() (*int32, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetAverageRateLimit(v int32)`

SetAverageRateLimit sets AverageRateLimit field to given value.


### GetMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetMaximumBurstSize() int32`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetMaximumBurstSizeOk() (*int32, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetMaximumBurstSize(v int32)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.

### GetStatusCode

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetContentType

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContentBody

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetContentBody() string`

GetContentBody returns the ContentBody field if non-nil, zero value otherwise.

### GetContentBodyOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetContentBodyOk() (*string, bool)`

GetContentBodyOk returns a tuple with the ContentBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBody

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetContentBody(v string)`

SetContentBody sets ContentBody field to given value.

### HasContentBody

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) HasContentBody() bool`

HasContentBody returns a boolean if a field has been set.

### GetId

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetMode

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetMode() ModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetModeOk() (*ModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EdgeFirewallBehaviorPolymorphicArgumentRequest) SetMode(v ModeEnum)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


