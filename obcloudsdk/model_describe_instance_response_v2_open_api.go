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

// checks if the DescribeInstanceResponseV2OpenAPI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeInstanceResponseV2OpenAPI{}

// DescribeInstanceResponseV2OpenAPI struct for DescribeInstanceResponseV2OpenAPI
type DescribeInstanceResponseV2OpenAPI struct {
	Region *string `json:"region,omitempty"`
	VpcId *string `json:"vpcId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	InstanceName *string `json:"instanceName,omitempty"`
	InstanceClass *string `json:"instanceClass,omitempty"`
	Series *string `json:"series,omitempty"`
	PayType *string `json:"payType,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty"`
	AvailableZones []string `json:"availableZones,omitempty"`
	Version *string `json:"version,omitempty"`
	ObRpmVersion *string `json:"obRpmVersion,omitempty"`
	DeployType *string `json:"deployType,omitempty"`
	DiskType *string `json:"diskType,omitempty"`
	DeployMode *string `json:"deployMode,omitempty"`
	ReplicaMode *string `json:"replicaMode,omitempty"`
	DataMergeTime *string `json:"dataMergeTime,omitempty"`
	Status *string `json:"status,omitempty"`
	Resource *ResourceDoV2 `json:"resource,omitempty"`
	CloudProvider *string `json:"cloudProvider,omitempty"`
	InstanceType *string `json:"instanceType,omitempty"`
	InstanceRole *string `json:"instanceRole,omitempty"`
	PrimaryInstanceId *string `json:"primaryInstanceId,omitempty"`
	StandbyInstanceIds []string `json:"standbyInstanceIds,omitempty"`
	SaleChannel *string `json:"saleChannel,omitempty"`
	ManagementMode *string `json:"managementMode,omitempty"`
	TagList []TagKeyValueDo `json:"tagList,omitempty"`
	NodeNum *int32 `json:"nodeNum,omitempty"`
	CpuArchitecture *string `json:"cpuArchitecture,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	StopTime *string `json:"stopTime,omitempty"`
}

// NewDescribeInstanceResponseV2OpenAPI instantiates a new DescribeInstanceResponseV2OpenAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeInstanceResponseV2OpenAPI() *DescribeInstanceResponseV2OpenAPI {
	this := DescribeInstanceResponseV2OpenAPI{}
	return &this
}

// NewDescribeInstanceResponseV2OpenAPIWithDefaults instantiates a new DescribeInstanceResponseV2OpenAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeInstanceResponseV2OpenAPIWithDefaults() *DescribeInstanceResponseV2OpenAPI {
	this := DescribeInstanceResponseV2OpenAPI{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DescribeInstanceResponseV2OpenAPI) SetRegion(v string) {
	o.Region = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *DescribeInstanceResponseV2OpenAPI) SetVpcId(v string) {
	o.VpcId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceName() bool {
	if o != nil && !IsNil(o.InstanceName) {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetInstanceClass returns the InstanceClass field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceClass() string {
	if o == nil || IsNil(o.InstanceClass) {
		var ret string
		return ret
	}
	return *o.InstanceClass
}

// GetInstanceClassOk returns a tuple with the InstanceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceClassOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceClass) {
		return nil, false
	}
	return o.InstanceClass, true
}

// HasInstanceClass returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceClass() bool {
	if o != nil && !IsNil(o.InstanceClass) {
		return true
	}

	return false
}

// SetInstanceClass gets a reference to the given string and assigns it to the InstanceClass field.
func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceClass(v string) {
	o.InstanceClass = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetSeries() string {
	if o == nil || IsNil(o.Series) {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *DescribeInstanceResponseV2OpenAPI) SetSeries(v string) {
	o.Series = &v
}

// GetPayType returns the PayType field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetPayType() string {
	if o == nil || IsNil(o.PayType) {
		var ret string
		return ret
	}
	return *o.PayType
}

// GetPayTypeOk returns a tuple with the PayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetPayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayType) {
		return nil, false
	}
	return o.PayType, true
}

