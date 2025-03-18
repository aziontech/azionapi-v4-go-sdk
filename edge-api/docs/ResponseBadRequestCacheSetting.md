# ResponseBadRequestCacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**BrowserCache** | Pointer to [**ResponseBadRequestSerializerMetaclassBrowserCacheField**](ResponseBadRequestSerializerMetaclassBrowserCacheField.md) |  | [optional] 
**EdgeCache** | Pointer to [**ResponseBadRequestSerializerMetaclassEdgeCacheField**](ResponseBadRequestSerializerMetaclassEdgeCacheField.md) |  | [optional] 
**ApplicationControls** | Pointer to [**ResponseBadRequestSerializerMetaclassApplicationControlsField**](ResponseBadRequestSerializerMetaclassApplicationControlsField.md) |  | [optional] 
**SliceControls** | Pointer to [**ResponseBadRequestSerializerMetaclassSliceControlsField**](ResponseBadRequestSerializerMetaclassSliceControlsField.md) |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestCacheSetting

`func NewResponseBadRequestCacheSetting() *ResponseBadRequestCacheSetting`

NewResponseBadRequestCacheSetting instantiates a new ResponseBadRequestCacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestCacheSettingWithDefaults

`func NewResponseBadRequestCacheSettingWithDefaults() *ResponseBadRequestCacheSetting`

NewResponseBadRequestCacheSettingWithDefaults instantiates a new ResponseBadRequestCacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseBadRequestCacheSetting) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBadRequestCacheSetting) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBadRequestCacheSetting) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBadRequestCacheSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResponseBadRequestCacheSetting) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestCacheSetting) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestCacheSetting) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestCacheSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBrowserCache

`func (o *ResponseBadRequestCacheSetting) GetBrowserCache() ResponseBadRequestSerializerMetaclassBrowserCacheField`

GetBrowserCache returns the BrowserCache field if non-nil, zero value otherwise.

### GetBrowserCacheOk

`func (o *ResponseBadRequestCacheSetting) GetBrowserCacheOk() (*ResponseBadRequestSerializerMetaclassBrowserCacheField, bool)`

GetBrowserCacheOk returns a tuple with the BrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCache

`func (o *ResponseBadRequestCacheSetting) SetBrowserCache(v ResponseBadRequestSerializerMetaclassBrowserCacheField)`

SetBrowserCache sets BrowserCache field to given value.

### HasBrowserCache

`func (o *ResponseBadRequestCacheSetting) HasBrowserCache() bool`

HasBrowserCache returns a boolean if a field has been set.

### GetEdgeCache

`func (o *ResponseBadRequestCacheSetting) GetEdgeCache() ResponseBadRequestSerializerMetaclassEdgeCacheField`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *ResponseBadRequestCacheSetting) GetEdgeCacheOk() (*ResponseBadRequestSerializerMetaclassEdgeCacheField, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *ResponseBadRequestCacheSetting) SetEdgeCache(v ResponseBadRequestSerializerMetaclassEdgeCacheField)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *ResponseBadRequestCacheSetting) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetApplicationControls

`func (o *ResponseBadRequestCacheSetting) GetApplicationControls() ResponseBadRequestSerializerMetaclassApplicationControlsField`

GetApplicationControls returns the ApplicationControls field if non-nil, zero value otherwise.

### GetApplicationControlsOk

`func (o *ResponseBadRequestCacheSetting) GetApplicationControlsOk() (*ResponseBadRequestSerializerMetaclassApplicationControlsField, bool)`

GetApplicationControlsOk returns a tuple with the ApplicationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationControls

`func (o *ResponseBadRequestCacheSetting) SetApplicationControls(v ResponseBadRequestSerializerMetaclassApplicationControlsField)`

SetApplicationControls sets ApplicationControls field to given value.

### HasApplicationControls

`func (o *ResponseBadRequestCacheSetting) HasApplicationControls() bool`

HasApplicationControls returns a boolean if a field has been set.

### GetSliceControls

`func (o *ResponseBadRequestCacheSetting) GetSliceControls() ResponseBadRequestSerializerMetaclassSliceControlsField`

GetSliceControls returns the SliceControls field if non-nil, zero value otherwise.

### GetSliceControlsOk

`func (o *ResponseBadRequestCacheSetting) GetSliceControlsOk() (*ResponseBadRequestSerializerMetaclassSliceControlsField, bool)`

GetSliceControlsOk returns a tuple with the SliceControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceControls

`func (o *ResponseBadRequestCacheSetting) SetSliceControls(v ResponseBadRequestSerializerMetaclassSliceControlsField)`

SetSliceControls sets SliceControls field to given value.

### HasSliceControls

`func (o *ResponseBadRequestCacheSetting) HasSliceControls() bool`

HasSliceControls returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestCacheSetting) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestCacheSetting) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestCacheSetting) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestCacheSetting) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


