# EdgeApplicationFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Args** | Pointer to **interface{}** |  | [optional] 
**EdgeFunction** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 

## Methods

### NewEdgeApplicationFunctionInstance

`func NewEdgeApplicationFunctionInstance(id int64, name string, edgeFunction int64, lastEditor string, lastModified time.Time, ) *EdgeApplicationFunctionInstance`

NewEdgeApplicationFunctionInstance instantiates a new EdgeApplicationFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationFunctionInstanceWithDefaults

`func NewEdgeApplicationFunctionInstanceWithDefaults() *EdgeApplicationFunctionInstance`

NewEdgeApplicationFunctionInstanceWithDefaults instantiates a new EdgeApplicationFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeApplicationFunctionInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeApplicationFunctionInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeApplicationFunctionInstance) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeApplicationFunctionInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationFunctionInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationFunctionInstance) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *EdgeApplicationFunctionInstance) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *EdgeApplicationFunctionInstance) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *EdgeApplicationFunctionInstance) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *EdgeApplicationFunctionInstance) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *EdgeApplicationFunctionInstance) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *EdgeApplicationFunctionInstance) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetEdgeFunction

`func (o *EdgeApplicationFunctionInstance) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *EdgeApplicationFunctionInstance) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *EdgeApplicationFunctionInstance) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.


### GetActive

`func (o *EdgeApplicationFunctionInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplicationFunctionInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplicationFunctionInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplicationFunctionInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *EdgeApplicationFunctionInstance) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeApplicationFunctionInstance) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeApplicationFunctionInstance) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeApplicationFunctionInstance) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeApplicationFunctionInstance) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeApplicationFunctionInstance) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


