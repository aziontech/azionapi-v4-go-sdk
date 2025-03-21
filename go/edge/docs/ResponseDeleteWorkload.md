# ResponseDeleteWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableWorkload**](Workload.md) |  | 

## Methods

### NewResponseDeleteWorkload

`func NewResponseDeleteWorkload(state string, data NullableWorkload, ) *ResponseDeleteWorkload`

NewResponseDeleteWorkload instantiates a new ResponseDeleteWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteWorkloadWithDefaults

`func NewResponseDeleteWorkloadWithDefaults() *ResponseDeleteWorkload`

NewResponseDeleteWorkloadWithDefaults instantiates a new ResponseDeleteWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteWorkload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteWorkload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteWorkload) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteWorkload) GetData() Workload`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteWorkload) GetDataOk() (*Workload, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteWorkload) SetData(v Workload)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteWorkload) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteWorkload) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