// HasPayType returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasPayType() bool {
	if o != nil && !IsNil(o.PayType) {
		return true
	}

	return false
}

// SetPayType gets a reference to the given string and assigns it to the PayType field.
func (o *DescribeInstanceResponseV2OpenAPI) SetPayType(v string) {
	o.PayType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *DescribeInstanceResponseV2OpenAPI) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetExpireTime() string {
	if o == nil || IsNil(o.ExpireTime) {
		var ret string
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given string and assigns it to the ExpireTime field.
func (o *DescribeInstanceResponseV2OpenAPI) SetExpireTime(v string) {
	o.ExpireTime = &v
}

// GetAvailableZones returns the AvailableZones field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetAvailableZones() []string {
	if o == nil || IsNil(o.AvailableZones) {
		var ret []string
		return ret
	}
	return o.AvailableZones
}

// GetAvailableZonesOk returns a tuple with the AvailableZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetAvailableZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableZones) {
		return nil, false
	}
	return o.AvailableZones, true
}

// HasAvailableZones returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasAvailableZones() bool {
	if o != nil && !IsNil(o.AvailableZones) {
		return true
	}

	return false
}

// SetAvailableZones gets a reference to the given []string and assigns it to the AvailableZones field.
func (o *DescribeInstanceResponseV2OpenAPI) SetAvailableZones(v []string) {
	o.AvailableZones = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DescribeInstanceResponseV2OpenAPI) SetVersion(v string) {
	o.Version = &v
}

// GetObRpmVersion returns the ObRpmVersion field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetObRpmVersion() string {
	if o == nil || IsNil(o.ObRpmVersion) {
		var ret string
		return ret
	}
	return *o.ObRpmVersion
}

// GetObRpmVersionOk returns a tuple with the ObRpmVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetObRpmVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ObRpmVersion) {
		return nil, false
	}
	return o.ObRpmVersion, true
}

// HasObRpmVersion returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasObRpmVersion() bool {
	if o != nil && !IsNil(o.ObRpmVersion) {
		return true
	}

	return false
}

// SetObRpmVersion gets a reference to the given string and assigns it to the ObRpmVersion field.
func (o *DescribeInstanceResponseV2OpenAPI) SetObRpmVersion(v string) {
	o.ObRpmVersion = &v
}

// GetDeployType returns the DeployType field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetDeployType() string {
	if o == nil || IsNil(o.DeployType) {
		var ret string
		return ret
	}
	return *o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetDeployTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeployType) {
		return nil, false
	}
	return o.DeployType, true
}

// HasDeployType returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasDeployType() bool {
	if o != nil && !IsNil(o.DeployType) {
		return true
	}

	return false
}

// SetDeployType gets a reference to the given string and assigns it to the DeployType field.
func (o *DescribeInstanceResponseV2OpenAPI) SetDeployType(v string) {
	o.DeployType = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetDiskType() string {
	if o == nil || IsNil(o.DiskType) {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetDiskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiskType) {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasDiskType() bool {
	if o != nil && !IsNil(o.DiskType) {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *DescribeInstanceResponseV2OpenAPI) SetDiskType(v string) {
	o.DiskType = &v
}

// GetDeployMode returns the DeployMode field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetDeployMode() string {
	if o == nil || IsNil(o.DeployMode) {
		var ret string
		return ret
	}
	return *o.DeployMode
}

// GetDeployModeOk returns a tuple with the DeployMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetDeployModeOk() (*string, bool) {
	if o == nil || IsNil(o.DeployMode) {
		return nil, false
	}
	return o.DeployMode, true
}

// HasDeployMode returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasDeployMode() bool {
	if o != nil && !IsNil(o.DeployMode) {
		return true
	}

	return false
}

// SetDeployMode gets a reference to the given string and assigns it to the DeployMode field.
func (o *DescribeInstanceResponseV2OpenAPI) SetDeployMode(v string) {
	o.DeployMode = &v
}

// GetReplicaMode returns the ReplicaMode field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetReplicaMode() string {
	if o == nil || IsNil(o.ReplicaMode) {
		var ret string
		return ret
	}
	return *o.ReplicaMode
}

