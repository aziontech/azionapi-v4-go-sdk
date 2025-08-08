# ResponseListCacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**BrowserCache** | Pointer to [**BrowserCacheModule**](BrowserCacheModule.md) |  | [optional] 
**Modules** | Pointer to [**CacheSettingsModules**](CacheSettingsModules.md) |  | [optional] 

## Methods

### NewResponseListCacheSetting

`func NewResponseListCacheSetting(id int64, name string, ) *ResponseListCacheSetting`

NewResponseListCacheSetting instantiates a new ResponseListCacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListCacheSettingWithDefaults

`func NewResponseListCacheSettingWithDefaults() *ResponseListCacheSetting`

NewResponseListCacheSettingWithDefaults instantiates a new ResponseListCacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListCacheSetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListCacheSetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListCacheSetting) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListCacheSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListCacheSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListCacheSetting) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCache

`func (o *ResponseListCacheSetting) GetBrowserCache() BrowserCacheModule`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *ResponseListCacheSetting) GetBrowserCacheOk() (*BrowserCacheModule, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *ResponseListCacheSetting) SetBrowserCache(v BrowserCacheModule)`

SetBrowserCache sets BrowserCache field to given value.

### HasBrowserCache

`func (o *ResponseListCacheSetting) HasBrowserCache() bool`

HasBrowserCache returns a boolean if a field has been set.

### GetModules

`func (o *ResponseListCacheSetting) GetModules() CacheSettingsModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ResponseListCacheSetting) GetModulesOk() (*CacheSettingsModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ResponseListCacheSetting) SetModules(v CacheSettingsModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ResponseListCacheSetting) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


