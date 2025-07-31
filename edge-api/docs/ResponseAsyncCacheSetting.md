# ResponseAsyncCacheSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**CacheSetting**](CacheSetting.md) |  | 

## Methods

### NewResponseAsyncCacheSetting

`func NewResponseAsyncCacheSetting(data CacheSetting, ) *ResponseAsyncCacheSetting`

NewResponseAsyncCacheSetting instantiates a new ResponseAsyncCacheSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncCacheSettingWithDefaults

`func NewResponseAsyncCacheSettingWithDefaults() *ResponseAsyncCacheSetting`

NewResponseAsyncCacheSettingWithDefaults instantiates a new ResponseAsyncCacheSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncCacheSetting) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncCacheSetting) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncCacheSetting) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncCacheSetting) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncCacheSetting) GetData() CacheSetting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncCacheSetting) GetDataOk() (*CacheSetting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncCacheSetting) SetData(v CacheSetting)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