// GetReplicaModeOk returns a tuple with the ReplicaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetReplicaModeOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaMode) {
		return nil, false
	}
	return o.ReplicaMode, true
}

// HasReplicaMode returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasReplicaMode() bool {
	if o != nil && !IsNil(o.ReplicaMode) {
		return true
	}

	return false
}

// SetReplicaMode gets a reference to the given string and assigns it to the ReplicaMode field.
func (o *DescribeInstanceResponseV2OpenAPI) SetReplicaMode(v string) {
	o.ReplicaMode = &v
}

// GetDataMergeTime returns the DataMergeTime field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetDataMergeTime() string {
	if o == nil || IsNil(o.DataMergeTime) {
		var ret string
		return ret
	}
	return *o.DataMergeTime
}

// GetDataMergeTimeOk returns a tuple with the DataMergeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetDataMergeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DataMergeTime) {
		return nil, false
	}
	return o.DataMergeTime, true
}

// HasDataMergeTime returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasDataMergeTime() bool {
	if o != nil && !IsNil(o.DataMergeTime) {
		return true
	}

	return false
}

// SetDataMergeTime gets a reference to the given string and assigns it to the DataMergeTime field.
func (o *DescribeInstanceResponseV2OpenAPI) SetDataMergeTime(v string) {
	o.DataMergeTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DescribeInstanceResponseV2OpenAPI) SetStatus(v string) {
	o.Status = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetResource() ResourceDoV2 {
	if o == nil || IsNil(o.Resource) {
		var ret ResourceDoV2
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetResourceOk() (*ResourceDoV2, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ResourceDoV2 and assigns it to the Resource field.
func (o *DescribeInstanceResponseV2OpenAPI) SetResource(v ResourceDoV2) {
	o.Resource = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DescribeInstanceResponseV2OpenAPI) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetInstanceRole returns the InstanceRole field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceRole() string {
	if o == nil || IsNil(o.InstanceRole) {
		var ret string
		return ret
	}
	return *o.InstanceRole
}

// GetInstanceRoleOk returns a tuple with the InstanceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceRoleOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceRole) {
		return nil, false
	}
	return o.InstanceRole, true
}

// HasInstanceRole returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceRole() bool {
	if o != nil && !IsNil(o.InstanceRole) {
		return true
	}

	return false
}

// SetInstanceRole gets a reference to the given string and assigns it to the InstanceRole field.
func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceRole(v string) {
	o.InstanceRole = &v
}

// GetPrimaryInstanceId returns the PrimaryInstanceId field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetPrimaryInstanceId() string {
	if o == nil || IsNil(o.PrimaryInstanceId) {
		var ret string
		return ret
	}
	return *o.PrimaryInstanceId
}

// GetPrimaryInstanceIdOk returns a tuple with the PrimaryInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetPrimaryInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryInstanceId) {
		return nil, false
	}
	return o.PrimaryInstanceId, true
}

// HasPrimaryInstanceId returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasPrimaryInstanceId() bool {
	if o != nil && !IsNil(o.PrimaryInstanceId) {
		return true
	}

	return false
}

// SetPrimaryInstanceId gets a reference to the given string and assigns it to the PrimaryInstanceId field.
func (o *DescribeInstanceResponseV2OpenAPI) SetPrimaryInstanceId(v string) {
	o.PrimaryInstanceId = &v
}

// GetStandbyInstanceIds returns the StandbyInstanceIds field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetStandbyInstanceIds() []string {
	if o == nil || IsNil(o.StandbyInstanceIds) {
		var ret []string
		return ret
	}
	return o.StandbyInstanceIds
}

