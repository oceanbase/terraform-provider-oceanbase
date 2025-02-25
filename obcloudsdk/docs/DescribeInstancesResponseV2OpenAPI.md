# DescribeInstancesResponseV2OpenAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**InstanceClass** | Pointer to **string** |  | [optional] 
**Series** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**ObRpmVersion** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Mem** | Pointer to **int64** |  | [optional] 
**UsedDiskSize** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**ResourceDoV2**](ResourceDoV2.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeployMode** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**PayType** | Pointer to **string** |  | [optional] 
**AvailableZones** | Pointer to **[]string** |  | [optional] 
**DeployType** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**ManagementMode** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**DiskType** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int64** |  | [optional] 
**InstanceRole** | Pointer to **string** |  | [optional] 
**PrimaryInstanceId** | Pointer to **string** |  | [optional] 
**StandbyInstanceIds** | Pointer to **[]string** |  | [optional] 
**StopTime** | Pointer to **string** |  | [optional] 
**SaleChannel** | Pointer to **string** |  | [optional] 
**TagList** | Pointer to [**[]TagKeyValueDo**](TagKeyValueDo.md) |  | [optional] 
**CpuArchitecture** | Pointer to **string** |  | [optional] 

## Methods

### NewDescribeInstancesResponseV2OpenAPI

`func NewDescribeInstancesResponseV2OpenAPI() *DescribeInstancesResponseV2OpenAPI`

NewDescribeInstancesResponseV2OpenAPI instantiates a new DescribeInstancesResponseV2OpenAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeInstancesResponseV2OpenAPIWithDefaults

`func NewDescribeInstancesResponseV2OpenAPIWithDefaults() *DescribeInstancesResponseV2OpenAPI`

NewDescribeInstancesResponseV2OpenAPIWithDefaults instantiates a new DescribeInstancesResponseV2OpenAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *DescribeInstancesResponseV2OpenAPI) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeInstancesResponseV2OpenAPI) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DescribeInstancesResponseV2OpenAPI) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetInstanceId

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DescribeInstancesResponseV2OpenAPI) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *DescribeInstancesResponseV2OpenAPI) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceName

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *DescribeInstancesResponseV2OpenAPI) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *DescribeInstancesResponseV2OpenAPI) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetInstanceClass

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *DescribeInstancesResponseV2OpenAPI) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.

### HasInstanceClass

`func (o *DescribeInstancesResponseV2OpenAPI) HasInstanceClass() bool`

HasInstanceClass returns a boolean if a field has been set.

### GetSeries

`func (o *DescribeInstancesResponseV2OpenAPI) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DescribeInstancesResponseV2OpenAPI) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *DescribeInstancesResponseV2OpenAPI) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetCreateTime

`func (o *DescribeInstancesResponseV2OpenAPI) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DescribeInstancesResponseV2OpenAPI) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DescribeInstancesResponseV2OpenAPI) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetObRpmVersion

`func (o *DescribeInstancesResponseV2OpenAPI) GetObRpmVersion() string`

GetObRpmVersion returns the ObRpmVersion field if non-nil, zero value otherwise.

### GetObRpmVersionOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetObRpmVersionOk() (*string, bool)`

GetObRpmVersionOk returns a tuple with the ObRpmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObRpmVersion

`func (o *DescribeInstancesResponseV2OpenAPI) SetObRpmVersion(v string)`

SetObRpmVersion sets ObRpmVersion field to given value.

### HasObRpmVersion

`func (o *DescribeInstancesResponseV2OpenAPI) HasObRpmVersion() bool`

HasObRpmVersion returns a boolean if a field has been set.

### GetVersion

`func (o *DescribeInstancesResponseV2OpenAPI) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DescribeInstancesResponseV2OpenAPI) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DescribeInstancesResponseV2OpenAPI) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCpu

`func (o *DescribeInstancesResponseV2OpenAPI) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DescribeInstancesResponseV2OpenAPI) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DescribeInstancesResponseV2OpenAPI) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMem

`func (o *DescribeInstancesResponseV2OpenAPI) GetMem() int64`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetMemOk() (*int64, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *DescribeInstancesResponseV2OpenAPI) SetMem(v int64)`

SetMem sets Mem field to given value.

### HasMem

`func (o *DescribeInstancesResponseV2OpenAPI) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetUsedDiskSize

