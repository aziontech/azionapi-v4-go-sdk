# PatchedWorkloadDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** |  | [optional] 
**Binds** | Pointer to [**WorkloadDeploymentBindsRequest**](WorkloadDeploymentBindsRequest.md) |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedWorkloadDeploymentRequest

`func NewPatchedWorkloadDeploymentRequest() *PatchedWorkloadDeploymentRequest`

NewPatchedWorkloadDeploymentRequest instantiates a new PatchedWorkloadDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWorkloadDeploymentRequestWithDefaults

`func NewPatchedWorkloadDeploymentRequestWithDefaults() *PatchedWorkloadDeploymentRequest`

NewPatchedWorkloadDeploymentRequestWithDefaults instantiates a new PatchedWorkloadDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *PatchedWorkloadDeploymentRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PatchedWorkloadDeploymentRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PatchedWorkloadDeploymentRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PatchedWorkloadDeploymentRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetBinds

`func (o *PatchedWorkloadDeploymentRequest) GetBinds() WorkloadDeploymentBindsRequest`

GetBinds returns the Binds field if non-nil, zero value otherwise.

### GetBindsOk

`func (o *PatchedWorkloadDeploymentRequest) GetBindsOk() (*WorkloadDeploymentBindsRequest, bool)`

GetBindsOk returns a tuple with the Binds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinds

`func (o *PatchedWorkloadDeploymentRequest) SetBinds(v WorkloadDeploymentBindsRequest)`

SetBinds sets Binds field to given value.

### HasBinds

`func (o *PatchedWorkloadDeploymentRequest) HasBinds() bool`

HasBinds returns a boolean if a field has been set.

### GetCurrent

`func (o *PatchedWorkloadDeploymentRequest) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *PatchedWorkloadDeploymentRequest) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *PatchedWorkloadDeploymentRequest) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *PatchedWorkloadDeploymentRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


