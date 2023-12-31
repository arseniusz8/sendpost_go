/*
SendPost API

Email API and SMTP relay to not just send and measure email sending, but also alert and optimise. We provide you with tools, expertise and support needed to reliably deliver emails to your customers inboxes on time, every time.

API version: 1.0.0
Contact: hello@sendpost.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sendpost

import (
	"encoding/json"
)

// checks if the City type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &City{}

// City struct for City
type City struct {
	CityID *int32 `json:"cityID,omitempty"`
	ContinentCode *string `json:"continentCode,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewCity instantiates a new City object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCity() *City {
	this := City{}
	return &this
}

// NewCityWithDefaults instantiates a new City object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityWithDefaults() *City {
	this := City{}
	return &this
}

// GetCityID returns the CityID field value if set, zero value otherwise.
func (o *City) GetCityID() int32 {
	if o == nil || IsNil(o.CityID) {
		var ret int32
		return ret
	}
	return *o.CityID
}

// GetCityIDOk returns a tuple with the CityID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetCityIDOk() (*int32, bool) {
	if o == nil || IsNil(o.CityID) {
		return nil, false
	}
	return o.CityID, true
}

// HasCityID returns a boolean if a field has been set.
func (o *City) HasCityID() bool {
	if o != nil && !IsNil(o.CityID) {
		return true
	}

	return false
}

// SetCityID gets a reference to the given int32 and assigns it to the CityID field.
func (o *City) SetCityID(v int32) {
	o.CityID = &v
}

// GetContinentCode returns the ContinentCode field value if set, zero value otherwise.
func (o *City) GetContinentCode() string {
	if o == nil || IsNil(o.ContinentCode) {
		var ret string
		return ret
	}
	return *o.ContinentCode
}

// GetContinentCodeOk returns a tuple with the ContinentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetContinentCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ContinentCode) {
		return nil, false
	}
	return o.ContinentCode, true
}

// HasContinentCode returns a boolean if a field has been set.
func (o *City) HasContinentCode() bool {
	if o != nil && !IsNil(o.ContinentCode) {
		return true
	}

	return false
}

// SetContinentCode gets a reference to the given string and assigns it to the ContinentCode field.
func (o *City) SetContinentCode(v string) {
	o.ContinentCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *City) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *City) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *City) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *City) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *City) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *City) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *City) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *City) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *City) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o City) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o City) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CityID) {
		toSerialize["cityID"] = o.CityID
	}
	if !IsNil(o.ContinentCode) {
		toSerialize["continentCode"] = o.ContinentCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableCity struct {
	value *City
	isSet bool
}

func (v NullableCity) Get() *City {
	return v.value
}

func (v *NullableCity) Set(val *City) {
	v.value = val
	v.isSet = true
}

func (v NullableCity) IsSet() bool {
	return v.isSet
}

func (v *NullableCity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCity(val *City) *NullableCity {
	return &NullableCity{value: val, isSet: true}
}

func (v NullableCity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


