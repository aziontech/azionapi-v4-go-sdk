# WAFThreatsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossSiteScripting** | Pointer to **bool** |  | [optional] 
**CrossSiteScriptingSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**DirectoryTraversal** | Pointer to **bool** |  | [optional] 
**DirectoryTraversalSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**EvadingTricks** | Pointer to **bool** |  | [optional] 
**EvadingTricksSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**FileUpload** | Pointer to **bool** |  | [optional] 
**FileUploadSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**IdentifiedAttack** | Pointer to **bool** |  | [optional] 
**IdentifiedAttackSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**RemoteFileInclusion** | Pointer to **bool** |  | [optional] 
**RemoteFileInclusionSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**SqlInjection** | Pointer to **bool** |  | [optional] 
**SqlInjectionSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 
**UnwantedAccess** | Pointer to **bool** |  | [optional] 
**UnwantedAccessSensitivity** | Pointer to **string** | * &#x60;lowest&#x60; - lowest * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;highest&#x60; - highest | [optional] 

## Methods

### NewWAFThreatsConfiguration

`func NewWAFThreatsConfiguration() *WAFThreatsConfiguration`

NewWAFThreatsConfiguration instantiates a new WAFThreatsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFThreatsConfigurationWithDefaults

`func NewWAFThreatsConfigurationWithDefaults() *WAFThreatsConfiguration`

NewWAFThreatsConfigurationWithDefaults instantiates a new WAFThreatsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossSiteScripting

`func (o *WAFThreatsConfiguration) GetCrossSiteScripting() bool`

GetCrossSiteScripting returns the CrossSiteScripting field if non-nil, zero value otherwise.

### GetCrossSiteScriptingOk

`func (o *WAFThreatsConfiguration) GetCrossSiteScriptingOk() (*bool, bool)`

GetCrossSiteScriptingOk returns a tuple with the CrossSiteScripting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScripting

`func (o *WAFThreatsConfiguration) SetCrossSiteScripting(v bool)`

SetCrossSiteScripting sets CrossSiteScripting field to given value.

### HasCrossSiteScripting

`func (o *WAFThreatsConfiguration) HasCrossSiteScripting() bool`

HasCrossSiteScripting returns a boolean if a field has been set.

### GetCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfiguration) GetCrossSiteScriptingSensitivity() string`

GetCrossSiteScriptingSensitivity returns the CrossSiteScriptingSensitivity field if non-nil, zero value otherwise.

### GetCrossSiteScriptingSensitivityOk

`func (o *WAFThreatsConfiguration) GetCrossSiteScriptingSensitivityOk() (*string, bool)`

GetCrossSiteScriptingSensitivityOk returns a tuple with the CrossSiteScriptingSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfiguration) SetCrossSiteScriptingSensitivity(v string)`

SetCrossSiteScriptingSensitivity sets CrossSiteScriptingSensitivity field to given value.

### HasCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfiguration) HasCrossSiteScriptingSensitivity() bool`

HasCrossSiteScriptingSensitivity returns a boolean if a field has been set.

### GetDirectoryTraversal

`func (o *WAFThreatsConfiguration) GetDirectoryTraversal() bool`

GetDirectoryTraversal returns the DirectoryTraversal field if non-nil, zero value otherwise.

### GetDirectoryTraversalOk

`func (o *WAFThreatsConfiguration) GetDirectoryTraversalOk() (*bool, bool)`

GetDirectoryTraversalOk returns a tuple with the DirectoryTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversal

`func (o *WAFThreatsConfiguration) SetDirectoryTraversal(v bool)`

SetDirectoryTraversal sets DirectoryTraversal field to given value.

### HasDirectoryTraversal

`func (o *WAFThreatsConfiguration) HasDirectoryTraversal() bool`

HasDirectoryTraversal returns a boolean if a field has been set.

### GetDirectoryTraversalSensitivity

`func (o *WAFThreatsConfiguration) GetDirectoryTraversalSensitivity() string`

GetDirectoryTraversalSensitivity returns the DirectoryTraversalSensitivity field if non-nil, zero value otherwise.

### GetDirectoryTraversalSensitivityOk

`func (o *WAFThreatsConfiguration) GetDirectoryTraversalSensitivityOk() (*string, bool)`

GetDirectoryTraversalSensitivityOk returns a tuple with the DirectoryTraversalSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversalSensitivity

`func (o *WAFThreatsConfiguration) SetDirectoryTraversalSensitivity(v string)`

SetDirectoryTraversalSensitivity sets DirectoryTraversalSensitivity field to given value.

### HasDirectoryTraversalSensitivity

`func (o *WAFThreatsConfiguration) HasDirectoryTraversalSensitivity() bool`

