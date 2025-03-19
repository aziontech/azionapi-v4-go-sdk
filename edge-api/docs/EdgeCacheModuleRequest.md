# EdgeCacheModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | **string** | * &#x60;honor&#x60; - Honor Origin Cache Settings * &#x60;override&#x60; - Override Cache Settings | 
**MaxAge** | **int64** |  | 
**CachingForPostEnabled** | **bool** |  | 
**CachingForOptionsEnabled** | **bool** |  | 
**StaleCacheEnabled** | **bool** |  | 
**TieredCacheEnabled** | **bool** |  | 
**TieredCacheRegion** | Pointer to **string** |  | [optional] 

## Methods

### NewEdgeCacheModuleRequest

`func NewEdgeCacheModuleRequest(behavior string, maxAge int64, cachingForPostEnabled bool, cachingForOptionsEnabled bool, staleCacheEnabled bool, tieredCacheEnabled bool, ) *EdgeCacheModuleRequest`

NewEdgeCacheModuleRequest instantiates a new EdgeCacheModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeCacheModuleRequestWithDefaults

`func NewEdgeCacheModuleRequestWithDefaults() *EdgeCacheModuleRequest`

NewEdgeCacheModuleRequestWithDefaults instantiates a new EdgeCacheModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *EdgeCacheModuleRequest) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *EdgeCacheModuleRequest) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *EdgeCacheModuleRequest) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.


### GetMaxAge

`func (o *EdgeCacheModuleRequest) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *EdgeCacheModuleRequest) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *EdgeCacheModuleRequest) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.


### GetCachingForPostEnabled

`func (o *EdgeCacheModuleRequest) GetCachingForPostEnabled() bool`

GetCachingForPostEnabled returns the CachingForPostEnabled field if non-nil, zero value otherwise.

### GetCachingForPostEnabledOk

`func (o *EdgeCacheModuleRequest) GetCachingForPostEnabledOk() (*bool, bool)`

GetCachingForPostEnabledOk returns a tuple with the CachingForPostEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingForPostEnabled

`func (o *EdgeCacheModuleRequest) SetCachingForPostEnabled(v bool)`

SetCachingForPostEnabled sets CachingForPostEnabled field to given value.


### GetCachingForOptionsEnabled

`func (o *EdgeCacheModuleRequest) GetCachingForOptionsEnabled() bool`

GetCachingForOptionsEnabled returns the CachingForOptionsEnabled field if non-nil, zero value otherwise.

### GetCachingForOptionsEnabledOk

`func (o *EdgeCacheModuleRequest) GetCachingForOptionsEnabledOk() (*bool, bool)`

GetCachingForOptionsEnabledOk returns a tuple with the CachingForOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachingForOptionsEnabled

`func (o *EdgeCacheModuleRequest) SetCachingForOptionsEnabled(v bool)`

SetCachingForOptionsEnabled sets CachingForOptionsEnabled field to given value.


### GetStaleCacheEnabled

`func (o *EdgeCacheModuleRequest) GetStaleCacheEnabled() bool`

GetStaleCacheEnabled returns the StaleCacheEnabled field if non-nil, zero value otherwise.

### GetStaleCacheEnabledOk

`func (o *EdgeCacheModuleRequest) GetStaleCacheEnabledOk() (*bool, bool)`

GetStaleCacheEnabledOk returns a tuple with the StaleCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCacheEnabled

`func (o *EdgeCacheModuleRequest) SetStaleCacheEnabled(v bool)`

SetStaleCacheEnabled sets StaleCacheEnabled field to given value.


### GetTieredCacheEnabled

`func (o *EdgeCacheModuleRequest) GetTieredCacheEnabled() bool`

GetTieredCacheEnabled returns the TieredCacheEnabled field if non-nil, zero value otherwise.

### GetTieredCacheEnabledOk

`func (o *EdgeCacheModuleRequest) GetTieredCacheEnabledOk() (*bool, bool)`

GetTieredCacheEnabledOk returns a tuple with the TieredCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCacheEnabled

`func (o *EdgeCacheModuleRequest) SetTieredCacheEnabled(v bool)`

SetTieredCacheEnabled sets TieredCacheEnabled field to given value.


### GetTieredCacheRegion

`func (o *EdgeCacheModuleRequest) GetTieredCacheRegion() string`

GetTieredCacheRegion returns the TieredCacheRegion field if non-nil, zero value otherwise.

### GetTieredCacheRegionOk

`func (o *EdgeCacheModuleRequest) GetTieredCacheRegionOk() (*string, bool)`

GetTieredCacheRegionOk returns a tuple with the TieredCacheRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCacheRegion

`func (o *EdgeCacheModuleRequest) SetTieredCacheRegion(v string)`

SetTieredCacheRegion sets TieredCacheRegion field to given value.

### HasTieredCacheRegion

`func (o *EdgeCacheModuleRequest) HasTieredCacheRegion() bool`

HasTieredCacheRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


