# OBCloudResultCreateInstanceResponseV2OpenApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CreateInstanceResponseV2OpenApi**](CreateInstanceResponseV2OpenApi.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultCreateInstanceResponseV2OpenApi

`func NewOBCloudResultCreateInstanceResponseV2OpenApi() *OBCloudResultCreateInstanceResponseV2OpenApi`

NewOBCloudResultCreateInstanceResponseV2OpenApi instantiates a new OBCloudResultCreateInstanceResponseV2OpenApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultCreateInstanceResponseV2OpenApiWithDefaults

`func NewOBCloudResultCreateInstanceResponseV2OpenApiWithDefaults() *OBCloudResultCreateInstanceResponseV2OpenApi`

NewOBCloudResultCreateInstanceResponseV2OpenApiWithDefaults instantiates a new OBCloudResultCreateInstanceResponseV2OpenApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetData() CreateInstanceResponseV2OpenApi`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetDataOk() (*CreateInstanceResponseV2OpenApi, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetData(v CreateInstanceResponseV2OpenApi)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultCreateInstanceResponseV2OpenApi) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


