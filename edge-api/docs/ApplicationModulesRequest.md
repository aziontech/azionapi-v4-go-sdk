# ApplicationModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCache** | Pointer to [**CacheModuleRequest**](CacheModuleRequest.md) |  | [optional] 
**Functions** | Pointer to [**EdgeFunctionModuleRequest**](EdgeFunctionModuleRequest.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**ApplicationAcceleratorModuleRequest**](ApplicationAcceleratorModuleRequest.md) |  | [optional] 
**ImageProcessor** | Pointer to [**ImageProcessorModuleRequest**](ImageProcessorModuleRequest.md) |  | [optional] 
**TieredCache** | Pointer to [**TieredCacheModuleRequest**](TieredCacheModuleRequest.md) |  | [optional] 

## Methods

### NewApplicationModulesRequest

`func NewApplicationModulesRequest() *ApplicationModulesRequest`

NewApplicationModulesRequest instantiates a new ApplicationModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationModulesRequestWithDefaults

`func NewApplicationModulesRequestWithDefaults() *ApplicationModulesRequest`

NewApplicationModulesRequestWithDefaults instantiates a new ApplicationModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCache

`func (o *ApplicationModulesRequest) GetEdgeCache() CacheModuleRequest`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *ApplicationModulesRequest) GetEdgeCacheOk() (*CacheModuleRequest, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *ApplicationModulesRequest) SetEdgeCache(v CacheModuleRequest)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *ApplicationModulesRequest) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetFunctions

`func (o *ApplicationModulesRequest) GetFunctions() EdgeFunctionModuleRequest`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ApplicationModulesRequest) GetFunctionsOk() (*EdgeFunctionModuleRequest, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ApplicationModulesRequest) SetFunctions(v EdgeFunctionModuleRequest)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *ApplicationModulesRequest) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetApplicationAccelerator

`func (o *ApplicationModulesRequest) GetApplicationAccelerator() ApplicationAcceleratorModuleRequest`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *ApplicationModulesRequest) GetApplicationAcceleratorOk() (*ApplicationAcceleratorModuleRequest, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *ApplicationModulesRequest) SetApplicationAccelerator(v ApplicationAcceleratorModuleRequest)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *ApplicationModulesRequest) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.

### GetImageProcessor

`func (o *ApplicationModulesRequest) GetImageProcessor() ImageProcessorModuleRequest`

GetImageProcessor returns the ImageProcessor field if non-nil, zero value otherwise.

### GetImageProcessorOk

`func (o *ApplicationModulesRequest) GetImageProcessorOk() (*ImageProcessorModuleRequest, bool)`

GetImageProcessorOk returns a tuple with the ImageProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProcessor

`func (o *ApplicationModulesRequest) SetImageProcessor(v ImageProcessorModuleRequest)`

SetImageProcessor sets ImageProcessor field to given value.

### HasImageProcessor

`func (o *ApplicationModulesRequest) HasImageProcessor() bool`

HasImageProcessor returns a boolean if a field has been set.

### GetTieredCache

`func (o *ApplicationModulesRequest) GetTieredCache() TieredCacheModuleRequest`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *ApplicationModulesRequest) GetTieredCacheOk() (*TieredCacheModuleRequest, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *ApplicationModulesRequest) SetTieredCache(v TieredCacheModuleRequest)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *ApplicationModulesRequest) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


