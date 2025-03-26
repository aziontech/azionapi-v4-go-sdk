# PatchedEdgeApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to [**EdgeApplicationModulesRequest**](EdgeApplicationModulesRequest.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedEdgeApplicationRequest

`func NewPatchedEdgeApplicationRequest() *PatchedEdgeApplicationRequest`

NewPatchedEdgeApplicationRequest instantiates a new PatchedEdgeApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeApplicationRequestWithDefaults

`func NewPatchedEdgeApplicationRequestWithDefaults() *PatchedEdgeApplicationRequest`

NewPatchedEdgeApplicationRequestWithDefaults instantiates a new PatchedEdgeApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeApplicationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeApplicationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModules

`func (o *PatchedEdgeApplicationRequest) GetModules() EdgeApplicationModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchedEdgeApplicationRequest) GetModulesOk() (*EdgeApplicationModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchedEdgeApplicationRequest) SetModules(v EdgeApplicationModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchedEdgeApplicationRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeApplicationRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeApplicationRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeApplicationRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeApplicationRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDebug

`func (o *PatchedEdgeApplicationRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *PatchedEdgeApplicationRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *PatchedEdgeApplicationRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *PatchedEdgeApplicationRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


