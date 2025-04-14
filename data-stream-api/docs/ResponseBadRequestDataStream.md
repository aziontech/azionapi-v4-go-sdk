# ResponseBadRequestDataStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **[]string** |  | [optional] 
**DataSource** | Pointer to **[]string** |  | [optional] 
**DataSetId** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**Filters** | Pointer to [**ResponseBadRequestSerializerMetaclassFiltersField**](ResponseBadRequestSerializerMetaclassFiltersField.md) |  | [optional] 
**ProductVersion** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestDataStream

`func NewResponseBadRequestDataStream() *ResponseBadRequestDataStream`

NewResponseBadRequestDataStream instantiates a new ResponseBadRequestDataStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestDataStreamWithDefaults

`func NewResponseBadRequestDataStreamWithDefaults() *ResponseBadRequestDataStream`

NewResponseBadRequestDataStreamWithDefaults instantiates a new ResponseBadRequestDataStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseBadRequestDataStream) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestDataStream) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestDataStream) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestDataStream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataSource

`func (o *ResponseBadRequestDataStream) GetDataSource() []string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ResponseBadRequestDataStream) GetDataSourceOk() (*[]string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ResponseBadRequestDataStream) SetDataSource(v []string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ResponseBadRequestDataStream) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSetId

`func (o *ResponseBadRequestDataStream) GetDataSetId() []string`

GetDataSetId returns the DataSetId field if non-nil, zero value otherwise.

### GetDataSetIdOk

`func (o *ResponseBadRequestDataStream) GetDataSetIdOk() (*[]string, bool)`

GetDataSetIdOk returns a tuple with the DataSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSetId

`func (o *ResponseBadRequestDataStream) SetDataSetId(v []string)`

SetDataSetId sets DataSetId field to given value.

### HasDataSetId

`func (o *ResponseBadRequestDataStream) HasDataSetId() bool`

HasDataSetId returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestDataStream) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestDataStream) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestDataStream) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestDataStream) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFilters

`func (o *ResponseBadRequestDataStream) GetFilters() ResponseBadRequestSerializerMetaclassFiltersField`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ResponseBadRequestDataStream) GetFiltersOk() (*ResponseBadRequestSerializerMetaclassFiltersField, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ResponseBadRequestDataStream) SetFilters(v ResponseBadRequestSerializerMetaclassFiltersField)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ResponseBadRequestDataStream) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseBadRequestDataStream) GetProductVersion() []string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseBadRequestDataStream) GetProductVersionOk() (*[]string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseBadRequestDataStream) SetProductVersion(v []string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ResponseBadRequestDataStream) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestDataStream) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestDataStream) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestDataStream) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestDataStream) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestDataStream) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestDataStream) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestDataStream) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestDataStream) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestDataStream) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestDataStream) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestDataStream) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestDataStream) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


