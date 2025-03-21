# ResponseBadRequestWAFRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **[]string** |  | [optional] 
**MatchZones** | Pointer to **[]string** |  | [optional] 
**UseRegex** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestWAFRule

`func NewResponseBadRequestWAFRule() *ResponseBadRequestWAFRule`

NewResponseBadRequestWAFRule instantiates a new ResponseBadRequestWAFRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestWAFRuleWithDefaults

`func NewResponseBadRequestWAFRuleWithDefaults() *ResponseBadRequestWAFRule`

NewResponseBadRequestWAFRuleWithDefaults instantiates a new ResponseBadRequestWAFRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *ResponseBadRequestWAFRule) GetRuleId() []string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ResponseBadRequestWAFRule) GetRuleIdOk() (*[]string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ResponseBadRequestWAFRule) SetRuleId(v []string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ResponseBadRequestWAFRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetName

`func (o *ResponseBadRequestWAFRule) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestWAFRule) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestWAFRule) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestWAFRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ResponseBadRequestWAFRule) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResponseBadRequestWAFRule) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResponseBadRequestWAFRule) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ResponseBadRequestWAFRule) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMatchZones

`func (o *ResponseBadRequestWAFRule) GetMatchZones() []string`

GetMatchZones returns the MatchZones field if non-nil, zero value otherwise.

### GetMatchZonesOk

`func (o *ResponseBadRequestWAFRule) GetMatchZonesOk() (*[]string, bool)`

GetMatchZonesOk returns a tuple with the MatchZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchZones

`func (o *ResponseBadRequestWAFRule) SetMatchZones(v []string)`

SetMatchZones sets MatchZones field to given value.

### HasMatchZones

`func (o *ResponseBadRequestWAFRule) HasMatchZones() bool`

HasMatchZones returns a boolean if a field has been set.

### GetUseRegex

`func (o *ResponseBadRequestWAFRule) GetUseRegex() []string`

GetUseRegex returns the UseRegex field if non-nil, zero value otherwise.

### GetUseRegexOk

`func (o *ResponseBadRequestWAFRule) GetUseRegexOk() (*[]string, bool)`

GetUseRegexOk returns a tuple with the UseRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRegex

`func (o *ResponseBadRequestWAFRule) SetUseRegex(v []string)`

SetUseRegex sets UseRegex field to given value.

### HasUseRegex

`func (o *ResponseBadRequestWAFRule) HasUseRegex() bool`

HasUseRegex returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestWAFRule) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestWAFRule) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestWAFRule) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestWAFRule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestWAFRule) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestWAFRule) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestWAFRule) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestWAFRule) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestWAFRule) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestWAFRule) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestWAFRule) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestWAFRule) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestWAFRule) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestWAFRule) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestWAFRule) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestWAFRule) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


