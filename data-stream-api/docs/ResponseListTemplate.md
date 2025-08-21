# ResponseListTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**Custom** | **bool** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**DataSet** | **string** |  | 

## Methods

### NewResponseListTemplate

`func NewResponseListTemplate(id int64, name string, lastEditor string, lastModified time.Time, custom bool, dataSet string, ) *ResponseListTemplate`

NewResponseListTemplate instantiates a new ResponseListTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListTemplateWithDefaults

`func NewResponseListTemplateWithDefaults() *ResponseListTemplate`

NewResponseListTemplateWithDefaults instantiates a new ResponseListTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListTemplate) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ResponseListTemplate) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListTemplate) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListTemplate) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListTemplate) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListTemplate) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListTemplate) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetCustom

`func (o *ResponseListTemplate) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ResponseListTemplate) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ResponseListTemplate) SetCustom(v bool)`

SetCustom sets Custom field to given value.


### GetActive

`func (o *ResponseListTemplate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListTemplate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListTemplate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListTemplate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDataSet

`func (o *ResponseListTemplate) GetDataSet() string`

GetDataSet returns the DataSet field if non-nil, zero value otherwise.

### GetDataSetOk

`func (o *ResponseListTemplate) GetDataSetOk() (*string, bool)`

GetDataSetOk returns a tuple with the DataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSet

`func (o *ResponseListTemplate) SetDataSet(v string)`

SetDataSet sets DataSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


