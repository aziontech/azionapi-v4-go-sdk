# WAF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **NullableString** |  | [readonly] 
**ThreatsConfiguration** | Pointer to [**WAFThreatsConfiguration**](WAFThreatsConfiguration.md) |  | [optional] 

## Methods

### NewWAF

`func NewWAF(id int64, name string, lastEditor string, lastModified time.Time, productVersion NullableString, ) *WAF`

NewWAF instantiates a new WAF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFWithDefaults

`func NewWAFWithDefaults() *WAF`

NewWAFWithDefaults instantiates a new WAF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WAF) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WAF) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WAF) SetId(v int64)`

SetId sets Id field to given value.


### GetActive

`func (o *WAF) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WAF) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WAF) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WAF) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *WAF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WAF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WAF) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *WAF) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *WAF) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *WAF) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *WAF) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *WAF) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *WAF) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *WAF) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *WAF) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *WAF) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### SetProductVersionNil

`func (o *WAF) SetProductVersionNil(b bool)`

 SetProductVersionNil sets the value for ProductVersion to be an explicit nil

### UnsetProductVersion
`func (o *WAF) UnsetProductVersion()`

UnsetProductVersion ensures that no value is present for ProductVersion, not even an explicit nil
### GetThreatsConfiguration

`func (o *WAF) GetThreatsConfiguration() WAFThreatsConfiguration`

GetThreatsConfiguration returns the ThreatsConfiguration field if non-nil, zero value otherwise.

### GetThreatsConfigurationOk

`func (o *WAF) GetThreatsConfigurationOk() (*WAFThreatsConfiguration, bool)`

GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatsConfiguration

`func (o *WAF) SetThreatsConfiguration(v WAFThreatsConfiguration)`

SetThreatsConfiguration sets ThreatsConfiguration field to given value.

### HasThreatsConfiguration

`func (o *WAF) HasThreatsConfiguration() bool`

HasThreatsConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


