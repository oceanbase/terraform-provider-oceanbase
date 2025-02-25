# MultiCloudCreateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **string** |  | [optional] 
**InstanceName** | Pointer to **string** |  | [optional] 
**CloudProvider** | **string** |  | 
**DiskSize** | **int32** |  | 
**Region** | **string** |  | 
**ObVersion** | **string** |  | 
**Zones** | **string** |  | 
**InstanceClass** | **string** |  | 
**ReplicaMode** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**SaleChannel** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 

## Methods

### NewMultiCloudCreateInstanceRequest

`func NewMultiCloudCreateInstanceRequest(cloudProvider string, diskSize int32, region string, obVersion string, zones string, instanceClass string, ) *MultiCloudCreateInstanceRequest`

NewMultiCloudCreateInstanceRequest instantiates a new MultiCloudCreateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudCreateInstanceRequestWithDefaults

`func NewMultiCloudCreateInstanceRequestWithDefaults() *MultiCloudCreateInstanceRequest`

NewMultiCloudCreateInstanceRequestWithDefaults instantiates a new MultiCloudCreateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *MultiCloudCreateInstanceRequest) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *MultiCloudCreateInstanceRequest) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *MultiCloudCreateInstanceRequest) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *MultiCloudCreateInstanceRequest) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetInstanceName

`func (o *MultiCloudCreateInstanceRequest) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MultiCloudCreateInstanceRequest) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MultiCloudCreateInstanceRequest) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *MultiCloudCreateInstanceRequest) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetCloudProvider

`func (o *MultiCloudCreateInstanceRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *MultiCloudCreateInstanceRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *MultiCloudCreateInstanceRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDiskSize

`func (o *MultiCloudCreateInstanceRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *MultiCloudCreateInstanceRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *MultiCloudCreateInstanceRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.


### GetRegion

`func (o *MultiCloudCreateInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MultiCloudCreateInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MultiCloudCreateInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetObVersion

`func (o *MultiCloudCreateInstanceRequest) GetObVersion() string`

GetObVersion returns the ObVersion field if non-nil, zero value otherwise.

### GetObVersionOk

`func (o *MultiCloudCreateInstanceRequest) GetObVersionOk() (*string, bool)`

GetObVersionOk returns a tuple with the ObVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObVersion

`func (o *MultiCloudCreateInstanceRequest) SetObVersion(v string)`

SetObVersion sets ObVersion field to given value.


### GetZones

`func (o *MultiCloudCreateInstanceRequest) GetZones() string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *MultiCloudCreateInstanceRequest) GetZonesOk() (*string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *MultiCloudCreateInstanceRequest) SetZones(v string)`

SetZones sets Zones field to given value.


### GetInstanceClass

`func (o *MultiCloudCreateInstanceRequest) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *MultiCloudCreateInstanceRequest) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *MultiCloudCreateInstanceRequest) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.


### GetReplicaMode

`func (o *MultiCloudCreateInstanceRequest) GetReplicaMode() string`

GetReplicaMode returns the ReplicaMode field if non-nil, zero value otherwise.

### GetReplicaModeOk

`func (o *MultiCloudCreateInstanceRequest) GetReplicaModeOk() (*string, bool)`

GetReplicaModeOk returns a tuple with the ReplicaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaMode

`func (o *MultiCloudCreateInstanceRequest) SetReplicaMode(v string)`

SetReplicaMode sets ReplicaMode field to given value.

### HasReplicaMode

`func (o *MultiCloudCreateInstanceRequest) HasReplicaMode() bool`

HasReplicaMode returns a boolean if a field has been set.

### GetDryRun

`func (o *MultiCloudCreateInstanceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *MultiCloudCreateInstanceRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *MultiCloudCreateInstanceRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *MultiCloudCreateInstanceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSaleChannel

`func (o *MultiCloudCreateInstanceRequest) GetSaleChannel() string`

GetSaleChannel returns the SaleChannel field if non-nil, zero value otherwise.

### GetSaleChannelOk

`func (o *MultiCloudCreateInstanceRequest) GetSaleChannelOk() (*string, bool)`

GetSaleChannelOk returns a tuple with the SaleChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleChannel

`func (o *MultiCloudCreateInstanceRequest) SetSaleChannel(v string)`

SetSaleChannel sets SaleChannel field to given value.

### HasSaleChannel

`func (o *MultiCloudCreateInstanceRequest) HasSaleChannel() bool`

HasSaleChannel returns a boolean if a field has been set.

### GetInstanceType

`func (o *MultiCloudCreateInstanceRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *MultiCloudCreateInstanceRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *MultiCloudCreateInstanceRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *MultiCloudCreateInstanceRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


