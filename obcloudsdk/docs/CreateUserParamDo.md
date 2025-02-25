# CreateUserParamDo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** |  | [optional] 
**CallerType** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecurityToken** | Pointer to **string** |  | [optional] 
**CallerUid** | Pointer to **string** |  | [optional] 
**CallerBid** | Pointer to **string** |  | [optional] 
**StsTokenCallerUid** | Pointer to **string** |  | [optional] 
**StsTokenCallerBid** | Pointer to **string** |  | [optional] 
**AcceptLanguage** | Pointer to **string** |  | [optional] 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**MergedCallerBid** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **string** |  | [optional] 
**EncryptionType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MergedCallerUid** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserParamDo

`func NewCreateUserParamDo() *CreateUserParamDo`

NewCreateUserParamDo instantiates a new CreateUserParamDo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserParamDoWithDefaults

`func NewCreateUserParamDoWithDefaults() *CreateUserParamDo`

NewCreateUserParamDoWithDefaults instantiates a new CreateUserParamDo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateUserParamDo) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateUserParamDo) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateUserParamDo) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateUserParamDo) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCallerType

`func (o *CreateUserParamDo) GetCallerType() string`

GetCallerType returns the CallerType field if non-nil, zero value otherwise.

### GetCallerTypeOk

`func (o *CreateUserParamDo) GetCallerTypeOk() (*string, bool)`

GetCallerTypeOk returns a tuple with the CallerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerType

`func (o *CreateUserParamDo) SetCallerType(v string)`

SetCallerType sets CallerType field to given value.

### HasCallerType

`func (o *CreateUserParamDo) HasCallerType() bool`

HasCallerType returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *CreateUserParamDo) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *CreateUserParamDo) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *CreateUserParamDo) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *CreateUserParamDo) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecurityToken

`func (o *CreateUserParamDo) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *CreateUserParamDo) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *CreateUserParamDo) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *CreateUserParamDo) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetCallerUid

`func (o *CreateUserParamDo) GetCallerUid() string`

GetCallerUid returns the CallerUid field if non-nil, zero value otherwise.

### GetCallerUidOk

`func (o *CreateUserParamDo) GetCallerUidOk() (*string, bool)`

GetCallerUidOk returns a tuple with the CallerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerUid

`func (o *CreateUserParamDo) SetCallerUid(v string)`

SetCallerUid sets CallerUid field to given value.

### HasCallerUid

`func (o *CreateUserParamDo) HasCallerUid() bool`

HasCallerUid returns a boolean if a field has been set.

### GetCallerBid

`func (o *CreateUserParamDo) GetCallerBid() string`

GetCallerBid returns the CallerBid field if non-nil, zero value otherwise.

### GetCallerBidOk

`func (o *CreateUserParamDo) GetCallerBidOk() (*string, bool)`

GetCallerBidOk returns a tuple with the CallerBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerBid

`func (o *CreateUserParamDo) SetCallerBid(v string)`

SetCallerBid sets CallerBid field to given value.

### HasCallerBid

`func (o *CreateUserParamDo) HasCallerBid() bool`

HasCallerBid returns a boolean if a field has been set.

### GetStsTokenCallerUid

`func (o *CreateUserParamDo) GetStsTokenCallerUid() string`

GetStsTokenCallerUid returns the StsTokenCallerUid field if non-nil, zero value otherwise.

### GetStsTokenCallerUidOk

`func (o *CreateUserParamDo) GetStsTokenCallerUidOk() (*string, bool)`

GetStsTokenCallerUidOk returns a tuple with the StsTokenCallerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsTokenCallerUid

`func (o *CreateUserParamDo) SetStsTokenCallerUid(v string)`

SetStsTokenCallerUid sets StsTokenCallerUid field to given value.

### HasStsTokenCallerUid

`func (o *CreateUserParamDo) HasStsTokenCallerUid() bool`

HasStsTokenCallerUid returns a boolean if a field has been set.

### GetStsTokenCallerBid

`func (o *CreateUserParamDo) GetStsTokenCallerBid() string`

GetStsTokenCallerBid returns the StsTokenCallerBid field if non-nil, zero value otherwise.

### GetStsTokenCallerBidOk

`func (o *CreateUserParamDo) GetStsTokenCallerBidOk() (*string, bool)`

GetStsTokenCallerBidOk returns a tuple with the StsTokenCallerBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsTokenCallerBid

`func (o *CreateUserParamDo) SetStsTokenCallerBid(v string)`

SetStsTokenCallerBid sets StsTokenCallerBid field to given value.

### HasStsTokenCallerBid

`func (o *CreateUserParamDo) HasStsTokenCallerBid() bool`

HasStsTokenCallerBid returns a boolean if a field has been set.

### GetAcceptLanguage

`func (o *CreateUserParamDo) GetAcceptLanguage() string`

GetAcceptLanguage returns the AcceptLanguage field if non-nil, zero value otherwise.

### GetAcceptLanguageOk

`func (o *CreateUserParamDo) GetAcceptLanguageOk() (*string, bool)`

GetAcceptLanguageOk returns a tuple with the AcceptLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptLanguage

`func (o *CreateUserParamDo) SetAcceptLanguage(v string)`

SetAcceptLanguage sets AcceptLanguage field to given value.

### HasAcceptLanguage

`func (o *CreateUserParamDo) HasAcceptLanguage() bool`

HasAcceptLanguage returns a boolean if a field has been set.

### GetPageNumber

`func (o *CreateUserParamDo) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *CreateUserParamDo) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *CreateUserParamDo) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *CreateUserParamDo) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *CreateUserParamDo) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CreateUserParamDo) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CreateUserParamDo) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CreateUserParamDo) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetMergedCallerBid

