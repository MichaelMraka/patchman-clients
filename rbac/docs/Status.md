# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **int64** |  | 
**Commit** | Pointer to **string** |  | [optional] 
**ServerAddress** | Pointer to **string** |  | [optional] 
**PlatformInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**PythonVersion** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStatus

`func NewStatus(apiVersion int64, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Status) GetApiVersion() int64`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Status) GetApiVersionOk() (*int64, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Status) SetApiVersion(v int64)`

SetApiVersion sets ApiVersion field to given value.


### GetCommit

`func (o *Status) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *Status) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *Status) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *Status) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetServerAddress

`func (o *Status) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *Status) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *Status) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *Status) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetPlatformInfo

`func (o *Status) GetPlatformInfo() map[string]interface{}`

GetPlatformInfo returns the PlatformInfo field if non-nil, zero value otherwise.

### GetPlatformInfoOk

`func (o *Status) GetPlatformInfoOk() (*map[string]interface{}, bool)`

GetPlatformInfoOk returns a tuple with the PlatformInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformInfo

`func (o *Status) SetPlatformInfo(v map[string]interface{})`

SetPlatformInfo sets PlatformInfo field to given value.

### HasPlatformInfo

`func (o *Status) HasPlatformInfo() bool`

HasPlatformInfo returns a boolean if a field has been set.

### GetPythonVersion

`func (o *Status) GetPythonVersion() string`

GetPythonVersion returns the PythonVersion field if non-nil, zero value otherwise.

### GetPythonVersionOk

`func (o *Status) GetPythonVersionOk() (*string, bool)`

GetPythonVersionOk returns a tuple with the PythonVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonVersion

`func (o *Status) SetPythonVersion(v string)`

SetPythonVersion sets PythonVersion field to given value.

### HasPythonVersion

`func (o *Status) HasPythonVersion() bool`

HasPythonVersion returns a boolean if a field has been set.

### GetModules

`func (o *Status) GetModules() map[string]interface{}`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *Status) GetModulesOk() (*map[string]interface{}, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *Status) SetModules(v map[string]interface{})`

SetModules sets Modules field to given value.

### HasModules

`func (o *Status) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


