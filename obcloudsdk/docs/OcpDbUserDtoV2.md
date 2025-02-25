# OcpDbUserDtoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**UserStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Databases** | Pointer to [**[]OcpUserPrivilegeV2**](OcpUserPrivilegeV2.md) |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewOcpDbUserDtoV2

`func NewOcpDbUserDtoV2() *OcpDbUserDtoV2`

NewOcpDbUserDtoV2 instantiates a new OcpDbUserDtoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcpDbUserDtoV2WithDefaults

`func NewOcpDbUserDtoV2WithDefaults() *OcpDbUserDtoV2`

NewOcpDbUserDtoV2WithDefaults instantiates a new OcpDbUserDtoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *OcpDbUserDtoV2) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *OcpDbUserDtoV2) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *OcpDbUserDtoV2) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *OcpDbUserDtoV2) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserType

`func (o *OcpDbUserDtoV2) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *OcpDbUserDtoV2) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *OcpDbUserDtoV2) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *OcpDbUserDtoV2) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserStatus

`func (o *OcpDbUserDtoV2) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *OcpDbUserDtoV2) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *OcpDbUserDtoV2) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *OcpDbUserDtoV2) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetDescription

`func (o *OcpDbUserDtoV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OcpDbUserDtoV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OcpDbUserDtoV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OcpDbUserDtoV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDatabases

`func (o *OcpDbUserDtoV2) GetDatabases() []OcpUserPrivilegeV2`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *OcpDbUserDtoV2) GetDatabasesOk() (*[]OcpUserPrivilegeV2, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *OcpDbUserDtoV2) SetDatabases(v []OcpUserPrivilegeV2)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *OcpDbUserDtoV2) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetPassword

`func (o *OcpDbUserDtoV2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OcpDbUserDtoV2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OcpDbUserDtoV2) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OcpDbUserDtoV2) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


