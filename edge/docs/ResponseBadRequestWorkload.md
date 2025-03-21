# ResponseBadRequestWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeApplication** | Pointer to **[]string** |  | [optional] 
**EdgeFirewall** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**AlternateDomains** | Pointer to [**ResponseBadRequestEdgeApplicationRuleEngineBehaviors**](ResponseBadRequestEdgeApplicationRuleEngineBehaviors.md) |  | [optional] 
**NetworkMap** | Pointer to **[]string** |  | [optional] 
**Tls** | Pointer to [**ResponseBadRequestSerializerMetaclassTlsField**](ResponseBadRequestSerializerMetaclassTlsField.md) |  | [optional] 
**Protocols** | Pointer to [**ResponseBadRequestSerializerMetaclassProtocolsField**](ResponseBadRequestSerializerMetaclassProtocolsField.md) |  | [optional] 
**Mtls** | Pointer to [**ResponseBadRequestSerializerMetaclassMtlsField**](ResponseBadRequestSerializerMetaclassMtlsField.md) |  | [optional] 
**Domains** | Pointer to [**ResponseBadRequestEdgeApplicationRuleEngineBehaviors**](ResponseBadRequestEdgeApplicationRuleEngineBehaviors.md) |  | [optional] 
**ProductVersion** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestWorkload

`func NewResponseBadRequestWorkload() *ResponseBadRequestWorkload`

NewResponseBadRequestWorkload instantiates a new ResponseBadRequestWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestWorkloadWithDefaults

`func NewResponseBadRequestWorkloadWithDefaults() *ResponseBadRequestWorkload`

NewResponseBadRequestWorkloadWithDefaults instantiates a new ResponseBadRequestWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeApplication

`func (o *ResponseBadRequestWorkload) GetEdgeApplication() []string`

GetEdgeApplication returns the EdgeApplication field if non-nil, zero value otherwise.

### GetEdgeApplicationOk

`func (o *ResponseBadRequestWorkload) GetEdgeApplicationOk() (*[]string, bool)`

GetEdgeApplicationOk returns a tuple with the EdgeApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplication

`func (o *ResponseBadRequestWorkload) SetEdgeApplication(v []string)`

SetEdgeApplication sets EdgeApplication field to given value.

### HasEdgeApplication

`func (o *ResponseBadRequestWorkload) HasEdgeApplication() bool`

HasEdgeApplication returns a boolean if a field has been set.

### GetEdgeFirewall

`func (o *ResponseBadRequestWorkload) GetEdgeFirewall() []string`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *ResponseBadRequestWorkload) GetEdgeFirewallOk() (*[]string, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *ResponseBadRequestWorkload) SetEdgeFirewall(v []string)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *ResponseBadRequestWorkload) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### GetId

`func (o *ResponseBadRequestWorkload) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBadRequestWorkload) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBadRequestWorkload) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBadRequestWorkload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResponseBadRequestWorkload) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestWorkload) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestWorkload) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestWorkload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestWorkload) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestWorkload) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestWorkload) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestWorkload) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestWorkload) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestWorkload) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestWorkload) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestWorkload) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestWorkload) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestWorkload) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestWorkload) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestWorkload) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAlternateDomains

`func (o *ResponseBadRequestWorkload) GetAlternateDomains() ResponseBadRequestEdgeApplicationRuleEngineBehaviors`

GetAlternateDomains returns the AlternateDomains field if non-nil, zero value otherwise.

### GetAlternateDomainsOk

`func (o *ResponseBadRequestWorkload) GetAlternateDomainsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool)`

GetAlternateDomainsOk returns a tuple with the AlternateDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateDomains

`func (o *ResponseBadRequestWorkload) SetAlternateDomains(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors)`

SetAlternateDomains sets AlternateDomains field to given value.

### HasAlternateDomains

`func (o *ResponseBadRequestWorkload) HasAlternateDomains() bool`

HasAlternateDomains returns a boolean if a field has been set.

### GetNetworkMap

`func (o *ResponseBadRequestWorkload) GetNetworkMap() []string`

GetNetworkMap returns the NetworkMap field if non-nil, zero value otherwise.

### GetNetworkMapOk

`func (o *ResponseBadRequestWorkload) GetNetworkMapOk() (*[]string, bool)`

GetNetworkMapOk returns a tuple with the NetworkMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMap

`func (o *ResponseBadRequestWorkload) SetNetworkMap(v []string)`

SetNetworkMap sets NetworkMap field to given value.

### HasNetworkMap

`func (o *ResponseBadRequestWorkload) HasNetworkMap() bool`

HasNetworkMap returns a boolean if a field has been set.

### GetTls

`func (o *ResponseBadRequestWorkload) GetTls() ResponseBadRequestSerializerMetaclassTlsField`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *ResponseBadRequestWorkload) GetTlsOk() (*ResponseBadRequestSerializerMetaclassTlsField, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *ResponseBadRequestWorkload) SetTls(v ResponseBadRequestSerializerMetaclassTlsField)`

SetTls sets Tls field to given value.

### HasTls

`func (o *ResponseBadRequestWorkload) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetProtocols

`func (o *ResponseBadRequestWorkload) GetProtocols() ResponseBadRequestSerializerMetaclassProtocolsField`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ResponseBadRequestWorkload) GetProtocolsOk() (*ResponseBadRequestSerializerMetaclassProtocolsField, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ResponseBadRequestWorkload) SetProtocols(v ResponseBadRequestSerializerMetaclassProtocolsField)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *ResponseBadRequestWorkload) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetMtls

`func (o *ResponseBadRequestWorkload) GetMtls() ResponseBadRequestSerializerMetaclassMtlsField`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *ResponseBadRequestWorkload) GetMtlsOk() (*ResponseBadRequestSerializerMetaclassMtlsField, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *ResponseBadRequestWorkload) SetMtls(v ResponseBadRequestSerializerMetaclassMtlsField)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *ResponseBadRequestWorkload) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetDomains

`func (o *ResponseBadRequestWorkload) GetDomains() ResponseBadRequestEdgeApplicationRuleEngineBehaviors`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ResponseBadRequestWorkload) GetDomainsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ResponseBadRequestWorkload) SetDomains(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ResponseBadRequestWorkload) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseBadRequestWorkload) GetProductVersion() []string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseBadRequestWorkload) GetProductVersionOk() (*[]string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseBadRequestWorkload) SetProductVersion(v []string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ResponseBadRequestWorkload) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestWorkload) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestWorkload) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestWorkload) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestWorkload) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


