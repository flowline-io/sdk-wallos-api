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

// checks if the ApiNotificationsGetNotificationSettingsPhpGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNotificationsGetNotificationSettingsPhpGet200Response{}

// ApiNotificationsGetNotificationSettingsPhpGet200Response struct for ApiNotificationsGetNotificationSettingsPhpGet200Response
type ApiNotificationsGetNotificationSettingsPhpGet200Response struct {
	Success              *bool                                                                         `json:"success,omitempty"`
	Title                *string                                                                       `json:"title,omitempty"`
	NotificationSettings *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings `json:"notification_settings,omitempty"`
	Notes                []string                                                                      `json:"notes,omitempty"`
}

// NewApiNotificationsGetNotificationSettingsPhpGet200Response instantiates a new ApiNotificationsGetNotificationSettingsPhpGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNotificationsGetNotificationSettingsPhpGet200Response() *ApiNotificationsGetNotificationSettingsPhpGet200Response {
	this := ApiNotificationsGetNotificationSettingsPhpGet200Response{}
	return &this
}

// NewApiNotificationsGetNotificationSettingsPhpGet200ResponseWithDefaults instantiates a new ApiNotificationsGetNotificationSettingsPhpGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNotificationsGetNotificationSettingsPhpGet200ResponseWithDefaults() *ApiNotificationsGetNotificationSettingsPhpGet200Response {
	this := ApiNotificationsGetNotificationSettingsPhpGet200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetNotificationSettings returns the NotificationSettings field value if set, zero value otherwise.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetNotificationSettings() ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings {
	if o == nil || IsNil(o.NotificationSettings) {
		var ret ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings
		return ret
	}
	return *o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetNotificationSettingsOk() (*ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings, bool) {
	if o == nil || IsNil(o.NotificationSettings) {
		return nil, false
	}
	return o.NotificationSettings, true
}

// HasNotificationSettings returns a boolean if a field has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) HasNotificationSettings() bool {
	if o != nil && !IsNil(o.NotificationSettings) {
		return true
	}

	return false
}

// SetNotificationSettings gets a reference to the given ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings and assigns it to the NotificationSettings field.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) SetNotificationSettings(v ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) {
	o.NotificationSettings = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200Response) SetNotes(v []string) {
	o.Notes = v
}

func (o ApiNotificationsGetNotificationSettingsPhpGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNotificationsGetNotificationSettingsPhpGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.NotificationSettings) {
		toSerialize["notification_settings"] = o.NotificationSettings
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableApiNotificationsGetNotificationSettingsPhpGet200Response struct {
	value *ApiNotificationsGetNotificationSettingsPhpGet200Response
	isSet bool
}

func (v NullableApiNotificationsGetNotificationSettingsPhpGet200Response) Get() *ApiNotificationsGetNotificationSettingsPhpGet200Response {
	return v.value
}

func (v *NullableApiNotificationsGetNotificationSettingsPhpGet200Response) Set(val *ApiNotificationsGetNotificationSettingsPhpGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNotificationsGetNotificationSettingsPhpGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNotificationsGetNotificationSettingsPhpGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNotificationsGetNotificationSettingsPhpGet200Response(val *ApiNotificationsGetNotificationSettingsPhpGet200Response) *NullableApiNotificationsGetNotificationSettingsPhpGet200Response {
	return &NullableApiNotificationsGetNotificationSettingsPhpGet200Response{value: val, isSet: true}
}

func (v NullableApiNotificationsGetNotificationSettingsPhpGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNotificationsGetNotificationSettingsPhpGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
