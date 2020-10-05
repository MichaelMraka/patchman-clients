# HostOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | A Red Hat Account number that owns the host. | 
**AnsibleHost** | Pointer to **NullableString** | The ansible host name for remediations | [optional] 
**BiosUuid** | Pointer to **NullableString** | A UUID of the host machine BIOS.  This field is considered to be a canonical fact. | [optional] 
**Created** | Pointer to **string** | A timestamp when the entry was created. | [optional] 
**CulledTimestamp** | Pointer to **NullableTime** | Timestamp from which the host is considered deleted. | [optional] 
**DisplayName** | Pointer to **NullableString** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 
**ExternalId** | Pointer to **NullableString** | Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact. | [optional] 
**Facts** | Pointer to [**[]FactSet**](FactSet.md) | A set of facts belonging to the host. | [optional] 
**Fqdn** | Pointer to **NullableString** | A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact. | [optional] 
**Id** | Pointer to **string** | A durable and reliable platform-wide host identifier. Applications should use this identifier to reference hosts. | [optional] 
**InsightsId** | Pointer to **NullableString** | An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact. | [optional] 
**IpAddresses** | Pointer to **[]string** | Host’s network IP addresses.  This field is considered to be a canonical fact. | [optional] 
**MacAddresses** | Pointer to **[]string** | Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact. | [optional] 
**Reporter** | Pointer to **NullableString** | Reporting source of the host. Used when updating the stale_timestamp. | [optional] 
**RhelMachineId** | Pointer to **NullableString** | A Machine ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SatelliteId** | Pointer to **NullableString** | A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**StaleTimestamp** | Pointer to **NullableTime** | Timestamp from which the host is considered stale. | [optional] 
**StaleWarningTimestamp** | Pointer to **NullableTime** | Timestamp from which the host is considered too stale to be listed without an explicit toggle. | [optional] 
**SubscriptionManagerId** | Pointer to **NullableString** | A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**Updated** | Pointer to **string** | A timestamp when the entry was last updated. | [optional] 

## Methods

### NewHostOut

`func NewHostOut(account string, ) *HostOut`

NewHostOut instantiates a new HostOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOutWithDefaults

`func NewHostOutWithDefaults() *HostOut`

NewHostOutWithDefaults instantiates a new HostOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *HostOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HostOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HostOut) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetAnsibleHost

`func (o *HostOut) GetAnsibleHost() string`

GetAnsibleHost returns the AnsibleHost field if non-nil, zero value otherwise.

### GetAnsibleHostOk

`func (o *HostOut) GetAnsibleHostOk() (*string, bool)`

GetAnsibleHostOk returns a tuple with the AnsibleHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleHost

`func (o *HostOut) SetAnsibleHost(v string)`

SetAnsibleHost sets AnsibleHost field to given value.

### HasAnsibleHost

`func (o *HostOut) HasAnsibleHost() bool`

HasAnsibleHost returns a boolean if a field has been set.

### SetAnsibleHostNil

`func (o *HostOut) SetAnsibleHostNil(b bool)`

 SetAnsibleHostNil sets the value for AnsibleHost to be an explicit nil

### UnsetAnsibleHost
`func (o *HostOut) UnsetAnsibleHost()`

UnsetAnsibleHost ensures that no value is present for AnsibleHost, not even an explicit nil
### GetBiosUuid

`func (o *HostOut) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *HostOut) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *HostOut) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *HostOut) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.

### SetBiosUuidNil

`func (o *HostOut) SetBiosUuidNil(b bool)`

 SetBiosUuidNil sets the value for BiosUuid to be an explicit nil

### UnsetBiosUuid
`func (o *HostOut) UnsetBiosUuid()`

UnsetBiosUuid ensures that no value is present for BiosUuid, not even an explicit nil
### GetCreated

`func (o *HostOut) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HostOut) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HostOut) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *HostOut) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCulledTimestamp

`func (o *HostOut) GetCulledTimestamp() string`

GetCulledTimestamp returns the CulledTimestamp field if non-nil, zero value otherwise.

### GetCulledTimestampOk

`func (o *HostOut) GetCulledTimestampOk() (*string, bool)`

GetCulledTimestampOk returns a tuple with the CulledTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulledTimestamp

`func (o *HostOut) SetCulledTimestamp(v string)`

