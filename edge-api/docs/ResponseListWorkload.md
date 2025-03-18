# ResponseListWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AlternateDomains** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**NetworkMap** | Pointer to **map[string]interface{}** | * &#x60;1&#x60; - Edge Global Network * &#x60;2&#x60; - Staging Network | [optional] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Tls** | Pointer to [**TLS**](TLS.md) |  | [optional] 
**Protocols** | Pointer to [**Protocols**](Protocols.md) |  | [optional] 
**Mtls** | Pointer to [**MTLS**](MTLS.md) |  | [optional] 
**Domains** | Pointer to [**[]DomainInfo**](DomainInfo.md) |  | [optional] 
**ProductVersion** | **string** |  | [readonly] 

## Methods

### NewResponseListWorkload

`func NewResponseListWorkload(id int32, name string, lastEditor string, lastModified time.Time, productVersion string, ) *ResponseListWorkload`

NewResponseListWorkload instantiates a new ResponseListWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListWorkloadWithDefaults

`func NewResponseListWorkloadWithDefaults() *ResponseListWorkload`

NewResponseListWorkloadWithDefaults instantiates a new ResponseListWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListWorkload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListWorkload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListWorkload) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListWorkload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListWorkload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListWorkload) SetName(v string)`

SetName sets Name field to given value.


### GetAlternateDomains

`func (o *ResponseListWorkload) GetAlternateDomains() []string`

GetAlternateDomains returns the AlternateDomains field if non-nil, zero value otherwise.

### GetAlternateDomainsOk

`func (o *ResponseListWorkload) GetAlternateDomainsOk() (*[]string, bool)`

GetAlternateDomainsOk returns a tuple with the AlternateDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateDomains

`func (o *ResponseListWorkload) SetAlternateDomains(v []string)`

SetAlternateDomains sets AlternateDomains field to given value.

### HasAlternateDomains

`func (o *ResponseListWorkload) HasAlternateDomains() bool`

HasAlternateDomains returns a boolean if a field has been set.

### GetActive

`func (o *ResponseListWorkload) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListWorkload) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListWorkload) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListWorkload) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworkMap

`func (o *ResponseListWorkload) GetNetworkMap() map[string]interface{}`

GetNetworkMap returns the NetworkMap field if non-nil, zero value otherwise.

### GetNetworkMapOk

`func (o *ResponseListWorkload) GetNetworkMapOk() (*map[string]interface{}, bool)`

GetNetworkMapOk returns a tuple with the NetworkMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMap

`func (o *ResponseListWorkload) SetNetworkMap(v map[string]interface{})`

SetNetworkMap sets NetworkMap field to given value.

### HasNetworkMap

`func (o *ResponseListWorkload) HasNetworkMap() bool`

HasNetworkMap returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseListWorkload) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListWorkload) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListWorkload) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListWorkload) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListWorkload) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListWorkload) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetTls

`func (o *ResponseListWorkload) GetTls() TLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *ResponseListWorkload) GetTlsOk() (*TLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *ResponseListWorkload) SetTls(v TLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *ResponseListWorkload) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetProtocols

`func (o *ResponseListWorkload) GetProtocols() Protocols`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ResponseListWorkload) GetProtocolsOk() (*Protocols, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ResponseListWorkload) SetProtocols(v Protocols)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *ResponseListWorkload) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetMtls

`func (o *ResponseListWorkload) GetMtls() MTLS`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *ResponseListWorkload) GetMtlsOk() (*MTLS, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *ResponseListWorkload) SetMtls(v MTLS)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *ResponseListWorkload) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetDomains

`func (o *ResponseListWorkload) GetDomains() []DomainInfo`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ResponseListWorkload) GetDomainsOk() (*[]DomainInfo, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ResponseListWorkload) SetDomains(v []DomainInfo)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ResponseListWorkload) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseListWorkload) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseListWorkload) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseListWorkload) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


