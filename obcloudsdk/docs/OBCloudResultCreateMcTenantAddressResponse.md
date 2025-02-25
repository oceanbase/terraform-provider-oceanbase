# OBCloudResultCreateMcTenantAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CreateMcTenantAddressResponse**](CreateMcTenantAddressResponse.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultCreateMcTenantAddressResponse

`func NewOBCloudResultCreateMcTenantAddressResponse() *OBCloudResultCreateMcTenantAddressResponse`

NewOBCloudResultCreateMcTenantAddressResponse instantiates a new OBCloudResultCreateMcTenantAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultCreateMcTenantAddressResponseWithDefaults

`func NewOBCloudResultCreateMcTenantAddressResponseWithDefaults() *OBCloudResultCreateMcTenantAddressResponse`

NewOBCloudResultCreateMcTenantAddressResponseWithDefaults instantiates a new OBCloudResultCreateMcTenantAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetData() CreateMcTenantAddressResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetDataOk() (*CreateMcTenantAddressResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetData(v CreateMcTenantAddressResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultCreateMcTenantAddressResponse) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultCreateMcTenantAddressResponse) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultCreateMcTenantAddressResponse) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


