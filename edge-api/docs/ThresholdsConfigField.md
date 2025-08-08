# ThresholdsConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threat** | **string** | * &#x60;cross_site_scripting&#x60; - cross_site_scripting * &#x60;directory_traversal&#x60; - directory_traversal * &#x60;evading_tricks&#x60; - evading_tricks * &#x60;file_upload&#x60; - file_upload * &#x60;identified_attack&#x60; - identified_attack * &#x60;remote_file_inclusion&#x60; - remote_file_inclusion * &#x60;sql_injection&#x60; - sql_injection * &#x60;unwanted_access&#x60; - unwanted_access | 
**Sensitivity** | Pointer to **string** | * &#x60;highest&#x60; - Highest * &#x60;high&#x60; - High * &#x60;medium&#x60; - Medium * &#x60;low&#x60; - Low * &#x60;lowest&#x60; - Lowest | [optional] 

## Methods

### NewThresholdsConfigField

`func NewThresholdsConfigField(threat string, ) *ThresholdsConfigField`

NewThresholdsConfigField instantiates a new ThresholdsConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdsConfigFieldWithDefaults

`func NewThresholdsConfigFieldWithDefaults() *ThresholdsConfigField`

NewThresholdsConfigFieldWithDefaults instantiates a new ThresholdsConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreat

`func (o *ThresholdsConfigField) GetThreat() string`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *ThresholdsConfigField) GetThreatOk() (*string, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *ThresholdsConfigField) SetThreat(v string)`

SetThreat sets Threat field to given value.


### GetSensitivity

`func (o *ThresholdsConfigField) GetSensitivity() string`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *ThresholdsConfigField) GetSensitivityOk() (*string, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *ThresholdsConfigField) SetSensitivity(v string)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *ThresholdsConfigField) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


