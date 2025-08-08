# CacheSettingsModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCache** | Pointer to [**CacheSettingsEdgeCacheModuleRequest**](CacheSettingsEdgeCacheModuleRequest.md) |  | [optional] 
**TieredCache** | Pointer to [**NullableCacheSettingsTieredCacheModuleRequest**](CacheSettingsTieredCacheModuleRequest.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**CacheSettingsApplicationAcceleratorModuleRequest**](CacheSettingsApplicationAcceleratorModuleRequest.md) |  | [optional] 

## Methods

### NewCacheSettingsModulesRequest

`func NewCacheSettingsModulesRequest() *CacheSettingsModulesRequest`

NewCacheSettingsModulesRequest instantiates a new CacheSettingsModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsModulesRequestWithDefaults

`func NewCacheSettingsModulesRequestWithDefaults() *CacheSettingsModulesRequest`

NewCacheSettingsModulesRequestWithDefaults instantiates a new CacheSettingsModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCache

`func (o *CacheSettingsModulesRequest) GetEdgeCache() CacheSettingsEdgeCacheModuleRequest`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *CacheSettingsModulesRequest) GetEdgeCacheOk() (*CacheSettingsEdgeCacheModuleRequest, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *CacheSettingsModulesRequest) SetEdgeCache(v CacheSettingsEdgeCacheModuleRequest)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *CacheSettingsModulesRequest) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetTieredCache

`func (o *CacheSettingsModulesRequest) GetTieredCache() CacheSettingsTieredCacheModuleRequest`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *CacheSettingsModulesRequest) GetTieredCacheOk() (*CacheSettingsTieredCacheModuleRequest, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *CacheSettingsModulesRequest) SetTieredCache(v CacheSettingsTieredCacheModuleRequest)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *CacheSettingsModulesRequest) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.

### SetTieredCacheNil

`func (o *CacheSettingsModulesRequest) SetTieredCacheNil(b bool)`

 SetTieredCacheNil sets the value for TieredCache to be an explicit nil

### UnsetTieredCache
`func (o *CacheSettingsModulesRequest) UnsetTieredCache()`

UnsetTieredCache ensures that no value is present for TieredCache, not even an explicit nil
### GetApplicationAccelerator

`func (o *CacheSettingsModulesRequest) GetApplicationAccelerator() CacheSettingsApplicationAcceleratorModuleRequest`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *CacheSettingsModulesRequest) GetApplicationAcceleratorOk() (*CacheSettingsApplicationAcceleratorModuleRequest, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *CacheSettingsModulesRequest) SetApplicationAccelerator(v CacheSettingsApplicationAcceleratorModuleRequest)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *CacheSettingsModulesRequest) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


