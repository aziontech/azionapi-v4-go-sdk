# ResponseAsyncTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**Template**](Template.md) |  | 

## Methods

### NewResponseAsyncTemplate

`func NewResponseAsyncTemplate(data Template, ) *ResponseAsyncTemplate`

NewResponseAsyncTemplate instantiates a new ResponseAsyncTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncTemplateWithDefaults

`func NewResponseAsyncTemplateWithDefaults() *ResponseAsyncTemplate`

NewResponseAsyncTemplateWithDefaults instantiates a new ResponseAsyncTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncTemplate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncTemplate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncTemplate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncTemplate) GetData() Template`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncTemplate) GetDataOk() (*Template, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncTemplate) SetData(v Template)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


