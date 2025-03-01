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

// checks if the ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings{}

// ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings struct for ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings
type ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings struct {
	EmailNotifications *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsEmailNotifications `json:"email_notifications,omitempty"`
	NtfyNotifications  *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsNtfyNotifications  `json:"ntfy_notifications,omitempty"`
}

// NewApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings instantiates a new ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings() *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings {
	this := ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings{}
	return &this
}

// NewApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsWithDefaults instantiates a new ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsWithDefaults() *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings {
	this := ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings{}
	return &this
}

// GetEmailNotifications returns the EmailNotifications field value if set, zero value otherwise.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) GetEmailNotifications() ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsEmailNotifications {
	if o == nil || IsNil(o.EmailNotifications) {
		var ret ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsEmailNotifications
		return ret
	}
	return *o.EmailNotifications
}

// GetEmailNotificationsOk returns a tuple with the EmailNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) GetEmailNotificationsOk() (*ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsEmailNotifications, bool) {
	if o == nil || IsNil(o.EmailNotifications) {
		return nil, false
	}
	return o.EmailNotifications, true
}

// HasEmailNotifications returns a boolean if a field has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) HasEmailNotifications() bool {
	if o != nil && !IsNil(o.EmailNotifications) {
		return true
	}

	return false
}

// SetEmailNotifications gets a reference to the given ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsEmailNotifications and assigns it to the EmailNotifications field.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) SetEmailNotifications(v ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsEmailNotifications) {
	o.EmailNotifications = &v
}

// GetNtfyNotifications returns the NtfyNotifications field value if set, zero value otherwise.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) GetNtfyNotifications() ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsNtfyNotifications {
	if o == nil || IsNil(o.NtfyNotifications) {
		var ret ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsNtfyNotifications
		return ret
	}
	return *o.NtfyNotifications
}

// GetNtfyNotificationsOk returns a tuple with the NtfyNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) GetNtfyNotificationsOk() (*ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsNtfyNotifications, bool) {
	if o == nil || IsNil(o.NtfyNotifications) {
		return nil, false
	}
	return o.NtfyNotifications, true
}

// HasNtfyNotifications returns a boolean if a field has been set.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) HasNtfyNotifications() bool {
	if o != nil && !IsNil(o.NtfyNotifications) {
		return true
	}

	return false
}

// SetNtfyNotifications gets a reference to the given ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsNtfyNotifications and assigns it to the NtfyNotifications field.
func (o *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) SetNtfyNotifications(v ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettingsNtfyNotifications) {
	o.NtfyNotifications = &v
}

func (o ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailNotifications) {
		toSerialize["email_notifications"] = o.EmailNotifications
	}
	if !IsNil(o.NtfyNotifications) {
		toSerialize["ntfy_notifications"] = o.NtfyNotifications
	}
	return toSerialize, nil
}

type NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings struct {
	value *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings
	isSet bool
}

func (v NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) Get() *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings {
	return v.value
}

func (v *NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) Set(val *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings(val *ApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) *NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings {
	return &NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings{value: val, isSet: true}
}

func (v NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNotificationsGetNotificationSettingsPhpGet200ResponseNotificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
