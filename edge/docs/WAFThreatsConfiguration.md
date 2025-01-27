# WAFThreatsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossSiteScripting** | Pointer to **bool** |  | [optional] [default to true]
**CrossSiteScriptingSensitivity** | Pointer to [**CrossSiteScriptingSensitivityEnum**](CrossSiteScriptingSensitivityEnum.md) |  | [optional] [default to medium]
**DirectoryTraversal** | Pointer to **bool** |  | [optional] [default to true]
**DirectoryTraversalSensitivity** | Pointer to [**DirectoryTraversalSensitivityEnum**](DirectoryTraversalSensitivityEnum.md) |  | [optional] [default to medium]
**EvadingTricks** | Pointer to **bool** |  | [optional] [default to true]
**EvadingTricksSensitivity** | Pointer to [**EvadingTricksSensitivityEnum**](EvadingTricksSensitivityEnum.md) |  | [optional] [default to medium]
**FileUpload** | Pointer to **bool** |  | [optional] [default to true]
**FileUploadSensitivity** | Pointer to [**FileUploadSensitivityEnum**](FileUploadSensitivityEnum.md) |  | [optional] [default to medium]
**IdentifiedAttack** | Pointer to **bool** |  | [optional] [default to true]
**IdentifiedAttackSensitivity** | Pointer to [**IdentifiedAttackSensitivityEnum**](IdentifiedAttackSensitivityEnum.md) |  | [optional] [default to medium]
**RemoteFileInclusion** | Pointer to **bool** |  | [optional] [default to true]
**RemoteFileInclusionSensitivity** | Pointer to [**RemoteFileInclusionSensitivityEnum**](RemoteFileInclusionSensitivityEnum.md) |  | [optional] [default to medium]
**SqlInjection** | Pointer to **bool** |  | [optional] [default to true]
**SqlInjectionSensitivity** | Pointer to [**SqlInjectionSensitivityEnum**](SqlInjectionSensitivityEnum.md) |  | [optional] [default to medium]
**UnwantedAccess** | Pointer to **bool** |  | [optional] [default to true]
**UnwantedAccessSensitivity** | Pointer to [**UnwantedAccessSensitivityEnum**](UnwantedAccessSensitivityEnum.md) |  | [optional] [default to medium]

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

`func (o *WAFThreatsConfiguration) GetCrossSiteScriptingSensitivity() CrossSiteScriptingSensitivityEnum`

GetCrossSiteScriptingSensitivity returns the CrossSiteScriptingSensitivity field if non-nil, zero value otherwise.

### GetCrossSiteScriptingSensitivityOk

`func (o *WAFThreatsConfiguration) GetCrossSiteScriptingSensitivityOk() (*CrossSiteScriptingSensitivityEnum, bool)`

GetCrossSiteScriptingSensitivityOk returns a tuple with the CrossSiteScriptingSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScriptingSensitivity

`func (o *WAFThreatsConfiguration) SetCrossSiteScriptingSensitivity(v CrossSiteScriptingSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetDirectoryTraversalSensitivity() DirectoryTraversalSensitivityEnum`

GetDirectoryTraversalSensitivity returns the DirectoryTraversalSensitivity field if non-nil, zero value otherwise.

### GetDirectoryTraversalSensitivityOk

`func (o *WAFThreatsConfiguration) GetDirectoryTraversalSensitivityOk() (*DirectoryTraversalSensitivityEnum, bool)`

GetDirectoryTraversalSensitivityOk returns a tuple with the DirectoryTraversalSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversalSensitivity

`func (o *WAFThreatsConfiguration) SetDirectoryTraversalSensitivity(v DirectoryTraversalSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetEvadingTricksSensitivity() EvadingTricksSensitivityEnum`

GetEvadingTricksSensitivity returns the EvadingTricksSensitivity field if non-nil, zero value otherwise.

### GetEvadingTricksSensitivityOk

`func (o *WAFThreatsConfiguration) GetEvadingTricksSensitivityOk() (*EvadingTricksSensitivityEnum, bool)`

GetEvadingTricksSensitivityOk returns a tuple with the EvadingTricksSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricksSensitivity

`func (o *WAFThreatsConfiguration) SetEvadingTricksSensitivity(v EvadingTricksSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetFileUploadSensitivity() FileUploadSensitivityEnum`

GetFileUploadSensitivity returns the FileUploadSensitivity field if non-nil, zero value otherwise.

### GetFileUploadSensitivityOk

`func (o *WAFThreatsConfiguration) GetFileUploadSensitivityOk() (*FileUploadSensitivityEnum, bool)`

GetFileUploadSensitivityOk returns a tuple with the FileUploadSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSensitivity

`func (o *WAFThreatsConfiguration) SetFileUploadSensitivity(v FileUploadSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetIdentifiedAttackSensitivity() IdentifiedAttackSensitivityEnum`

GetIdentifiedAttackSensitivity returns the IdentifiedAttackSensitivity field if non-nil, zero value otherwise.

### GetIdentifiedAttackSensitivityOk

`func (o *WAFThreatsConfiguration) GetIdentifiedAttackSensitivityOk() (*IdentifiedAttackSensitivityEnum, bool)`

GetIdentifiedAttackSensitivityOk returns a tuple with the IdentifiedAttackSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttackSensitivity

`func (o *WAFThreatsConfiguration) SetIdentifiedAttackSensitivity(v IdentifiedAttackSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetRemoteFileInclusionSensitivity() RemoteFileInclusionSensitivityEnum`

GetRemoteFileInclusionSensitivity returns the RemoteFileInclusionSensitivity field if non-nil, zero value otherwise.

### GetRemoteFileInclusionSensitivityOk

`func (o *WAFThreatsConfiguration) GetRemoteFileInclusionSensitivityOk() (*RemoteFileInclusionSensitivityEnum, bool)`

GetRemoteFileInclusionSensitivityOk returns a tuple with the RemoteFileInclusionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusionSensitivity

`func (o *WAFThreatsConfiguration) SetRemoteFileInclusionSensitivity(v RemoteFileInclusionSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetSqlInjectionSensitivity() SqlInjectionSensitivityEnum`

GetSqlInjectionSensitivity returns the SqlInjectionSensitivity field if non-nil, zero value otherwise.

### GetSqlInjectionSensitivityOk

`func (o *WAFThreatsConfiguration) GetSqlInjectionSensitivityOk() (*SqlInjectionSensitivityEnum, bool)`

GetSqlInjectionSensitivityOk returns a tuple with the SqlInjectionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjectionSensitivity

`func (o *WAFThreatsConfiguration) SetSqlInjectionSensitivity(v SqlInjectionSensitivityEnum)`

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

`func (o *WAFThreatsConfiguration) GetUnwantedAccessSensitivity() UnwantedAccessSensitivityEnum`

GetUnwantedAccessSensitivity returns the UnwantedAccessSensitivity field if non-nil, zero value otherwise.

### GetUnwantedAccessSensitivityOk

`func (o *WAFThreatsConfiguration) GetUnwantedAccessSensitivityOk() (*UnwantedAccessSensitivityEnum, bool)`

GetUnwantedAccessSensitivityOk returns a tuple with the UnwantedAccessSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccessSensitivity

`func (o *WAFThreatsConfiguration) SetUnwantedAccessSensitivity(v UnwantedAccessSensitivityEnum)`

SetUnwantedAccessSensitivity sets UnwantedAccessSensitivity field to given value.

### HasUnwantedAccessSensitivity

`func (o *WAFThreatsConfiguration) HasUnwantedAccessSensitivity() bool`

HasUnwantedAccessSensitivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


