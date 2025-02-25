# MultiCloudModifyInstanceSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceClass** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int32** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 

## Methods

### NewMultiCloudModifyInstanceSpecRequest

`func NewMultiCloudModifyInstanceSpecRequest() *MultiCloudModifyInstanceSpecRequest`

NewMultiCloudModifyInstanceSpecRequest instantiates a new MultiCloudModifyInstanceSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiCloudModifyInstanceSpecRequestWithDefaults

`func NewMultiCloudModifyInstanceSpecRequestWithDefaults() *MultiCloudModifyInstanceSpecRequest`

NewMultiCloudModifyInstanceSpecRequestWithDefaults instantiates a new MultiCloudModifyInstanceSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceClass

`func (o *MultiCloudModifyInstanceSpecRequest) GetInstanceClass() string`

GetInstanceClass returns the InstanceClass field if non-nil, zero value otherwise.

### GetInstanceClassOk

`func (o *MultiCloudModifyInstanceSpecRequest) GetInstanceClassOk() (*string, bool)`

GetInstanceClassOk returns a tuple with the InstanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceClass

`func (o *MultiCloudModifyInstanceSpecRequest) SetInstanceClass(v string)`

SetInstanceClass sets InstanceClass field to given value.

### HasInstanceClass

`func (o *MultiCloudModifyInstanceSpecRequest) HasInstanceClass() bool`

HasInstanceClass returns a boolean if a field has been set.

### GetDiskSize

`func (o *MultiCloudModifyInstanceSpecRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *MultiCloudModifyInstanceSpecRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *MultiCloudModifyInstanceSpecRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *MultiCloudModifyInstanceSpecRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetDryRun

`func (o *MultiCloudModifyInstanceSpecRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *MultiCloudModifyInstanceSpecRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *MultiCloudModifyInstanceSpecRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *MultiCloudModifyInstanceSpecRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


