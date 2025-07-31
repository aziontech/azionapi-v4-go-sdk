# CacheSettingsModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCache** | Pointer to [**CacheSettingsEdgeCacheModule**](CacheSettingsEdgeCacheModule.md) |  | [optional] 
**TieredCache** | Pointer to [**NullableCacheSettingsTieredCacheModule**](CacheSettingsTieredCacheModule.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**CacheSettingsApplicationAcceleratorModule**](CacheSettingsApplicationAcceleratorModule.md) |  | [optional] 

## Methods

### NewCacheSettingsModules

`func NewCacheSettingsModules() *CacheSettingsModules`

NewCacheSettingsModules instantiates a new CacheSettingsModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsModulesWithDefaults

`func NewCacheSettingsModulesWithDefaults() *CacheSettingsModules`

NewCacheSettingsModulesWithDefaults instantiates a new CacheSettingsModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCache

`func (o *CacheSettingsModules) GetEdgeCache() CacheSettingsEdgeCacheModule`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *CacheSettingsModules) GetEdgeCacheOk() (*CacheSettingsEdgeCacheModule, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *CacheSettingsModules) SetEdgeCache(v CacheSettingsEdgeCacheModule)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *CacheSettingsModules) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetTieredCache

`func (o *CacheSettingsModules) GetTieredCache() CacheSettingsTieredCacheModule`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *CacheSettingsModules) GetTieredCacheOk() (*CacheSettingsTieredCacheModule, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *CacheSettingsModules) SetTieredCache(v CacheSettingsTieredCacheModule)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *CacheSettingsModules) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.

### SetTieredCacheNil

`func (o *CacheSettingsModules) SetTieredCacheNil(b bool)`

 SetTieredCacheNil sets the value for TieredCache to be an explicit nil

### UnsetTieredCache
`func (o *CacheSettingsModules) UnsetTieredCache()`

UnsetTieredCache ensures that no value is present for TieredCache, not even an explicit nil
### GetApplicationAccelerator

`func (o *CacheSettingsModules) GetApplicationAccelerator() CacheSettingsApplicationAcceleratorModule`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *CacheSettingsModules) GetApplicationAcceleratorOk() (*CacheSettingsApplicationAcceleratorModule, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *CacheSettingsModules) SetApplicationAccelerator(v CacheSettingsApplicationAcceleratorModule)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *CacheSettingsModules) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


