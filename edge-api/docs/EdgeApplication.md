# EdgeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **NullableString** |  | [readonly] 
**Modules** | Pointer to [**EdgeApplicationModules**](EdgeApplicationModules.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdgeApplication

`func NewEdgeApplication(id int32, name string, lastEditor string, lastModified time.Time, productVersion NullableString, ) *EdgeApplication`

NewEdgeApplication instantiates a new EdgeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationWithDefaults

`func NewEdgeApplicationWithDefaults() *EdgeApplication`

NewEdgeApplicationWithDefaults instantiates a new EdgeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeApplication) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeApplication) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeApplication) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplication) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *EdgeApplication) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeApplication) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeApplication) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeApplication) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeApplication) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeApplication) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *EdgeApplication) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeApplication) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeApplication) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### SetProductVersionNil

`func (o *EdgeApplication) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *EdgeApplication) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
### GetModules

`func (o *EdgeApplication) GetModules() EdgeApplicationModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeApplication) GetModulesOk() (*EdgeApplicationModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeApplication) SetModules(v EdgeApplicationModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *EdgeApplication) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *EdgeApplication) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplication) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplication) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplication) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDebug

`func (o *EdgeApplication) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *EdgeApplication) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *EdgeApplication) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *EdgeApplication) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


