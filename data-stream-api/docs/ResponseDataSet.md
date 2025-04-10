# ResponseDataSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**DataSet**](DataSet.md) |  | 

## Methods

### NewResponseDataSet

`func NewResponseDataSet(state string, data DataSet, ) *ResponseDataSet`

NewResponseDataSet instantiates a new ResponseDataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDataSetWithDefaults

`func NewResponseDataSetWithDefaults() *ResponseDataSet`

NewResponseDataSetWithDefaults instantiates a new ResponseDataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDataSet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDataSet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDataSet) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDataSet) GetData() DataSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDataSet) GetDataOk() (*DataSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDataSet) SetData(v DataSet)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


