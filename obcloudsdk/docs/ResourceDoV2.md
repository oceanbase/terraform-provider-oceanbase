# ResourceDoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**CpuDoV2**](CpuDoV2.md) |  | [optional] 
**Memory** | Pointer to [**MemoryDoV2**](MemoryDoV2.md) |  | [optional] 
**Disk** | Pointer to [**DiskDoV2**](DiskDoV2.md) |  | [optional] 

## Methods

### NewResourceDoV2

`func NewResourceDoV2() *ResourceDoV2`

NewResourceDoV2 instantiates a new ResourceDoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDoV2WithDefaults

`func NewResourceDoV2WithDefaults() *ResourceDoV2`

NewResourceDoV2WithDefaults instantiates a new ResourceDoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceDoV2) GetCpu() CpuDoV2`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceDoV2) GetCpuOk() (*CpuDoV2, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceDoV2) SetCpu(v CpuDoV2)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ResourceDoV2) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ResourceDoV2) GetMemory() MemoryDoV2`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceDoV2) GetMemoryOk() (*MemoryDoV2, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceDoV2) SetMemory(v MemoryDoV2)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ResourceDoV2) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDisk

`func (o *ResourceDoV2) GetDisk() DiskDoV2`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ResourceDoV2) GetDiskOk() (*DiskDoV2, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ResourceDoV2) SetDisk(v DiskDoV2)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ResourceDoV2) HasDisk() bool`

HasDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


