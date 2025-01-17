/*
Wallos API

API documentation for Wallos

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss{}

// ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss struct for ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss
type ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss struct {
	Css *string `json:"css,omitempty"`
}

// NewApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss instantiates a new ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss() *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss {
	this := ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss{}
	return &this
}

// NewApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCssWithDefaults instantiates a new ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCssWithDefaults() *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss {
	this := ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss{}
	return &this
}

// GetCss returns the Css field value if set, zero value otherwise.
func (o *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) GetCss() string {
	if o == nil || IsNil(o.Css) {
		var ret string
		return ret
	}
	return *o.Css
}

// GetCssOk returns a tuple with the Css field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) GetCssOk() (*string, bool) {
	if o == nil || IsNil(o.Css) {
		return nil, false
	}
	return o.Css, true
}

// HasCss returns a boolean if a field has been set.
func (o *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) HasCss() bool {
	if o != nil && !IsNil(o.Css) {
		return true
	}

	return false
}

// SetCss gets a reference to the given string and assigns it to the Css field.
func (o *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) SetCss(v string) {
	o.Css = &v
}

func (o ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Css) {
		toSerialize["css"] = o.Css
	}
	return toSerialize, nil
}

type NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss struct {
	value *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss
	isSet bool
}

func (v NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) Get() *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss {
	return v.value
}

func (v *NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) Set(val *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss(val *ApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) *NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss {
	return &NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss{value: val, isSet: true}
}

func (v NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSettingsGetSettingsPhpGet200ResponseSettingsCustomCss) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

