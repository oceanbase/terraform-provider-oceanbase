# DescribeInstanceResponseV2OpenAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**InstanceClass** | Pointer to **string** |  | [optional] 
**Series** | Pointer to **string** |  | [optional] 
**PayType** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **string** |  | [optional] 
**AvailableZones** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ObRpmVersion** | Pointer to **string** |  | [optional] 
**DeployType** | Pointer to **string** |  | [optional] 
**DiskType** | Pointer to **string** |  | [optional] 
**DeployMode** | Pointer to **string** |  | [optional] 
**ReplicaMode** | Pointer to **string** |  | [optional] 
**DataMergeTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**ResourceDoV2**](ResourceDoV2.md) |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InstanceRole** | Pointer to **string** |  | [optional] 
**PrimaryInstanceId** | Pointer to **string** |  | [optional] 
**StandbyInstanceIds** | Pointer to **[]string** |  | [optional] 
**SaleChannel** | Pointer to **string** |  | [optional] 
**ManagementMode** | Pointer to **string** |  | [optional] 
**TagList** | Pointer to [**[]TagKeyValueDo**](TagKeyValueDo.md) |  | [optional] 
**NodeNum** | Pointer to **int32** |  | [optional] 
**CpuArchitecture** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**StopTime** | Pointer to **string** |  | [optional] 

## Methods

### NewDescribeInstanceResponseV2OpenAPI

`func NewDescribeInstanceResponseV2OpenAPI() *DescribeInstanceResponseV2OpenAPI`

NewDescribeInstanceResponseV2OpenAPI instantiates a new DescribeInstanceResponseV2OpenAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInstanceResponseV2OpenAPIWithDefaults

`func NewDescribeInstanceResponseV2OpenAPIWithDefaults() *DescribeInstanceResponseV2OpenAPI`

NewDescribeInstanceResponseV2OpenAPIWithDefaults instantiates a new DescribeInstanceResponseV2OpenAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *DescribeInstanceResponseV2OpenAPI) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeInstanceResponseV2OpenAPI) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DescribeInstanceResponseV2OpenAPI) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVpcId

`func (o *DescribeInstanceResponseV2OpenAPI) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *DescribeInstanceResponseV2OpenAPI) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *DescribeInstanceResponseV2OpenAPI) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetInstanceId

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceName

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetInstanceClass

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.

### HasInstanceClass

`func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceClass() bool`

HasInstanceClass returns a boolean if a field has been set.

### GetSeries

`func (o *DescribeInstanceResponseV2OpenAPI) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DescribeInstanceResponseV2OpenAPI) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *DescribeInstanceResponseV2OpenAPI) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetPayType

`func (o *DescribeInstanceResponseV2OpenAPI) GetPayType() string`

GetPayType returns the PayType field if non-nil, zero value otherwise.

### GetPayTypeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetPayTypeOk() (*string, bool)`

GetPayTypeOk returns a tuple with the PayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayType

`func (o *DescribeInstanceResponseV2OpenAPI) SetPayType(v string)`

SetPayType sets PayType field to given value.

### HasPayType

`func (o *DescribeInstanceResponseV2OpenAPI) HasPayType() bool`

HasPayType returns a boolean if a field has been set.

### GetCreateTime

`func (o *DescribeInstanceResponseV2OpenAPI) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DescribeInstanceResponseV2OpenAPI) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DescribeInstanceResponseV2OpenAPI) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *DescribeInstanceResponseV2OpenAPI) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *DescribeInstanceResponseV2OpenAPI) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *DescribeInstanceResponseV2OpenAPI) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetAvailableZones

`func (o *DescribeInstanceResponseV2OpenAPI) GetAvailableZones() []string`

GetAvailableZones returns the AvailableZones field if non-nil, zero value otherwise.

### GetAvailableZonesOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetAvailableZonesOk() (*[]string, bool)`

GetAvailableZonesOk returns a tuple with the AvailableZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableZones

`func (o *DescribeInstanceResponseV2OpenAPI) SetAvailableZones(v []string)`

SetAvailableZones sets AvailableZones field to given value.

### HasAvailableZones

`func (o *DescribeInstanceResponseV2OpenAPI) HasAvailableZones() bool`

HasAvailableZones returns a boolean if a field has been set.

### GetVersion

