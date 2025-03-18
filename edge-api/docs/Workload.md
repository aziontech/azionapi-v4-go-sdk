# Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AlternateDomains** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**NetworkMap** | Pointer to **string** |  | [optional] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Tls** | Pointer to [**TLS**](TLS.md) |  | [optional] 
**Protocols** | Pointer to [**Protocols**](Protocols.md) |  | [optional] 
**Mtls** | Pointer to [**MTLS**](MTLS.md) |  | [optional] 
**Domains** | Pointer to [**[]DomainInfo**](DomainInfo.md) |  | [optional] 
**ProductVersion** | **string** |  | [readonly] 

## Methods

### NewWorkload

`func NewWorkload(id int32, name string, lastEditor string, lastModified time.Time, productVersion string, ) *Workload`

NewWorkload instantiates a new Workload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWithDefaults

`func NewWorkloadWithDefaults() *Workload`

NewWorkloadWithDefaults instantiates a new Workload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workload) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Workload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workload) SetName(v string)`

SetName sets Name field to given value.


### GetAlternateDomains

`func (o *Workload) GetAlternateDomains() []string`

GetAlternateDomains returns the AlternateDomains field if non-nil, zero value otherwise.

### GetAlternateDomainsOk

`func (o *Workload) GetAlternateDomainsOk() (*[]string, bool)`

GetAlternateDomainsOk returns a tuple with the AlternateDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateDomains

`func (o *Workload) SetAlternateDomains(v []string)`

SetAlternateDomains sets AlternateDomains field to given value.

### HasAlternateDomains

`func (o *Workload) HasAlternateDomains() bool`

HasAlternateDomains returns a boolean if a field has been set.

### GetActive

`func (o *Workload) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Workload) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Workload) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Workload) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworkMap

`func (o *Workload) GetNetworkMap() string`

GetNetworkMap returns the NetworkMap field if non-nil, zero value otherwise.

### GetNetworkMapOk

`func (o *Workload) GetNetworkMapOk() (*string, bool)`

GetNetworkMapOk returns a tuple with the NetworkMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMap

`func (o *Workload) SetNetworkMap(v string)`

SetNetworkMap sets NetworkMap field to given value.

### HasNetworkMap

`func (o *Workload) HasNetworkMap() bool`

HasNetworkMap returns a boolean if a field has been set.

### GetLastEditor

`func (o *Workload) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *Workload) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *Workload) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *Workload) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Workload) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Workload) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetTls

`func (o *Workload) GetTls() TLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *Workload) GetTlsOk() (*TLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *Workload) SetTls(v TLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *Workload) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetProtocols

`func (o *Workload) GetProtocols() Protocols`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *Workload) GetProtocolsOk() (*Protocols, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *Workload) SetProtocols(v Protocols)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *Workload) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetMtls

`func (o *Workload) GetMtls() MTLS`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *Workload) GetMtlsOk() (*MTLS, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *Workload) SetMtls(v MTLS)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *Workload) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetDomains

`func (o *Workload) GetDomains() []DomainInfo`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Workload) GetDomainsOk() (*[]DomainInfo, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Workload) SetDomains(v []DomainInfo)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Workload) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetProductVersion

`func (o *Workload) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *Workload) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *Workload) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


