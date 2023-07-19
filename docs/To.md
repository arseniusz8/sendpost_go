# To

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Cc** | Pointer to [**[]CopyTo**](CopyTo.md) |  | [optional] 
**Bcc** | Pointer to [**[]CopyTo**](CopyTo.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTo

`func NewTo() *To`

NewTo instantiates a new To object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToWithDefaults

`func NewToWithDefaults() *To`

NewToWithDefaults instantiates a new To object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *To) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *To) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *To) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *To) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *To) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *To) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *To) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *To) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCc

`func (o *To) GetCc() []CopyTo`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *To) GetCcOk() (*[]CopyTo, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *To) SetCc(v []CopyTo)`

SetCc sets Cc field to given value.

### HasCc

`func (o *To) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *To) GetBcc() []CopyTo`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *To) GetBccOk() (*[]CopyTo, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *To) SetBcc(v []CopyTo)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *To) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCustomFields

`func (o *To) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *To) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *To) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *To) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


