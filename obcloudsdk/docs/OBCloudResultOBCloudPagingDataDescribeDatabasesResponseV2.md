# OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**OBCloudPagingDataDescribeDatabasesResponseV2**](OBCloudPagingDataDescribeDatabasesResponseV2.md) |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2

`func NewOBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2() *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2`

NewOBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2 instantiates a new OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2WithDefaults

`func NewOBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2WithDefaults() *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2`

NewOBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2WithDefaults instantiates a new OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrorCode

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetData

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetData() OBCloudPagingDataDescribeDatabasesResponseV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetDataOk() (*OBCloudPagingDataDescribeDatabasesResponseV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetData(v OBCloudPagingDataDescribeDatabasesResponseV2)`

SetData sets Data field to given value.

### HasData

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCost

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetServer

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetRequestId

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotalCount

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetExtra

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