// GetStandbyInstanceIdsOk returns a tuple with the StandbyInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetStandbyInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StandbyInstanceIds) {
		return nil, false
	}
	return o.StandbyInstanceIds, true
}

// HasStandbyInstanceIds returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasStandbyInstanceIds() bool {
	if o != nil && !IsNil(o.StandbyInstanceIds) {
		return true
	}

	return false
}

// SetStandbyInstanceIds gets a reference to the given []string and assigns it to the StandbyInstanceIds field.
func (o *DescribeInstanceResponseV2OpenAPI) SetStandbyInstanceIds(v []string) {
	o.StandbyInstanceIds = v
}

// GetSaleChannel returns the SaleChannel field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetSaleChannel() string {
	if o == nil || IsNil(o.SaleChannel) {
		var ret string
		return ret
	}
	return *o.SaleChannel
}

// GetSaleChannelOk returns a tuple with the SaleChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetSaleChannelOk() (*string, bool) {
	if o == nil || IsNil(o.SaleChannel) {
		return nil, false
	}
	return o.SaleChannel, true
}

// HasSaleChannel returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasSaleChannel() bool {
	if o != nil && !IsNil(o.SaleChannel) {
		return true
	}

	return false
}

// SetSaleChannel gets a reference to the given string and assigns it to the SaleChannel field.
func (o *DescribeInstanceResponseV2OpenAPI) SetSaleChannel(v string) {
	o.SaleChannel = &v
}

// GetManagementMode returns the ManagementMode field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetManagementMode() string {
	if o == nil || IsNil(o.ManagementMode) {
		var ret string
		return ret
	}
	return *o.ManagementMode
}

// GetManagementModeOk returns a tuple with the ManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetManagementModeOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementMode) {
		return nil, false
	}
	return o.ManagementMode, true
}

// HasManagementMode returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasManagementMode() bool {
	if o != nil && !IsNil(o.ManagementMode) {
		return true
	}

	return false
}

// SetManagementMode gets a reference to the given string and assigns it to the ManagementMode field.
func (o *DescribeInstanceResponseV2OpenAPI) SetManagementMode(v string) {
	o.ManagementMode = &v
}

// GetTagList returns the TagList field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetTagList() []TagKeyValueDo {
	if o == nil || IsNil(o.TagList) {
		var ret []TagKeyValueDo
		return ret
	}
	return o.TagList
}

// GetTagListOk returns a tuple with the TagList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetTagListOk() ([]TagKeyValueDo, bool) {
	if o == nil || IsNil(o.TagList) {
		return nil, false
	}
	return o.TagList, true
}

// HasTagList returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasTagList() bool {
	if o != nil && !IsNil(o.TagList) {
		return true
	}

	return false
}

// SetTagList gets a reference to the given []TagKeyValueDo and assigns it to the TagList field.
func (o *DescribeInstanceResponseV2OpenAPI) SetTagList(v []TagKeyValueDo) {
	o.TagList = v
}

// GetNodeNum returns the NodeNum field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetNodeNum() int32 {
	if o == nil || IsNil(o.NodeNum) {
		var ret int32
		return ret
	}
	return *o.NodeNum
}

// GetNodeNumOk returns a tuple with the NodeNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetNodeNumOk() (*int32, bool) {
	if o == nil || IsNil(o.NodeNum) {
		return nil, false
	}
	return o.NodeNum, true
}

// HasNodeNum returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasNodeNum() bool {
	if o != nil && !IsNil(o.NodeNum) {
		return true
	}

	return false
}

// SetNodeNum gets a reference to the given int32 and assigns it to the NodeNum field.
func (o *DescribeInstanceResponseV2OpenAPI) SetNodeNum(v int32) {
	o.NodeNum = &v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetCpuArchitecture() string {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret string
		return ret
	}
	return *o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetCpuArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// HasCpuArchitecture returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasCpuArchitecture() bool {
	if o != nil && !IsNil(o.CpuArchitecture) {
		return true
	}

	return false
}