`func (o *DescribeInstanceResponseV2OpenAPI) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DescribeInstanceResponseV2OpenAPI) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DescribeInstanceResponseV2OpenAPI) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetObRpmVersion

`func (o *DescribeInstanceResponseV2OpenAPI) GetObRpmVersion() string`

GetObRpmVersion returns the ObRpmVersion field if non-nil, zero value otherwise.

### GetObRpmVersionOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetObRpmVersionOk() (*string, bool)`

GetObRpmVersionOk returns a tuple with the ObRpmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObRpmVersion

`func (o *DescribeInstanceResponseV2OpenAPI) SetObRpmVersion(v string)`

SetObRpmVersion sets ObRpmVersion field to given value.

### HasObRpmVersion

`func (o *DescribeInstanceResponseV2OpenAPI) HasObRpmVersion() bool`

HasObRpmVersion returns a boolean if a field has been set.

### GetDeployType

`func (o *DescribeInstanceResponseV2OpenAPI) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DescribeInstanceResponseV2OpenAPI) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *DescribeInstanceResponseV2OpenAPI) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetDiskType

`func (o *DescribeInstanceResponseV2OpenAPI) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *DescribeInstanceResponseV2OpenAPI) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *DescribeInstanceResponseV2OpenAPI) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetDeployMode

`func (o *DescribeInstanceResponseV2OpenAPI) GetDeployMode() string`

GetDeployMode returns the DeployMode field if non-nil, zero value otherwise.

### GetDeployModeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetDeployModeOk() (*string, bool)`

GetDeployModeOk returns a tuple with the DeployMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployMode

`func (o *DescribeInstanceResponseV2OpenAPI) SetDeployMode(v string)`

SetDeployMode sets DeployMode field to given value.

### HasDeployMode

`func (o *DescribeInstanceResponseV2OpenAPI) HasDeployMode() bool`

HasDeployMode returns a boolean if a field has been set.

### GetReplicaMode

`func (o *DescribeInstanceResponseV2OpenAPI) GetReplicaMode() string`

GetReplicaMode returns the ReplicaMode field if non-nil, zero value otherwise.

### GetReplicaModeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetReplicaModeOk() (*string, bool)`

GetReplicaModeOk returns a tuple with the ReplicaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaMode

`func (o *DescribeInstanceResponseV2OpenAPI) SetReplicaMode(v string)`

SetReplicaMode sets ReplicaMode field to given value.

### HasReplicaMode

`func (o *DescribeInstanceResponseV2OpenAPI) HasReplicaMode() bool`

HasReplicaMode returns a boolean if a field has been set.

### GetDataMergeTime

`func (o *DescribeInstanceResponseV2OpenAPI) GetDataMergeTime() string`

GetDataMergeTime returns the DataMergeTime field if non-nil, zero value otherwise.

### GetDataMergeTimeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetDataMergeTimeOk() (*string, bool)`

GetDataMergeTimeOk returns a tuple with the DataMergeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMergeTime

`func (o *DescribeInstanceResponseV2OpenAPI) SetDataMergeTime(v string)`

SetDataMergeTime sets DataMergeTime field to given value.

### HasDataMergeTime

`func (o *DescribeInstanceResponseV2OpenAPI) HasDataMergeTime() bool`

HasDataMergeTime returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeInstanceResponseV2OpenAPI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeInstanceResponseV2OpenAPI) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeInstanceResponseV2OpenAPI) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResource

`func (o *DescribeInstanceResponseV2OpenAPI) GetResource() ResourceDoV2`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetResourceOk() (*ResourceDoV2, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeInstanceResponseV2OpenAPI) SetResource(v ResourceDoV2)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DescribeInstanceResponseV2OpenAPI) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetCloudProvider

`func (o *DescribeInstanceResponseV2OpenAPI) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeInstanceResponseV2OpenAPI) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DescribeInstanceResponseV2OpenAPI) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetInstanceType

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceRole

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceRole() string`

GetInstanceRole returns the InstanceRole field if non-nil, zero value otherwise.

### GetInstanceRoleOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetInstanceRoleOk() (*string, bool)`

GetInstanceRoleOk returns a tuple with the InstanceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceRole

`func (o *DescribeInstanceResponseV2OpenAPI) SetInstanceRole(v string)`

SetInstanceRole sets InstanceRole field to given value.

