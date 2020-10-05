# CvesResponseCveList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impact** | Pointer to **string** |  | [optional] 
**PublicDate** | Pointer to **string** |  | [optional] 
**Synopsis** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ModifiedDate** | Pointer to **string** |  | [optional] 
**RedhatUrl** | Pointer to **string** |  | [optional] 
**SecondaryUrl** | Pointer to **string** |  | [optional] 
**Cvss2Score** | Pointer to **string** |  | [optional] 
**Cvss2Metrics** | Pointer to **string** |  | [optional] 
**Cvss3Score** | Pointer to **string** |  | [optional] 
**Cvss3Metrics** | Pointer to **string** |  | [optional] 
**CweList** | Pointer to **[]string** |  | [optional] 
**ErrataList** | Pointer to **[]string** |  | [optional] 
**PackageList** | Pointer to **[]string** |  | [optional] 
**SourcePackageList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCvesResponseCveList

`func NewCvesResponseCveList() *CvesResponseCveList`

NewCvesResponseCveList instantiates a new CvesResponseCveList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvesResponseCveListWithDefaults

`func NewCvesResponseCveListWithDefaults() *CvesResponseCveList`

NewCvesResponseCveListWithDefaults instantiates a new CvesResponseCveList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpact

`func (o *CvesResponseCveList) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *CvesResponseCveList) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *CvesResponseCveList) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *CvesResponseCveList) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetPublicDate

`func (o *CvesResponseCveList) GetPublicDate() string`

GetPublicDate returns the PublicDate field if non-nil, zero value otherwise.

### GetPublicDateOk

`func (o *CvesResponseCveList) GetPublicDateOk() (*string, bool)`

GetPublicDateOk returns a tuple with the PublicDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDate

`func (o *CvesResponseCveList) SetPublicDate(v string)`

SetPublicDate sets PublicDate field to given value.

### HasPublicDate

`func (o *CvesResponseCveList) HasPublicDate() bool`

HasPublicDate returns a boolean if a field has been set.

### GetSynopsis

`func (o *CvesResponseCveList) GetSynopsis() string`

GetSynopsis returns the Synopsis field if non-nil, zero value otherwise.

### GetSynopsisOk

`func (o *CvesResponseCveList) GetSynopsisOk() (*string, bool)`

GetSynopsisOk returns a tuple with the Synopsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynopsis

`func (o *CvesResponseCveList) SetSynopsis(v string)`

SetSynopsis sets Synopsis field to given value.

### HasSynopsis

`func (o *CvesResponseCveList) HasSynopsis() bool`

HasSynopsis returns a boolean if a field has been set.

### GetDescription

`func (o *CvesResponseCveList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CvesResponseCveList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CvesResponseCveList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CvesResponseCveList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModifiedDate

`func (o *CvesResponseCveList) GetModifiedDate() string`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CvesResponseCveList) GetModifiedDateOk() (*string, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CvesResponseCveList) SetModifiedDate(v string)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *CvesResponseCveList) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetRedhatUrl

`func (o *CvesResponseCveList) GetRedhatUrl() string`

GetRedhatUrl returns the RedhatUrl field if non-nil, zero value otherwise.

### GetRedhatUrlOk

`func (o *CvesResponseCveList) GetRedhatUrlOk() (*string, bool)`

GetRedhatUrlOk returns a tuple with the RedhatUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedhatUrl

`func (o *CvesResponseCveList) SetRedhatUrl(v string)`

SetRedhatUrl sets RedhatUrl field to given value.

### HasRedhatUrl

`func (o *CvesResponseCveList) HasRedhatUrl() bool`

HasRedhatUrl returns a boolean if a field has been set.

### GetSecondaryUrl

`func (o *CvesResponseCveList) GetSecondaryUrl() string`

GetSecondaryUrl returns the SecondaryUrl field if non-nil, zero value otherwise.

### GetSecondaryUrlOk

`func (o *CvesResponseCveList) GetSecondaryUrlOk() (*string, bool)`

GetSecondaryUrlOk returns a tuple with the SecondaryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryUrl

`func (o *CvesResponseCveList) SetSecondaryUrl(v string)`

SetSecondaryUrl sets SecondaryUrl field to given value.

### HasSecondaryUrl

`func (o *CvesResponseCveList) HasSecondaryUrl() bool`

HasSecondaryUrl returns a boolean if a field has been set.

### GetCvss2Score

`func (o *CvesResponseCveList) GetCvss2Score() string`

GetCvss2Score returns the Cvss2Score field if non-nil, zero value otherwise.

### GetCvss2ScoreOk