// SetCpuArchitecture gets a reference to the given string and assigns it to the CpuArchitecture field.
func (o *DescribeInstanceResponseV2OpenAPI) SetCpuArchitecture(v string) {
	o.CpuArchitecture = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DescribeInstanceResponseV2OpenAPI) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *DescribeInstanceResponseV2OpenAPI) GetStopTime() string {
	if o == nil || IsNil(o.StopTime) {
		var ret string
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeInstanceResponseV2OpenAPI) GetStopTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StopTime) {
		return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *DescribeInstanceResponseV2OpenAPI) HasStopTime() bool {
	if o != nil && !IsNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given string and assigns it to the StopTime field.
func (o *DescribeInstanceResponseV2OpenAPI) SetStopTime(v string) {
	o.StopTime = &v
}

func (o DescribeInstanceResponseV2OpenAPI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeInstanceResponseV2OpenAPI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.VpcId) {
		toSerialize["vpcId"] = o.VpcId
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.InstanceName) {
		toSerialize["instanceName"] = o.InstanceName
	}
	if !IsNil(o.InstanceClass) {
		toSerialize["instanceClass"] = o.InstanceClass
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.PayType) {
		toSerialize["payType"] = o.PayType
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !IsNil(o.AvailableZones) {
		toSerialize["availableZones"] = o.AvailableZones
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ObRpmVersion) {
		toSerialize["obRpmVersion"] = o.ObRpmVersion
	}
	if !IsNil(o.DeployType) {
		toSerialize["deployType"] = o.DeployType
	}
	if !IsNil(o.DiskType) {
		toSerialize["diskType"] = o.DiskType
	}
	if !IsNil(o.DeployMode) {
		toSerialize["deployMode"] = o.DeployMode
	}
	if !IsNil(o.ReplicaMode) {
		toSerialize["replicaMode"] = o.ReplicaMode
	}
	if !IsNil(o.DataMergeTime) {
		toSerialize["dataMergeTime"] = o.DataMergeTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instanceType"] = o.InstanceType
	}
	if !IsNil(o.InstanceRole) {
		toSerialize["instanceRole"] = o.InstanceRole
	}
	if !IsNil(o.PrimaryInstanceId) {
		toSerialize["primaryInstanceId"] = o.PrimaryInstanceId
	}
	if !IsNil(o.StandbyInstanceIds) {
		toSerialize["standbyInstanceIds"] = o.StandbyInstanceIds
	}
	if !IsNil(o.SaleChannel) {
		toSerialize["saleChannel"] = o.SaleChannel
	}
	if !IsNil(o.ManagementMode) {
		toSerialize["managementMode"] = o.ManagementMode
	}
	if !IsNil(o.TagList) {
		toSerialize["tagList"] = o.TagList
	}
	if !IsNil(o.NodeNum) {
		toSerialize["nodeNum"] = o.NodeNum
	}
	if !IsNil(o.CpuArchitecture) {
		toSerialize["cpuArchitecture"] = o.CpuArchitecture
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.StopTime) {
		toSerialize["stopTime"] = o.StopTime
	}
	return toSerialize, nil
}

type NullableDescribeInstanceResponseV2OpenAPI struct {
	value *DescribeInstanceResponseV2OpenAPI
	isSet bool
}

func (v NullableDescribeInstanceResponseV2OpenAPI) Get() *DescribeInstanceResponseV2OpenAPI {
	return v.value
}

func (v *NullableDescribeInstanceResponseV2OpenAPI) Set(val *DescribeInstanceResponseV2OpenAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeInstanceResponseV2OpenAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeInstanceResponseV2OpenAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeInstanceResponseV2OpenAPI(val *DescribeInstanceResponseV2OpenAPI) *NullableDescribeInstanceResponseV2OpenAPI {
	return &NullableDescribeInstanceResponseV2OpenAPI{value: val, isSet: true}
}

func (v NullableDescribeInstanceResponseV2OpenAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeInstanceResponseV2OpenAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


