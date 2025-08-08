# CacheSettingsEdgeCacheModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to **string** | * &#x60;honor&#x60; - Honor Origin Cache Settings * &#x60;override&#x60; - Override Cache Settings | [optional] 
**MaxAge** | Pointer to **int64** | To use a value lower than 60s, the Application Acceleration module must be enabled on the Edge Application. | [optional] 
**StaleCache** | Pointer to [**StateCacheModuleRequest**](StateCacheModuleRequest.md) |  | [optional] 
**LargeFileCache** | Pointer to [**LargeFileCacheModuleRequest**](LargeFileCacheModuleRequest.md) |  | [optional] 

## Methods

### NewCacheSettingsEdgeCacheModuleRequest

`func NewCacheSettingsEdgeCacheModuleRequest() *CacheSettingsEdgeCacheModuleRequest`

NewCacheSettingsEdgeCacheModuleRequest instantiates a new CacheSettingsEdgeCacheModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsEdgeCacheModuleRequestWithDefaults

`func NewCacheSettingsEdgeCacheModuleRequestWithDefaults() *CacheSettingsEdgeCacheModuleRequest`

NewCacheSettingsEdgeCacheModuleRequestWithDefaults instantiates a new CacheSettingsEdgeCacheModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *CacheSettingsEdgeCacheModuleRequest) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *CacheSettingsEdgeCacheModuleRequest) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *CacheSettingsEdgeCacheModuleRequest) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *CacheSettingsEdgeCacheModuleRequest) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetMaxAge

`func (o *CacheSettingsEdgeCacheModuleRequest) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *CacheSettingsEdgeCacheModuleRequest) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *CacheSettingsEdgeCacheModuleRequest) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *CacheSettingsEdgeCacheModuleRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetStaleCache

`func (o *CacheSettingsEdgeCacheModuleRequest) GetStaleCache() StateCacheModuleRequest`

GetStaleCache returns the StaleCache field if non-nil, zero value otherwise.

### GetStaleCacheOk

`func (o *CacheSettingsEdgeCacheModuleRequest) GetStaleCacheOk() (*StateCacheModuleRequest, bool)`

GetStaleCacheOk returns a tuple with the StaleCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleCache

`func (o *CacheSettingsEdgeCacheModuleRequest) SetStaleCache(v StateCacheModuleRequest)`

SetStaleCache sets StaleCache field to given value.

### HasStaleCache

`func (o *CacheSettingsEdgeCacheModuleRequest) HasStaleCache() bool`

HasStaleCache returns a boolean if a field has been set.

### GetLargeFileCache

`func (o *CacheSettingsEdgeCacheModuleRequest) GetLargeFileCache() LargeFileCacheModuleRequest`

GetLargeFileCache returns the LargeFileCache field if non-nil, zero value otherwise.

### GetLargeFileCacheOk

`func (o *CacheSettingsEdgeCacheModuleRequest) GetLargeFileCacheOk() (*LargeFileCacheModuleRequest, bool)`

GetLargeFileCacheOk returns a tuple with the LargeFileCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFileCache

`func (o *CacheSettingsEdgeCacheModuleRequest) SetLargeFileCache(v LargeFileCacheModuleRequest)`

SetLargeFileCache sets LargeFileCache field to given value.

### HasLargeFileCache

`func (o *CacheSettingsEdgeCacheModuleRequest) HasLargeFileCache() bool`

HasLargeFileCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


