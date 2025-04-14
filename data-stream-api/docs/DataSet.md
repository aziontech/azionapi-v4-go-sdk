# DataSet

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

### NewDataSet

`func NewDataSet(id int64, name string, lastEditor string, lastModified time.Time, custom bool, dataSet string, ) *DataSet`

NewDataSet instantiates a new DataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSetWithDefaults

`func NewDataSetWithDefaults() *DataSet`

NewDataSetWithDefaults instantiates a new DataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSet) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DataSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSet) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *DataSet) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *DataSet) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *DataSet) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *DataSet) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DataSet) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DataSet) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetCustom

`func (o *DataSet) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *DataSet) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *DataSet) SetCustom(v bool)`

SetCustom sets Custom field to given value.


### GetActive

`func (o *DataSet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DataSet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DataSet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DataSet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDataSet

`func (o *DataSet) GetDataSet() string`

GetDataSet returns the DataSet field if non-nil, zero value otherwise.

### GetDataSetOk

`func (o *DataSet) GetDataSetOk() (*string, bool)`

GetDataSetOk returns a tuple with the DataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSet

`func (o *DataSet) SetDataSet(v string)`

SetDataSet sets DataSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


