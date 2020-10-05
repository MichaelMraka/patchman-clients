# SystemProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**BiosReleaseDate** | Pointer to **string** |  | [optional] 
**BiosVendor** | Pointer to **string** |  | [optional] 
**BiosVersion** | Pointer to **string** |  | [optional] 
**CapturedDate** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CoresPerSocket** | Pointer to **int32** |  | [optional] 
**CpuFlags** | Pointer to **[]string** |  | [optional] 
**DiskDevices** | Pointer to [**[]SystemProfileDiskDevices**](SystemProfile_disk_devices.md) |  | [optional] 
**DnfModules** | Pointer to [**[]SystemProfileDnfModules**](SystemProfile_dnf_modules.md) |  | [optional] 
**EnabledServices** | Pointer to **[]string** |  | [optional] 
**InfrastructureType** | Pointer to **string** |  | [optional] 
**InfrastructureVendor** | Pointer to **string** |  | [optional] 
**InsightsClientVersion** | Pointer to **string** |  | [optional] 
**InsightsEggVersion** | Pointer to **string** |  | [optional] 
**InstalledPackages** | Pointer to **[]string** |  | [optional] 
**InstalledProducts** | Pointer to [**[]SystemProfileInstalledProducts**](SystemProfile_installed_products.md) |  | [optional] 
**InstalledServices** | Pointer to **[]string** |  | [optional] 
**KatelloAgentRunning** | Pointer to **bool** |  | [optional] 
**KernelModules** | Pointer to **[]string** |  | [optional] 
**LastBootTime** | Pointer to **string** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]SystemProfileNetworkInterfaces**](SystemProfile_network_interfaces.md) |  | [optional] 
**NumberOfCpus** | Pointer to **int32** |  | [optional] 
**NumberOfSockets** | Pointer to **int32** |  | [optional] 
**OsKernelVersion** | Pointer to **string** |  | [optional] 
**OsRelease** | Pointer to **string** |  | [optional] 
**RunningProcesses** | Pointer to **[]string** |  | [optional] 
**SapSids** | Pointer to **[]string** | List of SAP SIDs | [optional] 
**SapSystem** | Pointer to **bool** | Indicates if SAP is installed on the system | [optional] 
**SatelliteManaged** | Pointer to **bool** |  | [optional] 
**SubscriptionAutoAttach** | Pointer to **string** |  | [optional] 
**SubscriptionStatus** | Pointer to **string** |  | [optional] 
**SystemMemoryBytes** | Pointer to **int64** |  | [optional] 
**TunedProfile** | Pointer to **string** | Current profile resulting from command tuned-adm active | [optional] 
**YumRepos** | Pointer to [**[]SystemProfileYumRepos**](SystemProfile_yum_repos.md) |  | [optional] 

## Methods

### NewSystemProfile

`func NewSystemProfile() *SystemProfile`

NewSystemProfile instantiates a new SystemProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileWithDefaults

`func NewSystemProfileWithDefaults() *SystemProfile`

NewSystemProfileWithDefaults instantiates a new SystemProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *SystemProfile) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *SystemProfile) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *SystemProfile) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *SystemProfile) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBiosReleaseDate

`func (o *SystemProfile) GetBiosReleaseDate() string`

GetBiosReleaseDate returns the BiosReleaseDate field if non-nil, zero value otherwise.

### GetBiosReleaseDateOk

`func (o *SystemProfile) GetBiosReleaseDateOk() (*string, bool)`

GetBiosReleaseDateOk returns a tuple with the BiosReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosReleaseDate

`func (o *SystemProfile) SetBiosReleaseDate(v string)`

SetBiosReleaseDate sets BiosReleaseDate field to given value.

### HasBiosReleaseDate

`func (o *SystemProfile) HasBiosReleaseDate() bool`

HasBiosReleaseDate returns a boolean if a field has been set.

### GetBiosVendor

`func (o *SystemProfile) GetBiosVendor() string`

GetBiosVendor returns the BiosVendor field if non-nil, zero value otherwise.

### GetBiosVendorOk

`func (o *SystemProfile) GetBiosVendorOk() (*string, bool)`

GetBiosVendorOk returns a tuple with the BiosVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVendor

`func (o *SystemProfile) SetBiosVendor(v string)`

SetBiosVendor sets BiosVendor field to given value.

### HasBiosVendor

`func (o *SystemProfile) HasBiosVendor() bool`

HasBiosVendor returns a boolean if a field has been set.

### GetBiosVersion

`func (o *SystemProfile) GetBiosVersion() string`

GetBiosVersion returns the BiosVersion field if non-nil, zero value otherwise.

