# ResponseDeleteDataSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableDataSet**](DataSet.md) |  | 

## Methods

### NewResponseDeleteDataSet

`func NewResponseDeleteDataSet(state string, data NullableDataSet, ) *ResponseDeleteDataSet`

NewResponseDeleteDataSet instantiates a new ResponseDeleteDataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteDataSetWithDefaults

`func NewResponseDeleteDataSetWithDefaults() *ResponseDeleteDataSet`

NewResponseDeleteDataSetWithDefaults instantiates a new ResponseDeleteDataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteDataSet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteDataSet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteDataSet) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteDataSet) GetData() DataSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteDataSet) GetDataOk() (*DataSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteDataSet) SetData(v DataSet)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteDataSet) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteDataSet) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


