# WAFThreatsConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossSiteScripting** | Pointer to **bool** |  | [optional] [default to true]
**CrossSiteScriptingSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**DirectoryTraversal** | Pointer to **bool** |  | [optional] [default to true]
**DirectoryTraversalSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**EvadingTricks** | Pointer to **bool** |  | [optional] [default to true]
**EvadingTricksSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**FileUpload** | Pointer to **bool** |  | [optional] [default to true]
**FileUploadSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**IdentifiedAttack** | Pointer to **bool** |  | [optional] [default to true]
**IdentifiedAttackSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**RemoteFileInclusion** | Pointer to **bool** |  | [optional] [default to true]
**RemoteFileInclusionSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**SqlInjection** | Pointer to **bool** |  | [optional] [default to true]
**SqlInjectionSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]
**UnwantedAccess** | Pointer to **bool** |  | [optional] [default to true]
**UnwantedAccessSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] [default to "medium"]

## Methods

### NewWAFThreatsConfigurationRequest

`func NewWAFThreatsConfigurationRequest() *WAFThreatsConfigurationRequest`

NewWAFThreatsConfigurationRequest instantiates a new WAFThreatsConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFThreatsConfigurationRequestWithDefaults

`func NewWAFThreatsConfigurationRequestWithDefaults() *WAFThreatsConfigurationRequest`

NewWAFThreatsConfigurationRequestWithDefaults instantiates a new WAFThreatsConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossSiteScripting

`func (o *WAFThreatsConfigurationRequest) GetCrossSiteScripting() bool`

GetCrossSiteScripting returns the CrossSiteScripting field if non-nil, zero value otherwise.

### GetCrossSiteScriptingOk

`func (o *WAFThreatsConfigurationRequest) GetCrossSiteScriptingOk() (*bool, bool)`

GetCrossSiteScriptingOk returns a tuple with the CrossSiteScripting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScripting

`func (o *WAFThreatsConfigurationRequest) SetCrossSiteScripting(v bool)`

SetCrossSiteScripting sets CrossSiteScripting field to given value.

### HasCrossSiteScripting

`func (o *WAFThreatsConfigurationRequest) HasCrossSiteScripting() bool`

HasCrossSiteScripting returns a boolean if a field has been set.

### GetCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfigurationRequest) GetCrossSiteScriptingSensitivity() string`

GetCrossSiteScriptingSensitivity returns the CrossSiteScriptingSensitivity field if non-nil, zero value otherwise.

### GetCrossSiteScriptingSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetCrossSiteScriptingSensitivityOk() (*string, bool)`

GetCrossSiteScriptingSensitivityOk returns a tuple with the CrossSiteScriptingSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfigurationRequest) SetCrossSiteScriptingSensitivity(v string)`

SetCrossSiteScriptingSensitivity sets CrossSiteScriptingSensitivity field to given value.

### HasCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfigurationRequest) HasCrossSiteScriptingSensitivity() bool`

HasCrossSiteScriptingSensitivity returns a boolean if a field has been set.

### GetDirectoryTraversal

`func (o *WAFThreatsConfigurationRequest) GetDirectoryTraversal() bool`

GetDirectoryTraversal returns the DirectoryTraversal field if non-nil, zero value otherwise.

### GetDirectoryTraversalOk

`func (o *WAFThreatsConfigurationRequest) GetDirectoryTraversalOk() (*bool, bool)`

GetDirectoryTraversalOk returns a tuple with the DirectoryTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversal

`func (o *WAFThreatsConfigurationRequest) SetDirectoryTraversal(v bool)`

SetDirectoryTraversal sets DirectoryTraversal field to given value.

### HasDirectoryTraversal

`func (o *WAFThreatsConfigurationRequest) HasDirectoryTraversal() bool`

HasDirectoryTraversal returns a boolean if a field has been set.

### GetDirectoryTraversalSensitivity

`func (o *WAFThreatsConfigurationRequest) GetDirectoryTraversalSensitivity() string`

GetDirectoryTraversalSensitivity returns the DirectoryTraversalSensitivity field if non-nil, zero value otherwise.

### GetDirectoryTraversalSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetDirectoryTraversalSensitivityOk() (*string, bool)`

GetDirectoryTraversalSensitivityOk returns a tuple with the DirectoryTraversalSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversalSensitivity

`func (o *WAFThreatsConfigurationRequest) SetDirectoryTraversalSensitivity(v string)`

SetDirectoryTraversalSensitivity sets DirectoryTraversalSensitivity field to given value.

### HasDirectoryTraversalSensitivity

`func (o *WAFThreatsConfigurationRequest) HasDirectoryTraversalSensitivity() bool`

