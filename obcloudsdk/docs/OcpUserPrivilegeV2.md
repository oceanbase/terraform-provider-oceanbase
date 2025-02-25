# OcpUserPrivilegeV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** |  | [optional] 
**Database** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Privileges** | Pointer to **string** |  | [optional] 
**IsSuccess** | Pointer to **bool** |  | [optional] 
**WithGrant** | Pointer to **int64** |  | [optional] 

## Methods

### NewOcpUserPrivilegeV2

`func NewOcpUserPrivilegeV2() *OcpUserPrivilegeV2`

NewOcpUserPrivilegeV2 instantiates a new OcpUserPrivilegeV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcpUserPrivilegeV2WithDefaults

`func NewOcpUserPrivilegeV2WithDefaults() *OcpUserPrivilegeV2`

NewOcpUserPrivilegeV2WithDefaults instantiates a new OcpUserPrivilegeV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *OcpUserPrivilegeV2) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OcpUserPrivilegeV2) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OcpUserPrivilegeV2) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OcpUserPrivilegeV2) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetDatabase

`func (o *OcpUserPrivilegeV2) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *OcpUserPrivilegeV2) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *OcpUserPrivilegeV2) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *OcpUserPrivilegeV2) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetTable

`func (o *OcpUserPrivilegeV2) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *OcpUserPrivilegeV2) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *OcpUserPrivilegeV2) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *OcpUserPrivilegeV2) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetRole

`func (o *OcpUserPrivilegeV2) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OcpUserPrivilegeV2) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OcpUserPrivilegeV2) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OcpUserPrivilegeV2) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPrivileges

`func (o *OcpUserPrivilegeV2) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *OcpUserPrivilegeV2) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *OcpUserPrivilegeV2) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *OcpUserPrivilegeV2) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetIsSuccess

`func (o *OcpUserPrivilegeV2) GetIsSuccess() bool`

GetIsSuccess returns the IsSuccess field if non-nil, zero value otherwise.

### GetIsSuccessOk

`func (o *OcpUserPrivilegeV2) GetIsSuccessOk() (*bool, bool)`

GetIsSuccessOk returns a tuple with the IsSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccess

`func (o *OcpUserPrivilegeV2) SetIsSuccess(v bool)`

SetIsSuccess sets IsSuccess field to given value.

### HasIsSuccess

`func (o *OcpUserPrivilegeV2) HasIsSuccess() bool`

HasIsSuccess returns a boolean if a field has been set.

### GetWithGrant

`func (o *OcpUserPrivilegeV2) GetWithGrant() int64`

GetWithGrant returns the WithGrant field if non-nil, zero value otherwise.

### GetWithGrantOk

`func (o *OcpUserPrivilegeV2) GetWithGrantOk() (*int64, bool)`

GetWithGrantOk returns a tuple with the WithGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithGrant

`func (o *OcpUserPrivilegeV2) SetWithGrant(v int64)`

SetWithGrant sets WithGrant field to given value.

### HasWithGrant

`func (o *OcpUserPrivilegeV2) HasWithGrant() bool`

HasWithGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


