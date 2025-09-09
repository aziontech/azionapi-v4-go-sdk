# ResponseBadRequestApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Modules** | Pointer to [**ResponseBadRequestSerializerMetaclassModulesField**](ResponseBadRequestSerializerMetaclassModulesField.md) |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**Debug** | Pointer to **[]string** |  | [optional] 
**ProductVersion** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestApplication

`func NewResponseBadRequestApplication() *ResponseBadRequestApplication`

NewResponseBadRequestApplication instantiates a new ResponseBadRequestApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestApplicationWithDefaults

`func NewResponseBadRequestApplicationWithDefaults() *ResponseBadRequestApplication`

NewResponseBadRequestApplicationWithDefaults instantiates a new ResponseBadRequestApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseBadRequestApplication) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestApplication) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestApplication) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ResponseBadRequestApplication) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBadRequestApplication) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBadRequestApplication) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBadRequestApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestApplication) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestApplication) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestApplication) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestApplication) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestApplication) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestApplication) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestApplication) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestApplication) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetModules

`func (o *ResponseBadRequestApplication) GetModules() ResponseBadRequestSerializerMetaclassModulesField`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ResponseBadRequestApplication) GetModulesOk() (*ResponseBadRequestSerializerMetaclassModulesField, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ResponseBadRequestApplication) SetModules(v ResponseBadRequestSerializerMetaclassModulesField)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ResponseBadRequestApplication) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestApplication) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestApplication) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestApplication) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestApplication) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDebug

`func (o *ResponseBadRequestApplication) GetDebug() []string`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *ResponseBadRequestApplication) GetDebugOk() (*[]string, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *ResponseBadRequestApplication) SetDebug(v []string)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *ResponseBadRequestApplication) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseBadRequestApplication) GetProductVersion() []string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseBadRequestApplication) GetProductVersionOk() (*[]string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseBadRequestApplication) SetProductVersion(v []string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ResponseBadRequestApplication) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestApplication) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestApplication) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestApplication) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestApplication) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


