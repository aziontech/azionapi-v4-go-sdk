# WorkloadDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** |  | 
**Binds** | [**WorkloadDeploymentBindsRequest**](WorkloadDeploymentBindsRequest.md) |  | 
**Current** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkloadDeploymentRequest

`func NewWorkloadDeploymentRequest(tag string, binds WorkloadDeploymentBindsRequest, ) *WorkloadDeploymentRequest`

NewWorkloadDeploymentRequest instantiates a new WorkloadDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeploymentRequestWithDefaults

`func NewWorkloadDeploymentRequestWithDefaults() *WorkloadDeploymentRequest`

NewWorkloadDeploymentRequestWithDefaults instantiates a new WorkloadDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *WorkloadDeploymentRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WorkloadDeploymentRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WorkloadDeploymentRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetBinds

`func (o *WorkloadDeploymentRequest) GetBinds() WorkloadDeploymentBindsRequest`

GetBinds returns the Binds field if non-nil, zero value otherwise.

### GetBindsOk

`func (o *WorkloadDeploymentRequest) GetBindsOk() (*WorkloadDeploymentBindsRequest, bool)`

GetBindsOk returns a tuple with the Binds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinds

`func (o *WorkloadDeploymentRequest) SetBinds(v WorkloadDeploymentBindsRequest)`

SetBinds sets Binds field to given value.


### GetCurrent

`func (o *WorkloadDeploymentRequest) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WorkloadDeploymentRequest) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WorkloadDeploymentRequest) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *WorkloadDeploymentRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


