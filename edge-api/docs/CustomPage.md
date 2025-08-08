# CustomPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | 
**Pages** | [**[]Page**](Page.md) |  | 

## Methods

### NewCustomPage

`func NewCustomPage(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, pages []Page, ) *CustomPage`

NewCustomPage instantiates a new CustomPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPageWithDefaults

`func NewCustomPageWithDefaults() *CustomPage`

NewCustomPageWithDefaults instantiates a new CustomPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomPage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomPage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomPage) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CustomPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomPage) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *CustomPage) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *CustomPage) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *CustomPage) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *CustomPage) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CustomPage) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CustomPage) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *CustomPage) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomPage) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomPage) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CustomPage) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *CustomPage) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *CustomPage) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *CustomPage) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetPages

`func (o *CustomPage) GetPages() []Page`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *CustomPage) GetPagesOk() (*[]Page, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *CustomPage) SetPages(v []Page)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


