# DataStreamFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamplingEnable** | Pointer to **bool** |  | [optional] 
**SamplingRate** | Pointer to **int64** |  | [optional] 
**Workloads** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewDataStreamFilterRequest

`func NewDataStreamFilterRequest() *DataStreamFilterRequest`

NewDataStreamFilterRequest instantiates a new DataStreamFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamFilterRequestWithDefaults

`func NewDataStreamFilterRequestWithDefaults() *DataStreamFilterRequest`

NewDataStreamFilterRequestWithDefaults instantiates a new DataStreamFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamplingEnable

`func (o *DataStreamFilterRequest) GetSamplingEnable() bool`

GetSamplingEnable returns the SamplingEnable field if non-nil, zero value otherwise.

### GetSamplingEnableOk

`func (o *DataStreamFilterRequest) GetSamplingEnableOk() (*bool, bool)`

GetSamplingEnableOk returns a tuple with the SamplingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingEnable

`func (o *DataStreamFilterRequest) SetSamplingEnable(v bool)`

SetSamplingEnable sets SamplingEnable field to given value.

### HasSamplingEnable

`func (o *DataStreamFilterRequest) HasSamplingEnable() bool`

HasSamplingEnable returns a boolean if a field has been set.

### GetSamplingRate

`func (o *DataStreamFilterRequest) GetSamplingRate() int64`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *DataStreamFilterRequest) GetSamplingRateOk() (*int64, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *DataStreamFilterRequest) SetSamplingRate(v int64)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *DataStreamFilterRequest) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetWorkloads

`func (o *DataStreamFilterRequest) GetWorkloads() []int64`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *DataStreamFilterRequest) GetWorkloadsOk() (*[]int64, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *DataStreamFilterRequest) SetWorkloads(v []int64)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *DataStreamFilterRequest) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


