# WAFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ThreatsConfiguration** | Pointer to [**WAFThreatsConfigurationRequest**](WAFThreatsConfigurationRequest.md) |  | [optional] 

## Methods

### NewWAFRequest

`func NewWAFRequest(name string, ) *WAFRequest`

NewWAFRequest instantiates a new WAFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFRequestWithDefaults

`func NewWAFRequestWithDefaults() *WAFRequest`

NewWAFRequestWithDefaults instantiates a new WAFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WAFRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WAFRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WAFRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WAFRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *WAFRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WAFRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WAFRequest) SetName(v string)`

SetName sets Name field to given value.


### GetThreatsConfiguration

`func (o *WAFRequest) GetThreatsConfiguration() WAFThreatsConfigurationRequest`

GetThreatsConfiguration returns the ThreatsConfiguration field if non-nil, zero value otherwise.

### GetThreatsConfigurationOk

`func (o *WAFRequest) GetThreatsConfigurationOk() (*WAFThreatsConfigurationRequest, bool)`

GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatsConfiguration

`func (o *WAFRequest) SetThreatsConfiguration(v WAFThreatsConfigurationRequest)`

SetThreatsConfiguration sets ThreatsConfiguration field to given value.

### HasThreatsConfiguration

`func (o *WAFRequest) HasThreatsConfiguration() bool`

HasThreatsConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


