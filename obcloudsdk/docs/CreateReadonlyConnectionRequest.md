# CreateReadonlyConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneIdList** | Pointer to **[]string** |  | [optional] 
**UserAccount** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateReadonlyConnectionRequest

`func NewCreateReadonlyConnectionRequest() *CreateReadonlyConnectionRequest`

NewCreateReadonlyConnectionRequest instantiates a new CreateReadonlyConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReadonlyConnectionRequestWithDefaults

`func NewCreateReadonlyConnectionRequestWithDefaults() *CreateReadonlyConnectionRequest`

NewCreateReadonlyConnectionRequestWithDefaults instantiates a new CreateReadonlyConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneIdList

`func (o *CreateReadonlyConnectionRequest) GetZoneIdList() []string`

GetZoneIdList returns the ZoneIdList field if non-nil, zero value otherwise.

### GetZoneIdListOk

`func (o *CreateReadonlyConnectionRequest) GetZoneIdListOk() (*[]string, bool)`

GetZoneIdListOk returns a tuple with the ZoneIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIdList

`func (o *CreateReadonlyConnectionRequest) SetZoneIdList(v []string)`

SetZoneIdList sets ZoneIdList field to given value.

### HasZoneIdList

`func (o *CreateReadonlyConnectionRequest) HasZoneIdList() bool`

HasZoneIdList returns a boolean if a field has been set.

### GetUserAccount

`func (o *CreateReadonlyConnectionRequest) GetUserAccount() string`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *CreateReadonlyConnectionRequest) GetUserAccountOk() (*string, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *CreateReadonlyConnectionRequest) SetUserAccount(v string)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *CreateReadonlyConnectionRequest) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


