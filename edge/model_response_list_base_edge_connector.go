/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ResponseListBaseEdgeConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListBaseEdgeConnector{}

// ResponseListBaseEdgeConnector struct for ResponseListBaseEdgeConnector
type ResponseListBaseEdgeConnector struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	Modules EdgeConnectorModules `json:"modules"`
	Active *bool `json:"active,omitempty"`
	ProductVersion string `json:"product_version" validate:"regexp=\\\\d+\\\\.\\\\d+"`
	// * `http` - HTTP * `s3` - S3 * `edge_storage` - Edge Storage * `live_ingest` - Live Ingest
	Type string `json:"type"`
	Addresses []Address `json:"addresses,omitempty"`
	Tls *TLSEdgeConnector `json:"tls,omitempty"`
	// * `off` - Off * `ip_hash` - IP Hash * `least_connections` - Least Connections * `round_robin` - Round Robin
	LoadBalanceMethod *string `json:"load_balance_method,omitempty"`
	ConnectionPreference []string `json:"connection_preference,omitempty"`
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`
	ReadWriteTimeout *int64 `json:"read_write_timeout,omitempty"`
	MaxRetries *int64 `json:"max_retries,omitempty"`
}

type _ResponseListBaseEdgeConnector ResponseListBaseEdgeConnector

// NewResponseListBaseEdgeConnector instantiates a new ResponseListBaseEdgeConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListBaseEdgeConnector(id int64, name string, lastEditor string, lastModified time.Time, modules EdgeConnectorModules, productVersion string, type_ string) *ResponseListBaseEdgeConnector {
	this := ResponseListBaseEdgeConnector{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.Modules = modules
	this.ProductVersion = productVersion
	this.Type = type_
	return &this
}

// NewResponseListBaseEdgeConnectorWithDefaults instantiates a new ResponseListBaseEdgeConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListBaseEdgeConnectorWithDefaults() *ResponseListBaseEdgeConnector {
	this := ResponseListBaseEdgeConnector{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListBaseEdgeConnector) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListBaseEdgeConnector) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListBaseEdgeConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListBaseEdgeConnector) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListBaseEdgeConnector) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListBaseEdgeConnector) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListBaseEdgeConnector) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListBaseEdgeConnector) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetModules returns the Modules field value
func (o *ResponseListBaseEdgeConnector) GetModules() EdgeConnectorModules {
	if o == nil {
		var ret EdgeConnectorModules
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetModulesOk() (*EdgeConnectorModules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modules, true
}

// SetModules sets field value
func (o *ResponseListBaseEdgeConnector) SetModules(v EdgeConnectorModules) {
	o.Modules = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListBaseEdgeConnector) SetActive(v bool) {
	o.Active = &v
}

// GetProductVersion returns the ProductVersion field value
func (o *ResponseListBaseEdgeConnector) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *ResponseListBaseEdgeConnector) SetProductVersion(v string) {
	o.ProductVersion = v
}

// GetType returns the Type field value
func (o *ResponseListBaseEdgeConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResponseListBaseEdgeConnector) SetType(v string) {
	o.Type = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetAddresses() []Address {
	if o == nil || IsNil(o.Addresses) {
		var ret []Address
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetAddressesOk() ([]Address, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []Address and assigns it to the Addresses field.
func (o *ResponseListBaseEdgeConnector) SetAddresses(v []Address) {
	o.Addresses = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetTls() TLSEdgeConnector {
	if o == nil || IsNil(o.Tls) {
		var ret TLSEdgeConnector
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetTlsOk() (*TLSEdgeConnector, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given TLSEdgeConnector and assigns it to the Tls field.
func (o *ResponseListBaseEdgeConnector) SetTls(v TLSEdgeConnector) {
	o.Tls = &v
}

// GetLoadBalanceMethod returns the LoadBalanceMethod field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetLoadBalanceMethod() string {
	if o == nil || IsNil(o.LoadBalanceMethod) {
		var ret string
		return ret
	}
	return *o.LoadBalanceMethod
}

// GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetLoadBalanceMethodOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceMethod) {
		return nil, false
	}
	return o.LoadBalanceMethod, true
}

// HasLoadBalanceMethod returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasLoadBalanceMethod() bool {
	if o != nil && !IsNil(o.LoadBalanceMethod) {
		return true
	}

	return false
}

// SetLoadBalanceMethod gets a reference to the given string and assigns it to the LoadBalanceMethod field.
func (o *ResponseListBaseEdgeConnector) SetLoadBalanceMethod(v string) {
	o.LoadBalanceMethod = &v
}

// GetConnectionPreference returns the ConnectionPreference field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetConnectionPreference() []string {
	if o == nil || IsNil(o.ConnectionPreference) {
		var ret []string
		return ret
	}
	return o.ConnectionPreference
}

// GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetConnectionPreferenceOk() ([]string, bool) {
	if o == nil || IsNil(o.ConnectionPreference) {
		return nil, false
	}
	return o.ConnectionPreference, true
}

// HasConnectionPreference returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasConnectionPreference() bool {
	if o != nil && !IsNil(o.ConnectionPreference) {
		return true
	}

	return false
}

// SetConnectionPreference gets a reference to the given []string and assigns it to the ConnectionPreference field.
func (o *ResponseListBaseEdgeConnector) SetConnectionPreference(v []string) {
	o.ConnectionPreference = v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *ResponseListBaseEdgeConnector) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetReadWriteTimeout returns the ReadWriteTimeout field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetReadWriteTimeout() int64 {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		var ret int64
		return ret
	}
	return *o.ReadWriteTimeout
}

// GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetReadWriteTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		return nil, false
	}
	return o.ReadWriteTimeout, true
}

// HasReadWriteTimeout returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasReadWriteTimeout() bool {
	if o != nil && !IsNil(o.ReadWriteTimeout) {
		return true
	}

	return false
}

// SetReadWriteTimeout gets a reference to the given int64 and assigns it to the ReadWriteTimeout field.
func (o *ResponseListBaseEdgeConnector) SetReadWriteTimeout(v int64) {
	o.ReadWriteTimeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *ResponseListBaseEdgeConnector) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int64
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListBaseEdgeConnector) GetMaxRetriesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *ResponseListBaseEdgeConnector) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int64 and assigns it to the MaxRetries field.
func (o *ResponseListBaseEdgeConnector) SetMaxRetries(v int64) {
	o.MaxRetries = &v
}

func (o ResponseListBaseEdgeConnector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListBaseEdgeConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["modules"] = o.Modules
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["product_version"] = o.ProductVersion
	toSerialize["type"] = o.Type
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	if !IsNil(o.LoadBalanceMethod) {
		toSerialize["load_balance_method"] = o.LoadBalanceMethod
	}
	if !IsNil(o.ConnectionPreference) {
		toSerialize["connection_preference"] = o.ConnectionPreference
	}
	if !IsNil(o.ConnectionTimeout) {
		toSerialize["connection_timeout"] = o.ConnectionTimeout
	}
	if !IsNil(o.ReadWriteTimeout) {
		toSerialize["read_write_timeout"] = o.ReadWriteTimeout
	}
	if !IsNil(o.MaxRetries) {
		toSerialize["max_retries"] = o.MaxRetries
	}
	return toSerialize, nil
}

func (o *ResponseListBaseEdgeConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"modules",
		"product_version",
		"type",
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

	varResponseListBaseEdgeConnector := _ResponseListBaseEdgeConnector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListBaseEdgeConnector)

	if err != nil {
		return err
	}

	*o = ResponseListBaseEdgeConnector(varResponseListBaseEdgeConnector)

	return err
}

type NullableResponseListBaseEdgeConnector struct {
	value *ResponseListBaseEdgeConnector
	isSet bool
}

func (v NullableResponseListBaseEdgeConnector) Get() *ResponseListBaseEdgeConnector {
	return v.value
}

func (v *NullableResponseListBaseEdgeConnector) Set(val *ResponseListBaseEdgeConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListBaseEdgeConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListBaseEdgeConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListBaseEdgeConnector(val *ResponseListBaseEdgeConnector) *NullableResponseListBaseEdgeConnector {
	return &NullableResponseListBaseEdgeConnector{value: val, isSet: true}
}

func (v NullableResponseListBaseEdgeConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListBaseEdgeConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


