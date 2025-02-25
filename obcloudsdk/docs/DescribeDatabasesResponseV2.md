# DescribeDatabasesResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 
**GmtCreate** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]OcpUserRoleDo**](OcpUserRoleDo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequiredSize** | Pointer to **float64** |  | [optional] 
**Tables** | Pointer to [**[]TableInfoDo**](TableInfoDo.md) |  | [optional] 
**Collation** | Pointer to **string** |  | [optional] 
**DbType** | Pointer to **string** |  | [optional] 

## Methods

### NewDescribeDatabasesResponseV2

`func NewDescribeDatabasesResponseV2() *DescribeDatabasesResponseV2`

NewDescribeDatabasesResponseV2 instantiates a new DescribeDatabasesResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDatabasesResponseV2WithDefaults

`func NewDescribeDatabasesResponseV2WithDefaults() *DescribeDatabasesResponseV2`

NewDescribeDatabasesResponseV2WithDefaults instantiates a new DescribeDatabasesResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DescribeDatabasesResponseV2) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DescribeDatabasesResponseV2) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DescribeDatabasesResponseV2) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DescribeDatabasesResponseV2) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetDatabaseName

`func (o *DescribeDatabasesResponseV2) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *DescribeDatabasesResponseV2) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *DescribeDatabasesResponseV2) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *DescribeDatabasesResponseV2) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetGmtCreate

`func (o *DescribeDatabasesResponseV2) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *DescribeDatabasesResponseV2) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *DescribeDatabasesResponseV2) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *DescribeDatabasesResponseV2) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetEncoding

`func (o *DescribeDatabasesResponseV2) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *DescribeDatabasesResponseV2) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *DescribeDatabasesResponseV2) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *DescribeDatabasesResponseV2) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeDatabasesResponseV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeDatabasesResponseV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeDatabasesResponseV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeDatabasesResponseV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsers

`func (o *DescribeDatabasesResponseV2) GetUsers() []OcpUserRoleDo`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DescribeDatabasesResponseV2) GetUsersOk() (*[]OcpUserRoleDo, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DescribeDatabasesResponseV2) SetUsers(v []OcpUserRoleDo)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DescribeDatabasesResponseV2) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetDescription

`func (o *DescribeDatabasesResponseV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeDatabasesResponseV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeDatabasesResponseV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DescribeDatabasesResponseV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredSize

`func (o *DescribeDatabasesResponseV2) GetRequiredSize() float64`

GetRequiredSize returns the RequiredSize field if non-nil, zero value otherwise.

### GetRequiredSizeOk

`func (o *DescribeDatabasesResponseV2) GetRequiredSizeOk() (*float64, bool)`

GetRequiredSizeOk returns a tuple with the RequiredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSize

`func (o *DescribeDatabasesResponseV2) SetRequiredSize(v float64)`

SetRequiredSize sets RequiredSize field to given value.

### HasRequiredSize

`func (o *DescribeDatabasesResponseV2) HasRequiredSize() bool`

HasRequiredSize returns a boolean if a field has been set.

### GetTables

`func (o *DescribeDatabasesResponseV2) GetTables() []TableInfoDo`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DescribeDatabasesResponseV2) GetTablesOk() (*[]TableInfoDo, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DescribeDatabasesResponseV2) SetTables(v []TableInfoDo)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DescribeDatabasesResponseV2) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetCollation

`func (o *DescribeDatabasesResponseV2) GetCollation() string`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *DescribeDatabasesResponseV2) GetCollationOk() (*string, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *DescribeDatabasesResponseV2) SetCollation(v string)`

SetCollation sets Collation field to given value.

### HasCollation

`func (o *DescribeDatabasesResponseV2) HasCollation() bool`

HasCollation returns a boolean if a field has been set.

### GetDbType

`func (o *DescribeDatabasesResponseV2) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *DescribeDatabasesResponseV2) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *DescribeDatabasesResponseV2) SetDbType(v string)`

SetDbType sets DbType field to given value.

### HasDbType

`func (o *DescribeDatabasesResponseV2) HasDbType() bool`

HasDbType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


