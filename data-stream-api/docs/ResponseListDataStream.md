# ResponseListDataStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] 
**Inputs** | [**[]InputPolymorphicInputDataSourceAttributes**](InputPolymorphicInputDataSourceAttributes.md) |  | 
**Transform** | [**[]TransformPolymorphic**](TransformPolymorphic.md) |  | 
**Outputs** | [**[]Output**](Output.md) |  | 

## Methods

### NewResponseListDataStream

`func NewResponseListDataStream(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, inputs []InputPolymorphicInputDataSourceAttributes, transform []TransformPolymorphic, outputs []Output, ) *ResponseListDataStream`

NewResponseListDataStream instantiates a new ResponseListDataStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListDataStreamWithDefaults

`func NewResponseListDataStreamWithDefaults() *ResponseListDataStream`

NewResponseListDataStreamWithDefaults instantiates a new ResponseListDataStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListDataStream) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListDataStream) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListDataStream) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListDataStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListDataStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListDataStream) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ResponseListDataStream) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListDataStream) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListDataStream) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListDataStream) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListDataStream) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListDataStream) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *ResponseListDataStream) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseListDataStream) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseListDataStream) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetActive

`func (o *ResponseListDataStream) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListDataStream) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListDataStream) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListDataStream) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInputs

`func (o *ResponseListDataStream) GetInputs() []InputPolymorphicInputDataSourceAttributes`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ResponseListDataStream) GetInputsOk() (*[]InputPolymorphicInputDataSourceAttributes, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ResponseListDataStream) SetInputs(v []InputPolymorphicInputDataSourceAttributes)`

SetInputs sets Inputs field to given value.


### GetTransform

`func (o *ResponseListDataStream) GetTransform() []TransformPolymorphic`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *ResponseListDataStream) GetTransformOk() (*[]TransformPolymorphic, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *ResponseListDataStream) SetTransform(v []TransformPolymorphic)`

SetTransform sets Transform field to given value.


### GetOutputs

`func (o *ResponseListDataStream) GetOutputs() []Output`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ResponseListDataStream) GetOutputsOk() (*[]Output, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ResponseListDataStream) SetOutputs(v []Output)`

SetOutputs sets Outputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


