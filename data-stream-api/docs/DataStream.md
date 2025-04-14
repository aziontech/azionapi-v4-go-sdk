# DataStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**DataSource** | **string** | * &#x60;http&#x60; - Edge Applications * &#x60;waf&#x60; - WAF Events * &#x60;cells_console&#x60; - Edge Functions * &#x60;rtm_activity&#x60; - Activity History | 
**DataSetId** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Filters** | [**DataStreamFilter**](DataStreamFilter.md) |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 
**Endpoint** | [**Endpoint**](Endpoint.md) |  | 

## Methods

### NewDataStream

`func NewDataStream(id int64, name string, dataSource string, dataSetId int64, filters DataStreamFilter, lastEditor string, lastModified time.Time, productVersion string, endpoint Endpoint, ) *DataStream`

NewDataStream instantiates a new DataStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamWithDefaults

`func NewDataStreamWithDefaults() *DataStream`

NewDataStreamWithDefaults instantiates a new DataStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataStream) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStream) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStream) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DataStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStream) SetName(v string)`

SetName sets Name field to given value.


### GetDataSource

`func (o *DataStream) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataStream) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataStream) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetDataSetId

`func (o *DataStream) GetDataSetId() int64`

GetDataSetId returns the DataSetId field if non-nil, zero value otherwise.

### GetDataSetIdOk

`func (o *DataStream) GetDataSetIdOk() (*int64, bool)`

GetDataSetIdOk returns a tuple with the DataSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSetId

`func (o *DataStream) SetDataSetId(v int64)`

SetDataSetId sets DataSetId field to given value.


### GetActive

`func (o *DataStream) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DataStream) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DataStream) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DataStream) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFilters

`func (o *DataStream) GetFilters() DataStreamFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DataStream) GetFiltersOk() (*DataStreamFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DataStream) SetFilters(v DataStreamFilter)`

SetFilters sets Filters field to given value.


### GetLastEditor

`func (o *DataStream) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *DataStream) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *DataStream) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *DataStream) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DataStream) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DataStream) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *DataStream) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *DataStream) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *DataStream) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetEndpoint

`func (o *DataStream) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DataStream) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DataStream) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


