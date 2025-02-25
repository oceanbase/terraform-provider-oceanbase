# OBCloudResultModifyDatabaseDescriptionResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**ModifyDatabaseDescriptionResponseV2**](ModifyDatabaseDescriptionResponseV2.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultModifyDatabaseDescriptionResponseV2

`func NewOBCloudResultModifyDatabaseDescriptionResponseV2() *OBCloudResultModifyDatabaseDescriptionResponseV2`

NewOBCloudResultModifyDatabaseDescriptionResponseV2 instantiates a new OBCloudResultModifyDatabaseDescriptionResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultModifyDatabaseDescriptionResponseV2WithDefaults

`func NewOBCloudResultModifyDatabaseDescriptionResponseV2WithDefaults() *OBCloudResultModifyDatabaseDescriptionResponseV2`

NewOBCloudResultModifyDatabaseDescriptionResponseV2WithDefaults instantiates a new OBCloudResultModifyDatabaseDescriptionResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetData() ModifyDatabaseDescriptionResponseV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetDataOk() (*ModifyDatabaseDescriptionResponseV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetData(v ModifyDatabaseDescriptionResponseV2)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultModifyDatabaseDescriptionResponseV2) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