### GetBiosVersionOk

`func (o *SystemProfile) GetBiosVersionOk() (*string, bool)`

GetBiosVersionOk returns a tuple with the BiosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVersion

`func (o *SystemProfile) SetBiosVersion(v string)`

SetBiosVersion sets BiosVersion field to given value.

### HasBiosVersion

`func (o *SystemProfile) HasBiosVersion() bool`

HasBiosVersion returns a boolean if a field has been set.

### GetCapturedDate

`func (o *SystemProfile) GetCapturedDate() string`

GetCapturedDate returns the CapturedDate field if non-nil, zero value otherwise.

### GetCapturedDateOk

`func (o *SystemProfile) GetCapturedDateOk() (*string, bool)`

GetCapturedDateOk returns a tuple with the CapturedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedDate

`func (o *SystemProfile) SetCapturedDate(v string)`

SetCapturedDate sets CapturedDate field to given value.

### HasCapturedDate

`func (o *SystemProfile) HasCapturedDate() bool`

HasCapturedDate returns a boolean if a field has been set.

### GetCloudProvider

`func (o *SystemProfile) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SystemProfile) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SystemProfile) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *SystemProfile) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *SystemProfile) GetCoresPerSocket() int32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *SystemProfile) GetCoresPerSocketOk() (*int32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *SystemProfile) SetCoresPerSocket(v int32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *SystemProfile) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCpuFlags

`func (o *SystemProfile) GetCpuFlags() []string`

GetCpuFlags returns the CpuFlags field if non-nil, zero value otherwise.

### GetCpuFlagsOk

`func (o *SystemProfile) GetCpuFlagsOk() (*[]string, bool)`

GetCpuFlagsOk returns a tuple with the CpuFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFlags

`func (o *SystemProfile) SetCpuFlags(v []string)`

SetCpuFlags sets CpuFlags field to given value.

### HasCpuFlags

`func (o *SystemProfile) HasCpuFlags() bool`

HasCpuFlags returns a boolean if a field has been set.

### GetDiskDevices

`func (o *SystemProfile) GetDiskDevices() []SystemProfileDiskDevices`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *SystemProfile) GetDiskDevicesOk() (*[]SystemProfileDiskDevices, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *SystemProfile) SetDiskDevices(v []SystemProfileDiskDevices)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *SystemProfile) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetDnfModules

`func (o *SystemProfile) GetDnfModules() []SystemProfileDnfModules`

GetDnfModules returns the DnfModules field if non-nil, zero value otherwise.

### GetDnfModulesOk

`func (o *SystemProfile) GetDnfModulesOk() (*[]SystemProfileDnfModules, bool)`

GetDnfModulesOk returns a tuple with the DnfModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnfModules

`func (o *SystemProfile) SetDnfModules(v []SystemProfileDnfModules)`

SetDnfModules sets DnfModules field to given value.

### HasDnfModules

`func (o *SystemProfile) HasDnfModules() bool`

HasDnfModules returns a boolean if a field has been set.

### GetEnabledServices

`func (o *SystemProfile) GetEnabledServices() []string`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *SystemProfile) GetEnabledServicesOk() (*[]string, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledServices

`func (o *SystemProfile) SetEnabledServices(v []string)`

SetEnabledServices sets EnabledServices field to given value.

### HasEnabledServices

`func (o *SystemProfile) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### GetInfrastructureType

`func (o *SystemProfile) GetInfrastructureType() string`

GetInfrastructureType returns the InfrastructureType field if non-nil, zero value otherwise.

### GetInfrastructureTypeOk

`func (o *SystemProfile) GetInfrastructureTypeOk() (*string, bool)`

GetInfrastructureTypeOk returns a tuple with the InfrastructureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureType

`func (o *SystemProfile) SetInfrastructureType(v string)`

SetInfrastructureType sets InfrastructureType field to given value.

### HasInfrastructureType

`func (o *SystemProfile) HasInfrastructureType() bool`

HasInfrastructureType returns a boolean if a field has been set.

### GetInfrastructureVendor

`func (o *SystemProfile) GetInfrastructureVendor() string`

GetInfrastructureVendor returns the InfrastructureVendor field if non-nil, zero value otherwise.

### GetInfrastructureVendorOk

`func (o *SystemProfile) GetInfrastructureVendorOk() (*string, bool)`

GetInfrastructureVendorOk returns a tuple with the InfrastructureVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureVendor

`func (o *SystemProfile) SetInfrastructureVendor(v string)`

SetInfrastructureVendor sets InfrastructureVendor field to given value.

### HasInfrastructureVendor

