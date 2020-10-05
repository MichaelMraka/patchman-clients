# FactSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facts** | **map[string]interface{}** | The facts themselves. | 
**Namespace** | **string** | A namespace the facts belong to. | 

## Methods

### NewFactSet

`func NewFactSet(facts map[string]interface{}, namespace string, ) *FactSet`

NewFactSet instantiates a new FactSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactSetWithDefaults

`func NewFactSetWithDefaults() *FactSet`

NewFactSetWithDefaults instantiates a new FactSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacts

`func (o *FactSet) GetFacts() map[string]interface{}`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *FactSet) GetFactsOk() (*map[string]interface{}, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *FactSet) SetFacts(v map[string]interface{})`

SetFacts sets Facts field to given value.


### GetNamespace

`func (o *FactSet) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FactSet) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FactSet) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