`func (o *CreateUserParamDo) GetMergedCallerBid() string`

GetMergedCallerBid returns the MergedCallerBid field if non-nil, zero value otherwise.

### GetMergedCallerBidOk

`func (o *CreateUserParamDo) GetMergedCallerBidOk() (*string, bool)`

GetMergedCallerBidOk returns a tuple with the MergedCallerBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedCallerBid

`func (o *CreateUserParamDo) SetMergedCallerBid(v string)`

SetMergedCallerBid sets MergedCallerBid field to given value.

### HasMergedCallerBid

`func (o *CreateUserParamDo) HasMergedCallerBid() bool`

HasMergedCallerBid returns a boolean if a field has been set.

### GetInstanceId

`func (o *CreateUserParamDo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateUserParamDo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateUserParamDo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CreateUserParamDo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetTenantId

`func (o *CreateUserParamDo) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateUserParamDo) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateUserParamDo) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CreateUserParamDo) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUserName

`func (o *CreateUserParamDo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateUserParamDo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateUserParamDo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CreateUserParamDo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *CreateUserParamDo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserParamDo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserParamDo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserParamDo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserType

`func (o *CreateUserParamDo) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *CreateUserParamDo) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *CreateUserParamDo) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *CreateUserParamDo) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetRoles

`func (o *CreateUserParamDo) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserParamDo) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserParamDo) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateUserParamDo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetEncryptionType

`func (o *CreateUserParamDo) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *CreateUserParamDo) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *CreateUserParamDo) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.

### HasEncryptionType

`func (o *CreateUserParamDo) HasEncryptionType() bool`

HasEncryptionType returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUserParamDo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUserParamDo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUserParamDo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUserParamDo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMergedCallerUid

`func (o *CreateUserParamDo) GetMergedCallerUid() string`

GetMergedCallerUid returns the MergedCallerUid field if non-nil, zero value otherwise.

### GetMergedCallerUidOk

`func (o *CreateUserParamDo) GetMergedCallerUidOk() (*string, bool)`

GetMergedCallerUidOk returns a tuple with the MergedCallerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedCallerUid

`func (o *CreateUserParamDo) SetMergedCallerUid(v string)`

SetMergedCallerUid sets MergedCallerUid field to given value.

### HasMergedCallerUid

`func (o *CreateUserParamDo) HasMergedCallerUid() bool`

HasMergedCallerUid returns a boolean if a field has been set.

### GetUserID

`func (o *CreateUserParamDo) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *CreateUserParamDo) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *CreateUserParamDo) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *CreateUserParamDo) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUid

`func (o *CreateUserParamDo) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreateUserParamDo) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreateUserParamDo) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CreateUserParamDo) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


