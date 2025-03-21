# CacheSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**BrowserCache** | [**BrowserCacheModuleRequest**](BrowserCacheModuleRequest.md) |  | 
**EdgeCache** | [**EdgeCacheModuleRequest**](EdgeCacheModuleRequest.md) |  | 
**ApplicationControls** | [**ApplicationControlsModuleRequest**](ApplicationControlsModuleRequest.md) |  | 
**SliceControls** | [**SliceControlsModuleRequest**](SliceControlsModuleRequest.md) |  | 

## Methods

### NewCacheSettingRequest

`func NewCacheSettingRequest(name string, browserCache BrowserCacheModuleRequest, edgeCache EdgeCacheModuleRequest, applicationControls ApplicationControlsModuleRequest, sliceControls SliceControlsModuleRequest, ) *CacheSettingRequest`

NewCacheSettingRequest instantiates a new CacheSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingRequestWithDefaults

`func NewCacheSettingRequestWithDefaults() *CacheSettingRequest`

NewCacheSettingRequestWithDefaults instantiates a new CacheSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CacheSettingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CacheSettingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CacheSettingRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCache

`func (o *CacheSettingRequest) GetBrowserCache() BrowserCacheModuleRequest`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *CacheSettingRequest) GetBrowserCacheOk() (*BrowserCacheModuleRequest, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *CacheSettingRequest) SetBrowserCache(v BrowserCacheModuleRequest)`

SetBrowserCache sets BrowserCache field to given value.


### GetEdgeCache

`func (o *CacheSettingRequest) GetEdgeCache() EdgeCacheModuleRequest`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *CacheSettingRequest) GetEdgeCacheOk() (*EdgeCacheModuleRequest, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *CacheSettingRequest) SetEdgeCache(v EdgeCacheModuleRequest)`

SetEdgeCache sets EdgeCache field to given value.


### GetApplicationControls

`func (o *CacheSettingRequest) GetApplicationControls() ApplicationControlsModuleRequest`

GetApplicationControls returns the ApplicationControls field if non-nil, zero value otherwise.

### GetApplicationControlsOk

`func (o *CacheSettingRequest) GetApplicationControlsOk() (*ApplicationControlsModuleRequest, bool)`

GetApplicationControlsOk returns a tuple with the ApplicationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControls

`func (o *CacheSettingRequest) SetApplicationControls(v ApplicationControlsModuleRequest)`

SetApplicationControls sets ApplicationControls field to given value.


### GetSliceControls

`func (o *CacheSettingRequest) GetSliceControls() SliceControlsModuleRequest`

GetSliceControls returns the SliceControls field if non-nil, zero value otherwise.

### GetSliceControlsOk

`func (o *CacheSettingRequest) GetSliceControlsOk() (*SliceControlsModuleRequest, bool)`

GetSliceControlsOk returns a tuple with the SliceControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceControls

`func (o *CacheSettingRequest) SetSliceControls(v SliceControlsModuleRequest)`

SetSliceControls sets SliceControls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


