# ResponseTOTPDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Data** | [**TOTPDevice**](TOTPDevice.md) |  | 

## Methods

### NewResponseTOTPDevice

`func NewResponseTOTPDevice(state string, data TOTPDevice, ) *ResponseTOTPDevice`

NewResponseTOTPDevice instantiates a new ResponseTOTPDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTOTPDeviceWithDefaults

`func NewResponseTOTPDeviceWithDefaults() *ResponseTOTPDevice`

NewResponseTOTPDeviceWithDefaults instantiates a new ResponseTOTPDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseTOTPDevice) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseTOTPDevice) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseTOTPDevice) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseTOTPDevice) GetData() TOTPDevice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseTOTPDevice) GetDataOk() (*TOTPDevice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseTOTPDevice) SetData(v TOTPDevice)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


