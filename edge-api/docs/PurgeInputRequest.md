# PurgeInputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]string** |  | 
**Layer** | Pointer to **string** | * &#x60;cache&#x60; - Cache * &#x60;tiered_cache&#x60; - Tiered Cache | [optional] 

## Methods

### NewPurgeInputRequest

`func NewPurgeInputRequest(items []string, ) *PurgeInputRequest`

NewPurgeInputRequest instantiates a new PurgeInputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeInputRequestWithDefaults

`func NewPurgeInputRequestWithDefaults() *PurgeInputRequest`

NewPurgeInputRequestWithDefaults instantiates a new PurgeInputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PurgeInputRequest) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PurgeInputRequest) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PurgeInputRequest) SetItems(v []string)`

SetItems sets Items field to given value.


### GetLayer

`func (o *PurgeInputRequest) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *PurgeInputRequest) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *PurgeInputRequest) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *PurgeInputRequest) HasLayer() bool`

HasLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


