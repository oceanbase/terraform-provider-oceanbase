# TenantConnectionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**UserVpcOwnerId** | Pointer to **string** |  | [optional] 
**AddressId** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**AddressType** | Pointer to **string** |  | [optional] 
**NetworkType** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**IntranetAddress** | Pointer to **string** |  | [optional] 
**IntranetPort** | Pointer to **int32** |  | [optional] 
**IntranetDomain** | Pointer to **string** |  | [optional] 
**IntranetAddressStatus** | Pointer to **string** |  | [optional] 
**InternetAddress** | Pointer to **string** |  | [optional] 
**InternetPort** | Pointer to **int32** |  | [optional] 
**InternetDomain** | Pointer to **string** |  | [optional] 
**InternetAddressStatus** | Pointer to **string** |  | [optional] 
**IntranetPeeringAddress** | Pointer to **string** |  | [optional] 
**IntranetPeeringPort** | Pointer to **int32** |  | [optional] 
**IntranetPeeringDomain** | Pointer to **string** |  | [optional] 
**IntranetPeeringAddressStatus** | Pointer to **string** |  | [optional] 
**IntranetAddressMasterZoneId** | Pointer to **string** |  | [optional] 
**IntranetAddressSlaveZoneId** | Pointer to **string** |  | [optional] 
**ConnectionZones** | Pointer to **[]string** |  | [optional] 
**AddressStatus** | Pointer to **string** |  | [optional] 
**UserNameFormat** | Pointer to **string** |  | [optional] 
**PrivateLinkList** | Pointer to [**[]TenantPrivateLinkDTO**](TenantPrivateLinkDTO.md) |  | [optional] 
**PeeringDstVpcId** | Pointer to **string** |  | [optional] 
**PeeringDstCidrList** | Pointer to **string** |  | [optional] 
**PeeringId** | Pointer to **string** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 
**ProxyClusterId** | Pointer to **string** |  | [optional] 
**MaxConnectionNum** | Pointer to **int32** |  | [optional] 
**InternetMaxConnectionNum** | Pointer to **int32** |  | [optional] 
**IntranetPeeringMaxConnectionNum** | Pointer to **int32** |  | [optional] 
**EnableRPc** | Pointer to **bool** |  | [optional] 
**IntranetRpcPort** | Pointer to **int32** |  | [optional] 
**InternetRpcPort** | Pointer to **int32** |  | [optional] 
**ProxyClusterInfo** | Pointer to [**ProxyClusterInfoDTO**](ProxyClusterInfoDTO.md) |  | [optional] 
**VswitchId** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantConnectionDTO

`func NewTenantConnectionDTO() *TenantConnectionDTO`

NewTenantConnectionDTO instantiates a new TenantConnectionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantConnectionDTOWithDefaults

`func NewTenantConnectionDTOWithDefaults() *TenantConnectionDTO`

NewTenantConnectionDTOWithDefaults instantiates a new TenantConnectionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantConnectionDTO) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantConnectionDTO) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantConnectionDTO) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantConnectionDTO) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetVpcId

`func (o *TenantConnectionDTO) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *TenantConnectionDTO) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *TenantConnectionDTO) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *TenantConnectionDTO) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetUserVpcOwnerId

`func (o *TenantConnectionDTO) GetUserVpcOwnerId() string`

GetUserVpcOwnerId returns the UserVpcOwnerId field if non-nil, zero value otherwise.

### GetUserVpcOwnerIdOk

`func (o *TenantConnectionDTO) GetUserVpcOwnerIdOk() (*string, bool)`

GetUserVpcOwnerIdOk returns a tuple with the UserVpcOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVpcOwnerId

`func (o *TenantConnectionDTO) SetUserVpcOwnerId(v string)`

SetUserVpcOwnerId sets UserVpcOwnerId field to given value.

### HasUserVpcOwnerId

`func (o *TenantConnectionDTO) HasUserVpcOwnerId() bool`

HasUserVpcOwnerId returns a boolean if a field has been set.

### GetAddressId

`func (o *TenantConnectionDTO) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *TenantConnectionDTO) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *TenantConnectionDTO) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *TenantConnectionDTO) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetServiceType

