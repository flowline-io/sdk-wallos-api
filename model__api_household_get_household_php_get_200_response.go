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

// checks if the ApiHouseholdGetHouseholdPhpGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHouseholdGetHouseholdPhpGet200Response{}

// ApiHouseholdGetHouseholdPhpGet200Response struct for ApiHouseholdGetHouseholdPhpGet200Response
type ApiHouseholdGetHouseholdPhpGet200Response struct {
	Success   *bool                                                     `json:"success,omitempty"`
	Title     *string                                                   `json:"title,omitempty"`
	Household []ApiHouseholdGetHouseholdPhpGet200ResponseHouseholdInner `json:"household,omitempty"`
	Notes     []string                                                  `json:"notes,omitempty"`
}

// NewApiHouseholdGetHouseholdPhpGet200Response instantiates a new ApiHouseholdGetHouseholdPhpGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHouseholdGetHouseholdPhpGet200Response() *ApiHouseholdGetHouseholdPhpGet200Response {
	this := ApiHouseholdGetHouseholdPhpGet200Response{}
	return &this
}

// NewApiHouseholdGetHouseholdPhpGet200ResponseWithDefaults instantiates a new ApiHouseholdGetHouseholdPhpGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHouseholdGetHouseholdPhpGet200ResponseWithDefaults() *ApiHouseholdGetHouseholdPhpGet200Response {
	this := ApiHouseholdGetHouseholdPhpGet200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetHousehold returns the Household field value if set, zero value otherwise.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetHousehold() []ApiHouseholdGetHouseholdPhpGet200ResponseHouseholdInner {
	if o == nil || IsNil(o.Household) {
		var ret []ApiHouseholdGetHouseholdPhpGet200ResponseHouseholdInner
		return ret
	}
	return o.Household
}

// GetHouseholdOk returns a tuple with the Household field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetHouseholdOk() ([]ApiHouseholdGetHouseholdPhpGet200ResponseHouseholdInner, bool) {
	if o == nil || IsNil(o.Household) {
		return nil, false
	}
	return o.Household, true
}

// HasHousehold returns a boolean if a field has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) HasHousehold() bool {
	if o != nil && !IsNil(o.Household) {
		return true
	}

	return false
}

// SetHousehold gets a reference to the given []ApiHouseholdGetHouseholdPhpGet200ResponseHouseholdInner and assigns it to the Household field.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) SetHousehold(v []ApiHouseholdGetHouseholdPhpGet200ResponseHouseholdInner) {
	o.Household = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ApiHouseholdGetHouseholdPhpGet200Response) SetNotes(v []string) {
	o.Notes = v
}

func (o ApiHouseholdGetHouseholdPhpGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHouseholdGetHouseholdPhpGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Household) {
		toSerialize["household"] = o.Household
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableApiHouseholdGetHouseholdPhpGet200Response struct {
	value *ApiHouseholdGetHouseholdPhpGet200Response
	isSet bool
}

func (v NullableApiHouseholdGetHouseholdPhpGet200Response) Get() *ApiHouseholdGetHouseholdPhpGet200Response {
	return v.value
}

func (v *NullableApiHouseholdGetHouseholdPhpGet200Response) Set(val *ApiHouseholdGetHouseholdPhpGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHouseholdGetHouseholdPhpGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHouseholdGetHouseholdPhpGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHouseholdGetHouseholdPhpGet200Response(val *ApiHouseholdGetHouseholdPhpGet200Response) *NullableApiHouseholdGetHouseholdPhpGet200Response {
	return &NullableApiHouseholdGetHouseholdPhpGet200Response{value: val, isSet: true}
}

func (v NullableApiHouseholdGetHouseholdPhpGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHouseholdGetHouseholdPhpGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
