# DataStreamFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamplingEnable** | Pointer to **bool** |  | [optional] 
**SamplingRate** | Pointer to **int64** |  | [optional] 
**Workloads** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewDataStreamFilter

`func NewDataStreamFilter() *DataStreamFilter`

NewDataStreamFilter instantiates a new DataStreamFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamFilterWithDefaults

`func NewDataStreamFilterWithDefaults() *DataStreamFilter`

NewDataStreamFilterWithDefaults instantiates a new DataStreamFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamplingEnable

`func (o *DataStreamFilter) GetSamplingEnable() bool`

GetSamplingEnable returns the SamplingEnable field if non-nil, zero value otherwise.

### GetSamplingEnableOk

`func (o *DataStreamFilter) GetSamplingEnableOk() (*bool, bool)`

GetSamplingEnableOk returns a tuple with the SamplingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingEnable

`func (o *DataStreamFilter) SetSamplingEnable(v bool)`

SetSamplingEnable sets SamplingEnable field to given value.

### HasSamplingEnable

`func (o *DataStreamFilter) HasSamplingEnable() bool`

HasSamplingEnable returns a boolean if a field has been set.

### GetSamplingRate

`func (o *DataStreamFilter) GetSamplingRate() int64`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *DataStreamFilter) GetSamplingRateOk() (*int64, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *DataStreamFilter) SetSamplingRate(v int64)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *DataStreamFilter) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetWorkloads

`func (o *DataStreamFilter) GetWorkloads() []int64`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *DataStreamFilter) GetWorkloadsOk() (*[]int64, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *DataStreamFilter) SetWorkloads(v []int64)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *DataStreamFilter) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