`func (o *TenantConnectionDTO) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *TenantConnectionDTO) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *TenantConnectionDTO) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *TenantConnectionDTO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetAddressType

`func (o *TenantConnectionDTO) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *TenantConnectionDTO) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *TenantConnectionDTO) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *TenantConnectionDTO) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetNetworkType

`func (o *TenantConnectionDTO) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *TenantConnectionDTO) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *TenantConnectionDTO) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *TenantConnectionDTO) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetRole

`func (o *TenantConnectionDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TenantConnectionDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TenantConnectionDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TenantConnectionDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetIntranetAddress

`func (o *TenantConnectionDTO) GetIntranetAddress() string`

GetIntranetAddress returns the IntranetAddress field if non-nil, zero value otherwise.

### GetIntranetAddressOk

`func (o *TenantConnectionDTO) GetIntranetAddressOk() (*string, bool)`

GetIntranetAddressOk returns a tuple with the IntranetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetAddress

`func (o *TenantConnectionDTO) SetIntranetAddress(v string)`

SetIntranetAddress sets IntranetAddress field to given value.

### HasIntranetAddress

`func (o *TenantConnectionDTO) HasIntranetAddress() bool`

HasIntranetAddress returns a boolean if a field has been set.

### GetIntranetPort

`func (o *TenantConnectionDTO) GetIntranetPort() int32`

GetIntranetPort returns the IntranetPort field if non-nil, zero value otherwise.

### GetIntranetPortOk

`func (o *TenantConnectionDTO) GetIntranetPortOk() (*int32, bool)`

GetIntranetPortOk returns a tuple with the IntranetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetPort

`func (o *TenantConnectionDTO) SetIntranetPort(v int32)`

SetIntranetPort sets IntranetPort field to given value.

### HasIntranetPort

`func (o *TenantConnectionDTO) HasIntranetPort() bool`

HasIntranetPort returns a boolean if a field has been set.

### GetIntranetDomain

`func (o *TenantConnectionDTO) GetIntranetDomain() string`

GetIntranetDomain returns the IntranetDomain field if non-nil, zero value otherwise.

### GetIntranetDomainOk

`func (o *TenantConnectionDTO) GetIntranetDomainOk() (*string, bool)`

GetIntranetDomainOk returns a tuple with the IntranetDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetDomain

`func (o *TenantConnectionDTO) SetIntranetDomain(v string)`

SetIntranetDomain sets IntranetDomain field to given value.

### HasIntranetDomain

`func (o *TenantConnectionDTO) HasIntranetDomain() bool`

HasIntranetDomain returns a boolean if a field has been set.

### GetIntranetAddressStatus

`func (o *TenantConnectionDTO) GetIntranetAddressStatus() string`

GetIntranetAddressStatus returns the IntranetAddressStatus field if non-nil, zero value otherwise.

### GetIntranetAddressStatusOk

`func (o *TenantConnectionDTO) GetIntranetAddressStatusOk() (*string, bool)`

GetIntranetAddressStatusOk returns a tuple with the IntranetAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetAddressStatus

`func (o *TenantConnectionDTO) SetIntranetAddressStatus(v string)`

SetIntranetAddressStatus sets IntranetAddressStatus field to given value.

### HasIntranetAddressStatus

`func (o *TenantConnectionDTO) HasIntranetAddressStatus() bool`

HasIntranetAddressStatus returns a boolean if a field has been set.

### GetInternetAddress

`func (o *TenantConnectionDTO) GetInternetAddress() string`

GetInternetAddress returns the InternetAddress field if non-nil, zero value otherwise.

### GetInternetAddressOk

`func (o *TenantConnectionDTO) GetInternetAddressOk() (*string, bool)`

GetInternetAddressOk returns a tuple with the InternetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAddress

`func (o *TenantConnectionDTO) SetInternetAddress(v string)`

SetInternetAddress sets InternetAddress field to given value.

### HasInternetAddress

`func (o *TenantConnectionDTO) HasInternetAddress() bool`

HasInternetAddress returns a boolean if a field has been set.

### GetInternetPort

`func (o *TenantConnectionDTO) GetInternetPort() int32`

GetInternetPort returns the InternetPort field if non-nil, zero value otherwise.

### GetInternetPortOk

`func (o *TenantConnectionDTO) GetInternetPortOk() (*int32, bool)`

GetInternetPortOk returns a tuple with the InternetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetPort

`func (o *TenantConnectionDTO) SetInternetPort(v int32)`

SetInternetPort sets InternetPort field to given value.

### HasInternetPort

`func (o *TenantConnectionDTO) HasInternetPort() bool`

HasInternetPort returns a boolean if a field has been set.

### GetInternetDomain

`func (o *TenantConnectionDTO) GetInternetDomain() string`

GetInternetDomain returns the InternetDomain field if non-nil, zero value otherwise.

### GetInternetDomainOk

`func (o *TenantConnectionDTO) GetInternetDomainOk() (*string, bool)`

GetInternetDomainOk returns a tuple with the InternetDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetDomain

`func (o *TenantConnectionDTO) SetInternetDomain(v string)`

SetInternetDomain sets InternetDomain field to given value.

### HasInternetDomain

`func (o *TenantConnectionDTO) HasInternetDomain() bool`

HasInternetDomain returns a boolean if a field has been set.

### GetInternetAddressStatus

`func (o *TenantConnectionDTO) GetInternetAddressStatus() string`

GetInternetAddressStatus returns the InternetAddressStatus field if non-nil, zero value otherwise.

### GetInternetAddressStatusOk

`func (o *TenantConnectionDTO) GetInternetAddressStatusOk() (*string, bool)`

GetInternetAddressStatusOk returns a tuple with the InternetAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAddressStatus

`func (o *TenantConnectionDTO) SetInternetAddressStatus(v string)`

SetInternetAddressStatus sets InternetAddressStatus field to given value.

### HasInternetAddressStatus

`func (o *TenantConnectionDTO) HasInternetAddressStatus() bool`

HasInternetAddressStatus returns a boolean if a field has been set.

### GetIntranetPeeringAddress

`func (o *TenantConnectionDTO) GetIntranetPeeringAddress() string`

GetIntranetPeeringAddress returns the IntranetPeeringAddress field if non-nil, zero value otherwise.

### GetIntranetPeeringAddressOk

`func (o *TenantConnectionDTO) GetIntranetPeeringAddressOk() (*string, bool)`

GetIntranetPeeringAddressOk returns a tuple with the IntranetPeeringAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetPeeringAddress

`func (o *TenantConnectionDTO) SetIntranetPeeringAddress(v string)`

SetIntranetPeeringAddress sets IntranetPeeringAddress field to given value.

### HasIntranetPeeringAddress

`func (o *TenantConnectionDTO) HasIntranetPeeringAddress() bool`

HasIntranetPeeringAddress returns a boolean if a field has been set.

### GetIntranetPeeringPort

`func (o *TenantConnectionDTO) GetIntranetPeeringPort() int32`

GetIntranetPeeringPort returns the IntranetPeeringPort field if non-nil, zero value otherwise.

### GetIntranetPeeringPortOk

`func (o *TenantConnectionDTO) GetIntranetPeeringPortOk() (*int32, bool)`

GetIntranetPeeringPortOk returns a tuple with the IntranetPeeringPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetPeeringPort

`func (o *TenantConnectionDTO) SetIntranetPeeringPort(v int32)`

SetIntranetPeeringPort sets IntranetPeeringPort field to given value.

### HasIntranetPeeringPort

`func (o *TenantConnectionDTO) HasIntranetPeeringPort() bool`

HasIntranetPeeringPort returns a boolean if a field has been set.

### GetIntranetPeeringDomain

`func (o *TenantConnectionDTO) GetIntranetPeeringDomain() string`

GetIntranetPeeringDomain returns the IntranetPeeringDomain field if non-nil, zero value otherwise.

### GetIntranetPeeringDomainOk

`func (o *TenantConnectionDTO) GetIntranetPeeringDomainOk() (*string, bool)`

GetIntranetPeeringDomainOk returns a tuple with the IntranetPeeringDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetPeeringDomain

`func (o *TenantConnectionDTO) SetIntranetPeeringDomain(v string)`

SetIntranetPeeringDomain sets IntranetPeeringDomain field to given value.

### HasIntranetPeeringDomain

`func (o *TenantConnectionDTO) HasIntranetPeeringDomain() bool`

HasIntranetPeeringDomain returns a boolean if a field has been set.

### GetIntranetPeeringAddressStatus

`func (o *TenantConnectionDTO) GetIntranetPeeringAddressStatus() string`

GetIntranetPeeringAddressStatus returns the IntranetPeeringAddressStatus field if non-nil, zero value otherwise.

### GetIntranetPeeringAddressStatusOk

`func (o *TenantConnectionDTO) GetIntranetPeeringAddressStatusOk() (*string, bool)`

GetIntranetPeeringAddressStatusOk returns a tuple with the IntranetPeeringAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetPeeringAddressStatus

`func (o *TenantConnectionDTO) SetIntranetPeeringAddressStatus(v string)`

SetIntranetPeeringAddressStatus sets IntranetPeeringAddressStatus field to given value.

### HasIntranetPeeringAddressStatus

`func (o *TenantConnectionDTO) HasIntranetPeeringAddressStatus() bool`

HasIntranetPeeringAddressStatus returns a boolean if a field has been set.

### GetIntranetAddressMasterZoneId

`func (o *TenantConnectionDTO) GetIntranetAddressMasterZoneId() string`

GetIntranetAddressMasterZoneId returns the IntranetAddressMasterZoneId field if non-nil, zero value otherwise.

### GetIntranetAddressMasterZoneIdOk

`func (o *TenantConnectionDTO) GetIntranetAddressMasterZoneIdOk() (*string, bool)`

GetIntranetAddressMasterZoneIdOk returns a tuple with the IntranetAddressMasterZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetAddressMasterZoneId

`func (o *TenantConnectionDTO) SetIntranetAddressMasterZoneId(v string)`

SetIntranetAddressMasterZoneId sets IntranetAddressMasterZoneId field to given value.

### HasIntranetAddressMasterZoneId

`func (o *TenantConnectionDTO) HasIntranetAddressMasterZoneId() bool`

HasIntranetAddressMasterZoneId returns a boolean if a field has been set.

### GetIntranetAddressSlaveZoneId

`func (o *TenantConnectionDTO) GetIntranetAddressSlaveZoneId() string`

GetIntranetAddressSlaveZoneId returns the IntranetAddressSlaveZoneId field if non-nil, zero value otherwise.

### GetIntranetAddressSlaveZoneIdOk

`func (o *TenantConnectionDTO) GetIntranetAddressSlaveZoneIdOk() (*string, bool)`

GetIntranetAddressSlaveZoneIdOk returns a tuple with the IntranetAddressSlaveZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetAddressSlaveZoneId

`func (o *TenantConnectionDTO) SetIntranetAddressSlaveZoneId(v string)`

SetIntranetAddressSlaveZoneId sets IntranetAddressSlaveZoneId field to given value.

### HasIntranetAddressSlaveZoneId

`func (o *TenantConnectionDTO) HasIntranetAddressSlaveZoneId() bool`

HasIntranetAddressSlaveZoneId returns a boolean if a field has been set.

### GetConnectionZones

`func (o *TenantConnectionDTO) GetConnectionZones() []string`

GetConnectionZones returns the ConnectionZones field if non-nil, zero value otherwise.

### GetConnectionZonesOk

`func (o *TenantConnectionDTO) GetConnectionZonesOk() (*[]string, bool)`

GetConnectionZonesOk returns a tuple with the ConnectionZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionZones

`func (o *TenantConnectionDTO) SetConnectionZones(v []string)`

SetConnectionZones sets ConnectionZones field to given value.

### HasConnectionZones

`func (o *TenantConnectionDTO) HasConnectionZones() bool`

HasConnectionZones returns a boolean if a field has been set.

### GetAddressStatus

`func (o *TenantConnectionDTO) GetAddressStatus() string`

GetAddressStatus returns the AddressStatus field if non-nil, zero value otherwise.

### GetAddressStatusOk

`func (o *TenantConnectionDTO) GetAddressStatusOk() (*string, bool)`

GetAddressStatusOk returns a tuple with the AddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStatus

`func (o *TenantConnectionDTO) SetAddressStatus(v string)`

SetAddressStatus sets AddressStatus field to given value.

### HasAddressStatus

`func (o *TenantConnectionDTO) HasAddressStatus() bool`

HasAddressStatus returns a boolean if a field has been set.

### GetUserNameFormat

`func (o *TenantConnectionDTO) GetUserNameFormat() string`

GetUserNameFormat returns the UserNameFormat field if non-nil, zero value otherwise.

### GetUserNameFormatOk

`func (o *TenantConnectionDTO) GetUserNameFormatOk() (*string, bool)`

GetUserNameFormatOk returns a tuple with the UserNameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameFormat

`func (o *TenantConnectionDTO) SetUserNameFormat(v string)`

SetUserNameFormat sets UserNameFormat field to given value.

### HasUserNameFormat

`func (o *TenantConnectionDTO) HasUserNameFormat() bool`

HasUserNameFormat returns a boolean if a field has been set.

### GetPrivateLinkList

`func (o *TenantConnectionDTO) GetPrivateLinkList() []TenantPrivateLinkDTO`

GetPrivateLinkList returns the PrivateLinkList field if non-nil, zero value otherwise.

### GetPrivateLinkListOk

`func (o *TenantConnectionDTO) GetPrivateLinkListOk() (*[]TenantPrivateLinkDTO, bool)`

GetPrivateLinkListOk returns a tuple with the PrivateLinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkList

`func (o *TenantConnectionDTO) SetPrivateLinkList(v []TenantPrivateLinkDTO)`

SetPrivateLinkList sets PrivateLinkList field to given value.

### HasPrivateLinkList

`func (o *TenantConnectionDTO) HasPrivateLinkList() bool`

HasPrivateLinkList returns a boolean if a field has been set.

### GetPeeringDstVpcId

`func (o *TenantConnectionDTO) GetPeeringDstVpcId() string`

GetPeeringDstVpcId returns the PeeringDstVpcId field if non-nil, zero value otherwise.

### GetPeeringDstVpcIdOk

`func (o *TenantConnectionDTO) GetPeeringDstVpcIdOk() (*string, bool)`

GetPeeringDstVpcIdOk returns a tuple with the PeeringDstVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringDstVpcId

`func (o *TenantConnectionDTO) SetPeeringDstVpcId(v string)`

SetPeeringDstVpcId sets PeeringDstVpcId field to given value.

### HasPeeringDstVpcId

`func (o *TenantConnectionDTO) HasPeeringDstVpcId() bool`

HasPeeringDstVpcId returns a boolean if a field has been set.

### GetPeeringDstCidrList

`func (o *TenantConnectionDTO) GetPeeringDstCidrList() string`

GetPeeringDstCidrList returns the PeeringDstCidrList field if non-nil, zero value otherwise.

### GetPeeringDstCidrListOk

`func (o *TenantConnectionDTO) GetPeeringDstCidrListOk() (*string, bool)`

GetPeeringDstCidrListOk returns a tuple with the PeeringDstCidrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringDstCidrList

`func (o *TenantConnectionDTO) SetPeeringDstCidrList(v string)`

SetPeeringDstCidrList sets PeeringDstCidrList field to given value.

### HasPeeringDstCidrList

`func (o *TenantConnectionDTO) HasPeeringDstCidrList() bool`

HasPeeringDstCidrList returns a boolean if a field has been set.

### GetPeeringId

`func (o *TenantConnectionDTO) GetPeeringId() string`

GetPeeringId returns the PeeringId field if non-nil, zero value otherwise.

### GetPeeringIdOk

`func (o *TenantConnectionDTO) GetPeeringIdOk() (*string, bool)`

GetPeeringIdOk returns a tuple with the PeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringId

`func (o *TenantConnectionDTO) SetPeeringId(v string)`

SetPeeringId sets PeeringId field to given value.

### HasPeeringId

`func (o *TenantConnectionDTO) HasPeeringId() bool`

HasPeeringId returns a boolean if a field has been set.

### GetUseSSL

`func (o *TenantConnectionDTO) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *TenantConnectionDTO) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *TenantConnectionDTO) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *TenantConnectionDTO) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetProxyClusterId

`func (o *TenantConnectionDTO) GetProxyClusterId() string`

GetProxyClusterId returns the ProxyClusterId field if non-nil, zero value otherwise.

### GetProxyClusterIdOk

`func (o *TenantConnectionDTO) GetProxyClusterIdOk() (*string, bool)`

GetProxyClusterIdOk returns a tuple with the ProxyClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyClusterId

`func (o *TenantConnectionDTO) SetProxyClusterId(v string)`

SetProxyClusterId sets ProxyClusterId field to given value.

### HasProxyClusterId

`func (o *TenantConnectionDTO) HasProxyClusterId() bool`

HasProxyClusterId returns a boolean if a field has been set.

### GetMaxConnectionNum

`func (o *TenantConnectionDTO) GetMaxConnectionNum() int32`

GetMaxConnectionNum returns the MaxConnectionNum field if non-nil, zero value otherwise.

### GetMaxConnectionNumOk

`func (o *TenantConnectionDTO) GetMaxConnectionNumOk() (*int32, bool)`

GetMaxConnectionNumOk returns a tuple with the MaxConnectionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionNum

`func (o *TenantConnectionDTO) SetMaxConnectionNum(v int32)`

SetMaxConnectionNum sets MaxConnectionNum field to given value.

### HasMaxConnectionNum

`func (o *TenantConnectionDTO) HasMaxConnectionNum() bool`

HasMaxConnectionNum returns a boolean if a field has been set.

### GetInternetMaxConnectionNum

`func (o *TenantConnectionDTO) GetInternetMaxConnectionNum() int32`

GetInternetMaxConnectionNum returns the InternetMaxConnectionNum field if non-nil, zero value otherwise.

### GetInternetMaxConnectionNumOk

`func (o *TenantConnectionDTO) GetInternetMaxConnectionNumOk() (*int32, bool)`

GetInternetMaxConnectionNumOk returns a tuple with the InternetMaxConnectionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMaxConnectionNum

`func (o *TenantConnectionDTO) SetInternetMaxConnectionNum(v int32)`

SetInternetMaxConnectionNum sets InternetMaxConnectionNum field to given value.

### HasInternetMaxConnectionNum

`func (o *TenantConnectionDTO) HasInternetMaxConnectionNum() bool`

HasInternetMaxConnectionNum returns a boolean if a field has been set.

### GetIntranetPeeringMaxConnectionNum

`func (o *TenantConnectionDTO) GetIntranetPeeringMaxConnectionNum() int32`

GetIntranetPeeringMaxConnectionNum returns the IntranetPeeringMaxConnectionNum field if non-nil, zero value otherwise.

### GetIntranetPeeringMaxConnectionNumOk

`func (o *TenantConnectionDTO) GetIntranetPeeringMaxConnectionNumOk() (*int32, bool)`

GetIntranetPeeringMaxConnectionNumOk returns a tuple with the IntranetPeeringMaxConnectionNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetPeeringMaxConnectionNum

`func (o *TenantConnectionDTO) SetIntranetPeeringMaxConnectionNum(v int32)`

SetIntranetPeeringMaxConnectionNum sets IntranetPeeringMaxConnectionNum field to given value.

### HasIntranetPeeringMaxConnectionNum

`func (o *TenantConnectionDTO) HasIntranetPeeringMaxConnectionNum() bool`

HasIntranetPeeringMaxConnectionNum returns a boolean if a field has been set.

### GetEnableRPc

`func (o *TenantConnectionDTO) GetEnableRPc() bool`

GetEnableRPc returns the EnableRPc field if non-nil, zero value otherwise.

### GetEnableRPcOk

`func (o *TenantConnectionDTO) GetEnableRPcOk() (*bool, bool)`

GetEnableRPcOk returns a tuple with the EnableRPc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRPc

`func (o *TenantConnectionDTO) SetEnableRPc(v bool)`

SetEnableRPc sets EnableRPc field to given value.

### HasEnableRPc

`func (o *TenantConnectionDTO) HasEnableRPc() bool`

HasEnableRPc returns a boolean if a field has been set.

### GetIntranetRpcPort

`func (o *TenantConnectionDTO) GetIntranetRpcPort() int32`

GetIntranetRpcPort returns the IntranetRpcPort field if non-nil, zero value otherwise.

### GetIntranetRpcPortOk

`func (o *TenantConnectionDTO) GetIntranetRpcPortOk() (*int32, bool)`

GetIntranetRpcPortOk returns a tuple with the IntranetRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntranetRpcPort

`func (o *TenantConnectionDTO) SetIntranetRpcPort(v int32)`

SetIntranetRpcPort sets IntranetRpcPort field to given value.

### HasIntranetRpcPort

`func (o *TenantConnectionDTO) HasIntranetRpcPort() bool`

HasIntranetRpcPort returns a boolean if a field has been set.

### GetInternetRpcPort

`func (o *TenantConnectionDTO) GetInternetRpcPort() int32`

GetInternetRpcPort returns the InternetRpcPort field if non-nil, zero value otherwise.

### GetInternetRpcPortOk

`func (o *TenantConnectionDTO) GetInternetRpcPortOk() (*int32, bool)`

GetInternetRpcPortOk returns a tuple with the InternetRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetRpcPort

`func (o *TenantConnectionDTO) SetInternetRpcPort(v int32)`

SetInternetRpcPort sets InternetRpcPort field to given value.

### HasInternetRpcPort

`func (o *TenantConnectionDTO) HasInternetRpcPort() bool`

HasInternetRpcPort returns a boolean if a field has been set.

### GetProxyClusterInfo

`func (o *TenantConnectionDTO) GetProxyClusterInfo() ProxyClusterInfoDTO`

GetProxyClusterInfo returns the ProxyClusterInfo field if non-nil, zero value otherwise.

### GetProxyClusterInfoOk

`func (o *TenantConnectionDTO) GetProxyClusterInfoOk() (*ProxyClusterInfoDTO, bool)`

GetProxyClusterInfoOk returns a tuple with the ProxyClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyClusterInfo

`func (o *TenantConnectionDTO) SetProxyClusterInfo(v ProxyClusterInfoDTO)`

SetProxyClusterInfo sets ProxyClusterInfo field to given value.

### HasProxyClusterInfo

`func (o *TenantConnectionDTO) HasProxyClusterInfo() bool`

HasProxyClusterInfo returns a boolean if a field has been set.

### GetVswitchId

`func (o *TenantConnectionDTO) GetVswitchId() string`

GetVswitchId returns the VswitchId field if non-nil, zero value otherwise.

### GetVswitchIdOk

`func (o *TenantConnectionDTO) GetVswitchIdOk() (*string, bool)`

GetVswitchIdOk returns a tuple with the VswitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchId

`func (o *TenantConnectionDTO) SetVswitchId(v string)`

SetVswitchId sets VswitchId field to given value.

### HasVswitchId

`func (o *TenantConnectionDTO) HasVswitchId() bool`

HasVswitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


