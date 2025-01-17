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

// checks if the ApiAdminGetAdminSettingsPhpGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAdminGetAdminSettingsPhpGet200Response{}

// ApiAdminGetAdminSettingsPhpGet200Response struct for ApiAdminGetAdminSettingsPhpGet200Response
type ApiAdminGetAdminSettingsPhpGet200Response struct {
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty"`
	// The title of the response.
	Title *string `json:"title,omitempty"`
	AdminSettings *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings `json:"admin_settings,omitempty"`
	// Warning messages or additional information.
	Notes []string `json:"notes,omitempty"`
}

// NewApiAdminGetAdminSettingsPhpGet200Response instantiates a new ApiAdminGetAdminSettingsPhpGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAdminGetAdminSettingsPhpGet200Response() *ApiAdminGetAdminSettingsPhpGet200Response {
	this := ApiAdminGetAdminSettingsPhpGet200Response{}
	return &this
}

// NewApiAdminGetAdminSettingsPhpGet200ResponseWithDefaults instantiates a new ApiAdminGetAdminSettingsPhpGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAdminGetAdminSettingsPhpGet200ResponseWithDefaults() *ApiAdminGetAdminSettingsPhpGet200Response {
	this := ApiAdminGetAdminSettingsPhpGet200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetAdminSettings returns the AdminSettings field value if set, zero value otherwise.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetAdminSettings() ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings {
	if o == nil || IsNil(o.AdminSettings) {
		var ret ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings
		return ret
	}
	return *o.AdminSettings
}

// GetAdminSettingsOk returns a tuple with the AdminSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetAdminSettingsOk() (*ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings, bool) {
	if o == nil || IsNil(o.AdminSettings) {
		return nil, false
	}
	return o.AdminSettings, true
}

// HasAdminSettings returns a boolean if a field has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasAdminSettings() bool {
	if o != nil && !IsNil(o.AdminSettings) {
		return true
	}

	return false
}

// SetAdminSettings gets a reference to the given ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings and assigns it to the AdminSettings field.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetAdminSettings(v ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) {
	o.AdminSettings = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetNotes(v []string) {
	o.Notes = v
}

func (o ApiAdminGetAdminSettingsPhpGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAdminGetAdminSettingsPhpGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.AdminSettings) {
		toSerialize["admin_settings"] = o.AdminSettings
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableApiAdminGetAdminSettingsPhpGet200Response struct {
	value *ApiAdminGetAdminSettingsPhpGet200Response
	isSet bool
}

func (v NullableApiAdminGetAdminSettingsPhpGet200Response) Get() *ApiAdminGetAdminSettingsPhpGet200Response {
	return v.value
}

func (v *NullableApiAdminGetAdminSettingsPhpGet200Response) Set(val *ApiAdminGetAdminSettingsPhpGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAdminGetAdminSettingsPhpGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAdminGetAdminSettingsPhpGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAdminGetAdminSettingsPhpGet200Response(val *ApiAdminGetAdminSettingsPhpGet200Response) *NullableApiAdminGetAdminSettingsPhpGet200Response {
	return &NullableApiAdminGetAdminSettingsPhpGet200Response{value: val, isSet: true}
}

func (v NullableApiAdminGetAdminSettingsPhpGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAdminGetAdminSettingsPhpGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


