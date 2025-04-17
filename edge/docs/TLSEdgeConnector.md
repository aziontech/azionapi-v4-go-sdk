# TLSEdgeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | policy of security to access the origin Possible values: off, on, preserve. Default: preserve  * &#x60;off&#x60; - off * &#x60;on&#x60; - on * &#x60;preserve&#x60; - preserve | [optional] 

## Methods

### NewTLSEdgeConnector

`func NewTLSEdgeConnector() *TLSEdgeConnector`

NewTLSEdgeConnector instantiates a new TLSEdgeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSEdgeConnectorWithDefaults

`func NewTLSEdgeConnectorWithDefaults() *TLSEdgeConnector`

NewTLSEdgeConnectorWithDefaults instantiates a new TLSEdgeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *TLSEdgeConnector) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *TLSEdgeConnector) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *TLSEdgeConnector) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *TLSEdgeConnector) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