### HasInstanceRole

`func (o *DescribeInstanceResponseV2OpenAPI) HasInstanceRole() bool`

HasInstanceRole returns a boolean if a field has been set.

### GetPrimaryInstanceId

`func (o *DescribeInstanceResponseV2OpenAPI) GetPrimaryInstanceId() string`

GetPrimaryInstanceId returns the PrimaryInstanceId field if non-nil, zero value otherwise.

### GetPrimaryInstanceIdOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetPrimaryInstanceIdOk() (*string, bool)`

GetPrimaryInstanceIdOk returns a tuple with the PrimaryInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceId

`func (o *DescribeInstanceResponseV2OpenAPI) SetPrimaryInstanceId(v string)`

SetPrimaryInstanceId sets PrimaryInstanceId field to given value.

### HasPrimaryInstanceId

`func (o *DescribeInstanceResponseV2OpenAPI) HasPrimaryInstanceId() bool`

HasPrimaryInstanceId returns a boolean if a field has been set.

### GetStandbyInstanceIds

`func (o *DescribeInstanceResponseV2OpenAPI) GetStandbyInstanceIds() []string`

GetStandbyInstanceIds returns the StandbyInstanceIds field if non-nil, zero value otherwise.

### GetStandbyInstanceIdsOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetStandbyInstanceIdsOk() (*[]string, bool)`

GetStandbyInstanceIdsOk returns a tuple with the StandbyInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInstanceIds

`func (o *DescribeInstanceResponseV2OpenAPI) SetStandbyInstanceIds(v []string)`

SetStandbyInstanceIds sets StandbyInstanceIds field to given value.

### HasStandbyInstanceIds

`func (o *DescribeInstanceResponseV2OpenAPI) HasStandbyInstanceIds() bool`

HasStandbyInstanceIds returns a boolean if a field has been set.

### GetSaleChannel

`func (o *DescribeInstanceResponseV2OpenAPI) GetSaleChannel() string`

GetSaleChannel returns the SaleChannel field if non-nil, zero value otherwise.

### GetSaleChannelOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetSaleChannelOk() (*string, bool)`

GetSaleChannelOk returns a tuple with the SaleChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleChannel

`func (o *DescribeInstanceResponseV2OpenAPI) SetSaleChannel(v string)`

SetSaleChannel sets SaleChannel field to given value.

### HasSaleChannel

`func (o *DescribeInstanceResponseV2OpenAPI) HasSaleChannel() bool`

HasSaleChannel returns a boolean if a field has been set.

### GetManagementMode

`func (o *DescribeInstanceResponseV2OpenAPI) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *DescribeInstanceResponseV2OpenAPI) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *DescribeInstanceResponseV2OpenAPI) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetTagList

`func (o *DescribeInstanceResponseV2OpenAPI) GetTagList() []TagKeyValueDo`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetTagListOk() (*[]TagKeyValueDo, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *DescribeInstanceResponseV2OpenAPI) SetTagList(v []TagKeyValueDo)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *DescribeInstanceResponseV2OpenAPI) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetNodeNum

`func (o *DescribeInstanceResponseV2OpenAPI) GetNodeNum() int32`

GetNodeNum returns the NodeNum field if non-nil, zero value otherwise.

### GetNodeNumOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetNodeNumOk() (*int32, bool)`

GetNodeNumOk returns a tuple with the NodeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNum

`func (o *DescribeInstanceResponseV2OpenAPI) SetNodeNum(v int32)`

SetNodeNum sets NodeNum field to given value.

### HasNodeNum

`func (o *DescribeInstanceResponseV2OpenAPI) HasNodeNum() bool`

HasNodeNum returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *DescribeInstanceResponseV2OpenAPI) GetCpuArchitecture() string`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetCpuArchitectureOk() (*string, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *DescribeInstanceResponseV2OpenAPI) SetCpuArchitecture(v string)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *DescribeInstanceResponseV2OpenAPI) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeInstanceResponseV2OpenAPI) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeInstanceResponseV2OpenAPI) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DescribeInstanceResponseV2OpenAPI) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStopTime

`func (o *DescribeInstanceResponseV2OpenAPI) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *DescribeInstanceResponseV2OpenAPI) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *DescribeInstanceResponseV2OpenAPI) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *DescribeInstanceResponseV2OpenAPI) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


