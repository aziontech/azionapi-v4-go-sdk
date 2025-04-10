# DataStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DataSource** | **string** | * &#x60;http&#x60; - Edge Applications * &#x60;waf&#x60; - WAF Events * &#x60;cells_console&#x60; - Edge Functions * &#x60;rtm_activity&#x60; - Activity History | 
**DataSetId** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Filters** | [**DataStreamFilterRequest**](DataStreamFilterRequest.md) |  | 
**Endpoint** | [**EndpointRequest**](EndpointRequest.md) |  | 

## Methods

### NewDataStreamRequest

`func NewDataStreamRequest(name string, dataSource string, dataSetId int64, filters DataStreamFilterRequest, endpoint EndpointRequest, ) *DataStreamRequest`

NewDataStreamRequest instantiates a new DataStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamRequestWithDefaults

`func NewDataStreamRequestWithDefaults() *DataStreamRequest`

NewDataStreamRequestWithDefaults instantiates a new DataStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataStreamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStreamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStreamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDataSource

`func (o *DataStreamRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataStreamRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataStreamRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetDataSetId

`func (o *DataStreamRequest) GetDataSetId() int64`

GetDataSetId returns the DataSetId field if non-nil, zero value otherwise.

### GetDataSetIdOk

`func (o *DataStreamRequest) GetDataSetIdOk() (*int64, bool)`

GetDataSetIdOk returns a tuple with the DataSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSetId

`func (o *DataStreamRequest) SetDataSetId(v int64)`

SetDataSetId sets DataSetId field to given value.


### GetActive

`func (o *DataStreamRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DataStreamRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DataStreamRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DataStreamRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFilters

`func (o *DataStreamRequest) GetFilters() DataStreamFilterRequest`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DataStreamRequest) GetFiltersOk() (*DataStreamFilterRequest, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DataStreamRequest) SetFilters(v DataStreamFilterRequest)`

SetFilters sets Filters field to given value.


### GetEndpoint

`func (o *DataStreamRequest) GetEndpoint() EndpointRequest`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DataStreamRequest) GetEndpointOk() (*EndpointRequest, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DataStreamRequest) SetEndpoint(v EndpointRequest)`

SetEndpoint sets Endpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


