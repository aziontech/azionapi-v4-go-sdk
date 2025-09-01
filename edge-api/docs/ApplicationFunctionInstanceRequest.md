# ApplicationFunctionInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Args** | Pointer to **interface{}** |  | [optional] 
**AzionForm** | Pointer to [**ApplicationFunctionInstanceAzionForm**](ApplicationFunctionInstanceAzionForm.md) |  | [optional] 
**Function** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationFunctionInstanceRequest

`func NewApplicationFunctionInstanceRequest(name string, function int64, ) *ApplicationFunctionInstanceRequest`

NewApplicationFunctionInstanceRequest instantiates a new ApplicationFunctionInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFunctionInstanceRequestWithDefaults

`func NewApplicationFunctionInstanceRequestWithDefaults() *ApplicationFunctionInstanceRequest`

NewApplicationFunctionInstanceRequestWithDefaults instantiates a new ApplicationFunctionInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationFunctionInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationFunctionInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationFunctionInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *ApplicationFunctionInstanceRequest) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplicationFunctionInstanceRequest) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplicationFunctionInstanceRequest) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ApplicationFunctionInstanceRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *ApplicationFunctionInstanceRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *ApplicationFunctionInstanceRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetAzionForm

`func (o *ApplicationFunctionInstanceRequest) GetAzionForm() ApplicationFunctionInstanceAzionForm`

GetAzionForm returns the AzionForm field if non-nil, zero value otherwise.

### GetAzionFormOk

`func (o *ApplicationFunctionInstanceRequest) GetAzionFormOk() (*ApplicationFunctionInstanceAzionForm, bool)`

GetAzionFormOk returns a tuple with the AzionForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionForm

`func (o *ApplicationFunctionInstanceRequest) SetAzionForm(v ApplicationFunctionInstanceAzionForm)`

SetAzionForm sets AzionForm field to given value.

### HasAzionForm

`func (o *ApplicationFunctionInstanceRequest) HasAzionForm() bool`

HasAzionForm returns a boolean if a field has been set.

### GetFunction

`func (o *ApplicationFunctionInstanceRequest) GetFunction() int64`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ApplicationFunctionInstanceRequest) GetFunctionOk() (*int64, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ApplicationFunctionInstanceRequest) SetFunction(v int64)`

SetFunction sets Function field to given value.


### GetActive

`func (o *ApplicationFunctionInstanceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationFunctionInstanceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationFunctionInstanceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationFunctionInstanceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


