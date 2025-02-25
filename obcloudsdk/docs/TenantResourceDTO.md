# TenantResourceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**TenantCpuDTO**](TenantCpuDTO.md) |  | [optional] 
**Memory** | Pointer to [**TenantMemoryDTO**](TenantMemoryDTO.md) |  | [optional] 
**Disk** | Pointer to [**TenantDiskDTO**](TenantDiskDTO.md) |  | [optional] 
**UnitNum** | Pointer to **int32** |  | [optional] 

## Methods

### NewTenantResourceDTO

`func NewTenantResourceDTO() *TenantResourceDTO`

NewTenantResourceDTO instantiates a new TenantResourceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantResourceDTOWithDefaults

`func NewTenantResourceDTOWithDefaults() *TenantResourceDTO`

NewTenantResourceDTOWithDefaults instantiates a new TenantResourceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *TenantResourceDTO) GetCpu() TenantCpuDTO`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TenantResourceDTO) GetCpuOk() (*TenantCpuDTO, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TenantResourceDTO) SetCpu(v TenantCpuDTO)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TenantResourceDTO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *TenantResourceDTO) GetMemory() TenantMemoryDTO`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TenantResourceDTO) GetMemoryOk() (*TenantMemoryDTO, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TenantResourceDTO) SetMemory(v TenantMemoryDTO)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TenantResourceDTO) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDisk

`func (o *TenantResourceDTO) GetDisk() TenantDiskDTO`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *TenantResourceDTO) GetDiskOk() (*TenantDiskDTO, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *TenantResourceDTO) SetDisk(v TenantDiskDTO)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *TenantResourceDTO) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetUnitNum

`func (o *TenantResourceDTO) GetUnitNum() int32`

GetUnitNum returns the UnitNum field if non-nil, zero value otherwise.

### GetUnitNumOk

`func (o *TenantResourceDTO) GetUnitNumOk() (*int32, bool)`

GetUnitNumOk returns a tuple with the UnitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNum

`func (o *TenantResourceDTO) SetUnitNum(v int32)`

SetUnitNum sets UnitNum field to given value.

### HasUnitNum

`func (o *TenantResourceDTO) HasUnitNum() bool`

HasUnitNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


