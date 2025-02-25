# OBCloudResultCreateTenantPlServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CreateTenantPlServiceResponse**](CreateTenantPlServiceResponse.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultCreateTenantPlServiceResponse

`func NewOBCloudResultCreateTenantPlServiceResponse() *OBCloudResultCreateTenantPlServiceResponse`

NewOBCloudResultCreateTenantPlServiceResponse instantiates a new OBCloudResultCreateTenantPlServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultCreateTenantPlServiceResponseWithDefaults

`func NewOBCloudResultCreateTenantPlServiceResponseWithDefaults() *OBCloudResultCreateTenantPlServiceResponse`

NewOBCloudResultCreateTenantPlServiceResponseWithDefaults instantiates a new OBCloudResultCreateTenantPlServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetData() CreateTenantPlServiceResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetDataOk() (*CreateTenantPlServiceResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetData(v CreateTenantPlServiceResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultCreateTenantPlServiceResponse) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultCreateTenantPlServiceResponse) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultCreateTenantPlServiceResponse) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