HasDirectoryTraversalSensitivity returns a boolean if a field has been set.

### GetEvadingTricks

`func (o *WAFThreatsConfiguration) GetEvadingTricks() bool`

GetEvadingTricks returns the EvadingTricks field if non-nil, zero value otherwise.

### GetEvadingTricksOk

`func (o *WAFThreatsConfiguration) GetEvadingTricksOk() (*bool, bool)`

GetEvadingTricksOk returns a tuple with the EvadingTricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricks

`func (o *WAFThreatsConfiguration) SetEvadingTricks(v bool)`

SetEvadingTricks sets EvadingTricks field to given value.

### HasEvadingTricks

`func (o *WAFThreatsConfiguration) HasEvadingTricks() bool`

HasEvadingTricks returns a boolean if a field has been set.

### GetEvadingTricksSensitivity

`func (o *WAFThreatsConfiguration) GetEvadingTricksSensitivity() string`

GetEvadingTricksSensitivity returns the EvadingTricksSensitivity field if non-nil, zero value otherwise.

### GetEvadingTricksSensitivityOk

`func (o *WAFThreatsConfiguration) GetEvadingTricksSensitivityOk() (*string, bool)`

GetEvadingTricksSensitivityOk returns a tuple with the EvadingTricksSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricksSensitivity

`func (o *WAFThreatsConfiguration) SetEvadingTricksSensitivity(v string)`

SetEvadingTricksSensitivity sets EvadingTricksSensitivity field to given value.

### HasEvadingTricksSensitivity

`func (o *WAFThreatsConfiguration) HasEvadingTricksSensitivity() bool`

HasEvadingTricksSensitivity returns a boolean if a field has been set.

### GetFileUpload

`func (o *WAFThreatsConfiguration) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *WAFThreatsConfiguration) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *WAFThreatsConfiguration) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.

### HasFileUpload

`func (o *WAFThreatsConfiguration) HasFileUpload() bool`

HasFileUpload returns a boolean if a field has been set.

### GetFileUploadSensitivity

`func (o *WAFThreatsConfiguration) GetFileUploadSensitivity() string`

GetFileUploadSensitivity returns the FileUploadSensitivity field if non-nil, zero value otherwise.

### GetFileUploadSensitivityOk

`func (o *WAFThreatsConfiguration) GetFileUploadSensitivityOk() (*string, bool)`

GetFileUploadSensitivityOk returns a tuple with the FileUploadSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSensitivity

`func (o *WAFThreatsConfiguration) SetFileUploadSensitivity(v string)`

SetFileUploadSensitivity sets FileUploadSensitivity field to given value.

### HasFileUploadSensitivity

`func (o *WAFThreatsConfiguration) HasFileUploadSensitivity() bool`

HasFileUploadSensitivity returns a boolean if a field has been set.

### GetIdentifiedAttack

`func (o *WAFThreatsConfiguration) GetIdentifiedAttack() bool`

GetIdentifiedAttack returns the IdentifiedAttack field if non-nil, zero value otherwise.

### GetIdentifiedAttackOk

`func (o *WAFThreatsConfiguration) GetIdentifiedAttackOk() (*bool, bool)`

GetIdentifiedAttackOk returns a tuple with the IdentifiedAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttack

`func (o *WAFThreatsConfiguration) SetIdentifiedAttack(v bool)`

SetIdentifiedAttack sets IdentifiedAttack field to given value.

### HasIdentifiedAttack

`func (o *WAFThreatsConfiguration) HasIdentifiedAttack() bool`

HasIdentifiedAttack returns a boolean if a field has been set.

### GetIdentifiedAttackSensitivity

`func (o *WAFThreatsConfiguration) GetIdentifiedAttackSensitivity() string`

GetIdentifiedAttackSensitivity returns the IdentifiedAttackSensitivity field if non-nil, zero value otherwise.

### GetIdentifiedAttackSensitivityOk

`func (o *WAFThreatsConfiguration) GetIdentifiedAttackSensitivityOk() (*string, bool)`

GetIdentifiedAttackSensitivityOk returns a tuple with the IdentifiedAttackSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttackSensitivity

`func (o *WAFThreatsConfiguration) SetIdentifiedAttackSensitivity(v string)`

SetIdentifiedAttackSensitivity sets IdentifiedAttackSensitivity field to given value.

### HasIdentifiedAttackSensitivity

`func (o *WAFThreatsConfiguration) HasIdentifiedAttackSensitivity() bool`

HasIdentifiedAttackSensitivity returns a boolean if a field has been set.

### GetRemoteFileInclusion

`func (o *WAFThreatsConfiguration) GetRemoteFileInclusion() bool`

GetRemoteFileInclusion returns the RemoteFileInclusion field if non-nil, zero value otherwise.

### GetRemoteFileInclusionOk

`func (o *WAFThreatsConfiguration) GetRemoteFileInclusionOk() (*bool, bool)`

GetRemoteFileInclusionOk returns a tuple with the RemoteFileInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusion

`func (o *WAFThreatsConfiguration) SetRemoteFileInclusion(v bool)`

SetRemoteFileInclusion sets RemoteFileInclusion field to given value.

### HasRemoteFileInclusion

`func (o *WAFThreatsConfiguration) HasRemoteFileInclusion() bool`

HasRemoteFileInclusion returns a boolean if a field has been set.

### GetRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfiguration) GetRemoteFileInclusionSensitivity() string`

