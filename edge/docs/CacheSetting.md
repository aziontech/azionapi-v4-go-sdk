# CacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**BrowserCache** | [**BrowserCacheModule**](BrowserCacheModule.md) |  | 
**EdgeCache** | [**EdgeCacheModule**](EdgeCacheModule.md) |  | 
**ApplicationControls** | [**ApplicationControlsModule**](ApplicationControlsModule.md) |  | 
**SliceControls** | [**SliceControlsModule**](SliceControlsModule.md) |  | 

## Methods

### NewCacheSetting

`func NewCacheSetting(id int64, name string, browserCache BrowserCacheModule, edgeCache EdgeCacheModule, applicationControls ApplicationControlsModule, sliceControls SliceControlsModule, ) *CacheSetting`

NewCacheSetting instantiates a new CacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingWithDefaults

`func NewCacheSettingWithDefaults() *CacheSetting`

NewCacheSettingWithDefaults instantiates a new CacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CacheSetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CacheSetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CacheSetting) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CacheSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSetting) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCache

`func (o *CacheSetting) GetBrowserCache() BrowserCacheModule`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *CacheSetting) GetBrowserCacheOk() (*BrowserCacheModule, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *CacheSetting) SetBrowserCache(v BrowserCacheModule)`

SetBrowserCache sets BrowserCache field to given value.


### GetEdgeCache

`func (o *CacheSetting) GetEdgeCache() EdgeCacheModule`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *CacheSetting) GetEdgeCacheOk() (*EdgeCacheModule, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *CacheSetting) SetEdgeCache(v EdgeCacheModule)`

SetEdgeCache sets EdgeCache field to given value.


### GetApplicationControls

`func (o *CacheSetting) GetApplicationControls() ApplicationControlsModule`

GetApplicationControls returns the ApplicationControls field if non-nil, zero value otherwise.

### GetApplicationControlsOk

`func (o *CacheSetting) GetApplicationControlsOk() (*ApplicationControlsModule, bool)`

GetApplicationControlsOk returns a tuple with the ApplicationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControls

`func (o *CacheSetting) SetApplicationControls(v ApplicationControlsModule)`

SetApplicationControls sets ApplicationControls field to given value.


### GetSliceControls

`func (o *CacheSetting) GetSliceControls() SliceControlsModule`

GetSliceControls returns the SliceControls field if non-nil, zero value otherwise.

### GetSliceControlsOk

`func (o *CacheSetting) GetSliceControlsOk() (*SliceControlsModule, bool)`

GetSliceControlsOk returns a tuple with the SliceControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceControls

`func (o *CacheSetting) SetSliceControls(v SliceControlsModule)`

SetSliceControls sets SliceControls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


