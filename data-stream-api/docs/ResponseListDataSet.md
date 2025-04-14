# ResponseListDataSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Custom** | **bool** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] 
**DataSet** | **string** |  | 

## Methods

### NewResponseListDataSet

`func NewResponseListDataSet(id int64, name string, lastEditor string, lastModified time.Time, custom bool, dataSet string, ) *ResponseListDataSet`

NewResponseListDataSet instantiates a new ResponseListDataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListDataSetWithDefaults

`func NewResponseListDataSetWithDefaults() *ResponseListDataSet`

NewResponseListDataSetWithDefaults instantiates a new ResponseListDataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListDataSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListDataSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListDataSet) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListDataSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListDataSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListDataSet) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ResponseListDataSet) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListDataSet) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListDataSet) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListDataSet) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListDataSet) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListDataSet) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetCustom

`func (o *ResponseListDataSet) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ResponseListDataSet) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ResponseListDataSet) SetCustom(v bool)`

SetCustom sets Custom field to given value.


### GetActive

`func (o *ResponseListDataSet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListDataSet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListDataSet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListDataSet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDataSet

`func (o *ResponseListDataSet) GetDataSet() string`

GetDataSet returns the DataSet field if non-nil, zero value otherwise.

### GetDataSetOk

`func (o *ResponseListDataSet) GetDataSetOk() (*string, bool)`

GetDataSetOk returns a tuple with the DataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSet

`func (o *ResponseListDataSet) SetDataSet(v string)`

SetDataSet sets DataSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