HasDirectoryTraversalSensitivity returns a boolean if a field has been set.

### GetEvadingTricks

`func (o *WAFThreatsConfigurationRequest) GetEvadingTricks() bool`

GetEvadingTricks returns the EvadingTricks field if non-nil, zero value otherwise.

### GetEvadingTricksOk

`func (o *WAFThreatsConfigurationRequest) GetEvadingTricksOk() (*bool, bool)`

GetEvadingTricksOk returns a tuple with the EvadingTricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricks

`func (o *WAFThreatsConfigurationRequest) SetEvadingTricks(v bool)`

SetEvadingTricks sets EvadingTricks field to given value.

### HasEvadingTricks

`func (o *WAFThreatsConfigurationRequest) HasEvadingTricks() bool`

HasEvadingTricks returns a boolean if a field has been set.

### GetEvadingTricksSensitivity

`func (o *WAFThreatsConfigurationRequest) GetEvadingTricksSensitivity() string`

GetEvadingTricksSensitivity returns the EvadingTricksSensitivity field if non-nil, zero value otherwise.

### GetEvadingTricksSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetEvadingTricksSensitivityOk() (*string, bool)`

GetEvadingTricksSensitivityOk returns a tuple with the EvadingTricksSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricksSensitivity

`func (o *WAFThreatsConfigurationRequest) SetEvadingTricksSensitivity(v string)`

SetEvadingTricksSensitivity sets EvadingTricksSensitivity field to given value.

### HasEvadingTricksSensitivity

`func (o *WAFThreatsConfigurationRequest) HasEvadingTricksSensitivity() bool`

HasEvadingTricksSensitivity returns a boolean if a field has been set.

### GetFileUpload

`func (o *WAFThreatsConfigurationRequest) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *WAFThreatsConfigurationRequest) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *WAFThreatsConfigurationRequest) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.

### HasFileUpload

`func (o *WAFThreatsConfigurationRequest) HasFileUpload() bool`

HasFileUpload returns a boolean if a field has been set.

### GetFileUploadSensitivity

`func (o *WAFThreatsConfigurationRequest) GetFileUploadSensitivity() string`

GetFileUploadSensitivity returns the FileUploadSensitivity field if non-nil, zero value otherwise.

### GetFileUploadSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetFileUploadSensitivityOk() (*string, bool)`

GetFileUploadSensitivityOk returns a tuple with the FileUploadSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSensitivity

`func (o *WAFThreatsConfigurationRequest) SetFileUploadSensitivity(v string)`

SetFileUploadSensitivity sets FileUploadSensitivity field to given value.

### HasFileUploadSensitivity

`func (o *WAFThreatsConfigurationRequest) HasFileUploadSensitivity() bool`

HasFileUploadSensitivity returns a boolean if a field has been set.

### GetIdentifiedAttack

`func (o *WAFThreatsConfigurationRequest) GetIdentifiedAttack() bool`

GetIdentifiedAttack returns the IdentifiedAttack field if non-nil, zero value otherwise.

### GetIdentifiedAttackOk

`func (o *WAFThreatsConfigurationRequest) GetIdentifiedAttackOk() (*bool, bool)`

GetIdentifiedAttackOk returns a tuple with the IdentifiedAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttack

`func (o *WAFThreatsConfigurationRequest) SetIdentifiedAttack(v bool)`

SetIdentifiedAttack sets IdentifiedAttack field to given value.

### HasIdentifiedAttack

`func (o *WAFThreatsConfigurationRequest) HasIdentifiedAttack() bool`

HasIdentifiedAttack returns a boolean if a field has been set.

### GetIdentifiedAttackSensitivity

`func (o *WAFThreatsConfigurationRequest) GetIdentifiedAttackSensitivity() string`

GetIdentifiedAttackSensitivity returns the IdentifiedAttackSensitivity field if non-nil, zero value otherwise.

### GetIdentifiedAttackSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetIdentifiedAttackSensitivityOk() (*string, bool)`

GetIdentifiedAttackSensitivityOk returns a tuple with the IdentifiedAttackSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttackSensitivity

`func (o *WAFThreatsConfigurationRequest) SetIdentifiedAttackSensitivity(v string)`

SetIdentifiedAttackSensitivity sets IdentifiedAttackSensitivity field to given value.

### HasIdentifiedAttackSensitivity

`func (o *WAFThreatsConfigurationRequest) HasIdentifiedAttackSensitivity() bool`

HasIdentifiedAttackSensitivity returns a boolean if a field has been set.

### GetRemoteFileInclusion

`func (o *WAFThreatsConfigurationRequest) GetRemoteFileInclusion() bool`

GetRemoteFileInclusion returns the RemoteFileInclusion field if non-nil, zero value otherwise.

