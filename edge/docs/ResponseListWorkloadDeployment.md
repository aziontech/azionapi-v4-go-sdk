# ResponseListWorkloadDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tag** | **string** |  | 
**Binds** | [**WorkloadDeploymentBinds**](WorkloadDeploymentBinds.md) |  | 
**Current** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewResponseListWorkloadDeployment

`func NewResponseListWorkloadDeployment(id int32, tag string, binds WorkloadDeploymentBinds, ) *ResponseListWorkloadDeployment`

NewResponseListWorkloadDeployment instantiates a new ResponseListWorkloadDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListWorkloadDeploymentWithDefaults

`func NewResponseListWorkloadDeploymentWithDefaults() *ResponseListWorkloadDeployment`

NewResponseListWorkloadDeploymentWithDefaults instantiates a new ResponseListWorkloadDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListWorkloadDeployment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListWorkloadDeployment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListWorkloadDeployment) SetId(v int32)`

SetId sets Id field to given value.


### GetTag

`func (o *ResponseListWorkloadDeployment) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ResponseListWorkloadDeployment) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ResponseListWorkloadDeployment) SetTag(v string)`

SetTag sets Tag field to given value.


### GetBinds

`func (o *ResponseListWorkloadDeployment) GetBinds() WorkloadDeploymentBinds`

GetBinds returns the Binds field if non-nil, zero value otherwise.

### GetBindsOk

`func (o *ResponseListWorkloadDeployment) GetBindsOk() (*WorkloadDeploymentBinds, bool)`

GetBindsOk returns a tuple with the Binds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinds

`func (o *ResponseListWorkloadDeployment) SetBinds(v WorkloadDeploymentBinds)`

SetBinds sets Binds field to given value.


### GetCurrent

`func (o *ResponseListWorkloadDeployment) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ResponseListWorkloadDeployment) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ResponseListWorkloadDeployment) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *ResponseListWorkloadDeployment) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


