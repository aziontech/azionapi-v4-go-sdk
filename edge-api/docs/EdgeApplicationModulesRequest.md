# EdgeApplicationModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCache** | Pointer to [**CacheModuleRequest**](CacheModuleRequest.md) |  | [optional] 
**EdgeFunctions** | Pointer to [**EdgeFunctionModuleRequest**](EdgeFunctionModuleRequest.md) |  | [optional] 
**ApplicationAccelerator** | Pointer to [**ApplicationAcceleratorModuleRequest**](ApplicationAcceleratorModuleRequest.md) |  | [optional] 
**ImageProcessor** | Pointer to [**ImageProcessorModuleRequest**](ImageProcessorModuleRequest.md) |  | [optional] 
**TieredCache** | Pointer to [**TieredCacheModuleRequest**](TieredCacheModuleRequest.md) |  | [optional] 

## Methods

### NewEdgeApplicationModulesRequest

`func NewEdgeApplicationModulesRequest() *EdgeApplicationModulesRequest`

NewEdgeApplicationModulesRequest instantiates a new EdgeApplicationModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationModulesRequestWithDefaults

`func NewEdgeApplicationModulesRequestWithDefaults() *EdgeApplicationModulesRequest`

NewEdgeApplicationModulesRequestWithDefaults instantiates a new EdgeApplicationModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCache

`func (o *EdgeApplicationModulesRequest) GetEdgeCache() CacheModuleRequest`

GetEdgeCache returns the EdgeCache field if non-nil, zero value otherwise.

### GetEdgeCacheOk

`func (o *EdgeApplicationModulesRequest) GetEdgeCacheOk() (*CacheModuleRequest, bool)`

GetEdgeCacheOk returns a tuple with the EdgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCache

`func (o *EdgeApplicationModulesRequest) SetEdgeCache(v CacheModuleRequest)`

SetEdgeCache sets EdgeCache field to given value.

### HasEdgeCache

`func (o *EdgeApplicationModulesRequest) HasEdgeCache() bool`

HasEdgeCache returns a boolean if a field has been set.

### GetEdgeFunctions

`func (o *EdgeApplicationModulesRequest) GetEdgeFunctions() EdgeFunctionModuleRequest`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *EdgeApplicationModulesRequest) GetEdgeFunctionsOk() (*EdgeFunctionModuleRequest, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *EdgeApplicationModulesRequest) SetEdgeFunctions(v EdgeFunctionModuleRequest)`

SetEdgeFunctions sets EdgeFunctions field to given value.

### HasEdgeFunctions

`func (o *EdgeApplicationModulesRequest) HasEdgeFunctions() bool`

HasEdgeFunctions returns a boolean if a field has been set.

### GetApplicationAccelerator

`func (o *EdgeApplicationModulesRequest) GetApplicationAccelerator() ApplicationAcceleratorModuleRequest`

GetApplicationAccelerator returns the ApplicationAccelerator field if non-nil, zero value otherwise.

### GetApplicationAcceleratorOk

`func (o *EdgeApplicationModulesRequest) GetApplicationAcceleratorOk() (*ApplicationAcceleratorModuleRequest, bool)`

GetApplicationAcceleratorOk returns a tuple with the ApplicationAccelerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccelerator

`func (o *EdgeApplicationModulesRequest) SetApplicationAccelerator(v ApplicationAcceleratorModuleRequest)`

SetApplicationAccelerator sets ApplicationAccelerator field to given value.

### HasApplicationAccelerator

`func (o *EdgeApplicationModulesRequest) HasApplicationAccelerator() bool`

HasApplicationAccelerator returns a boolean if a field has been set.

### GetImageProcessor

`func (o *EdgeApplicationModulesRequest) GetImageProcessor() ImageProcessorModuleRequest`

GetImageProcessor returns the ImageProcessor field if non-nil, zero value otherwise.

### GetImageProcessorOk

`func (o *EdgeApplicationModulesRequest) GetImageProcessorOk() (*ImageProcessorModuleRequest, bool)`

GetImageProcessorOk returns a tuple with the ImageProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProcessor

`func (o *EdgeApplicationModulesRequest) SetImageProcessor(v ImageProcessorModuleRequest)`

SetImageProcessor sets ImageProcessor field to given value.

### HasImageProcessor

`func (o *EdgeApplicationModulesRequest) HasImageProcessor() bool`

HasImageProcessor returns a boolean if a field has been set.

### GetTieredCache

`func (o *EdgeApplicationModulesRequest) GetTieredCache() TieredCacheModuleRequest`

GetTieredCache returns the TieredCache field if non-nil, zero value otherwise.

### GetTieredCacheOk

`func (o *EdgeApplicationModulesRequest) GetTieredCacheOk() (*TieredCacheModuleRequest, bool)`

GetTieredCacheOk returns a tuple with the TieredCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredCache

`func (o *EdgeApplicationModulesRequest) SetTieredCache(v TieredCacheModuleRequest)`

SetTieredCache sets TieredCache field to given value.

### HasTieredCache

`func (o *EdgeApplicationModulesRequest) HasTieredCache() bool`

HasTieredCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


