# ApplicationModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCache** | Pointer to [**CacheModule**](CacheModule.md) |  | [optional] 
**Functions** | Pointer to [**EdgeFunctionModule**](EdgeFunctionModule.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**ApplicationAcceleratorModule**](ApplicationAcceleratorModule.md) |  | [optional] 
**ImageProcessor** | Pointer to [**ImageProcessorModule**](ImageProcessorModule.md) |  | [optional] 
**TieredCache** | Pointer to [**TieredCacheModule**](TieredCacheModule.md) |  | [optional] 

## Methods

### NewApplicationModules

`func NewApplicationModules() *ApplicationModules`

NewApplicationModules instantiates a new ApplicationModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationModulesWithDefaults

`func NewApplicationModulesWithDefaults() *ApplicationModules`

NewApplicationModulesWithDefaults instantiates a new ApplicationModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCache

`func (o *ApplicationModules) GetEdgeCache() CacheModule`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *ApplicationModules) GetEdgeCacheOk() (*CacheModule, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *ApplicationModules) SetEdgeCache(v CacheModule)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *ApplicationModules) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetFunctions

`func (o *ApplicationModules) GetFunctions() EdgeFunctionModule`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ApplicationModules) GetFunctionsOk() (*EdgeFunctionModule, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ApplicationModules) SetFunctions(v EdgeFunctionModule)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *ApplicationModules) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetApplicationAccelerator

`func (o *ApplicationModules) GetApplicationAccelerator() ApplicationAcceleratorModule`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *ApplicationModules) GetApplicationAcceleratorOk() (*ApplicationAcceleratorModule, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *ApplicationModules) SetApplicationAccelerator(v ApplicationAcceleratorModule)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *ApplicationModules) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.

### GetImageProcessor

`func (o *ApplicationModules) GetImageProcessor() ImageProcessorModule`

GetImageProcessor returns the ImageProcessor field if non-nil, zero value otherwise.

### GetImageProcessorOk

`func (o *ApplicationModules) GetImageProcessorOk() (*ImageProcessorModule, bool)`

GetImageProcessorOk returns a tuple with the ImageProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProcessor

`func (o *ApplicationModules) SetImageProcessor(v ImageProcessorModule)`

SetImageProcessor sets ImageProcessor field to given value.

### HasImageProcessor

`func (o *ApplicationModules) HasImageProcessor() bool`

HasImageProcessor returns a boolean if a field has been set.

### GetTieredCache

`func (o *ApplicationModules) GetTieredCache() TieredCacheModule`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *ApplicationModules) GetTieredCacheOk() (*TieredCacheModule, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *ApplicationModules) SetTieredCache(v TieredCacheModule)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *ApplicationModules) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