`func (o *SystemProfile) HasInfrastructureVendor() bool`

HasInfrastructureVendor returns a boolean if a field has been set.

### GetInsightsClientVersion

`func (o *SystemProfile) GetInsightsClientVersion() string`

GetInsightsClientVersion returns the InsightsClientVersion field if non-nil, zero value otherwise.

### GetInsightsClientVersionOk

`func (o *SystemProfile) GetInsightsClientVersionOk() (*string, bool)`

GetInsightsClientVersionOk returns a tuple with the InsightsClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsClientVersion

`func (o *SystemProfile) SetInsightsClientVersion(v string)`

SetInsightsClientVersion sets InsightsClientVersion field to given value.

### HasInsightsClientVersion

`func (o *SystemProfile) HasInsightsClientVersion() bool`

HasInsightsClientVersion returns a boolean if a field has been set.

### GetInsightsEggVersion

`func (o *SystemProfile) GetInsightsEggVersion() string`

GetInsightsEggVersion returns the InsightsEggVersion field if non-nil, zero value otherwise.

### GetInsightsEggVersionOk

`func (o *SystemProfile) GetInsightsEggVersionOk() (*string, bool)`

GetInsightsEggVersionOk returns a tuple with the InsightsEggVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsEggVersion

`func (o *SystemProfile) SetInsightsEggVersion(v string)`

SetInsightsEggVersion sets InsightsEggVersion field to given value.

### HasInsightsEggVersion

`func (o *SystemProfile) HasInsightsEggVersion() bool`

HasInsightsEggVersion returns a boolean if a field has been set.

### GetInstalledPackages

`func (o *SystemProfile) GetInstalledPackages() []string`

GetInstalledPackages returns the InstalledPackages field if non-nil, zero value otherwise.

### GetInstalledPackagesOk

`func (o *SystemProfile) GetInstalledPackagesOk() (*[]string, bool)`

GetInstalledPackagesOk returns a tuple with the InstalledPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPackages

`func (o *SystemProfile) SetInstalledPackages(v []string)`

SetInstalledPackages sets InstalledPackages field to given value.

### HasInstalledPackages

`func (o *SystemProfile) HasInstalledPackages() bool`

HasInstalledPackages returns a boolean if a field has been set.

### GetInstalledProducts

`func (o *SystemProfile) GetInstalledProducts() []SystemProfileInstalledProducts`

GetInstalledProducts returns the InstalledProducts field if non-nil, zero value otherwise.

### GetInstalledProductsOk

`func (o *SystemProfile) GetInstalledProductsOk() (*[]SystemProfileInstalledProducts, bool)`

GetInstalledProductsOk returns a tuple with the InstalledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProducts

`func (o *SystemProfile) SetInstalledProducts(v []SystemProfileInstalledProducts)`

SetInstalledProducts sets InstalledProducts field to given value.

### HasInstalledProducts

`func (o *SystemProfile) HasInstalledProducts() bool`

HasInstalledProducts returns a boolean if a field has been set.

### GetInstalledServices

`func (o *SystemProfile) GetInstalledServices() []string`

GetInstalledServices returns the InstalledServices field if non-nil, zero value otherwise.

### GetInstalledServicesOk

`func (o *SystemProfile) GetInstalledServicesOk() (*[]string, bool)`

GetInstalledServicesOk returns a tuple with the InstalledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledServices

`func (o *SystemProfile) SetInstalledServices(v []string)`

SetInstalledServices sets InstalledServices field to given value.

### HasInstalledServices

`func (o *SystemProfile) HasInstalledServices() bool`

HasInstalledServices returns a boolean if a field has been set.

### GetKatelloAgentRunning

`func (o *SystemProfile) GetKatelloAgentRunning() bool`

GetKatelloAgentRunning returns the KatelloAgentRunning field if non-nil, zero value otherwise.

### GetKatelloAgentRunningOk

`func (o *SystemProfile) GetKatelloAgentRunningOk() (*bool, bool)`

GetKatelloAgentRunningOk returns a tuple with the KatelloAgentRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKatelloAgentRunning

`func (o *SystemProfile) SetKatelloAgentRunning(v bool)`

SetKatelloAgentRunning sets KatelloAgentRunning field to given value.

### HasKatelloAgentRunning

`func (o *SystemProfile) HasKatelloAgentRunning() bool`

HasKatelloAgentRunning returns a boolean if a field has been set.

### GetKernelModules

`func (o *SystemProfile) GetKernelModules() []string`

GetKernelModules returns the KernelModules field if non-nil, zero value otherwise.

### GetKernelModulesOk

`func (o *SystemProfile) GetKernelModulesOk() (*[]string, bool)`

