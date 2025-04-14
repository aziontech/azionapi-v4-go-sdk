# WorkloadDeploymentBindsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeApplication** | **int64** |  | 
**EdgeFirewall** | Pointer to **NullableInt64** |  | [optional] 
**CustomPage** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewWorkloadDeploymentBindsRequest

`func NewWorkloadDeploymentBindsRequest(edgeApplication int64, ) *WorkloadDeploymentBindsRequest`

NewWorkloadDeploymentBindsRequest instantiates a new WorkloadDeploymentBindsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeploymentBindsRequestWithDefaults

`func NewWorkloadDeploymentBindsRequestWithDefaults() *WorkloadDeploymentBindsRequest`

NewWorkloadDeploymentBindsRequestWithDefaults instantiates a new WorkloadDeploymentBindsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeApplication

`func (o *WorkloadDeploymentBindsRequest) GetEdgeApplication() int64`

GetEdgeApplication returns the EdgeApplication field if non-nil, zero value otherwise.

### GetEdgeApplicationOk

`func (o *WorkloadDeploymentBindsRequest) GetEdgeApplicationOk() (*int64, bool)`

GetEdgeApplicationOk returns a tuple with the EdgeApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplication

`func (o *WorkloadDeploymentBindsRequest) SetEdgeApplication(v int64)`

SetEdgeApplication sets EdgeApplication field to given value.


### GetEdgeFirewall

`func (o *WorkloadDeploymentBindsRequest) GetEdgeFirewall() int64`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *WorkloadDeploymentBindsRequest) GetEdgeFirewallOk() (*int64, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *WorkloadDeploymentBindsRequest) SetEdgeFirewall(v int64)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *WorkloadDeploymentBindsRequest) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### SetEdgeFirewallNil

`func (o *WorkloadDeploymentBindsRequest) SetEdgeFirewallNil(b bool)`

 SetEdgeFirewallNil sets the value for EdgeFirewall to be an explicit nil

### UnsetEdgeFirewall
`func (o *WorkloadDeploymentBindsRequest) UnsetEdgeFirewall()`

UnsetEdgeFirewall ensures that no value is present for EdgeFirewall, not even an explicit nil
### GetCustomPage

`func (o *WorkloadDeploymentBindsRequest) GetCustomPage() int64`

GetCustomPage returns the CustomPage field if non-nil, zero value otherwise.

### GetCustomPageOk

`func (o *WorkloadDeploymentBindsRequest) GetCustomPageOk() (*int64, bool)`

GetCustomPageOk returns a tuple with the CustomPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPage

`func (o *WorkloadDeploymentBindsRequest) SetCustomPage(v int64)`

SetCustomPage sets CustomPage field to given value.

### HasCustomPage

`func (o *WorkloadDeploymentBindsRequest) HasCustomPage() bool`

HasCustomPage returns a boolean if a field has been set.

### SetCustomPageNil

`func (o *WorkloadDeploymentBindsRequest) SetCustomPageNil(b bool)`

 SetCustomPageNil sets the value for CustomPage to be an explicit nil

### UnsetCustomPage
`func (o *WorkloadDeploymentBindsRequest) UnsetCustomPage()`

UnsetCustomPage ensures that no value is present for CustomPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


