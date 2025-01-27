# ResponseListWAF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **NullableString** |  | [readonly] 
**ThreatsConfiguration** | Pointer to [**WAFThreatsConfiguration**](WAFThreatsConfiguration.md) |  | [optional] 

## Methods

### NewResponseListWAF

`func NewResponseListWAF(id int32, name string, lastEditor string, lastModified time.Time, productVersion NullableString, ) *ResponseListWAF`

NewResponseListWAF instantiates a new ResponseListWAF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListWAFWithDefaults

`func NewResponseListWAFWithDefaults() *ResponseListWAF`

NewResponseListWAFWithDefaults instantiates a new ResponseListWAF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListWAF) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListWAF) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListWAF) SetId(v int32)`

SetId sets Id field to given value.


### GetActive

`func (o *ResponseListWAF) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListWAF) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListWAF) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListWAF) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *ResponseListWAF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListWAF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListWAF) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ResponseListWAF) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListWAF) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListWAF) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListWAF) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListWAF) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListWAF) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *ResponseListWAF) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseListWAF) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseListWAF) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### SetProductVersionNil

`func (o *ResponseListWAF) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *ResponseListWAF) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
### GetThreatsConfiguration

`func (o *ResponseListWAF) GetThreatsConfiguration() WAFThreatsConfiguration`

GetThreatsConfiguration returns the ThreatsConfiguration field if non-nil, zero value otherwise.

### GetThreatsConfigurationOk

`func (o *ResponseListWAF) GetThreatsConfigurationOk() (*WAFThreatsConfiguration, bool)`

GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatsConfiguration

`func (o *ResponseListWAF) SetThreatsConfiguration(v WAFThreatsConfiguration)`

SetThreatsConfiguration sets ThreatsConfiguration field to given value.

### HasThreatsConfiguration

`func (o *ResponseListWAF) HasThreatsConfiguration() bool`

HasThreatsConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


