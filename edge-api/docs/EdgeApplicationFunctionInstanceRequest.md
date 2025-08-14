# EdgeApplicationFunctionInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Args** | Pointer to **interface{}** |  | [optional] 
**EdgeFunction** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdgeApplicationFunctionInstanceRequest

`func NewEdgeApplicationFunctionInstanceRequest(name string, edgeFunction int64, ) *EdgeApplicationFunctionInstanceRequest`

NewEdgeApplicationFunctionInstanceRequest instantiates a new EdgeApplicationFunctionInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationFunctionInstanceRequestWithDefaults

`func NewEdgeApplicationFunctionInstanceRequestWithDefaults() *EdgeApplicationFunctionInstanceRequest`

NewEdgeApplicationFunctionInstanceRequestWithDefaults instantiates a new EdgeApplicationFunctionInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeApplicationFunctionInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationFunctionInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationFunctionInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *EdgeApplicationFunctionInstanceRequest) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *EdgeApplicationFunctionInstanceRequest) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *EdgeApplicationFunctionInstanceRequest) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *EdgeApplicationFunctionInstanceRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *EdgeApplicationFunctionInstanceRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *EdgeApplicationFunctionInstanceRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetEdgeFunction

`func (o *EdgeApplicationFunctionInstanceRequest) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *EdgeApplicationFunctionInstanceRequest) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *EdgeApplicationFunctionInstanceRequest) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.


### GetActive

`func (o *EdgeApplicationFunctionInstanceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplicationFunctionInstanceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplicationFunctionInstanceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplicationFunctionInstanceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


