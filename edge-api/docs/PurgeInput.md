# PurgeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]string** |  | 
**Layer** | Pointer to **string** | * &#x60;edge_cache&#x60; - Edge Cache * &#x60;tiered_cache&#x60; - Tiered Cache | [optional] 

## Methods

### NewPurgeInput

`func NewPurgeInput(items []string, ) *PurgeInput`

NewPurgeInput instantiates a new PurgeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeInputWithDefaults

`func NewPurgeInputWithDefaults() *PurgeInput`

NewPurgeInputWithDefaults instantiates a new PurgeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PurgeInput) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PurgeInput) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PurgeInput) SetItems(v []string)`

SetItems sets Items field to given value.


### GetLayer

`func (o *PurgeInput) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *PurgeInput) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *PurgeInput) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *PurgeInput) HasLayer() bool`

HasLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