`func (o *CvesResponseCveList) GetCvss2ScoreOk() (*string, bool)`

GetCvss2ScoreOk returns a tuple with the Cvss2Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss2Score

`func (o *CvesResponseCveList) SetCvss2Score(v string)`

SetCvss2Score sets Cvss2Score field to given value.

### HasCvss2Score

`func (o *CvesResponseCveList) HasCvss2Score() bool`

HasCvss2Score returns a boolean if a field has been set.

### GetCvss2Metrics

`func (o *CvesResponseCveList) GetCvss2Metrics() string`

GetCvss2Metrics returns the Cvss2Metrics field if non-nil, zero value otherwise.

### GetCvss2MetricsOk

`func (o *CvesResponseCveList) GetCvss2MetricsOk() (*string, bool)`

GetCvss2MetricsOk returns a tuple with the Cvss2Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss2Metrics

`func (o *CvesResponseCveList) SetCvss2Metrics(v string)`

SetCvss2Metrics sets Cvss2Metrics field to given value.

### HasCvss2Metrics

`func (o *CvesResponseCveList) HasCvss2Metrics() bool`

HasCvss2Metrics returns a boolean if a field has been set.

### GetCvss3Score

`func (o *CvesResponseCveList) GetCvss3Score() string`

GetCvss3Score returns the Cvss3Score field if non-nil, zero value otherwise.

### GetCvss3ScoreOk

`func (o *CvesResponseCveList) GetCvss3ScoreOk() (*string, bool)`

GetCvss3ScoreOk returns a tuple with the Cvss3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Score

`func (o *CvesResponseCveList) SetCvss3Score(v string)`

SetCvss3Score sets Cvss3Score field to given value.

### HasCvss3Score

`func (o *CvesResponseCveList) HasCvss3Score() bool`

HasCvss3Score returns a boolean if a field has been set.

### GetCvss3Metrics

`func (o *CvesResponseCveList) GetCvss3Metrics() string`

GetCvss3Metrics returns the Cvss3Metrics field if non-nil, zero value otherwise.

### GetCvss3MetricsOk

`func (o *CvesResponseCveList) GetCvss3MetricsOk() (*string, bool)`

GetCvss3MetricsOk returns a tuple with the Cvss3Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Metrics

`func (o *CvesResponseCveList) SetCvss3Metrics(v string)`

SetCvss3Metrics sets Cvss3Metrics field to given value.

### HasCvss3Metrics

`func (o *CvesResponseCveList) HasCvss3Metrics() bool`

HasCvss3Metrics returns a boolean if a field has been set.

### GetCweList

`func (o *CvesResponseCveList) GetCweList() []string`

GetCweList returns the CweList field if non-nil, zero value otherwise.

### GetCweListOk

`func (o *CvesResponseCveList) GetCweListOk() (*[]string, bool)`

GetCweListOk returns a tuple with the CweList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCweList

`func (o *CvesResponseCveList) SetCweList(v []string)`

SetCweList sets CweList field to given value.

### HasCweList

`func (o *CvesResponseCveList) HasCweList() bool`

HasCweList returns a boolean if a field has been set.

### GetErrataList

`func (o *CvesResponseCveList) GetErrataList() []string`

GetErrataList returns the ErrataList field if non-nil, zero value otherwise.

### GetErrataListOk

`func (o *CvesResponseCveList) GetErrataListOk() (*[]string, bool)`

GetErrataListOk returns a tuple with the ErrataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrataList

`func (o *CvesResponseCveList) SetErrataList(v []string)`

SetErrataList sets ErrataList field to given value.

### HasErrataList

`func (o *CvesResponseCveList) HasErrataList() bool`

HasErrataList returns a boolean if a field has been set.

### GetPackageList

`func (o *CvesResponseCveList) GetPackageList() []string`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *CvesResponseCveList) GetPackageListOk() (*[]string, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *CvesResponseCveList) SetPackageList(v []string)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *CvesResponseCveList) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.

### GetSourcePackageList

`func (o *CvesResponseCveList) GetSourcePackageList() []string`

GetSourcePackageList returns the SourcePackageList field if non-nil, zero value otherwise.

### GetSourcePackageListOk

`func (o *CvesResponseCveList) GetSourcePackageListOk() (*[]string, bool)`

GetSourcePackageListOk returns a tuple with the SourcePackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePackageList

`func (o *CvesResponseCveList) SetSourcePackageList(v []string)`

SetSourcePackageList sets SourcePackageList field to given value.

### HasSourcePackageList

`func (o *CvesResponseCveList) HasSourcePackageList() bool`

HasSourcePackageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