`func (o *DescribeInstancesResponseV2OpenAPI) GetUsedDiskSize() int64`

GetUsedDiskSize returns the UsedDiskSize field if non-nil, zero value otherwise.

### GetUsedDiskSizeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetUsedDiskSizeOk() (*int64, bool)`

GetUsedDiskSizeOk returns a tuple with the UsedDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDiskSize

`func (o *DescribeInstancesResponseV2OpenAPI) SetUsedDiskSize(v int64)`

SetUsedDiskSize sets UsedDiskSize field to given value.

### HasUsedDiskSize

`func (o *DescribeInstancesResponseV2OpenAPI) HasUsedDiskSize() bool`

HasUsedDiskSize returns a boolean if a field has been set.

### GetState

`func (o *DescribeInstancesResponseV2OpenAPI) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DescribeInstancesResponseV2OpenAPI) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DescribeInstancesResponseV2OpenAPI) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResource

`func (o *DescribeInstancesResponseV2OpenAPI) GetResource() ResourceDoV2`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetResourceOk() (*ResourceDoV2, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeInstancesResponseV2OpenAPI) SetResource(v ResourceDoV2)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DescribeInstancesResponseV2OpenAPI) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeInstancesResponseV2OpenAPI) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeInstancesResponseV2OpenAPI) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeInstancesResponseV2OpenAPI) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployMode

`func (o *DescribeInstancesResponseV2OpenAPI) GetDeployMode() string`

GetDeployMode returns the DeployMode field if non-nil, zero value otherwise.

### GetDeployModeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetDeployModeOk() (*string, bool)`

GetDeployModeOk returns a tuple with the DeployMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployMode

`func (o *DescribeInstancesResponseV2OpenAPI) SetDeployMode(v string)`

SetDeployMode sets DeployMode field to given value.

### HasDeployMode

`func (o *DescribeInstancesResponseV2OpenAPI) HasDeployMode() bool`

HasDeployMode returns a boolean if a field has been set.

### GetVpcId

`func (o *DescribeInstancesResponseV2OpenAPI) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *DescribeInstancesResponseV2OpenAPI) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *DescribeInstancesResponseV2OpenAPI) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetPayType

`func (o *DescribeInstancesResponseV2OpenAPI) GetPayType() string`

GetPayType returns the PayType field if non-nil, zero value otherwise.

### GetPayTypeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetPayTypeOk() (*string, bool)`

GetPayTypeOk returns a tuple with the PayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayType

`func (o *DescribeInstancesResponseV2OpenAPI) SetPayType(v string)`

SetPayType sets PayType field to given value.

### HasPayType

`func (o *DescribeInstancesResponseV2OpenAPI) HasPayType() bool`

HasPayType returns a boolean if a field has been set.

### GetAvailableZones

`func (o *DescribeInstancesResponseV2OpenAPI) GetAvailableZones() []string`

GetAvailableZones returns the AvailableZones field if non-nil, zero value otherwise.

### GetAvailableZonesOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetAvailableZonesOk() (*[]string, bool)`

GetAvailableZonesOk returns a tuple with the AvailableZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableZones

`func (o *DescribeInstancesResponseV2OpenAPI) SetAvailableZones(v []string)`

SetAvailableZones sets AvailableZones field to given value.

### HasAvailableZones

`func (o *DescribeInstancesResponseV2OpenAPI) HasAvailableZones() bool`

HasAvailableZones returns a boolean if a field has been set.

### GetDeployType

`func (o *DescribeInstancesResponseV2OpenAPI) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DescribeInstancesResponseV2OpenAPI) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *DescribeInstancesResponseV2OpenAPI) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetExpireTime

`func (o *DescribeInstancesResponseV2OpenAPI) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *DescribeInstancesResponseV2OpenAPI) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *DescribeInstancesResponseV2OpenAPI) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetCloudProvider

`func (o *DescribeInstancesResponseV2OpenAPI) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeInstancesResponseV2OpenAPI) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DescribeInstancesResponseV2OpenAPI) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetManagementMode

`func (o *DescribeInstancesResponseV2OpenAPI) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *DescribeInstancesResponseV2OpenAPI) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *DescribeInstancesResponseV2OpenAPI) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetInstanceType

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DescribeInstancesResponseV2OpenAPI) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DescribeInstancesResponseV2OpenAPI) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetDiskType

