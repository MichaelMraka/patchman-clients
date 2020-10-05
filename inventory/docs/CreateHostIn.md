# CreateHostIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | A Red Hat Account number that owns the host. | 
**AnsibleHost** | Pointer to **string** | The ansible host name for remediations | [optional] 
**BiosUuid** | Pointer to **string** | A UUID of the host machine BIOS.  This field is considered to be a canonical fact. | [optional] 
**DisplayName** | Pointer to **string** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 
**ExternalId** | Pointer to **string** | Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact. | [optional] 
**Facts** | Pointer to [**[]FactSet**](FactSet.md) | A set of facts belonging to the host. | [optional] 
**Fqdn** | Pointer to **string** | A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact. | [optional] 
**InsightsId** | Pointer to **string** | An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact. | [optional] 
**IpAddresses** | Pointer to **[]string** | Host’s network IP addresses.  This field is considered to be a canonical fact. | [optional] 
**MacAddresses** | Pointer to **[]string** | Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact. | [optional] 
**Reporter** | **string** | Reporting source of the host. Used when updating the stale_timestamp. | 
**RhelMachineId** | Pointer to **string** | A Machine ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SatelliteId** | Pointer to **string** | A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**StaleTimestamp** | **string** | Timestamp from which the host is considered stale. | 
**SubscriptionManagerId** | Pointer to **string** | A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SystemProfile** | Pointer to [**SystemProfile**](SystemProfile.md) |  | [optional] 

## Methods

### NewCreateHostIn

`func NewCreateHostIn(account string, reporter string, staleTimestamp string, ) *CreateHostIn`

NewCreateHostIn instantiates a new CreateHostIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHostInWithDefaults

`func NewCreateHostInWithDefaults() *CreateHostIn`

NewCreateHostInWithDefaults instantiates a new CreateHostIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateHostIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateHostIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateHostIn) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetAnsibleHost

`func (o *CreateHostIn) GetAnsibleHost() string`

GetAnsibleHost returns the AnsibleHost field if non-nil, zero value otherwise.

### GetAnsibleHostOk

`func (o *CreateHostIn) GetAnsibleHostOk() (*string, bool)`

GetAnsibleHostOk returns a tuple with the AnsibleHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHost

`func (o *CreateHostIn) SetAnsibleHost(v string)`

SetAnsibleHost sets AnsibleHost field to given value.

### HasAnsibleHost

`func (o *CreateHostIn) HasAnsibleHost() bool`

HasAnsibleHost returns a boolean if a field has been set.

### GetBiosUuid

`func (o *CreateHostIn) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *CreateHostIn) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *CreateHostIn) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *CreateHostIn) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateHostIn) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateHostIn) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateHostIn) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateHostIn) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateHostIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateHostIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateHostIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateHostIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFacts

`func (o *CreateHostIn) GetFacts() []FactSet`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CreateHostIn) GetFactsOk() (*[]FactSet, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CreateHostIn) SetFacts(v []FactSet)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *CreateHostIn) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetFqdn

`func (o *CreateHostIn) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateHostIn) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateHostIn) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CreateHostIn) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInsightsId

`func (o *CreateHostIn) GetInsightsId() string`

GetInsightsId returns the InsightsId field if non-nil, zero value otherwise.

### GetInsightsIdOk

`func (o *CreateHostIn) GetInsightsIdOk() (*string, bool)`

GetInsightsIdOk returns a tuple with the InsightsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsId

`func (o *CreateHostIn) SetInsightsId(v string)`

SetInsightsId sets InsightsId field to given value.

### HasInsightsId

`func (o *CreateHostIn) HasInsightsId() bool`

HasInsightsId returns a boolean if a field has been set.

### GetIpAddresses

`func (o *CreateHostIn) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *CreateHostIn) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *CreateHostIn) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *CreateHostIn) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMacAddresses

`func (o *CreateHostIn) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *CreateHostIn) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *CreateHostIn) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *CreateHostIn) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### GetReporter

`func (o *CreateHostIn) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *CreateHostIn) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *CreateHostIn) SetReporter(v string)`

SetReporter sets Reporter field to given value.


### GetRhelMachineId

`func (o *CreateHostIn) GetRhelMachineId() string`

GetRhelMachineId returns the RhelMachineId field if non-nil, zero value otherwise.

### GetRhelMachineIdOk

`func (o *CreateHostIn) GetRhelMachineIdOk() (*string, bool)`

GetRhelMachineIdOk returns a tuple with the RhelMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhelMachineId

`func (o *CreateHostIn) SetRhelMachineId(v string)`

SetRhelMachineId sets RhelMachineId field to given value.

### HasRhelMachineId

`func (o *CreateHostIn) HasRhelMachineId() bool`

HasRhelMachineId returns a boolean if a field has been set.

### GetSatelliteId

`func (o *CreateHostIn) GetSatelliteId() string`

GetSatelliteId returns the SatelliteId field if non-nil, zero value otherwise.

### GetSatelliteIdOk

`func (o *CreateHostIn) GetSatelliteIdOk() (*string, bool)`

GetSatelliteIdOk returns a tuple with the SatelliteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteId

`func (o *CreateHostIn) SetSatelliteId(v string)`

SetSatelliteId sets SatelliteId field to given value.

### HasSatelliteId

`func (o *CreateHostIn) HasSatelliteId() bool`

HasSatelliteId returns a boolean if a field has been set.

### GetStaleTimestamp

`func (o *CreateHostIn) GetStaleTimestamp() string`

GetStaleTimestamp returns the StaleTimestamp field if non-nil, zero value otherwise.

### GetStaleTimestampOk

`func (o *CreateHostIn) GetStaleTimestampOk() (*string, bool)`

GetStaleTimestampOk returns a tuple with the StaleTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTimestamp

`func (o *CreateHostIn) SetStaleTimestamp(v string)`

SetStaleTimestamp sets StaleTimestamp field to given value.


### GetSubscriptionManagerId

`func (o *CreateHostIn) GetSubscriptionManagerId() string`

GetSubscriptionManagerId returns the SubscriptionManagerId field if non-nil, zero value otherwise.

### GetSubscriptionManagerIdOk

`func (o *CreateHostIn) GetSubscriptionManagerIdOk() (*string, bool)`

GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionManagerId

`func (o *CreateHostIn) SetSubscriptionManagerId(v string)`

SetSubscriptionManagerId sets SubscriptionManagerId field to given value.

### HasSubscriptionManagerId

`func (o *CreateHostIn) HasSubscriptionManagerId() bool`

HasSubscriptionManagerId returns a boolean if a field has been set.

### GetSystemProfile

`func (o *CreateHostIn) GetSystemProfile() SystemProfile`

GetSystemProfile returns the SystemProfile field if non-nil, zero value otherwise.

### GetSystemProfileOk

`func (o *CreateHostIn) GetSystemProfileOk() (*SystemProfile, bool)`

GetSystemProfileOk returns a tuple with the SystemProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProfile

`func (o *CreateHostIn) SetSystemProfile(v SystemProfile)`

SetSystemProfile sets SystemProfile field to given value.

### HasSystemProfile

`func (o *CreateHostIn) HasSystemProfile() bool`

HasSystemProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


