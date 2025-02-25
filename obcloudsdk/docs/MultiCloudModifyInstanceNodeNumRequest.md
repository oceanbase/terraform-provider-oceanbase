# MultiCloudModifyInstanceNodeNumRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeNum** | **int32** |  | 
**DryRun** | Pointer to **bool** |  | [optional] 

## Methods

### NewMultiCloudModifyInstanceNodeNumRequest

`func NewMultiCloudModifyInstanceNodeNumRequest(nodeNum int32, ) *MultiCloudModifyInstanceNodeNumRequest`

NewMultiCloudModifyInstanceNodeNumRequest instantiates a new MultiCloudModifyInstanceNodeNumRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudModifyInstanceNodeNumRequestWithDefaults

`func NewMultiCloudModifyInstanceNodeNumRequestWithDefaults() *MultiCloudModifyInstanceNodeNumRequest`

NewMultiCloudModifyInstanceNodeNumRequestWithDefaults instantiates a new MultiCloudModifyInstanceNodeNumRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeNum

`func (o *MultiCloudModifyInstanceNodeNumRequest) GetNodeNum() int32`

GetNodeNum returns the NodeNum field if non-nil, zero value otherwise.

### GetNodeNumOk

`func (o *MultiCloudModifyInstanceNodeNumRequest) GetNodeNumOk() (*int32, bool)`

GetNodeNumOk returns a tuple with the NodeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNum

`func (o *MultiCloudModifyInstanceNodeNumRequest) SetNodeNum(v int32)`

SetNodeNum sets NodeNum field to given value.


### GetDryRun

`func (o *MultiCloudModifyInstanceNodeNumRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *MultiCloudModifyInstanceNodeNumRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *MultiCloudModifyInstanceNodeNumRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *MultiCloudModifyInstanceNodeNumRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


