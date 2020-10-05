# DBChangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dbchange** | Pointer to [**DBChangeResponseDbchange**](DBChangeResponse_dbchange.md) |  | [optional] 

## Methods

### NewDBChangeResponse

`func NewDBChangeResponse() *DBChangeResponse`

NewDBChangeResponse instantiates a new DBChangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBChangeResponseWithDefaults

`func NewDBChangeResponseWithDefaults() *DBChangeResponse`

NewDBChangeResponseWithDefaults instantiates a new DBChangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbchange

`func (o *DBChangeResponse) GetDbchange() DBChangeResponseDbchange`

GetDbchange returns the Dbchange field if non-nil, zero value otherwise.

### GetDbchangeOk

`func (o *DBChangeResponse) GetDbchangeOk() (*DBChangeResponseDbchange, bool)`

GetDbchangeOk returns a tuple with the Dbchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbchange

`func (o *DBChangeResponse) SetDbchange(v DBChangeResponseDbchange)`

SetDbchange sets Dbchange field to given value.

### HasDbchange

`func (o *DBChangeResponse) HasDbchange() bool`

HasDbchange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


