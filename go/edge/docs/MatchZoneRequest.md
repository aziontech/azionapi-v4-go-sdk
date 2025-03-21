# MatchZoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | * &#x60;conditional_query_string&#x60; - Conditional Query String * &#x60;conditional_request_body&#x60; - Conditional Request Body * &#x60;conditional_request_header&#x60; - Conditional Request Header * &#x60;file_name&#x60; - File Name (Multipart Body) * &#x60;path&#x60; - Path * &#x60;query_string&#x60; - Query String * &#x60;raw_body&#x60; - Raw Body * &#x60;request_body&#x60; - Request Body * &#x60;request_header&#x60; - Request Header | 
**ZoneInput** | Pointer to **NullableString** |  | [optional] 
**MatchesOn** | Pointer to [**NullableMatchZoneMatchesOn**](MatchZoneMatchesOn.md) |  | [optional] 

## Methods

### NewMatchZoneRequest

`func NewMatchZoneRequest(zone string, ) *MatchZoneRequest`

NewMatchZoneRequest instantiates a new MatchZoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchZoneRequestWithDefaults

`func NewMatchZoneRequestWithDefaults() *MatchZoneRequest`

NewMatchZoneRequestWithDefaults instantiates a new MatchZoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *MatchZoneRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *MatchZoneRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *MatchZoneRequest) SetZone(v string)`

SetZone sets Zone field to given value.


### GetZoneInput

`func (o *MatchZoneRequest) GetZoneInput() string`

GetZoneInput returns the ZoneInput field if non-nil, zero value otherwise.

### GetZoneInputOk

`func (o *MatchZoneRequest) GetZoneInputOk() (*string, bool)`

GetZoneInputOk returns a tuple with the ZoneInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInput

`func (o *MatchZoneRequest) SetZoneInput(v string)`

SetZoneInput sets ZoneInput field to given value.

### HasZoneInput

`func (o *MatchZoneRequest) HasZoneInput() bool`

HasZoneInput returns a boolean if a field has been set.

### SetZoneInputNil

`func (o *MatchZoneRequest) SetZoneInputNil(b bool)`

 SetZoneInputNil sets the value for ZoneInput to be an explicit nil

### UnsetZoneInput
`func (o *MatchZoneRequest) UnsetZoneInput()`

UnsetZoneInput ensures that no value is present for ZoneInput, not even an explicit nil
### GetMatchesOn

`func (o *MatchZoneRequest) GetMatchesOn() MatchZoneMatchesOn`

GetMatchesOn returns the MatchesOn field if non-nil, zero value otherwise.

### GetMatchesOnOk

`func (o *MatchZoneRequest) GetMatchesOnOk() (*MatchZoneMatchesOn, bool)`

GetMatchesOnOk returns a tuple with the MatchesOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesOn

`func (o *MatchZoneRequest) SetMatchesOn(v MatchZoneMatchesOn)`

SetMatchesOn sets MatchesOn field to given value.

### HasMatchesOn

`func (o *MatchZoneRequest) HasMatchesOn() bool`

HasMatchesOn returns a boolean if a field has been set.

### SetMatchesOnNil

`func (o *MatchZoneRequest) SetMatchesOnNil(b bool)`

 SetMatchesOnNil sets the value for MatchesOn to be an explicit nil

### UnsetMatchesOn
`func (o *MatchZoneRequest) UnsetMatchesOn()`

UnsetMatchesOn ensures that no value is present for MatchesOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


