# EdgeFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] 
**Runtime** | Pointer to **string** | * &#x60;azion_js&#x60; - Azion JavaScript | [optional] 
**ExecutionEnvironment** | Pointer to **string** | * &#x60;firewall&#x60; - Firewall * &#x60;application&#x60; - Application | [optional] 
**Code** | **string** | String containing the function code. Maximum size: 20MB. | 
**DefaultArgs** | Pointer to **interface{}** |  | [optional] 
**ReferenceCount** | **int64** |  | [readonly] 
**Version** | **string** | Installed version, which may not be the latest if the vendor has released updates since installation. | [readonly] 
**Vendor** | **string** |  | [readonly] 

## Methods

### NewEdgeFunctions

`func NewEdgeFunctions(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, code string, referenceCount int64, version string, vendor string, ) *EdgeFunctions`

NewEdgeFunctions instantiates a new EdgeFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionsWithDefaults

`func NewEdgeFunctionsWithDefaults() *EdgeFunctions`

NewEdgeFunctionsWithDefaults instantiates a new EdgeFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeFunctions) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFunctions) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFunctions) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeFunctions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFunctions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFunctions) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *EdgeFunctions) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeFunctions) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeFunctions) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeFunctions) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeFunctions) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeFunctions) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *EdgeFunctions) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeFunctions) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeFunctions) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetActive

`func (o *EdgeFunctions) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFunctions) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFunctions) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFunctions) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRuntime

`func (o *EdgeFunctions) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *EdgeFunctions) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *EdgeFunctions) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *EdgeFunctions) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetExecutionEnvironment

`func (o *EdgeFunctions) GetExecutionEnvironment() string`

GetExecutionEnvironment returns the ExecutionEnvironment field if non-nil, zero value otherwise.

### GetExecutionEnvironmentOk

`func (o *EdgeFunctions) GetExecutionEnvironmentOk() (*string, bool)`

GetExecutionEnvironmentOk returns a tuple with the ExecutionEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnvironment

`func (o *EdgeFunctions) SetExecutionEnvironment(v string)`

SetExecutionEnvironment sets ExecutionEnvironment field to given value.

### HasExecutionEnvironment

`func (o *EdgeFunctions) HasExecutionEnvironment() bool`

HasExecutionEnvironment returns a boolean if a field has been set.

### GetCode

`func (o *EdgeFunctions) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EdgeFunctions) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EdgeFunctions) SetCode(v string)`

SetCode sets Code field to given value.


### GetDefaultArgs

`func (o *EdgeFunctions) GetDefaultArgs() interface{}`

GetDefaultArgs returns the DefaultArgs field if non-nil, zero value otherwise.

### GetDefaultArgsOk

`func (o *EdgeFunctions) GetDefaultArgsOk() (*interface{}, bool)`

GetDefaultArgsOk returns a tuple with the DefaultArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultArgs

`func (o *EdgeFunctions) SetDefaultArgs(v interface{})`

SetDefaultArgs sets DefaultArgs field to given value.

### HasDefaultArgs

`func (o *EdgeFunctions) HasDefaultArgs() bool`

HasDefaultArgs returns a boolean if a field has been set.

### SetDefaultArgsNil

`func (o *EdgeFunctions) SetDefaultArgsNil(b bool)`

 SetDefaultArgsNil sets the value for DefaultArgs to be an explicit nil

### UnsetDefaultArgs
`func (o *EdgeFunctions) UnsetDefaultArgs()`

UnsetDefaultArgs ensures that no value is present for DefaultArgs, not even an explicit nil
### GetReferenceCount

`func (o *EdgeFunctions) GetReferenceCount() int64`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *EdgeFunctions) GetReferenceCountOk() (*int64, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *EdgeFunctions) SetReferenceCount(v int64)`

SetReferenceCount sets ReferenceCount field to given value.


### GetVersion

`func (o *EdgeFunctions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EdgeFunctions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EdgeFunctions) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVendor

`func (o *EdgeFunctions) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EdgeFunctions) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EdgeFunctions) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


