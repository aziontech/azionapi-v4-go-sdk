# ResponseBadRequestBaseEdgeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Modules** | Pointer to [**ResponseBadRequestSerializerMetaclassModulesField**](ResponseBadRequestSerializerMetaclassModulesField.md) |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**ProductVersion** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 
**Addresses** | Pointer to **[]string** |  | [optional] 
**Tls** | Pointer to [**ResponseBadRequestSerializerMetaclassTlsField**](ResponseBadRequestSerializerMetaclassTlsField.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **[]string** |  | [optional] 
**ConnectionPreference** | Pointer to [**ResponseBadRequestBaseEdgeConnectorConnectionPreference**](ResponseBadRequestBaseEdgeConnectorConnectionPreference.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **[]string** |  | [optional] 
**ReadWriteTimeout** | Pointer to **[]string** |  | [optional] 
**MaxRetries** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestBaseEdgeConnector

`func NewResponseBadRequestBaseEdgeConnector() *ResponseBadRequestBaseEdgeConnector`

NewResponseBadRequestBaseEdgeConnector instantiates a new ResponseBadRequestBaseEdgeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestBaseEdgeConnectorWithDefaults

`func NewResponseBadRequestBaseEdgeConnectorWithDefaults() *ResponseBadRequestBaseEdgeConnector`

NewResponseBadRequestBaseEdgeConnectorWithDefaults instantiates a new ResponseBadRequestBaseEdgeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseBadRequestBaseEdgeConnector) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseBadRequestBaseEdgeConnector) SetId(v []string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponseBadRequestBaseEdgeConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResponseBadRequestBaseEdgeConnector) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestBaseEdgeConnector) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestBaseEdgeConnector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestBaseEdgeConnector) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestBaseEdgeConnector) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestBaseEdgeConnector) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestBaseEdgeConnector) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestBaseEdgeConnector) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestBaseEdgeConnector) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetModules

`func (o *ResponseBadRequestBaseEdgeConnector) GetModules() ResponseBadRequestSerializerMetaclassModulesField`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetModulesOk() (*ResponseBadRequestSerializerMetaclassModulesField, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ResponseBadRequestBaseEdgeConnector) SetModules(v ResponseBadRequestSerializerMetaclassModulesField)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ResponseBadRequestBaseEdgeConnector) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestBaseEdgeConnector) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestBaseEdgeConnector) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestBaseEdgeConnector) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *ResponseBadRequestBaseEdgeConnector) GetProductVersion() []string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetProductVersionOk() (*[]string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseBadRequestBaseEdgeConnector) SetProductVersion(v []string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ResponseBadRequestBaseEdgeConnector) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetType

`func (o *ResponseBadRequestBaseEdgeConnector) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseBadRequestBaseEdgeConnector) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseBadRequestBaseEdgeConnector) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddresses

`func (o *ResponseBadRequestBaseEdgeConnector) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ResponseBadRequestBaseEdgeConnector) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ResponseBadRequestBaseEdgeConnector) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *ResponseBadRequestBaseEdgeConnector) GetTls() ResponseBadRequestSerializerMetaclassTlsField`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetTlsOk() (*ResponseBadRequestSerializerMetaclassTlsField, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *ResponseBadRequestBaseEdgeConnector) SetTls(v ResponseBadRequestSerializerMetaclassTlsField)`

SetTls sets Tls field to given value.

### HasTls

`func (o *ResponseBadRequestBaseEdgeConnector) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *ResponseBadRequestBaseEdgeConnector) GetLoadBalanceMethod() []string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetLoadBalanceMethodOk() (*[]string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *ResponseBadRequestBaseEdgeConnector) SetLoadBalanceMethod(v []string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *ResponseBadRequestBaseEdgeConnector) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *ResponseBadRequestBaseEdgeConnector) GetConnectionPreference() ResponseBadRequestBaseEdgeConnectorConnectionPreference`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetConnectionPreferenceOk() (*ResponseBadRequestBaseEdgeConnectorConnectionPreference, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *ResponseBadRequestBaseEdgeConnector) SetConnectionPreference(v ResponseBadRequestBaseEdgeConnectorConnectionPreference)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *ResponseBadRequestBaseEdgeConnector) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *ResponseBadRequestBaseEdgeConnector) GetConnectionTimeout() []string`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetConnectionTimeoutOk() (*[]string, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *ResponseBadRequestBaseEdgeConnector) SetConnectionTimeout(v []string)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *ResponseBadRequestBaseEdgeConnector) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *ResponseBadRequestBaseEdgeConnector) GetReadWriteTimeout() []string`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetReadWriteTimeoutOk() (*[]string, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *ResponseBadRequestBaseEdgeConnector) SetReadWriteTimeout(v []string)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *ResponseBadRequestBaseEdgeConnector) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *ResponseBadRequestBaseEdgeConnector) GetMaxRetries() []string`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetMaxRetriesOk() (*[]string, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *ResponseBadRequestBaseEdgeConnector) SetMaxRetries(v []string)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *ResponseBadRequestBaseEdgeConnector) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestBaseEdgeConnector) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestBaseEdgeConnector) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestBaseEdgeConnector) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestBaseEdgeConnector) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


