# WorkloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AlternateDomains** | Pointer to **[]string** |  | [optional] [default to []]
**EdgeApplication** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**NetworkMap** | Pointer to **string** | * &#x60;1&#x60; - Edge Global Network * &#x60;2&#x60; - Staging Network | [optional] [default to "1"]
**EdgeFirewall** | Pointer to **NullableInt64** |  | [optional] 
**Tls** | Pointer to [**TLSRequest**](TLSRequest.md) |  | [optional] [default to {"certificate":null,"ciphers":null,"minimum_version":"tls_1_2"}]
**Protocols** | Pointer to [**ProtocolsRequest**](ProtocolsRequest.md) |  | [optional] [default to {"http":{"versions":["http1","http2"],"http_ports":[80],"https_ports":[443],"quic_ports":null}}]
**Mtls** | Pointer to [**MTLSRequest**](MTLSRequest.md) |  | [optional] 
**Domains** | Pointer to [**[]DomainInfoRequest**](DomainInfoRequest.md) |  | [optional] 

## Methods

### NewWorkloadRequest

`func NewWorkloadRequest(name string, edgeApplication int64, ) *WorkloadRequest`

NewWorkloadRequest instantiates a new WorkloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadRequestWithDefaults

`func NewWorkloadRequestWithDefaults() *WorkloadRequest`

NewWorkloadRequestWithDefaults instantiates a new WorkloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkloadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAlternateDomains

`func (o *WorkloadRequest) GetAlternateDomains() []string`

GetAlternateDomains returns the AlternateDomains field if non-nil, zero value otherwise.

### GetAlternateDomainsOk

`func (o *WorkloadRequest) GetAlternateDomainsOk() (*[]string, bool)`

GetAlternateDomainsOk returns a tuple with the AlternateDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateDomains

`func (o *WorkloadRequest) SetAlternateDomains(v []string)`

SetAlternateDomains sets AlternateDomains field to given value.

### HasAlternateDomains

`func (o *WorkloadRequest) HasAlternateDomains() bool`

HasAlternateDomains returns a boolean if a field has been set.

### GetEdgeApplication

`func (o *WorkloadRequest) GetEdgeApplication() int64`

GetEdgeApplication returns the EdgeApplication field if non-nil, zero value otherwise.

### GetEdgeApplicationOk

`func (o *WorkloadRequest) GetEdgeApplicationOk() (*int64, bool)`

GetEdgeApplicationOk returns a tuple with the EdgeApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplication

`func (o *WorkloadRequest) SetEdgeApplication(v int64)`

SetEdgeApplication sets EdgeApplication field to given value.


### GetActive

`func (o *WorkloadRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkloadRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkloadRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WorkloadRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworkMap

`func (o *WorkloadRequest) GetNetworkMap() string`

GetNetworkMap returns the NetworkMap field if non-nil, zero value otherwise.

### GetNetworkMapOk

`func (o *WorkloadRequest) GetNetworkMapOk() (*string, bool)`

GetNetworkMapOk returns a tuple with the NetworkMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMap

`func (o *WorkloadRequest) SetNetworkMap(v string)`

SetNetworkMap sets NetworkMap field to given value.

### HasNetworkMap

`func (o *WorkloadRequest) HasNetworkMap() bool`

HasNetworkMap returns a boolean if a field has been set.

### GetEdgeFirewall

`func (o *WorkloadRequest) GetEdgeFirewall() int64`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *WorkloadRequest) GetEdgeFirewallOk() (*int64, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *WorkloadRequest) SetEdgeFirewall(v int64)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *WorkloadRequest) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### SetEdgeFirewallNil

`func (o *WorkloadRequest) SetEdgeFirewallNil(b bool)`

 SetEdgeFirewallNil sets the value for EdgeFirewall to be an explicit nil

### UnsetEdgeFirewall
`func (o *WorkloadRequest) UnsetEdgeFirewall()`

UnsetEdgeFirewall ensures that no value is present for EdgeFirewall, not even an explicit nil
### GetTls

`func (o *WorkloadRequest) GetTls() TLSRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *WorkloadRequest) GetTlsOk() (*TLSRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *WorkloadRequest) SetTls(v TLSRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *WorkloadRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetProtocols

`func (o *WorkloadRequest) GetProtocols() ProtocolsRequest`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *WorkloadRequest) GetProtocolsOk() (*ProtocolsRequest, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *WorkloadRequest) SetProtocols(v ProtocolsRequest)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *WorkloadRequest) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetMtls

`func (o *WorkloadRequest) GetMtls() MTLSRequest`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *WorkloadRequest) GetMtlsOk() (*MTLSRequest, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *WorkloadRequest) SetMtls(v MTLSRequest)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *WorkloadRequest) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetDomains

`func (o *WorkloadRequest) GetDomains() []DomainInfoRequest`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *WorkloadRequest) GetDomainsOk() (*[]DomainInfoRequest, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *WorkloadRequest) SetDomains(v []DomainInfoRequest)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *WorkloadRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