### GetRemoteFileInclusionOk

`func (o *WAFThreatsConfigurationRequest) GetRemoteFileInclusionOk() (*bool, bool)`

GetRemoteFileInclusionOk returns a tuple with the RemoteFileInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusion

`func (o *WAFThreatsConfigurationRequest) SetRemoteFileInclusion(v bool)`

SetRemoteFileInclusion sets RemoteFileInclusion field to given value.

### HasRemoteFileInclusion

`func (o *WAFThreatsConfigurationRequest) HasRemoteFileInclusion() bool`

HasRemoteFileInclusion returns a boolean if a field has been set.

### GetRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfigurationRequest) GetRemoteFileInclusionSensitivity() string`

GetRemoteFileInclusionSensitivity returns the RemoteFileInclusionSensitivity field if non-nil, zero value otherwise.

### GetRemoteFileInclusionSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetRemoteFileInclusionSensitivityOk() (*string, bool)`

GetRemoteFileInclusionSensitivityOk returns a tuple with the RemoteFileInclusionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfigurationRequest) SetRemoteFileInclusionSensitivity(v string)`

SetRemoteFileInclusionSensitivity sets RemoteFileInclusionSensitivity field to given value.

### HasRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfigurationRequest) HasRemoteFileInclusionSensitivity() bool`

HasRemoteFileInclusionSensitivity returns a boolean if a field has been set.

### GetSqlInjection

`func (o *WAFThreatsConfigurationRequest) GetSqlInjection() bool`

GetSqlInjection returns the SqlInjection field if non-nil, zero value otherwise.

### GetSqlInjectionOk

`func (o *WAFThreatsConfigurationRequest) GetSqlInjectionOk() (*bool, bool)`

GetSqlInjectionOk returns a tuple with the SqlInjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjection

`func (o *WAFThreatsConfigurationRequest) SetSqlInjection(v bool)`

SetSqlInjection sets SqlInjection field to given value.

### HasSqlInjection

`func (o *WAFThreatsConfigurationRequest) HasSqlInjection() bool`

HasSqlInjection returns a boolean if a field has been set.

### GetSqlInjectionSensitivity

`func (o *WAFThreatsConfigurationRequest) GetSqlInjectionSensitivity() string`

GetSqlInjectionSensitivity returns the SqlInjectionSensitivity field if non-nil, zero value otherwise.

### GetSqlInjectionSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetSqlInjectionSensitivityOk() (*string, bool)`

GetSqlInjectionSensitivityOk returns a tuple with the SqlInjectionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjectionSensitivity

`func (o *WAFThreatsConfigurationRequest) SetSqlInjectionSensitivity(v string)`

SetSqlInjectionSensitivity sets SqlInjectionSensitivity field to given value.

### HasSqlInjectionSensitivity

`func (o *WAFThreatsConfigurationRequest) HasSqlInjectionSensitivity() bool`

HasSqlInjectionSensitivity returns a boolean if a field has been set.

### GetUnwantedAccess

`func (o *WAFThreatsConfigurationRequest) GetUnwantedAccess() bool`

GetUnwantedAccess returns the UnwantedAccess field if non-nil, zero value otherwise.

### GetUnwantedAccessOk

`func (o *WAFThreatsConfigurationRequest) GetUnwantedAccessOk() (*bool, bool)`

GetUnwantedAccessOk returns a tuple with the UnwantedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccess

`func (o *WAFThreatsConfigurationRequest) SetUnwantedAccess(v bool)`

SetUnwantedAccess sets UnwantedAccess field to given value.

### HasUnwantedAccess

`func (o *WAFThreatsConfigurationRequest) HasUnwantedAccess() bool`

HasUnwantedAccess returns a boolean if a field has been set.

### GetUnwantedAccessSensitivity

`func (o *WAFThreatsConfigurationRequest) GetUnwantedAccessSensitivity() string`

GetUnwantedAccessSensitivity returns the UnwantedAccessSensitivity field if non-nil, zero value otherwise.

### GetUnwantedAccessSensitivityOk

`func (o *WAFThreatsConfigurationRequest) GetUnwantedAccessSensitivityOk() (*string, bool)`

GetUnwantedAccessSensitivityOk returns a tuple with the UnwantedAccessSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccessSensitivity

`func (o *WAFThreatsConfigurationRequest) SetUnwantedAccessSensitivity(v string)`

SetUnwantedAccessSensitivity sets UnwantedAccessSensitivity field to given value.

### HasUnwantedAccessSensitivity

`func (o *WAFThreatsConfigurationRequest) HasUnwantedAccessSensitivity() bool`

HasUnwantedAccessSensitivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


