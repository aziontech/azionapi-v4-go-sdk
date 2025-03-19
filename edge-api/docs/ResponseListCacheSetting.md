# ResponseListCacheSetting

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

### NewResponseListCacheSetting

`func NewResponseListCacheSetting(id int64, name string, browserCache BrowserCacheModule, edgeCache EdgeCacheModule, applicationControls ApplicationControlsModule, sliceControls SliceControlsModule, ) *ResponseListCacheSetting`

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


### GetEdgeCache

`func (o *ResponseListCacheSetting) GetEdgeCache() EdgeCacheModule`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *ResponseListCacheSetting) GetEdgeCacheOk() (*EdgeCacheModule, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *ResponseListCacheSetting) SetEdgeCache(v EdgeCacheModule)`

SetEdgeCache sets EdgeCache field to given value.


### GetApplicationControls

`func (o *ResponseListCacheSetting) GetApplicationControls() ApplicationControlsModule`

GetApplicationControls returns the ApplicationControls field if non-nil, zero value otherwise.

### GetApplicationControlsOk

`func (o *ResponseListCacheSetting) GetApplicationControlsOk() (*ApplicationControlsModule, bool)`

GetApplicationControlsOk returns a tuple with the ApplicationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControls

`func (o *ResponseListCacheSetting) SetApplicationControls(v ApplicationControlsModule)`

SetApplicationControls sets ApplicationControls field to given value.


### GetSliceControls

`func (o *ResponseListCacheSetting) GetSliceControls() SliceControlsModule`

GetSliceControls returns the SliceControls field if non-nil, zero value otherwise.

### GetSliceControlsOk

`func (o *ResponseListCacheSetting) GetSliceControlsOk() (*SliceControlsModule, bool)`

GetSliceControlsOk returns a tuple with the SliceControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceControls

`func (o *ResponseListCacheSetting) SetSliceControls(v SliceControlsModule)`

SetSliceControls sets SliceControls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


