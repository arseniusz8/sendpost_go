# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** |  | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *Device) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Device) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Device) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Device) HasFamily() bool`

HasFamily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