GetKernelModulesOk returns a tuple with the KernelModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelModules

`func (o *SystemProfile) SetKernelModules(v []string)`

SetKernelModules sets KernelModules field to given value.

### HasKernelModules

`func (o *SystemProfile) HasKernelModules() bool`

HasKernelModules returns a boolean if a field has been set.

### GetLastBootTime

`func (o *SystemProfile) GetLastBootTime() string`

GetLastBootTime returns the LastBootTime field if non-nil, zero value otherwise.

### GetLastBootTimeOk

`func (o *SystemProfile) GetLastBootTimeOk() (*string, bool)`

GetLastBootTimeOk returns a tuple with the LastBootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootTime

`func (o *SystemProfile) SetLastBootTime(v string)`

SetLastBootTime sets LastBootTime field to given value.

### HasLastBootTime

`func (o *SystemProfile) HasLastBootTime() bool`

HasLastBootTime returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *SystemProfile) GetNetworkInterfaces() []SystemProfileNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *SystemProfile) GetNetworkInterfacesOk() (*[]SystemProfileNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *SystemProfile) SetNetworkInterfaces(v []SystemProfileNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *SystemProfile) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetNumberOfCpus

`func (o *SystemProfile) GetNumberOfCpus() int32`

GetNumberOfCpus returns the NumberOfCpus field if non-nil, zero value otherwise.

### GetNumberOfCpusOk

`func (o *SystemProfile) GetNumberOfCpusOk() (*int32, bool)`

GetNumberOfCpusOk returns a tuple with the NumberOfCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCpus

`func (o *SystemProfile) SetNumberOfCpus(v int32)`

SetNumberOfCpus sets NumberOfCpus field to given value.

### HasNumberOfCpus

`func (o *SystemProfile) HasNumberOfCpus() bool`

HasNumberOfCpus returns a boolean if a field has been set.

### GetNumberOfSockets

`func (o *SystemProfile) GetNumberOfSockets() int32`

GetNumberOfSockets returns the NumberOfSockets field if non-nil, zero value otherwise.

### GetNumberOfSocketsOk

`func (o *SystemProfile) GetNumberOfSocketsOk() (*int32, bool)`

GetNumberOfSocketsOk returns a tuple with the NumberOfSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSockets

`func (o *SystemProfile) SetNumberOfSockets(v int32)`

SetNumberOfSockets sets NumberOfSockets field to given value.

### HasNumberOfSockets

`func (o *SystemProfile) HasNumberOfSockets() bool`

HasNumberOfSockets returns a boolean if a field has been set.

### GetOsKernelVersion

`func (o *SystemProfile) GetOsKernelVersion() string`

GetOsKernelVersion returns the OsKernelVersion field if non-nil, zero value otherwise.

### GetOsKernelVersionOk

`func (o *SystemProfile) GetOsKernelVersionOk() (*string, bool)`

GetOsKernelVersionOk returns a tuple with the OsKernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsKernelVersion

`func (o *SystemProfile) SetOsKernelVersion(v string)`

SetOsKernelVersion sets OsKernelVersion field to given value.

### HasOsKernelVersion

`func (o *SystemProfile) HasOsKernelVersion() bool`

HasOsKernelVersion returns a boolean if a field has been set.

### GetOsRelease

`func (o *SystemProfile) GetOsRelease() string`

GetOsRelease returns the OsRelease field if non-nil, zero value otherwise.

### GetOsReleaseOk

`func (o *SystemProfile) GetOsReleaseOk() (*string, bool)`

GetOsReleaseOk returns a tuple with the OsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRelease

`func (o *SystemProfile) SetOsRelease(v string)`

SetOsRelease sets OsRelease field to given value.

### HasOsRelease

`func (o *SystemProfile) HasOsRelease() bool`

HasOsRelease returns a boolean if a field has been set.

### GetRunningProcesses

`func (o *SystemProfile) GetRunningProcesses() []string`

GetRunningProcesses returns the RunningProcesses field if non-nil, zero value otherwise.

### GetRunningProcessesOk

`func (o *SystemProfile) GetRunningProcessesOk() (*[]string, bool)`

GetRunningProcessesOk returns a tuple with the RunningProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningProcesses

`func (o *SystemProfile) SetRunningProcesses(v []string)`

SetRunningProcesses sets RunningProcesses field to given value.

### HasRunningProcesses

`func (o *SystemProfile) HasRunningProcesses() bool`

HasRunningProcesses returns a boolean if a field has been set.

### GetSapSids

`func (o *SystemProfile) GetSapSids() []string`

