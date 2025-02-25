/*
OceanBase Cloud API

API Documentation for OceanBase Cloud

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package obcloudsdk

import (
	"encoding/json"
)

// checks if the ModifyTenantNameRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyTenantNameRequestV2{}

// ModifyTenantNameRequestV2 struct for ModifyTenantNameRequestV2
type ModifyTenantNameRequestV2 struct {
	RequestId *string `json:"requestId,omitempty"`
	TenantName *string `json:"tenantName,omitempty"`
	PatternCheck *bool `json:"patternCheck,omitempty"`
}

// NewModifyTenantNameRequestV2 instantiates a new ModifyTenantNameRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyTenantNameRequestV2() *ModifyTenantNameRequestV2 {
	this := ModifyTenantNameRequestV2{}
	return &this
}

// NewModifyTenantNameRequestV2WithDefaults instantiates a new ModifyTenantNameRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyTenantNameRequestV2WithDefaults() *ModifyTenantNameRequestV2 {
	this := ModifyTenantNameRequestV2{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyTenantNameRequestV2) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTenantNameRequestV2) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyTenantNameRequestV2) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyTenantNameRequestV2) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *ModifyTenantNameRequestV2) GetTenantName() string {
	if o == nil || IsNil(o.TenantName) {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTenantNameRequestV2) GetTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.TenantName) {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *ModifyTenantNameRequestV2) HasTenantName() bool {
	if o != nil && !IsNil(o.TenantName) {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *ModifyTenantNameRequestV2) SetTenantName(v string) {
	o.TenantName = &v
}

// GetPatternCheck returns the PatternCheck field value if set, zero value otherwise.
func (o *ModifyTenantNameRequestV2) GetPatternCheck() bool {
	if o == nil || IsNil(o.PatternCheck) {
		var ret bool
		return ret
	}
	return *o.PatternCheck
}

// GetPatternCheckOk returns a tuple with the PatternCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTenantNameRequestV2) GetPatternCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.PatternCheck) {
		return nil, false
	}
	return o.PatternCheck, true
}

// HasPatternCheck returns a boolean if a field has been set.
func (o *ModifyTenantNameRequestV2) HasPatternCheck() bool {
	if o != nil && !IsNil(o.PatternCheck) {
		return true
	}

	return false
}

// SetPatternCheck gets a reference to the given bool and assigns it to the PatternCheck field.
func (o *ModifyTenantNameRequestV2) SetPatternCheck(v bool) {
	o.PatternCheck = &v
}

func (o ModifyTenantNameRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyTenantNameRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.TenantName) {
		toSerialize["tenantName"] = o.TenantName
	}
	if !IsNil(o.PatternCheck) {
		toSerialize["patternCheck"] = o.PatternCheck
	}
	return toSerialize, nil
}

type NullableModifyTenantNameRequestV2 struct {
	value *ModifyTenantNameRequestV2
	isSet bool
}

func (v NullableModifyTenantNameRequestV2) Get() *ModifyTenantNameRequestV2 {
	return v.value
}

func (v *NullableModifyTenantNameRequestV2) Set(val *ModifyTenantNameRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyTenantNameRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyTenantNameRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyTenantNameRequestV2(val *ModifyTenantNameRequestV2) *NullableModifyTenantNameRequestV2 {
	return &NullableModifyTenantNameRequestV2{value: val, isSet: true}
}

func (v NullableModifyTenantNameRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyTenantNameRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