GetRemoteFileInclusionSensitivity returns the RemoteFileInclusionSensitivity field if non-nil, zero value otherwise.

### GetRemoteFileInclusionSensitivityOk

`func (o *WAFThreatsConfiguration) GetRemoteFileInclusionSensitivityOk() (*string, bool)`

GetRemoteFileInclusionSensitivityOk returns a tuple with the RemoteFileInclusionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfiguration) SetRemoteFileInclusionSensitivity(v string)`

SetRemoteFileInclusionSensitivity sets RemoteFileInclusionSensitivity field to given value.

### HasRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfiguration) HasRemoteFileInclusionSensitivity() bool`

HasRemoteFileInclusionSensitivity returns a boolean if a field has been set.

### GetSqlInjection

`func (o *WAFThreatsConfiguration) GetSqlInjection() bool`

GetSqlInjection returns the SqlInjection field if non-nil, zero value otherwise.

### GetSqlInjectionOk

`func (o *WAFThreatsConfiguration) GetSqlInjectionOk() (*bool, bool)`

GetSqlInjectionOk returns a tuple with the SqlInjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjection

`func (o *WAFThreatsConfiguration) SetSqlInjection(v bool)`

SetSqlInjection sets SqlInjection field to given value.

### HasSqlInjection

`func (o *WAFThreatsConfiguration) HasSqlInjection() bool`

HasSqlInjection returns a boolean if a field has been set.

### GetSqlInjectionSensitivity

`func (o *WAFThreatsConfiguration) GetSqlInjectionSensitivity() string`

GetSqlInjectionSensitivity returns the SqlInjectionSensitivity field if non-nil, zero value otherwise.

### GetSqlInjectionSensitivityOk

`func (o *WAFThreatsConfiguration) GetSqlInjectionSensitivityOk() (*string, bool)`

GetSqlInjectionSensitivityOk returns a tuple with the SqlInjectionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjectionSensitivity

`func (o *WAFThreatsConfiguration) SetSqlInjectionSensitivity(v string)`

SetSqlInjectionSensitivity sets SqlInjectionSensitivity field to given value.

### HasSqlInjectionSensitivity

`func (o *WAFThreatsConfiguration) HasSqlInjectionSensitivity() bool`

HasSqlInjectionSensitivity returns a boolean if a field has been set.

### GetUnwantedAccess

`func (o *WAFThreatsConfiguration) GetUnwantedAccess() bool`

GetUnwantedAccess returns the UnwantedAccess field if non-nil, zero value otherwise.

### GetUnwantedAccessOk

`func (o *WAFThreatsConfiguration) GetUnwantedAccessOk() (*bool, bool)`

GetUnwantedAccessOk returns a tuple with the UnwantedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccess

`func (o *WAFThreatsConfiguration) SetUnwantedAccess(v bool)`

SetUnwantedAccess sets UnwantedAccess field to given value.

### HasUnwantedAccess

`func (o *WAFThreatsConfiguration) HasUnwantedAccess() bool`

HasUnwantedAccess returns a boolean if a field has been set.

### GetUnwantedAccessSensitivity

`func (o *WAFThreatsConfiguration) GetUnwantedAccessSensitivity() string`

GetUnwantedAccessSensitivity returns the UnwantedAccessSensitivity field if non-nil, zero value otherwise.

### GetUnwantedAccessSensitivityOk

`func (o *WAFThreatsConfiguration) GetUnwantedAccessSensitivityOk() (*string, bool)`

GetUnwantedAccessSensitivityOk returns a tuple with the UnwantedAccessSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccessSensitivity

`func (o *WAFThreatsConfiguration) SetUnwantedAccessSensitivity(v string)`

SetUnwantedAccessSensitivity sets UnwantedAccessSensitivity field to given value.

### HasUnwantedAccessSensitivity

`func (o *WAFThreatsConfiguration) HasUnwantedAccessSensitivity() bool`

HasUnwantedAccessSensitivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


