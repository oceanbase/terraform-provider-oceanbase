# TableInfoDo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableId** | Pointer to **int64** |  | [optional] 
**TableName** | Pointer to **string** |  | [optional] 
**DatabaseId** | Pointer to **int64** |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 
**TenantOBId** | Pointer to **int64** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TableGroupName** | Pointer to **string** |  | [optional] 
**ZoneList** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **int32** |  | [optional] 

## Methods

### NewTableInfoDo

`func NewTableInfoDo() *TableInfoDo`

NewTableInfoDo instantiates a new TableInfoDo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableInfoDoWithDefaults

`func NewTableInfoDoWithDefaults() *TableInfoDo`

NewTableInfoDoWithDefaults instantiates a new TableInfoDo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableId

`func (o *TableInfoDo) GetTableId() int64`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *TableInfoDo) GetTableIdOk() (*int64, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *TableInfoDo) SetTableId(v int64)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *TableInfoDo) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetTableName

`func (o *TableInfoDo) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableInfoDo) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableInfoDo) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableInfoDo) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetDatabaseId

`func (o *TableInfoDo) GetDatabaseId() int64`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *TableInfoDo) GetDatabaseIdOk() (*int64, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *TableInfoDo) SetDatabaseId(v int64)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *TableInfoDo) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetDatabaseName

`func (o *TableInfoDo) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *TableInfoDo) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *TableInfoDo) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *TableInfoDo) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetTenantOBId

`func (o *TableInfoDo) GetTenantOBId() int64`

GetTenantOBId returns the TenantOBId field if non-nil, zero value otherwise.

### GetTenantOBIdOk

`func (o *TableInfoDo) GetTenantOBIdOk() (*int64, bool)`

GetTenantOBIdOk returns a tuple with the TenantOBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantOBId

`func (o *TableInfoDo) SetTenantOBId(v int64)`

SetTenantOBId sets TenantOBId field to given value.

### HasTenantOBId

`func (o *TableInfoDo) HasTenantOBId() bool`

HasTenantOBId returns a boolean if a field has been set.

### GetTenantId

`func (o *TableInfoDo) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TableInfoDo) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TableInfoDo) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TableInfoDo) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTableGroupName

`func (o *TableInfoDo) GetTableGroupName() string`

GetTableGroupName returns the TableGroupName field if non-nil, zero value otherwise.

### GetTableGroupNameOk

`func (o *TableInfoDo) GetTableGroupNameOk() (*string, bool)`

GetTableGroupNameOk returns a tuple with the TableGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableGroupName

`func (o *TableInfoDo) SetTableGroupName(v string)`

SetTableGroupName sets TableGroupName field to given value.

### HasTableGroupName

`func (o *TableInfoDo) HasTableGroupName() bool`

HasTableGroupName returns a boolean if a field has been set.

### GetZoneList

`func (o *TableInfoDo) GetZoneList() string`

GetZoneList returns the ZoneList field if non-nil, zero value otherwise.

### GetZoneListOk

`func (o *TableInfoDo) GetZoneListOk() (*string, bool)`

GetZoneListOk returns a tuple with the ZoneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneList

`func (o *TableInfoDo) SetZoneList(v string)`

SetZoneList sets ZoneList field to given value.

### HasZoneList

`func (o *TableInfoDo) HasZoneList() bool`

HasZoneList returns a boolean if a field has been set.

### GetReadOnly

`func (o *TableInfoDo) GetReadOnly() int32`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *TableInfoDo) GetReadOnlyOk() (*int32, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *TableInfoDo) SetReadOnly(v int32)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *TableInfoDo) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


