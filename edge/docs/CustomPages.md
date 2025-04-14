# CustomPages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | [readonly] 
**ConnectorCustomPages** | [**ConnectorCustomPages**](ConnectorCustomPages.md) |  | 

## Methods

### NewCustomPages

`func NewCustomPages(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, connectorCustomPages ConnectorCustomPages, ) *CustomPages`

NewCustomPages instantiates a new CustomPages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPagesWithDefaults

`func NewCustomPagesWithDefaults() *CustomPages`

NewCustomPagesWithDefaults instantiates a new CustomPages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomPages) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomPages) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomPages) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CustomPages) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPages) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPages) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *CustomPages) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *CustomPages) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *CustomPages) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *CustomPages) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CustomPages) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CustomPages) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *CustomPages) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomPages) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomPages) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CustomPages) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *CustomPages) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *CustomPages) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *CustomPages) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetConnectorCustomPages

`func (o *CustomPages) GetConnectorCustomPages() ConnectorCustomPages`

GetConnectorCustomPages returns the ConnectorCustomPages field if non-nil, zero value otherwise.

### GetConnectorCustomPagesOk

`func (o *CustomPages) GetConnectorCustomPagesOk() (*ConnectorCustomPages, bool)`

GetConnectorCustomPagesOk returns a tuple with the ConnectorCustomPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCustomPages

`func (o *CustomPages) SetConnectorCustomPages(v ConnectorCustomPages)`

SetConnectorCustomPages sets ConnectorCustomPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


