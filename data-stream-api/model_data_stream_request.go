/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data-stream-api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DataStreamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStreamRequest{}

// DataStreamRequest struct for DataStreamRequest
type DataStreamRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	// * `http` - Edge Applications * `waf` - WAF Events * `cells_console` - Edge Functions * `rtm_activity` - Activity History
	DataSource string `json:"data_source"`
	DataSetId int64 `json:"data_set_id"`
	Active *bool `json:"active,omitempty"`
	Filters DataStreamFilterRequest `json:"filters"`
	Endpoint EndpointRequest `json:"endpoint"`
}

type _DataStreamRequest DataStreamRequest

// NewDataStreamRequest instantiates a new DataStreamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStreamRequest(name string, dataSource string, dataSetId int64, filters DataStreamFilterRequest, endpoint EndpointRequest) *DataStreamRequest {
	this := DataStreamRequest{}
	this.Name = name
	this.DataSource = dataSource
	this.DataSetId = dataSetId
	this.Filters = filters
	this.Endpoint = endpoint
	return &this
}

// NewDataStreamRequestWithDefaults instantiates a new DataStreamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStreamRequestWithDefaults() *DataStreamRequest {
	this := DataStreamRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DataStreamRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataStreamRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataStreamRequest) SetName(v string) {
	o.Name = v
}

// GetDataSource returns the DataSource field value
func (o *DataStreamRequest) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *DataStreamRequest) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *DataStreamRequest) SetDataSource(v string) {
	o.DataSource = v
}

// GetDataSetId returns the DataSetId field value
func (o *DataStreamRequest) GetDataSetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value
// and a boolean to check if the value has been set.
func (o *DataStreamRequest) GetDataSetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSetId, true
}

// SetDataSetId sets field value
func (o *DataStreamRequest) SetDataSetId(v int64) {
	o.DataSetId = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DataStreamRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DataStreamRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DataStreamRequest) SetActive(v bool) {
	o.Active = &v
}

// GetFilters returns the Filters field value
func (o *DataStreamRequest) GetFilters() DataStreamFilterRequest {
	if o == nil {
		var ret DataStreamFilterRequest
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *DataStreamRequest) GetFiltersOk() (*DataStreamFilterRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *DataStreamRequest) SetFilters(v DataStreamFilterRequest) {
	o.Filters = v
}

// GetEndpoint returns the Endpoint field value
func (o *DataStreamRequest) GetEndpoint() EndpointRequest {
	if o == nil {
		var ret EndpointRequest
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *DataStreamRequest) GetEndpointOk() (*EndpointRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *DataStreamRequest) SetEndpoint(v EndpointRequest) {
	o.Endpoint = v
}

func (o DataStreamRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStreamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["data_source"] = o.DataSource
	toSerialize["data_set_id"] = o.DataSetId
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["filters"] = o.Filters
	toSerialize["endpoint"] = o.Endpoint
	return toSerialize, nil
}

func (o *DataStreamRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"data_source",
		"data_set_id",
		"filters",
		"endpoint",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDataStreamRequest := _DataStreamRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataStreamRequest)

	if err != nil {
		return err
	}

	*o = DataStreamRequest(varDataStreamRequest)

	return err
}

type NullableDataStreamRequest struct {
	value *DataStreamRequest
	isSet bool
}

func (v NullableDataStreamRequest) Get() *DataStreamRequest {
	return v.value
}

func (v *NullableDataStreamRequest) Set(val *DataStreamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStreamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStreamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStreamRequest(val *DataStreamRequest) *NullableDataStreamRequest {
	return &NullableDataStreamRequest{value: val, isSet: true}
}

func (v NullableDataStreamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStreamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


