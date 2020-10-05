# SystemProfileNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Addresses** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **string** | MAC address (with or without colons) | [optional] 
**Mtu** | Pointer to **int32** | MTU (Maximum transmission unit) | [optional] 
**Name** | Pointer to **string** | Name of interface | [optional] 
**State** | Pointer to **string** | Interface state (UP, DOWN, UNKNOWN) | [optional] 
**Type** | Pointer to **string** | Interface type (ether, loopback) | [optional] 

## Methods

### NewSystemProfileNetworkInterfaces

`func NewSystemProfileNetworkInterfaces() *SystemProfileNetworkInterfaces`

NewSystemProfileNetworkInterfaces instantiates a new SystemProfileNetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileNetworkInterfacesWithDefaults

`func NewSystemProfileNetworkInterfacesWithDefaults() *SystemProfileNetworkInterfaces`

NewSystemProfileNetworkInterfacesWithDefaults instantiates a new SystemProfileNetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addresses

`func (o *SystemProfileNetworkInterfaces) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *SystemProfileNetworkInterfaces) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *SystemProfileNetworkInterfaces) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *SystemProfileNetworkInterfaces) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *SystemProfileNetworkInterfaces) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *SystemProfileNetworkInterfaces) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *SystemProfileNetworkInterfaces) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *SystemProfileNetworkInterfaces) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetMacAddress

`func (o *SystemProfileNetworkInterfaces) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *SystemProfileNetworkInterfaces) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *SystemProfileNetworkInterfaces) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *SystemProfileNetworkInterfaces) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMtu

`func (o *SystemProfileNetworkInterfaces) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SystemProfileNetworkInterfaces) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SystemProfileNetworkInterfaces) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SystemProfileNetworkInterfaces) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *SystemProfileNetworkInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemProfileNetworkInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemProfileNetworkInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemProfileNetworkInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *SystemProfileNetworkInterfaces) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SystemProfileNetworkInterfaces) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SystemProfileNetworkInterfaces) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SystemProfileNetworkInterfaces) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *SystemProfileNetworkInterfaces) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemProfileNetworkInterfaces) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemProfileNetworkInterfaces) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemProfileNetworkInterfaces) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