SetCulledTimestamp sets CulledTimestamp field to given value.

### HasCulledTimestamp

`func (o *HostOut) HasCulledTimestamp() bool`

HasCulledTimestamp returns a boolean if a field has been set.

### SetCulledTimestampNil

`func (o *HostOut) SetCulledTimestampNil(b bool)`

 SetCulledTimestampNil sets the value for CulledTimestamp to be an explicit nil

### UnsetCulledTimestamp
`func (o *HostOut) UnsetCulledTimestamp()`

UnsetCulledTimestamp ensures that no value is present for CulledTimestamp, not even an explicit nil
### GetDisplayName

`func (o *HostOut) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HostOut) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HostOut) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HostOut) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *HostOut) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *HostOut) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExternalId

`func (o *HostOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *HostOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *HostOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *HostOut) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *HostOut) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *HostOut) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetFacts

`func (o *HostOut) GetFacts() []FactSet`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *HostOut) GetFactsOk() (*[]FactSet, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *HostOut) SetFacts(v []FactSet)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *HostOut) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetFqdn

`func (o *HostOut) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *HostOut) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *HostOut) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *HostOut) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *HostOut) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *HostOut) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetId

`func (o *HostOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsightsId

`func (o *HostOut) GetInsightsId() string`

GetInsightsId returns the InsightsId field if non-nil, zero value otherwise.

### GetInsightsIdOk

`func (o *HostOut) GetInsightsIdOk() (*string, bool)`

GetInsightsIdOk returns a tuple with the InsightsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsId

`func (o *HostOut) SetInsightsId(v string)`

SetInsightsId sets InsightsId field to given value.

### HasInsightsId

`func (o *HostOut) HasInsightsId() bool`

HasInsightsId returns a boolean if a field has been set.

### SetInsightsIdNil

`func (o *HostOut) SetInsightsIdNil(b bool)`

 SetInsightsIdNil sets the value for InsightsId to be an explicit nil

### UnsetInsightsId
`func (o *HostOut) UnsetInsightsId()`

UnsetInsightsId ensures that no value is present for InsightsId, not even an explicit nil
### GetIpAddresses

`func (o *HostOut) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *HostOut) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *HostOut) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *HostOut) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *HostOut) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *HostOut) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil
### GetMacAddresses

`func (o *HostOut) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *HostOut) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *HostOut) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *HostOut) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### SetMacAddressesNil

`func (o *HostOut) SetMacAddressesNil(b bool)`

 SetMacAddressesNil sets the value for MacAddresses to be an explicit nil

### UnsetMacAddresses
`func (o *HostOut) UnsetMacAddresses()`

UnsetMacAddresses ensures that no value is present for MacAddresses, not even an explicit nil
### GetReporter

