# SliceControlsModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SliceConfigurationEnabled** | **bool** |  | 
**SliceEdgeCachingEnabled** | **bool** |  | 
**SliceTieredCachingEnabled** | **bool** |  | 
**SliceConfigurationRange** | Pointer to **int64** |  | [optional] 

## Methods

### NewSliceControlsModuleRequest

`func NewSliceControlsModuleRequest(sliceConfigurationEnabled bool, sliceEdgeCachingEnabled bool, sliceTieredCachingEnabled bool, ) *SliceControlsModuleRequest`

NewSliceControlsModuleRequest instantiates a new SliceControlsModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceControlsModuleRequestWithDefaults

`func NewSliceControlsModuleRequestWithDefaults() *SliceControlsModuleRequest`

NewSliceControlsModuleRequestWithDefaults instantiates a new SliceControlsModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSliceConfigurationEnabled

`func (o *SliceControlsModuleRequest) GetSliceConfigurationEnabled() bool`

GetSliceConfigurationEnabled returns the SliceConfigurationEnabled field if non-nil, zero value otherwise.

### GetSliceConfigurationEnabledOk

`func (o *SliceControlsModuleRequest) GetSliceConfigurationEnabledOk() (*bool, bool)`

GetSliceConfigurationEnabledOk returns a tuple with the SliceConfigurationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceConfigurationEnabled

`func (o *SliceControlsModuleRequest) SetSliceConfigurationEnabled(v bool)`

SetSliceConfigurationEnabled sets SliceConfigurationEnabled field to given value.


### GetSliceEdgeCachingEnabled

`func (o *SliceControlsModuleRequest) GetSliceEdgeCachingEnabled() bool`

GetSliceEdgeCachingEnabled returns the SliceEdgeCachingEnabled field if non-nil, zero value otherwise.

### GetSliceEdgeCachingEnabledOk

`func (o *SliceControlsModuleRequest) GetSliceEdgeCachingEnabledOk() (*bool, bool)`

GetSliceEdgeCachingEnabledOk returns a tuple with the SliceEdgeCachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceEdgeCachingEnabled

`func (o *SliceControlsModuleRequest) SetSliceEdgeCachingEnabled(v bool)`

SetSliceEdgeCachingEnabled sets SliceEdgeCachingEnabled field to given value.


### GetSliceTieredCachingEnabled

`func (o *SliceControlsModuleRequest) GetSliceTieredCachingEnabled() bool`

GetSliceTieredCachingEnabled returns the SliceTieredCachingEnabled field if non-nil, zero value otherwise.

### GetSliceTieredCachingEnabledOk

`func (o *SliceControlsModuleRequest) GetSliceTieredCachingEnabledOk() (*bool, bool)`

GetSliceTieredCachingEnabledOk returns a tuple with the SliceTieredCachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceTieredCachingEnabled

`func (o *SliceControlsModuleRequest) SetSliceTieredCachingEnabled(v bool)`

SetSliceTieredCachingEnabled sets SliceTieredCachingEnabled field to given value.


### GetSliceConfigurationRange

`func (o *SliceControlsModuleRequest) GetSliceConfigurationRange() int64`

GetSliceConfigurationRange returns the SliceConfigurationRange field if non-nil, zero value otherwise.

### GetSliceConfigurationRangeOk

`func (o *SliceControlsModuleRequest) GetSliceConfigurationRangeOk() (*int64, bool)`

GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceConfigurationRange

`func (o *SliceControlsModuleRequest) SetSliceConfigurationRange(v int64)`

SetSliceConfigurationRange sets SliceConfigurationRange field to given value.

### HasSliceConfigurationRange

`func (o *SliceControlsModuleRequest) HasSliceConfigurationRange() bool`

HasSliceConfigurationRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


