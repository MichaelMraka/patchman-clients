# ActiveTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **NullableInt32** | The number of hosts with the given tag. If the value is null this indicates that the count is unknown. | 
**Tag** | [**StructuredTag**](StructuredTag.md) |  | 

## Methods

### NewActiveTag

`func NewActiveTag(count NullableInt32, tag StructuredTag, ) *ActiveTag`

NewActiveTag instantiates a new ActiveTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveTagWithDefaults

`func NewActiveTagWithDefaults() *ActiveTag`

NewActiveTagWithDefaults instantiates a new ActiveTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ActiveTag) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ActiveTag) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ActiveTag) SetCount(v int32)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *ActiveTag) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ActiveTag) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetTag

`func (o *ActiveTag) GetTag() StructuredTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ActiveTag) GetTagOk() (*StructuredTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ActiveTag) SetTag(v StructuredTag)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


