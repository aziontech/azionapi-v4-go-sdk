# WorkloadDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Tag** | **string** |  | 
**Binds** | [**WorkloadDeploymentBinds**](WorkloadDeploymentBinds.md) |  | 
**Current** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkloadDeployment

`func NewWorkloadDeployment(id int64, tag string, binds WorkloadDeploymentBinds, ) *WorkloadDeployment`

NewWorkloadDeployment instantiates a new WorkloadDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeploymentWithDefaults

`func NewWorkloadDeploymentWithDefaults() *WorkloadDeployment`

NewWorkloadDeploymentWithDefaults instantiates a new WorkloadDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkloadDeployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadDeployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadDeployment) SetId(v int64)`

SetId sets Id field to given value.


### GetTag

`func (o *WorkloadDeployment) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WorkloadDeployment) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WorkloadDeployment) SetTag(v string)`

SetTag sets Tag field to given value.


### GetBinds

`func (o *WorkloadDeployment) GetBinds() WorkloadDeploymentBinds`

GetBinds returns the Binds field if non-nil, zero value otherwise.

### GetBindsOk

`func (o *WorkloadDeployment) GetBindsOk() (*WorkloadDeploymentBinds, bool)`

GetBindsOk returns a tuple with the Binds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinds

`func (o *WorkloadDeployment) SetBinds(v WorkloadDeploymentBinds)`

SetBinds sets Binds field to given value.


### GetCurrent

`func (o *WorkloadDeployment) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WorkloadDeployment) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WorkloadDeployment) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *WorkloadDeployment) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


