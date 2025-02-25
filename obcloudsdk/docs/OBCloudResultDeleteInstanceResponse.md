# OBCloudResultDeleteInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultDeleteInstanceResponse

`func NewOBCloudResultDeleteInstanceResponse() *OBCloudResultDeleteInstanceResponse`

NewOBCloudResultDeleteInstanceResponse instantiates a new OBCloudResultDeleteInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultDeleteInstanceResponseWithDefaults

`func NewOBCloudResultDeleteInstanceResponseWithDefaults() *OBCloudResultDeleteInstanceResponse`

NewOBCloudResultDeleteInstanceResponseWithDefaults instantiates a new OBCloudResultDeleteInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultDeleteInstanceResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultDeleteInstanceResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultDeleteInstanceResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultDeleteInstanceResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultDeleteInstanceResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultDeleteInstanceResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultDeleteInstanceResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultDeleteInstanceResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultDeleteInstanceResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultDeleteInstanceResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultDeleteInstanceResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultDeleteInstanceResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultDeleteInstanceResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultDeleteInstanceResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultDeleteInstanceResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultDeleteInstanceResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultDeleteInstanceResponse) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultDeleteInstanceResponse) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultDeleteInstanceResponse) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultDeleteInstanceResponse) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultDeleteInstanceResponse) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultDeleteInstanceResponse) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultDeleteInstanceResponse) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultDeleteInstanceResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultDeleteInstanceResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultDeleteInstanceResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultDeleteInstanceResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultDeleteInstanceResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultDeleteInstanceResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultDeleteInstanceResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultDeleteInstanceResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultDeleteInstanceResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultDeleteInstanceResponse) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultDeleteInstanceResponse) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultDeleteInstanceResponse) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultDeleteInstanceResponse) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