`func (o *HostOut) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *HostOut) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *HostOut) SetReporter(v string)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *HostOut) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### SetReporterNil

`func (o *HostOut) SetReporterNil(b bool)`

 SetReporterNil sets the value for Reporter to be an explicit nil

### UnsetReporter
`func (o *HostOut) UnsetReporter()`

UnsetReporter ensures that no value is present for Reporter, not even an explicit nil
### GetRhelMachineId

`func (o *HostOut) GetRhelMachineId() string`

GetRhelMachineId returns the RhelMachineId field if non-nil, zero value otherwise.

### GetRhelMachineIdOk

`func (o *HostOut) GetRhelMachineIdOk() (*string, bool)`

GetRhelMachineIdOk returns a tuple with the RhelMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhelMachineId

`func (o *HostOut) SetRhelMachineId(v string)`

SetRhelMachineId sets RhelMachineId field to given value.

### HasRhelMachineId

`func (o *HostOut) HasRhelMachineId() bool`

HasRhelMachineId returns a boolean if a field has been set.

### SetRhelMachineIdNil

`func (o *HostOut) SetRhelMachineIdNil(b bool)`

 SetRhelMachineIdNil sets the value for RhelMachineId to be an explicit nil

### UnsetRhelMachineId
`func (o *HostOut) UnsetRhelMachineId()`

UnsetRhelMachineId ensures that no value is present for RhelMachineId, not even an explicit nil
### GetSatelliteId

`func (o *HostOut) GetSatelliteId() string`

GetSatelliteId returns the SatelliteId field if non-nil, zero value otherwise.

### GetSatelliteIdOk

`func (o *HostOut) GetSatelliteIdOk() (*string, bool)`

GetSatelliteIdOk returns a tuple with the SatelliteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteId

`func (o *HostOut) SetSatelliteId(v string)`

SetSatelliteId sets SatelliteId field to given value.

### HasSatelliteId

`func (o *HostOut) HasSatelliteId() bool`

HasSatelliteId returns a boolean if a field has been set.

### SetSatelliteIdNil

`func (o *HostOut) SetSatelliteIdNil(b bool)`

 SetSatelliteIdNil sets the value for SatelliteId to be an explicit nil

### UnsetSatelliteId
`func (o *HostOut) UnsetSatelliteId()`

UnsetSatelliteId ensures that no value is present for SatelliteId, not even an explicit nil
### GetStaleTimestamp

`func (o *HostOut) GetStaleTimestamp() string`

GetStaleTimestamp returns the StaleTimestamp field if non-nil, zero value otherwise.

### GetStaleTimestampOk

`func (o *HostOut) GetStaleTimestampOk() (*string, bool)`

GetStaleTimestampOk returns a tuple with the StaleTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTimestamp

`func (o *HostOut) SetStaleTimestamp(v string)`

SetStaleTimestamp sets StaleTimestamp field to given value.

### HasStaleTimestamp

`func (o *HostOut) HasStaleTimestamp() bool`

HasStaleTimestamp returns a boolean if a field has been set.

### SetStaleTimestampNil

`func (o *HostOut) SetStaleTimestampNil(b bool)`

 SetStaleTimestampNil sets the value for StaleTimestamp to be an explicit nil

### UnsetStaleTimestamp
`func (o *HostOut) UnsetStaleTimestamp()`

UnsetStaleTimestamp ensures that no value is present for StaleTimestamp, not even an explicit nil
### GetStaleWarningTimestamp

`func (o *HostOut) GetStaleWarningTimestamp() string`

GetStaleWarningTimestamp returns the StaleWarningTimestamp field if non-nil, zero value otherwise.

### GetStaleWarningTimestampOk

`func (o *HostOut) GetStaleWarningTimestampOk() (*string, bool)`

GetStaleWarningTimestampOk returns a tuple with the StaleWarningTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleWarningTimestamp

`func (o *HostOut) SetStaleWarningTimestamp(v string)`

SetStaleWarningTimestamp sets StaleWarningTimestamp field to given value.

### HasStaleWarningTimestamp

`func (o *HostOut) HasStaleWarningTimestamp() bool`

HasStaleWarningTimestamp returns a boolean if a field has been set.

### SetStaleWarningTimestampNil

`func (o *HostOut) SetStaleWarningTimestampNil(b bool)`

 SetStaleWarningTimestampNil sets the value for StaleWarningTimestamp to be an explicit nil

### UnsetStaleWarningTimestamp
`func (o *HostOut) UnsetStaleWarningTimestamp()`

UnsetStaleWarningTimestamp ensures that no value is present for StaleWarningTimestamp, not even an explicit nil
### GetSubscriptionManagerId

`func (o *HostOut) GetSubscriptionManagerId() string`

GetSubscriptionManagerId returns the SubscriptionManagerId field if non-nil, zero value otherwise.

### GetSubscriptionManagerIdOk

`func (o *HostOut) GetSubscriptionManagerIdOk() (*string, bool)`

GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionManagerId

`func (o *HostOut) SetSubscriptionManagerId(v string)`

SetSubscriptionManagerId sets SubscriptionManagerId field to given value.

### HasSubscriptionManagerId

`func (o *HostOut) HasSubscriptionManagerId() bool`

HasSubscriptionManagerId returns a boolean if a field has been set.

### SetSubscriptionManagerIdNil

`func (o *HostOut) SetSubscriptionManagerIdNil(b bool)`

 SetSubscriptionManagerIdNil sets the value for SubscriptionManagerId to be an explicit nil

### UnsetSubscriptionManagerId
`func (o *HostOut) UnsetSubscriptionManagerId()`

UnsetSubscriptionManagerId ensures that no value is present for SubscriptionManagerId, not even an explicit nil
### GetUpdated

`func (o *HostOut) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *HostOut) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *HostOut) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *HostOut) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


