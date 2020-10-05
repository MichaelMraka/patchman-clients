# ErrataResponseErrataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**ReferenceList** | Pointer to **[]string** |  | [optional] 
**Issued** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Synopsis** | Pointer to **string** |  | [optional] 
**CveList** | Pointer to **[]string** |  | [optional] 
**BugzillaList** | Pointer to **[]string** |  | [optional] 
**PackageList** | Pointer to **[]string** |  | [optional] 
**SourcePackageList** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewErrataResponseErrataList

`func NewErrataResponseErrataList() *ErrataResponseErrataList`

NewErrataResponseErrataList instantiates a new ErrataResponseErrataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrataResponseErrataListWithDefaults

`func NewErrataResponseErrataListWithDefaults() *ErrataResponseErrataList`

NewErrataResponseErrataListWithDefaults instantiates a new ErrataResponseErrataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *ErrataResponseErrataList) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ErrataResponseErrataList) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ErrataResponseErrataList) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ErrataResponseErrataList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetSeverity

`func (o *ErrataResponseErrataList) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ErrataResponseErrataList) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ErrataResponseErrataList) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ErrataResponseErrataList) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *ErrataResponseErrataList) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *ErrataResponseErrataList) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetReferenceList

`func (o *ErrataResponseErrataList) GetReferenceList() []string`

GetReferenceList returns the ReferenceList field if non-nil, zero value otherwise.

### GetReferenceListOk

`func (o *ErrataResponseErrataList) GetReferenceListOk() (*[]string, bool)`

GetReferenceListOk returns a tuple with the ReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceList

`func (o *ErrataResponseErrataList) SetReferenceList(v []string)`

SetReferenceList sets ReferenceList field to given value.

### HasReferenceList

`func (o *ErrataResponseErrataList) HasReferenceList() bool`

HasReferenceList returns a boolean if a field has been set.

### GetIssued

`func (o *ErrataResponseErrataList) GetIssued() string`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *ErrataResponseErrataList) GetIssuedOk() (*string, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *ErrataResponseErrataList) SetIssued(v string)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *ErrataResponseErrataList) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetDescription

`func (o *ErrataResponseErrataList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrataResponseErrataList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrataResponseErrataList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrataResponseErrataList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSolution

`func (o *ErrataResponseErrataList) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *ErrataResponseErrataList) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *ErrataResponseErrataList) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *ErrataResponseErrataList) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetSummary

`func (o *ErrataResponseErrataList) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ErrataResponseErrataList) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ErrataResponseErrataList) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ErrataResponseErrataList) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUrl

`func (o *ErrataResponseErrataList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ErrataResponseErrataList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ErrataResponseErrataList) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ErrataResponseErrataList) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSynopsis

`func (o *ErrataResponseErrataList) GetSynopsis() string`

GetSynopsis returns the Synopsis field if non-nil, zero value otherwise.

### GetSynopsisOk

`func (o *ErrataResponseErrataList) GetSynopsisOk() (*string, bool)`

GetSynopsisOk returns a tuple with the Synopsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynopsis

`func (o *ErrataResponseErrataList) SetSynopsis(v string)`

SetSynopsis sets Synopsis field to given value.

### HasSynopsis

`func (o *ErrataResponseErrataList) HasSynopsis() bool`

HasSynopsis returns a boolean if a field has been set.

### GetCveList

`func (o *ErrataResponseErrataList) GetCveList() []string`

GetCveList returns the CveList field if non-nil, zero value otherwise.

### GetCveListOk

`func (o *ErrataResponseErrataList) GetCveListOk() (*[]string, bool)`

GetCveListOk returns a tuple with the CveList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveList

`func (o *ErrataResponseErrataList) SetCveList(v []string)`

SetCveList sets CveList field to given value.

### HasCveList

`func (o *ErrataResponseErrataList) HasCveList() bool`

HasCveList returns a boolean if a field has been set.

### GetBugzillaList

`func (o *ErrataResponseErrataList) GetBugzillaList() []string`

GetBugzillaList returns the BugzillaList field if non-nil, zero value otherwise.

### GetBugzillaListOk

`func (o *ErrataResponseErrataList) GetBugzillaListOk() (*[]string, bool)`

GetBugzillaListOk returns a tuple with the BugzillaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugzillaList

`func (o *ErrataResponseErrataList) SetBugzillaList(v []string)`

SetBugzillaList sets BugzillaList field to given value.

### HasBugzillaList

`func (o *ErrataResponseErrataList) HasBugzillaList() bool`

HasBugzillaList returns a boolean if a field has been set.

### GetPackageList

`func (o *ErrataResponseErrataList) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *ErrataResponseErrataList) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *ErrataResponseErrataList) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *ErrataResponseErrataList) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.

### GetSourcePackageList

`func (o *ErrataResponseErrataList) GetSourcePackageList() []string`

GetSourcePackageList returns the SourcePackageList field if non-nil, zero value otherwise.

### GetSourcePackageListOk

`func (o *ErrataResponseErrataList) GetSourcePackageListOk() (*[]string, bool)`

GetSourcePackageListOk returns a tuple with the SourcePackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePackageList

`func (o *ErrataResponseErrataList) SetSourcePackageList(v []string)`

SetSourcePackageList sets SourcePackageList field to given value.

### HasSourcePackageList

`func (o *ErrataResponseErrataList) HasSourcePackageList() bool`

HasSourcePackageList returns a boolean if a field has been set.

### GetType

`func (o *ErrataResponseErrataList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrataResponseErrataList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrataResponseErrataList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrataResponseErrataList) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


