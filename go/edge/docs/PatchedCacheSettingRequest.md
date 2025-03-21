# PatchedCacheSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**BrowserCache** | Pointer to [**BrowserCacheModuleRequest**](BrowserCacheModuleRequest.md) |  | [optional] 
**EdgeCache** | Pointer to [**EdgeCacheModuleRequest**](EdgeCacheModuleRequest.md) |  | [optional] 
**ApplicationControls** | Pointer to [**ApplicationControlsModuleRequest**](ApplicationControlsModuleRequest.md) |  | [optional] 
**SliceControls** | Pointer to [**SliceControlsModuleRequest**](SliceControlsModuleRequest.md) |  | [optional] 

## Methods

### NewPatchedCacheSettingRequest

`func NewPatchedCacheSettingRequest() *PatchedCacheSettingRequest`

NewPatchedCacheSettingRequest instantiates a new PatchedCacheSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCacheSettingRequestWithDefaults

`func NewPatchedCacheSettingRequestWithDefaults() *PatchedCacheSettingRequest`

NewPatchedCacheSettingRequestWithDefaults instantiates a new PatchedCacheSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCacheSettingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCacheSettingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCacheSettingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCacheSettingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBrowserCache

`func (o *PatchedCacheSettingRequest) GetBrowserCache() BrowserCacheModuleRequest`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *PatchedCacheSettingRequest) GetBrowserCacheOk() (*BrowserCacheModuleRequest, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *PatchedCacheSettingRequest) SetBrowserCache(v BrowserCacheModuleRequest)`

SetBrowserCache sets BrowserCache field to given value.

### HasBrowserCache

`func (o *PatchedCacheSettingRequest) HasBrowserCache() bool`

HasBrowserCache returns a boolean if a field has been set.

### GetEdgeCache

`func (o *PatchedCacheSettingRequest) GetEdgeCache() EdgeCacheModuleRequest`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *PatchedCacheSettingRequest) GetEdgeCacheOk() (*EdgeCacheModuleRequest, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *PatchedCacheSettingRequest) SetEdgeCache(v EdgeCacheModuleRequest)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *PatchedCacheSettingRequest) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetApplicationControls

`func (o *PatchedCacheSettingRequest) GetApplicationControls() ApplicationControlsModuleRequest`

GetApplicationControls returns the ApplicationControls field if non-nil, zero value otherwise.

### GetApplicationControlsOk

`func (o *PatchedCacheSettingRequest) GetApplicationControlsOk() (*ApplicationControlsModuleRequest, bool)`

GetApplicationControlsOk returns a tuple with the ApplicationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControls

`func (o *PatchedCacheSettingRequest) SetApplicationControls(v ApplicationControlsModuleRequest)`

SetApplicationControls sets ApplicationControls field to given value.

### HasApplicationControls

`func (o *PatchedCacheSettingRequest) HasApplicationControls() bool`

HasApplicationControls returns a boolean if a field has been set.

### GetSliceControls

`func (o *PatchedCacheSettingRequest) GetSliceControls() SliceControlsModuleRequest`

GetSliceControls returns the SliceControls field if non-nil, zero value otherwise.

### GetSliceControlsOk

`func (o *PatchedCacheSettingRequest) GetSliceControlsOk() (*SliceControlsModuleRequest, bool)`

GetSliceControlsOk returns a tuple with the SliceControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceControls

`func (o *PatchedCacheSettingRequest) SetSliceControls(v SliceControlsModuleRequest)`

SetSliceControls sets SliceControls field to given value.

### HasSliceControls

`func (o *PatchedCacheSettingRequest) HasSliceControls() bool`

HasSliceControls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


