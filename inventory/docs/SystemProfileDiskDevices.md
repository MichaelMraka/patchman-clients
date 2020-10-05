# SystemProfileDiskDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | User-defined mount label | [optional] 
**MountPoint** | Pointer to **string** | The mount point | [optional] 
**Options** | Pointer to **map[string]interface{}** | Mount options for nested object | [optional] 
**Type** | Pointer to **string** | The mount type | [optional] 

## Methods

### NewSystemProfileDiskDevices

`func NewSystemProfileDiskDevices() *SystemProfileDiskDevices`

NewSystemProfileDiskDevices instantiates a new SystemProfileDiskDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileDiskDevicesWithDefaults

`func NewSystemProfileDiskDevicesWithDefaults() *SystemProfileDiskDevices`

NewSystemProfileDiskDevicesWithDefaults instantiates a new SystemProfileDiskDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SystemProfileDiskDevices) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SystemProfileDiskDevices) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SystemProfileDiskDevices) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SystemProfileDiskDevices) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLabel

`func (o *SystemProfileDiskDevices) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SystemProfileDiskDevices) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SystemProfileDiskDevices) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SystemProfileDiskDevices) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMountPoint

`func (o *SystemProfileDiskDevices) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *SystemProfileDiskDevices) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *SystemProfileDiskDevices) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *SystemProfileDiskDevices) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.

### GetOptions

`func (o *SystemProfileDiskDevices) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemProfileDiskDevices) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemProfileDiskDevices) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemProfileDiskDevices) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetType

`func (o *SystemProfileDiskDevices) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemProfileDiskDevices) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemProfileDiskDevices) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemProfileDiskDevices) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


