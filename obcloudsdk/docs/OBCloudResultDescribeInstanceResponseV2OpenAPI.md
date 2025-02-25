# OBCloudResultDescribeInstanceResponseV2OpenAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**DescribeInstanceResponseV2OpenAPI**](DescribeInstanceResponseV2OpenAPI.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultDescribeInstanceResponseV2OpenAPI

`func NewOBCloudResultDescribeInstanceResponseV2OpenAPI() *OBCloudResultDescribeInstanceResponseV2OpenAPI`

NewOBCloudResultDescribeInstanceResponseV2OpenAPI instantiates a new OBCloudResultDescribeInstanceResponseV2OpenAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultDescribeInstanceResponseV2OpenAPIWithDefaults

`func NewOBCloudResultDescribeInstanceResponseV2OpenAPIWithDefaults() *OBCloudResultDescribeInstanceResponseV2OpenAPI`

NewOBCloudResultDescribeInstanceResponseV2OpenAPIWithDefaults instantiates a new OBCloudResultDescribeInstanceResponseV2OpenAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetData() DescribeInstanceResponseV2OpenAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetDataOk() (*DescribeInstanceResponseV2OpenAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetData(v DescribeInstanceResponseV2OpenAPI)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultDescribeInstanceResponseV2OpenAPI) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


