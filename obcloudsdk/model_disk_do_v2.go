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

// checks if the DiskDoV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskDoV2{}

// DiskDoV2 struct for DiskDoV2
type DiskDoV2 struct {
	TotalDiskSize *float64 `json:"totalDiskSize,omitempty"`
	UsedDiskSize *float64 `json:"usedDiskSize,omitempty"`
}

// NewDiskDoV2 instantiates a new DiskDoV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskDoV2() *DiskDoV2 {
	this := DiskDoV2{}
	return &this
}

// NewDiskDoV2WithDefaults instantiates a new DiskDoV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskDoV2WithDefaults() *DiskDoV2 {
	this := DiskDoV2{}
	return &this
}

// GetTotalDiskSize returns the TotalDiskSize field value if set, zero value otherwise.
func (o *DiskDoV2) GetTotalDiskSize() float64 {
	if o == nil || IsNil(o.TotalDiskSize) {
		var ret float64
		return ret
	}
	return *o.TotalDiskSize
}

// GetTotalDiskSizeOk returns a tuple with the TotalDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskDoV2) GetTotalDiskSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDiskSize) {
		return nil, false
	}
	return o.TotalDiskSize, true
}

// HasTotalDiskSize returns a boolean if a field has been set.
func (o *DiskDoV2) HasTotalDiskSize() bool {
	if o != nil && !IsNil(o.TotalDiskSize) {
		return true
	}

	return false
}

// SetTotalDiskSize gets a reference to the given float64 and assigns it to the TotalDiskSize field.
func (o *DiskDoV2) SetTotalDiskSize(v float64) {
	o.TotalDiskSize = &v
}

// GetUsedDiskSize returns the UsedDiskSize field value if set, zero value otherwise.
func (o *DiskDoV2) GetUsedDiskSize() float64 {
	if o == nil || IsNil(o.UsedDiskSize) {
		var ret float64
		return ret
	}
	return *o.UsedDiskSize
}

// GetUsedDiskSizeOk returns a tuple with the UsedDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskDoV2) GetUsedDiskSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.UsedDiskSize) {
		return nil, false
	}
	return o.UsedDiskSize, true
}

// HasUsedDiskSize returns a boolean if a field has been set.
func (o *DiskDoV2) HasUsedDiskSize() bool {
	if o != nil && !IsNil(o.UsedDiskSize) {
		return true
	}

	return false
}

// SetUsedDiskSize gets a reference to the given float64 and assigns it to the UsedDiskSize field.
func (o *DiskDoV2) SetUsedDiskSize(v float64) {
	o.UsedDiskSize = &v
}

func (o DiskDoV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskDoV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalDiskSize) {
		toSerialize["totalDiskSize"] = o.TotalDiskSize
	}
	if !IsNil(o.UsedDiskSize) {
		toSerialize["usedDiskSize"] = o.UsedDiskSize
	}
	return toSerialize, nil
}

type NullableDiskDoV2 struct {
	value *DiskDoV2
	isSet bool
}

func (v NullableDiskDoV2) Get() *DiskDoV2 {
	return v.value
}

func (v *NullableDiskDoV2) Set(val *DiskDoV2) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskDoV2) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskDoV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskDoV2(val *DiskDoV2) *NullableDiskDoV2 {
	return &NullableDiskDoV2{value: val, isSet: true}
}

func (v NullableDiskDoV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskDoV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