GetSapSids returns the SapSids field if non-nil, zero value otherwise.

### GetSapSidsOk

`func (o *SystemProfile) GetSapSidsOk() (*[]string, bool)`

GetSapSidsOk returns a tuple with the SapSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapSids

`func (o *SystemProfile) SetSapSids(v []string)`

SetSapSids sets SapSids field to given value.

### HasSapSids

`func (o *SystemProfile) HasSapSids() bool`

HasSapSids returns a boolean if a field has been set.

### GetSapSystem

`func (o *SystemProfile) GetSapSystem() bool`

GetSapSystem returns the SapSystem field if non-nil, zero value otherwise.

### GetSapSystemOk

`func (o *SystemProfile) GetSapSystemOk() (*bool, bool)`

GetSapSystemOk returns a tuple with the SapSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapSystem

`func (o *SystemProfile) SetSapSystem(v bool)`

SetSapSystem sets SapSystem field to given value.

### HasSapSystem

`func (o *SystemProfile) HasSapSystem() bool`

HasSapSystem returns a boolean if a field has been set.

### GetSatelliteManaged

`func (o *SystemProfile) GetSatelliteManaged() bool`

GetSatelliteManaged returns the SatelliteManaged field if non-nil, zero value otherwise.

### GetSatelliteManagedOk

`func (o *SystemProfile) GetSatelliteManagedOk() (*bool, bool)`

GetSatelliteManagedOk returns a tuple with the SatelliteManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteManaged

`func (o *SystemProfile) SetSatelliteManaged(v bool)`

SetSatelliteManaged sets SatelliteManaged field to given value.

### HasSatelliteManaged

`func (o *SystemProfile) HasSatelliteManaged() bool`

HasSatelliteManaged returns a boolean if a field has been set.

### GetSubscriptionAutoAttach

`func (o *SystemProfile) GetSubscriptionAutoAttach() string`

GetSubscriptionAutoAttach returns the SubscriptionAutoAttach field if non-nil, zero value otherwise.

### GetSubscriptionAutoAttachOk

`func (o *SystemProfile) GetSubscriptionAutoAttachOk() (*string, bool)`

GetSubscriptionAutoAttachOk returns a tuple with the SubscriptionAutoAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAutoAttach

`func (o *SystemProfile) SetSubscriptionAutoAttach(v string)`

SetSubscriptionAutoAttach sets SubscriptionAutoAttach field to given value.

### HasSubscriptionAutoAttach

`func (o *SystemProfile) HasSubscriptionAutoAttach() bool`

HasSubscriptionAutoAttach returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *SystemProfile) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *SystemProfile) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *SystemProfile) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *SystemProfile) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetSystemMemoryBytes

`func (o *SystemProfile) GetSystemMemoryBytes() int64`

GetSystemMemoryBytes returns the SystemMemoryBytes field if non-nil, zero value otherwise.

### GetSystemMemoryBytesOk

`func (o *SystemProfile) GetSystemMemoryBytesOk() (*int64, bool)`

GetSystemMemoryBytesOk returns a tuple with the SystemMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMemoryBytes

`func (o *SystemProfile) SetSystemMemoryBytes(v int64)`

SetSystemMemoryBytes sets SystemMemoryBytes field to given value.

### HasSystemMemoryBytes

`func (o *SystemProfile) HasSystemMemoryBytes() bool`

HasSystemMemoryBytes returns a boolean if a field has been set.

### GetTunedProfile

`func (o *SystemProfile) GetTunedProfile() string`

GetTunedProfile returns the TunedProfile field if non-nil, zero value otherwise.

### GetTunedProfileOk

`func (o *SystemProfile) GetTunedProfileOk() (*string, bool)`

GetTunedProfileOk returns a tuple with the TunedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunedProfile

`func (o *SystemProfile) SetTunedProfile(v string)`

SetTunedProfile sets TunedProfile field to given value.

### HasTunedProfile

`func (o *SystemProfile) HasTunedProfile() bool`

HasTunedProfile returns a boolean if a field has been set.

### GetYumRepos

`func (o *SystemProfile) GetYumRepos() []SystemProfileYumRepos`

GetYumRepos returns the YumRepos field if non-nil, zero value otherwise.

### GetYumReposOk

`func (o *SystemProfile) GetYumReposOk() (*[]SystemProfileYumRepos, bool)`

GetYumReposOk returns a tuple with the YumRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumRepos

`func (o *SystemProfile) SetYumRepos(v []SystemProfileYumRepos)`

SetYumRepos sets YumRepos field to given value.

### HasYumRepos

`func (o *SystemProfile) HasYumRepos() bool`

HasYumRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


