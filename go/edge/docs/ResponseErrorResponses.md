# ResponseErrorResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**ErrorResponses**](ErrorResponses.md) |  | 

## Methods

### NewResponseErrorResponses

`func NewResponseErrorResponses(state string, data ErrorResponses, ) *ResponseErrorResponses`

NewResponseErrorResponses instantiates a new ResponseErrorResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseErrorResponsesWithDefaults

`func NewResponseErrorResponsesWithDefaults() *ResponseErrorResponses`

NewResponseErrorResponsesWithDefaults instantiates a new ResponseErrorResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseErrorResponses) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseErrorResponses) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseErrorResponses) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseErrorResponses) GetData() ErrorResponses`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseErrorResponses) GetDataOk() (*ErrorResponses, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseErrorResponses) SetData(v ErrorResponses)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


