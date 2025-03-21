# ResponseListEdgeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **NullableString** |  | [readonly] 
**Modules** | Pointer to [**EdgeApplicationModules**](EdgeApplicationModules.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Debug** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewResponseListEdgeApplication

`func NewResponseListEdgeApplication(id int64, name string, lastEditor string, lastModified time.Time, productVersion NullableString, ) *ResponseListEdgeApplication`

NewResponseListEdgeApplication instantiates a new ResponseListEdgeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListEdgeApplicationWithDefaults

`func NewResponseListEdgeApplicationWithDefaults() *ResponseListEdgeApplication`

NewResponseListEdgeApplicationWithDefaults instantiates a new ResponseListEdgeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListEdgeApplication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListEdgeApplication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListEdgeApplication) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListEdgeApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListEdgeApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListEdgeApplication) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ResponseListEdgeApplication) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListEdgeApplication) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListEdgeApplication) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListEdgeApplication) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListEdgeApplication) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListEdgeApplication) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *ResponseListEdgeApplication) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseListEdgeApplication) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseListEdgeApplication) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### SetProductVersionNil

`func (o *ResponseListEdgeApplication) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *ResponseListEdgeApplication) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
### GetModules

`func (o *ResponseListEdgeApplication) GetModules() EdgeApplicationModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ResponseListEdgeApplication) GetModulesOk() (*EdgeApplicationModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ResponseListEdgeApplication) SetModules(v EdgeApplicationModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ResponseListEdgeApplication) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *ResponseListEdgeApplication) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListEdgeApplication) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListEdgeApplication) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListEdgeApplication) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDebug

`func (o *ResponseListEdgeApplication) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *ResponseListEdgeApplication) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *ResponseListEdgeApplication) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *ResponseListEdgeApplication) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


