# TLSEdgeConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | policy of security to access the origin Possible values: off, on, preserve. Default: preserve  * &#x60;off&#x60; - off * &#x60;on&#x60; - on * &#x60;preserve&#x60; - preserve | [optional] 

## Methods

### NewTLSEdgeConnectorRequest

`func NewTLSEdgeConnectorRequest() *TLSEdgeConnectorRequest`

NewTLSEdgeConnectorRequest instantiates a new TLSEdgeConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSEdgeConnectorRequestWithDefaults

`func NewTLSEdgeConnectorRequestWithDefaults() *TLSEdgeConnectorRequest`

NewTLSEdgeConnectorRequestWithDefaults instantiates a new TLSEdgeConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *TLSEdgeConnectorRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *TLSEdgeConnectorRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *TLSEdgeConnectorRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *TLSEdgeConnectorRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


