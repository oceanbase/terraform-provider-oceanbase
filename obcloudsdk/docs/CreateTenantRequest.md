# CreateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** |  | [optional] 
**TenantName** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TenantMode** | Pointer to **string** |  | [optional] 
**PrimaryZone** | Pointer to **string** |  | [optional] 
**Charset** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Whitelist** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**DiskSize** | Pointer to **int32** |  | [optional] 
**UnitNum** | Pointer to **int32** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**EnablePublicLink** | Pointer to **bool** |  | [optional] 
**CreateParams** | Pointer to **map[string]string** |  | [optional] 
**TenantCompatibilityModeEnum** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTenantRequest

`func NewCreateTenantRequest() *CreateTenantRequest`

NewCreateTenantRequest instantiates a new CreateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantRequestWithDefaults

`func NewCreateTenantRequestWithDefaults() *CreateTenantRequest`

NewCreateTenantRequestWithDefaults instantiates a new CreateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *CreateTenantRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateTenantRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateTenantRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CreateTenantRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetTenantName

`func (o *CreateTenantRequest) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *CreateTenantRequest) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *CreateTenantRequest) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *CreateTenantRequest) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetTenantId

`func (o *CreateTenantRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateTenantRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateTenantRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateTenantRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantMode

`func (o *CreateTenantRequest) GetTenantMode() string`

GetTenantMode returns the TenantMode field if non-nil, zero value otherwise.

### GetTenantModeOk

`func (o *CreateTenantRequest) GetTenantModeOk() (*string, bool)`

GetTenantModeOk returns a tuple with the TenantMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMode

`func (o *CreateTenantRequest) SetTenantMode(v string)`

SetTenantMode sets TenantMode field to given value.

### HasTenantMode

`func (o *CreateTenantRequest) HasTenantMode() bool`

HasTenantMode returns a boolean if a field has been set.

### GetPrimaryZone

`func (o *CreateTenantRequest) GetPrimaryZone() string`

GetPrimaryZone returns the PrimaryZone field if non-nil, zero value otherwise.

### GetPrimaryZoneOk

`func (o *CreateTenantRequest) GetPrimaryZoneOk() (*string, bool)`

GetPrimaryZoneOk returns a tuple with the PrimaryZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryZone

`func (o *CreateTenantRequest) SetPrimaryZone(v string)`

SetPrimaryZone sets PrimaryZone field to given value.

### HasPrimaryZone

`func (o *CreateTenantRequest) HasPrimaryZone() bool`

HasPrimaryZone returns a boolean if a field has been set.

### GetCharset

`func (o *CreateTenantRequest) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *CreateTenantRequest) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *CreateTenantRequest) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *CreateTenantRequest) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTenantRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTenantRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTenantRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTenantRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWhitelist

`func (o *CreateTenantRequest) GetWhitelist() string`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *CreateTenantRequest) GetWhitelistOk() (*string, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *CreateTenantRequest) SetWhitelist(v string)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *CreateTenantRequest) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.

### GetTimeZone

`func (o *CreateTenantRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateTenantRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateTenantRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CreateTenantRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCpu

`func (o *CreateTenantRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CreateTenantRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CreateTenantRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *CreateTenantRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *CreateTenantRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CreateTenantRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CreateTenantRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CreateTenantRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetDiskSize

`func (o *CreateTenantRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *CreateTenantRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *CreateTenantRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *CreateTenantRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetUnitNum

`func (o *CreateTenantRequest) GetUnitNum() int32`

GetUnitNum returns the UnitNum field if non-nil, zero value otherwise.

### GetUnitNumOk

`func (o *CreateTenantRequest) GetUnitNumOk() (*int32, bool)`

GetUnitNumOk returns a tuple with the UnitNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNum

`func (o *CreateTenantRequest) SetUnitNum(v int32)`

SetUnitNum sets UnitNum field to given value.

### HasUnitNum

`func (o *CreateTenantRequest) HasUnitNum() bool`

HasUnitNum returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateTenantRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateTenantRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateTenantRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateTenantRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnablePublicLink

`func (o *CreateTenantRequest) GetEnablePublicLink() bool`

GetEnablePublicLink returns the EnablePublicLink field if non-nil, zero value otherwise.

### GetEnablePublicLinkOk

`func (o *CreateTenantRequest) GetEnablePublicLinkOk() (*bool, bool)`

GetEnablePublicLinkOk returns a tuple with the EnablePublicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicLink

`func (o *CreateTenantRequest) SetEnablePublicLink(v bool)`

SetEnablePublicLink sets EnablePublicLink field to given value.

### HasEnablePublicLink

`func (o *CreateTenantRequest) HasEnablePublicLink() bool`

HasEnablePublicLink returns a boolean if a field has been set.

### GetCreateParams

`func (o *CreateTenantRequest) GetCreateParams() map[string]string`

GetCreateParams returns the CreateParams field if non-nil, zero value otherwise.

### GetCreateParamsOk

`func (o *CreateTenantRequest) GetCreateParamsOk() (*map[string]string, bool)`

GetCreateParamsOk returns a tuple with the CreateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateParams

`func (o *CreateTenantRequest) SetCreateParams(v map[string]string)`

SetCreateParams sets CreateParams field to given value.

### HasCreateParams

`func (o *CreateTenantRequest) HasCreateParams() bool`

HasCreateParams returns a boolean if a field has been set.

### GetTenantCompatibilityModeEnum

`func (o *CreateTenantRequest) GetTenantCompatibilityModeEnum() string`

GetTenantCompatibilityModeEnum returns the TenantCompatibilityModeEnum field if non-nil, zero value otherwise.

### GetTenantCompatibilityModeEnumOk

`func (o *CreateTenantRequest) GetTenantCompatibilityModeEnumOk() (*string, bool)`

GetTenantCompatibilityModeEnumOk returns a tuple with the TenantCompatibilityModeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCompatibilityModeEnum

`func (o *CreateTenantRequest) SetTenantCompatibilityModeEnum(v string)`

SetTenantCompatibilityModeEnum sets TenantCompatibilityModeEnum field to given value.

### HasTenantCompatibilityModeEnum

`func (o *CreateTenantRequest) HasTenantCompatibilityModeEnum() bool`

HasTenantCompatibilityModeEnum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