`func (o *DescribeInstancesResponseV2OpenAPI) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *DescribeInstancesResponseV2OpenAPI) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *DescribeInstancesResponseV2OpenAPI) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetDiskSize

`func (o *DescribeInstancesResponseV2OpenAPI) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *DescribeInstancesResponseV2OpenAPI) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *DescribeInstancesResponseV2OpenAPI) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetInstanceRole

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceRole() string`

GetInstanceRole returns the InstanceRole field if non-nil, zero value otherwise.

### GetInstanceRoleOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetInstanceRoleOk() (*string, bool)`

GetInstanceRoleOk returns a tuple with the InstanceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceRole

`func (o *DescribeInstancesResponseV2OpenAPI) SetInstanceRole(v string)`

SetInstanceRole sets InstanceRole field to given value.

### HasInstanceRole

`func (o *DescribeInstancesResponseV2OpenAPI) HasInstanceRole() bool`

HasInstanceRole returns a boolean if a field has been set.

### GetPrimaryInstanceId

`func (o *DescribeInstancesResponseV2OpenAPI) GetPrimaryInstanceId() string`

GetPrimaryInstanceId returns the PrimaryInstanceId field if non-nil, zero value otherwise.

### GetPrimaryInstanceIdOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetPrimaryInstanceIdOk() (*string, bool)`

GetPrimaryInstanceIdOk returns a tuple with the PrimaryInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceId

`func (o *DescribeInstancesResponseV2OpenAPI) SetPrimaryInstanceId(v string)`

SetPrimaryInstanceId sets PrimaryInstanceId field to given value.

### HasPrimaryInstanceId

`func (o *DescribeInstancesResponseV2OpenAPI) HasPrimaryInstanceId() bool`

HasPrimaryInstanceId returns a boolean if a field has been set.

### GetStandbyInstanceIds

`func (o *DescribeInstancesResponseV2OpenAPI) GetStandbyInstanceIds() []string`

GetStandbyInstanceIds returns the StandbyInstanceIds field if non-nil, zero value otherwise.

### GetStandbyInstanceIdsOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetStandbyInstanceIdsOk() (*[]string, bool)`

GetStandbyInstanceIdsOk returns a tuple with the StandbyInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInstanceIds

`func (o *DescribeInstancesResponseV2OpenAPI) SetStandbyInstanceIds(v []string)`

SetStandbyInstanceIds sets StandbyInstanceIds field to given value.

### HasStandbyInstanceIds

`func (o *DescribeInstancesResponseV2OpenAPI) HasStandbyInstanceIds() bool`

HasStandbyInstanceIds returns a boolean if a field has been set.

### GetStopTime

`func (o *DescribeInstancesResponseV2OpenAPI) GetStopTime() string`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetStopTimeOk() (*string, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *DescribeInstancesResponseV2OpenAPI) SetStopTime(v string)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *DescribeInstancesResponseV2OpenAPI) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.

### GetSaleChannel

`func (o *DescribeInstancesResponseV2OpenAPI) GetSaleChannel() string`

GetSaleChannel returns the SaleChannel field if non-nil, zero value otherwise.

### GetSaleChannelOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetSaleChannelOk() (*string, bool)`

GetSaleChannelOk returns a tuple with the SaleChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleChannel

`func (o *DescribeInstancesResponseV2OpenAPI) SetSaleChannel(v string)`

SetSaleChannel sets SaleChannel field to given value.

### HasSaleChannel

`func (o *DescribeInstancesResponseV2OpenAPI) HasSaleChannel() bool`

HasSaleChannel returns a boolean if a field has been set.

### GetTagList

`func (o *DescribeInstancesResponseV2OpenAPI) GetTagList() []TagKeyValueDo`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetTagListOk() (*[]TagKeyValueDo, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *DescribeInstancesResponseV2OpenAPI) SetTagList(v []TagKeyValueDo)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *DescribeInstancesResponseV2OpenAPI) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *DescribeInstancesResponseV2OpenAPI) GetCpuArchitecture() string`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *DescribeInstancesResponseV2OpenAPI) GetCpuArchitectureOk() (*string, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *DescribeInstancesResponseV2OpenAPI) SetCpuArchitecture(v string)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *DescribeInstancesResponseV2OpenAPI) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


