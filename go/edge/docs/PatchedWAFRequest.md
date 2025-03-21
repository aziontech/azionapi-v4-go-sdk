# PatchedWAFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Name** | Pointer to **string** |  | [optional] 
**ThreatsConfiguration** | Pointer to [**WAFThreatsConfigurationRequest**](WAFThreatsConfigurationRequest.md) |  | [optional] 

## Methods

### NewPatchedWAFRequest

`func NewPatchedWAFRequest() *PatchedWAFRequest`

NewPatchedWAFRequest instantiates a new PatchedWAFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWAFRequestWithDefaults

`func NewPatchedWAFRequestWithDefaults() *PatchedWAFRequest`

NewPatchedWAFRequestWithDefaults instantiates a new PatchedWAFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PatchedWAFRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedWAFRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedWAFRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedWAFRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *PatchedWAFRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWAFRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWAFRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWAFRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetThreatsConfiguration

`func (o *PatchedWAFRequest) GetThreatsConfiguration() WAFThreatsConfigurationRequest`

GetThreatsConfiguration returns the ThreatsConfiguration field if non-nil, zero value otherwise.

### GetThreatsConfigurationOk

`func (o *PatchedWAFRequest) GetThreatsConfigurationOk() (*WAFThreatsConfigurationRequest, bool)`

GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatsConfiguration

`func (o *PatchedWAFRequest) SetThreatsConfiguration(v WAFThreatsConfigurationRequest)`

SetThreatsConfiguration sets ThreatsConfiguration field to given value.

### HasThreatsConfiguration

`func (o *PatchedWAFRequest) HasThreatsConfiguration() bool`

HasThreatsConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


