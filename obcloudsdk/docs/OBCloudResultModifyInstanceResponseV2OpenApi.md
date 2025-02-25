# OBCloudResultModifyInstanceResponseV2OpenApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**ModifyInstanceResponseV2OpenApi**](ModifyInstanceResponseV2OpenApi.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultModifyInstanceResponseV2OpenApi

`func NewOBCloudResultModifyInstanceResponseV2OpenApi() *OBCloudResultModifyInstanceResponseV2OpenApi`

NewOBCloudResultModifyInstanceResponseV2OpenApi instantiates a new OBCloudResultModifyInstanceResponseV2OpenApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultModifyInstanceResponseV2OpenApiWithDefaults

`func NewOBCloudResultModifyInstanceResponseV2OpenApiWithDefaults() *OBCloudResultModifyInstanceResponseV2OpenApi`

NewOBCloudResultModifyInstanceResponseV2OpenApiWithDefaults instantiates a new OBCloudResultModifyInstanceResponseV2OpenApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetData() ModifyInstanceResponseV2OpenApi`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetDataOk() (*ModifyInstanceResponseV2OpenApi, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetData(v ModifyInstanceResponseV2OpenApi)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultModifyInstanceResponseV2OpenApi) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


