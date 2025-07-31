# ResponseTOTPDeviceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**TOTPDeviceCreate**](TOTPDeviceCreate.md) |  | 

## Methods

### NewResponseTOTPDeviceCreate

`func NewResponseTOTPDeviceCreate(data TOTPDeviceCreate, ) *ResponseTOTPDeviceCreate`

NewResponseTOTPDeviceCreate instantiates a new ResponseTOTPDeviceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTOTPDeviceCreateWithDefaults

`func NewResponseTOTPDeviceCreateWithDefaults() *ResponseTOTPDeviceCreate`

NewResponseTOTPDeviceCreateWithDefaults instantiates a new ResponseTOTPDeviceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseTOTPDeviceCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseTOTPDeviceCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseTOTPDeviceCreate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseTOTPDeviceCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseTOTPDeviceCreate) GetData() TOTPDeviceCreate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseTOTPDeviceCreate) GetDataOk() (*TOTPDeviceCreate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseTOTPDeviceCreate) SetData(v TOTPDeviceCreate)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


