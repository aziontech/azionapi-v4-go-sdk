# ResponseApplicationDeviceGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**ApplicationDeviceGroups**](ApplicationDeviceGroups.md) |  | 

## Methods

### NewResponseApplicationDeviceGroups

`func NewResponseApplicationDeviceGroups(data ApplicationDeviceGroups, ) *ResponseApplicationDeviceGroups`

NewResponseApplicationDeviceGroups instantiates a new ResponseApplicationDeviceGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseApplicationDeviceGroupsWithDefaults

`func NewResponseApplicationDeviceGroupsWithDefaults() *ResponseApplicationDeviceGroups`

NewResponseApplicationDeviceGroupsWithDefaults instantiates a new ResponseApplicationDeviceGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseApplicationDeviceGroups) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseApplicationDeviceGroups) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseApplicationDeviceGroups) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseApplicationDeviceGroups) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseApplicationDeviceGroups) GetData() ApplicationDeviceGroups`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseApplicationDeviceGroups) GetDataOk() (*ApplicationDeviceGroups, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseApplicationDeviceGroups) SetData(v ApplicationDeviceGroups)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


