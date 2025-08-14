# EdgeFirewallFunctionInstance

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

### NewEdgeFirewallFunctionInstance

`func NewEdgeFirewallFunctionInstance(id int64, name string, edgeFunction int64, lastEditor string, lastModified time.Time, ) *EdgeFirewallFunctionInstance`

NewEdgeFirewallFunctionInstance instantiates a new EdgeFirewallFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallFunctionInstanceWithDefaults

`func NewEdgeFirewallFunctionInstanceWithDefaults() *EdgeFirewallFunctionInstance`

NewEdgeFirewallFunctionInstanceWithDefaults instantiates a new EdgeFirewallFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeFirewallFunctionInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFirewallFunctionInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFirewallFunctionInstance) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeFirewallFunctionInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewallFunctionInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewallFunctionInstance) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *EdgeFirewallFunctionInstance) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *EdgeFirewallFunctionInstance) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *EdgeFirewallFunctionInstance) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *EdgeFirewallFunctionInstance) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *EdgeFirewallFunctionInstance) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *EdgeFirewallFunctionInstance) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetEdgeFunction

`func (o *EdgeFirewallFunctionInstance) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *EdgeFirewallFunctionInstance) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *EdgeFirewallFunctionInstance) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.


### GetActive

`func (o *EdgeFirewallFunctionInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFirewallFunctionInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFirewallFunctionInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFirewallFunctionInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *EdgeFirewallFunctionInstance) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeFirewallFunctionInstance) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeFirewallFunctionInstance) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeFirewallFunctionInstance) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeFirewallFunctionInstance) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeFirewallFunctionInstance) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


