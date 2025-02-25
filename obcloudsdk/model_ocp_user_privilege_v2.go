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

// checks if the OcpUserPrivilegeV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OcpUserPrivilegeV2{}

// OcpUserPrivilegeV2 struct for OcpUserPrivilegeV2
type OcpUserPrivilegeV2 struct {
	UserName *string `json:"userName,omitempty"`
	Database *string `json:"database,omitempty"`
	Table *string `json:"table,omitempty"`
	Role *string `json:"role,omitempty"`
	Privileges *string `json:"privileges,omitempty"`
	IsSuccess *bool `json:"isSuccess,omitempty"`
	WithGrant *int64 `json:"withGrant,omitempty"`
}

// NewOcpUserPrivilegeV2 instantiates a new OcpUserPrivilegeV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOcpUserPrivilegeV2() *OcpUserPrivilegeV2 {
	this := OcpUserPrivilegeV2{}
	return &this
}

// NewOcpUserPrivilegeV2WithDefaults instantiates a new OcpUserPrivilegeV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOcpUserPrivilegeV2WithDefaults() *OcpUserPrivilegeV2 {
	this := OcpUserPrivilegeV2{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *OcpUserPrivilegeV2) SetUserName(v string) {
	o.UserName = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *OcpUserPrivilegeV2) SetDatabase(v string) {
	o.Database = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *OcpUserPrivilegeV2) SetTable(v string) {
	o.Table = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OcpUserPrivilegeV2) SetRole(v string) {
	o.Role = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetPrivileges() string {
	if o == nil || IsNil(o.Privileges) {
		var ret string
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetPrivilegesOk() (*string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given string and assigns it to the Privileges field.
func (o *OcpUserPrivilegeV2) SetPrivileges(v string) {
	o.Privileges = &v
}

// GetIsSuccess returns the IsSuccess field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetIsSuccess() bool {
	if o == nil || IsNil(o.IsSuccess) {
		var ret bool
		return ret
	}
	return *o.IsSuccess
}

// GetIsSuccessOk returns a tuple with the IsSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetIsSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuccess) {
		return nil, false
	}
	return o.IsSuccess, true
}

// HasIsSuccess returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasIsSuccess() bool {
	if o != nil && !IsNil(o.IsSuccess) {
		return true
	}

	return false
}

// SetIsSuccess gets a reference to the given bool and assigns it to the IsSuccess field.
func (o *OcpUserPrivilegeV2) SetIsSuccess(v bool) {
	o.IsSuccess = &v
}

// GetWithGrant returns the WithGrant field value if set, zero value otherwise.
func (o *OcpUserPrivilegeV2) GetWithGrant() int64 {
	if o == nil || IsNil(o.WithGrant) {
		var ret int64
		return ret
	}
	return *o.WithGrant
}

// GetWithGrantOk returns a tuple with the WithGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcpUserPrivilegeV2) GetWithGrantOk() (*int64, bool) {
	if o == nil || IsNil(o.WithGrant) {
		return nil, false
	}
	return o.WithGrant, true
}

// HasWithGrant returns a boolean if a field has been set.
func (o *OcpUserPrivilegeV2) HasWithGrant() bool {
	if o != nil && !IsNil(o.WithGrant) {
		return true
	}

	return false
}

// SetWithGrant gets a reference to the given int64 and assigns it to the WithGrant field.
func (o *OcpUserPrivilegeV2) SetWithGrant(v int64) {
	o.WithGrant = &v
}

func (o OcpUserPrivilegeV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OcpUserPrivilegeV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !IsNil(o.IsSuccess) {
		toSerialize["isSuccess"] = o.IsSuccess
	}
	if !IsNil(o.WithGrant) {
		toSerialize["withGrant"] = o.WithGrant
	}
	return toSerialize, nil
}

type NullableOcpUserPrivilegeV2 struct {
	value *OcpUserPrivilegeV2
	isSet bool
}

func (v NullableOcpUserPrivilegeV2) Get() *OcpUserPrivilegeV2 {
	return v.value
}

func (v *NullableOcpUserPrivilegeV2) Set(val *OcpUserPrivilegeV2) {
	v.value = val
	v.isSet = true
}

func (v NullableOcpUserPrivilegeV2) IsSet() bool {
	return v.isSet
}

func (v *NullableOcpUserPrivilegeV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOcpUserPrivilegeV2(val *OcpUserPrivilegeV2) *NullableOcpUserPrivilegeV2 {
	return &NullableOcpUserPrivilegeV2{value: val, isSet: true}
}

func (v NullableOcpUserPrivilegeV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOcpUserPrivilegeV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


