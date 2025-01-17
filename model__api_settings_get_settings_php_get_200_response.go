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

// checks if the ApiSettingsGetSettingsPhpGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSettingsGetSettingsPhpGet200Response{}

// ApiSettingsGetSettingsPhpGet200Response struct for ApiSettingsGetSettingsPhpGet200Response
type ApiSettingsGetSettingsPhpGet200Response struct {
	Success *bool `json:"success,omitempty"`
	Title *string `json:"title,omitempty"`
	Settings *ApiSettingsGetSettingsPhpGet200ResponseSettings `json:"settings,omitempty"`
	Notes []string `json:"notes,omitempty"`
}

// NewApiSettingsGetSettingsPhpGet200Response instantiates a new ApiSettingsGetSettingsPhpGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSettingsGetSettingsPhpGet200Response() *ApiSettingsGetSettingsPhpGet200Response {
	this := ApiSettingsGetSettingsPhpGet200Response{}
	return &this
}

// NewApiSettingsGetSettingsPhpGet200ResponseWithDefaults instantiates a new ApiSettingsGetSettingsPhpGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSettingsGetSettingsPhpGet200ResponseWithDefaults() *ApiSettingsGetSettingsPhpGet200Response {
	this := ApiSettingsGetSettingsPhpGet200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiSettingsGetSettingsPhpGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiSettingsGetSettingsPhpGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetSettings() ApiSettingsGetSettingsPhpGet200ResponseSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ApiSettingsGetSettingsPhpGet200ResponseSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetSettingsOk() (*ApiSettingsGetSettingsPhpGet200ResponseSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ApiSettingsGetSettingsPhpGet200ResponseSettings and assigns it to the Settings field.
func (o *ApiSettingsGetSettingsPhpGet200Response) SetSettings(v ApiSettingsGetSettingsPhpGet200ResponseSettings) {
	o.Settings = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiSettingsGetSettingsPhpGet200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ApiSettingsGetSettingsPhpGet200Response) SetNotes(v []string) {
	o.Notes = v
}

func (o ApiSettingsGetSettingsPhpGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSettingsGetSettingsPhpGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableApiSettingsGetSettingsPhpGet200Response struct {
	value *ApiSettingsGetSettingsPhpGet200Response
	isSet bool
}

func (v NullableApiSettingsGetSettingsPhpGet200Response) Get() *ApiSettingsGetSettingsPhpGet200Response {
	return v.value
}

func (v *NullableApiSettingsGetSettingsPhpGet200Response) Set(val *ApiSettingsGetSettingsPhpGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSettingsGetSettingsPhpGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSettingsGetSettingsPhpGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSettingsGetSettingsPhpGet200Response(val *ApiSettingsGetSettingsPhpGet200Response) *NullableApiSettingsGetSettingsPhpGet200Response {
	return &NullableApiSettingsGetSettingsPhpGet200Response{value: val, isSet: true}
}

func (v NullableApiSettingsGetSettingsPhpGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSettingsGetSettingsPhpGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


