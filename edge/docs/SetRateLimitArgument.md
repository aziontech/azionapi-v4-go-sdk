# SetRateLimitArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SetRateLimitArgumentTypeEnum**](SetRateLimitArgumentTypeEnum.md) |  | [optional] [default to second]
**LimitBy** | [**LimitByEnum**](LimitByEnum.md) |  | 
**AverageRateLimit** | **int32** |  | 
**MaximumBurstSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewSetRateLimitArgument

`func NewSetRateLimitArgument(limitBy LimitByEnum, averageRateLimit int32, ) *SetRateLimitArgument`

NewSetRateLimitArgument instantiates a new SetRateLimitArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRateLimitArgumentWithDefaults

`func NewSetRateLimitArgumentWithDefaults() *SetRateLimitArgument`

NewSetRateLimitArgumentWithDefaults instantiates a new SetRateLimitArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetRateLimitArgument) GetType() SetRateLimitArgumentTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetRateLimitArgument) GetTypeOk() (*SetRateLimitArgumentTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetRateLimitArgument) SetType(v SetRateLimitArgumentTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *SetRateLimitArgument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *SetRateLimitArgument) GetLimitBy() LimitByEnum`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *SetRateLimitArgument) GetLimitByOk() (*LimitByEnum, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *SetRateLimitArgument) SetLimitBy(v LimitByEnum)`

SetLimitBy sets LimitBy field to given value.


### GetAverageRateLimit

`func (o *SetRateLimitArgument) GetAverageRateLimit() int32`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *SetRateLimitArgument) GetAverageRateLimitOk() (*int32, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *SetRateLimitArgument) SetAverageRateLimit(v int32)`

SetAverageRateLimit sets AverageRateLimit field to given value.


### GetMaximumBurstSize

`func (o *SetRateLimitArgument) GetMaximumBurstSize() int32`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *SetRateLimitArgument) GetMaximumBurstSizeOk() (*int32, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *SetRateLimitArgument) SetMaximumBurstSize(v int32)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *SetRateLimitArgument) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


