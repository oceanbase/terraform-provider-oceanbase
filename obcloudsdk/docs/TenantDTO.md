# TenantDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**OcpTenantId** | Pointer to **int64** |  | [optional] 
**TenantName** | Pointer to **string** |  | [optional] 
**ManagementMode** | Pointer to **string** |  | [optional] 
**GmtCreate** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**TenantMode** | Pointer to **string** |  | [optional] 
**PrimaryRegion** | Pointer to **string** |  | [optional] 
**PrimaryZone** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Mem** | Pointer to **int32** |  | [optional] 
**UnitCpu** | Pointer to **int32** |  | [optional] 
**UnitMem** | Pointer to **int32** |  | [optional] 
**UnitNum** | Pointer to **int32** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**DeployType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TenantConnections** | Pointer to [**[]TenantConnectionDTO**](TenantConnectionDTO.md) |  | [optional] 
**TenantZones** | Pointer to [**[]TenantZoneDTO**](TenantZoneDTO.md) |  | [optional] 
**TenantResource** | Pointer to [**TenantResourceDTO**](TenantResourceDTO.md) |  | [optional] 
**EnableBinlogService** | Pointer to **bool** |  | [optional] 
**ReadonlyConnectionAddable** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Charset** | Pointer to **string** |  | [optional] 
**LowerCaseTableNamesParam** | Pointer to **int32** |  | [optional] 
**TagList** | Pointer to [**[]TagKeyValueDo**](TagKeyValueDo.md) |  | [optional] 
**TenantCompatibilityMode** | Pointer to **string** |  | [optional] 
**DataMergeTime** | Pointer to **string** |  | [optional] 
**EnableProxyPublicLink** | Pointer to **bool** |  | [optional] 

## Methods

### NewTenantDTO

`func NewTenantDTO() *TenantDTO`

NewTenantDTO instantiates a new TenantDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDTOWithDefaults

`func NewTenantDTOWithDefaults() *TenantDTO`

NewTenantDTOWithDefaults instantiates a new TenantDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantDTO) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantDTO) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantDTO) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantDTO) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetOcpTenantId

`func (o *TenantDTO) GetOcpTenantId() int64`

GetOcpTenantId returns the OcpTenantId field if non-nil, zero value otherwise.

### GetOcpTenantIdOk

`func (o *TenantDTO) GetOcpTenantIdOk() (*int64, bool)`

GetOcpTenantIdOk returns a tuple with the OcpTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcpTenantId

`func (o *TenantDTO) SetOcpTenantId(v int64)`

SetOcpTenantId sets OcpTenantId field to given value.

### HasOcpTenantId

`func (o *TenantDTO) HasOcpTenantId() bool`

HasOcpTenantId returns a boolean if a field has been set.

### GetTenantName

`func (o *TenantDTO) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *TenantDTO) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *TenantDTO) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *TenantDTO) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetManagementMode

`func (o *TenantDTO) GetManagementMode() string`

GetManagementMode returns the ManagementMode field if non-nil, zero value otherwise.

### GetManagementModeOk

`func (o *TenantDTO) GetManagementModeOk() (*string, bool)`

GetManagementModeOk returns a tuple with the ManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMode

`func (o *TenantDTO) SetManagementMode(v string)`

SetManagementMode sets ManagementMode field to given value.

### HasManagementMode

`func (o *TenantDTO) HasManagementMode() bool`

HasManagementMode returns a boolean if a field has been set.

### GetGmtCreate

`func (o *TenantDTO) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *TenantDTO) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *TenantDTO) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *TenantDTO) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetCreateTime

`func (o *TenantDTO) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *TenantDTO) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *TenantDTO) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *TenantDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetTenantMode

`func (o *TenantDTO) GetTenantMode() string`

GetTenantMode returns the TenantMode field if non-nil, zero value otherwise.

### GetTenantModeOk

`func (o *TenantDTO) GetTenantModeOk() (*string, bool)`

GetTenantModeOk returns a tuple with the TenantMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMode

`func (o *TenantDTO) SetTenantMode(v string)`

SetTenantMode sets TenantMode field to given value.

### HasTenantMode

`func (o *TenantDTO) HasTenantMode() bool`

HasTenantMode returns a boolean if a field has been set.

### GetPrimaryRegion

`func (o *TenantDTO) GetPrimaryRegion() string`

GetPrimaryRegion returns the PrimaryRegion field if non-nil, zero value otherwise.

### GetPrimaryRegionOk

`func (o *TenantDTO) GetPrimaryRegionOk() (*string, bool)`

GetPrimaryRegionOk returns a tuple with the PrimaryRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRegion

`func (o *TenantDTO) SetPrimaryRegion(v string)`

SetPrimaryRegion sets PrimaryRegion field to given value.

### HasPrimaryRegion

`func (o *TenantDTO) HasPrimaryRegion() bool`

HasPrimaryRegion returns a boolean if a field has been set.

### GetPrimaryZone

`func (o *TenantDTO) GetPrimaryZone() string`

GetPrimaryZone returns the PrimaryZone field if non-nil, zero value otherwise.

### GetPrimaryZoneOk

`func (o *TenantDTO) GetPrimaryZoneOk() (*string, bool)`

GetPrimaryZoneOk returns a tuple with the PrimaryZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryZone

`func (o *TenantDTO) SetPrimaryZone(v string)`

SetPrimaryZone sets PrimaryZone field to given value.

### HasPrimaryZone

`func (o *TenantDTO) HasPrimaryZone() bool`

HasPrimaryZone returns a boolean if a field has been set.

### GetStatus

`func (o *TenantDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TenantDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCpu

`func (o *TenantDTO) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TenantDTO) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TenantDTO) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TenantDTO) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMem

`func (o *TenantDTO) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *TenantDTO) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *TenantDTO) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *TenantDTO) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetUnitCpu

`func (o *TenantDTO) GetUnitCpu() int32`

GetUnitCpu returns the UnitCpu field if non-nil, zero value otherwise.

### GetUnitCpuOk

`func (o *TenantDTO) GetUnitCpuOk() (*int32, bool)`

GetUnitCpuOk returns a tuple with the UnitCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCpu

`func (o *TenantDTO) SetUnitCpu(v int32)`

SetUnitCpu sets UnitCpu field to given value.

### HasUnitCpu

`func (o *TenantDTO) HasUnitCpu() bool`

HasUnitCpu returns a boolean if a field has been set.

### GetUnitMem

`func (o *TenantDTO) GetUnitMem() int32`

GetUnitMem returns the UnitMem field if non-nil, zero value otherwise.

### GetUnitMemOk

`func (o *TenantDTO) GetUnitMemOk() (*int32, bool)`

GetUnitMemOk returns a tuple with the UnitMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitMem

`func (o *TenantDTO) SetUnitMem(v int32)`

SetUnitMem sets UnitMem field to given value.

### HasUnitMem

`func (o *TenantDTO) HasUnitMem() bool`

HasUnitMem returns a boolean if a field has been set.

### GetUnitNum

`func (o *TenantDTO) GetUnitNum() int32`

GetUnitNum returns the UnitNum field if non-nil, zero value otherwise.

### GetUnitNumOk

`func (o *TenantDTO) GetUnitNumOk() (*int32, bool)`

GetUnitNumOk returns a tuple with the UnitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNum

`func (o *TenantDTO) SetUnitNum(v int32)`

SetUnitNum sets UnitNum field to given value.

### HasUnitNum

`func (o *TenantDTO) HasUnitNum() bool`

HasUnitNum returns a boolean if a field has been set.

### GetVpcId

`func (o *TenantDTO) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *TenantDTO) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *TenantDTO) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *TenantDTO) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetDeployType

`func (o *TenantDTO) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *TenantDTO) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *TenantDTO) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *TenantDTO) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetDescription

`func (o *TenantDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenantConnections

`func (o *TenantDTO) GetTenantConnections() []TenantConnectionDTO`

GetTenantConnections returns the TenantConnections field if non-nil, zero value otherwise.

### GetTenantConnectionsOk

`func (o *TenantDTO) GetTenantConnectionsOk() (*[]TenantConnectionDTO, bool)`

GetTenantConnectionsOk returns a tuple with the TenantConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantConnections

`func (o *TenantDTO) SetTenantConnections(v []TenantConnectionDTO)`

SetTenantConnections sets TenantConnections field to given value.

### HasTenantConnections

`func (o *TenantDTO) HasTenantConnections() bool`

HasTenantConnections returns a boolean if a field has been set.

### GetTenantZones

`func (o *TenantDTO) GetTenantZones() []TenantZoneDTO`

GetTenantZones returns the TenantZones field if non-nil, zero value otherwise.

### GetTenantZonesOk

`func (o *TenantDTO) GetTenantZonesOk() (*[]TenantZoneDTO, bool)`

GetTenantZonesOk returns a tuple with the TenantZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantZones

`func (o *TenantDTO) SetTenantZones(v []TenantZoneDTO)`

SetTenantZones sets TenantZones field to given value.

### HasTenantZones

`func (o *TenantDTO) HasTenantZones() bool`

HasTenantZones returns a boolean if a field has been set.

### GetTenantResource

`func (o *TenantDTO) GetTenantResource() TenantResourceDTO`

GetTenantResource returns the TenantResource field if non-nil, zero value otherwise.

### GetTenantResourceOk

`func (o *TenantDTO) GetTenantResourceOk() (*TenantResourceDTO, bool)`

GetTenantResourceOk returns a tuple with the TenantResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantResource

`func (o *TenantDTO) SetTenantResource(v TenantResourceDTO)`

SetTenantResource sets TenantResource field to given value.

### HasTenantResource

`func (o *TenantDTO) HasTenantResource() bool`

HasTenantResource returns a boolean if a field has been set.

### GetEnableBinlogService

`func (o *TenantDTO) GetEnableBinlogService() bool`

GetEnableBinlogService returns the EnableBinlogService field if non-nil, zero value otherwise.

### GetEnableBinlogServiceOk

`func (o *TenantDTO) GetEnableBinlogServiceOk() (*bool, bool)`

GetEnableBinlogServiceOk returns a tuple with the EnableBinlogService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBinlogService

`func (o *TenantDTO) SetEnableBinlogService(v bool)`

SetEnableBinlogService sets EnableBinlogService field to given value.

### HasEnableBinlogService

`func (o *TenantDTO) HasEnableBinlogService() bool`

HasEnableBinlogService returns a boolean if a field has been set.

### GetReadonlyConnectionAddable

`func (o *TenantDTO) GetReadonlyConnectionAddable() string`

GetReadonlyConnectionAddable returns the ReadonlyConnectionAddable field if non-nil, zero value otherwise.

### GetReadonlyConnectionAddableOk

`func (o *TenantDTO) GetReadonlyConnectionAddableOk() (*string, bool)`

GetReadonlyConnectionAddableOk returns a tuple with the ReadonlyConnectionAddable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyConnectionAddable

`func (o *TenantDTO) SetReadonlyConnectionAddable(v string)`

SetReadonlyConnectionAddable sets ReadonlyConnectionAddable field to given value.

### HasReadonlyConnectionAddable

`func (o *TenantDTO) HasReadonlyConnectionAddable() bool`

HasReadonlyConnectionAddable returns a boolean if a field has been set.

### GetTimeZone

`func (o *TenantDTO) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TenantDTO) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TenantDTO) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TenantDTO) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCharset

`func (o *TenantDTO) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *TenantDTO) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *TenantDTO) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *TenantDTO) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetLowerCaseTableNamesParam

`func (o *TenantDTO) GetLowerCaseTableNamesParam() int32`

GetLowerCaseTableNamesParam returns the LowerCaseTableNamesParam field if non-nil, zero value otherwise.

### GetLowerCaseTableNamesParamOk

`func (o *TenantDTO) GetLowerCaseTableNamesParamOk() (*int32, bool)`

GetLowerCaseTableNamesParamOk returns a tuple with the LowerCaseTableNamesParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerCaseTableNamesParam

`func (o *TenantDTO) SetLowerCaseTableNamesParam(v int32)`

SetLowerCaseTableNamesParam sets LowerCaseTableNamesParam field to given value.

### HasLowerCaseTableNamesParam

`func (o *TenantDTO) HasLowerCaseTableNamesParam() bool`

HasLowerCaseTableNamesParam returns a boolean if a field has been set.

### GetTagList

`func (o *TenantDTO) GetTagList() []TagKeyValueDo`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *TenantDTO) GetTagListOk() (*[]TagKeyValueDo, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *TenantDTO) SetTagList(v []TagKeyValueDo)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *TenantDTO) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetTenantCompatibilityMode

`func (o *TenantDTO) GetTenantCompatibilityMode() string`

GetTenantCompatibilityMode returns the TenantCompatibilityMode field if non-nil, zero value otherwise.

### GetTenantCompatibilityModeOk

`func (o *TenantDTO) GetTenantCompatibilityModeOk() (*string, bool)`

GetTenantCompatibilityModeOk returns a tuple with the TenantCompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCompatibilityMode

`func (o *TenantDTO) SetTenantCompatibilityMode(v string)`

SetTenantCompatibilityMode sets TenantCompatibilityMode field to given value.

### HasTenantCompatibilityMode

`func (o *TenantDTO) HasTenantCompatibilityMode() bool`

HasTenantCompatibilityMode returns a boolean if a field has been set.

### GetDataMergeTime

`func (o *TenantDTO) GetDataMergeTime() string`

GetDataMergeTime returns the DataMergeTime field if non-nil, zero value otherwise.

### GetDataMergeTimeOk

`func (o *TenantDTO) GetDataMergeTimeOk() (*string, bool)`

GetDataMergeTimeOk returns a tuple with the DataMergeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMergeTime

`func (o *TenantDTO) SetDataMergeTime(v string)`

SetDataMergeTime sets DataMergeTime field to given value.

### HasDataMergeTime

`func (o *TenantDTO) HasDataMergeTime() bool`

HasDataMergeTime returns a boolean if a field has been set.

### GetEnableProxyPublicLink

`func (o *TenantDTO) GetEnableProxyPublicLink() bool`

GetEnableProxyPublicLink returns the EnableProxyPublicLink field if non-nil, zero value otherwise.

### GetEnableProxyPublicLinkOk

`func (o *TenantDTO) GetEnableProxyPublicLinkOk() (*bool, bool)`

GetEnableProxyPublicLinkOk returns a tuple with the EnableProxyPublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyPublicLink

`func (o *TenantDTO) SetEnableProxyPublicLink(v bool)`

SetEnableProxyPublicLink sets EnableProxyPublicLink field to given value.

### HasEnableProxyPublicLink

`func (o *TenantDTO) HasEnableProxyPublicLink() bool`

HasEnableProxyPublicLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


