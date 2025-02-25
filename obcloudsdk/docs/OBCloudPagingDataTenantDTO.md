# OBCloudPagingDataTenantDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataList** | Pointer to [**[]TenantDTO**](TenantDTO.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewOBCloudPagingDataTenantDTO

`func NewOBCloudPagingDataTenantDTO() *OBCloudPagingDataTenantDTO`

NewOBCloudPagingDataTenantDTO instantiates a new OBCloudPagingDataTenantDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudPagingDataTenantDTOWithDefaults

`func NewOBCloudPagingDataTenantDTOWithDefaults() *OBCloudPagingDataTenantDTO`

NewOBCloudPagingDataTenantDTOWithDefaults instantiates a new OBCloudPagingDataTenantDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataList

`func (o *OBCloudPagingDataTenantDTO) GetDataList() []TenantDTO`

GetDataList returns the DataList field if non-nil, zero value otherwise.

### GetDataListOk

`func (o *OBCloudPagingDataTenantDTO) GetDataListOk() (*[]TenantDTO, bool)`

GetDataListOk returns a tuple with the DataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataList

`func (o *OBCloudPagingDataTenantDTO) SetDataList(v []TenantDTO)`

SetDataList sets DataList field to given value.

### HasDataList

`func (o *OBCloudPagingDataTenantDTO) HasDataList() bool`

HasDataList returns a boolean if a field has been set.

### GetTotal

`func (o *OBCloudPagingDataTenantDTO) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OBCloudPagingDataTenantDTO) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OBCloudPagingDataTenantDTO) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OBCloudPagingDataTenantDTO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *OBCloudPagingDataTenantDTO) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *OBCloudPagingDataTenantDTO) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *OBCloudPagingDataTenantDTO) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *OBCloudPagingDataTenantDTO) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


