# ResponseListEdgeApplicationFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**EdgeFunction** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 

## Methods

### NewResponseListEdgeApplicationFunctionInstance

`func NewResponseListEdgeApplicationFunctionInstance(id int64, name string, edgeFunction int64, lastEditor string, lastModified time.Time, ) *ResponseListEdgeApplicationFunctionInstance`

NewResponseListEdgeApplicationFunctionInstance instantiates a new ResponseListEdgeApplicationFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListEdgeApplicationFunctionInstanceWithDefaults

`func NewResponseListEdgeApplicationFunctionInstanceWithDefaults() *ResponseListEdgeApplicationFunctionInstance`

NewResponseListEdgeApplicationFunctionInstanceWithDefaults instantiates a new ResponseListEdgeApplicationFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListEdgeApplicationFunctionInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListEdgeApplicationFunctionInstance) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListEdgeApplicationFunctionInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListEdgeApplicationFunctionInstance) SetName(v string)`

SetName sets Name field to given value.


### GetJsonArgs

`func (o *ResponseListEdgeApplicationFunctionInstance) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *ResponseListEdgeApplicationFunctionInstance) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *ResponseListEdgeApplicationFunctionInstance) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *ResponseListEdgeApplicationFunctionInstance) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *ResponseListEdgeApplicationFunctionInstance) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetEdgeFunction

`func (o *ResponseListEdgeApplicationFunctionInstance) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *ResponseListEdgeApplicationFunctionInstance) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.


### GetActive

`func (o *ResponseListEdgeApplicationFunctionInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListEdgeApplicationFunctionInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListEdgeApplicationFunctionInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseListEdgeApplicationFunctionInstance) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListEdgeApplicationFunctionInstance) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListEdgeApplicationFunctionInstance) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListEdgeApplicationFunctionInstance) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListEdgeApplicationFunctionInstance) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


