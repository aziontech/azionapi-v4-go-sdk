# PatchedDataStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to **string** | * &#x60;http&#x60; - Edge Applications * &#x60;waf&#x60; - WAF Events * &#x60;cells_console&#x60; - Edge Functions * &#x60;rtm_activity&#x60; - Activity History | [optional] 
**DataSetId** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to [**DataStreamFilterRequest**](DataStreamFilterRequest.md) |  | [optional] 

## Methods

### NewPatchedDataStreamRequest

`func NewPatchedDataStreamRequest() *PatchedDataStreamRequest`

NewPatchedDataStreamRequest instantiates a new PatchedDataStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDataStreamRequestWithDefaults

`func NewPatchedDataStreamRequestWithDefaults() *PatchedDataStreamRequest`

NewPatchedDataStreamRequestWithDefaults instantiates a new PatchedDataStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedDataStreamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDataStreamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDataStreamRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDataStreamRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataSource

`func (o *PatchedDataStreamRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PatchedDataStreamRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PatchedDataStreamRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PatchedDataStreamRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSetId

`func (o *PatchedDataStreamRequest) GetDataSetId() int64`

GetDataSetId returns the DataSetId field if non-nil, zero value otherwise.

### GetDataSetIdOk

`func (o *PatchedDataStreamRequest) GetDataSetIdOk() (*int64, bool)`

GetDataSetIdOk returns a tuple with the DataSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSetId

`func (o *PatchedDataStreamRequest) SetDataSetId(v int64)`

SetDataSetId sets DataSetId field to given value.

### HasDataSetId

`func (o *PatchedDataStreamRequest) HasDataSetId() bool`

HasDataSetId returns a boolean if a field has been set.

### GetActive

`func (o *PatchedDataStreamRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedDataStreamRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedDataStreamRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedDataStreamRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedDataStreamRequest) GetFilters() DataStreamFilterRequest`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedDataStreamRequest) GetFiltersOk() (*DataStreamFilterRequest, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedDataStreamRequest) SetFilters(v DataStreamFilterRequest)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedDataStreamRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


