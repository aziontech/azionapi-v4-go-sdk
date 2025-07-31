# ResponseAsyncDNSSEC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**DNSSEC**](DNSSEC.md) |  | 

## Methods

### NewResponseAsyncDNSSEC

`func NewResponseAsyncDNSSEC(data DNSSEC, ) *ResponseAsyncDNSSEC`

NewResponseAsyncDNSSEC instantiates a new ResponseAsyncDNSSEC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncDNSSECWithDefaults

`func NewResponseAsyncDNSSECWithDefaults() *ResponseAsyncDNSSEC`

NewResponseAsyncDNSSECWithDefaults instantiates a new ResponseAsyncDNSSEC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncDNSSEC) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncDNSSEC) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncDNSSEC) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncDNSSEC) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncDNSSEC) GetData() DNSSEC`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncDNSSEC) GetDataOk() (*DNSSEC, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncDNSSEC) SetData(v DNSSEC)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


