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

// checks if the ApiSubscriptionsGetMonthlyCostPhpGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubscriptionsGetMonthlyCostPhpGet200Response{}

// ApiSubscriptionsGetMonthlyCostPhpGet200Response struct for ApiSubscriptionsGetMonthlyCostPhpGet200Response
type ApiSubscriptionsGetMonthlyCostPhpGet200Response struct {
	// Whether the request was successful.
	Success *bool `json:"success,omitempty"`
	// A string representing the month and year (e.g., March 2025).
	Title *string `json:"title,omitempty"`
	// The total cost for the specified month.
	MonthlyCost *float32 `json:"monthly_cost,omitempty"`
	// The total cost formatted according to the user's locale and currency.
	LocalizedMonthlyCost *string `json:"localized_monthly_cost,omitempty"`
	// The currency code of the user's main currency.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The currency symbol of the user's main currency.
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Warning messages or additional information.
	Notes []string `json:"notes,omitempty"`
}

// NewApiSubscriptionsGetMonthlyCostPhpGet200Response instantiates a new ApiSubscriptionsGetMonthlyCostPhpGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubscriptionsGetMonthlyCostPhpGet200Response() *ApiSubscriptionsGetMonthlyCostPhpGet200Response {
	this := ApiSubscriptionsGetMonthlyCostPhpGet200Response{}
	return &this
}

// NewApiSubscriptionsGetMonthlyCostPhpGet200ResponseWithDefaults instantiates a new ApiSubscriptionsGetMonthlyCostPhpGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubscriptionsGetMonthlyCostPhpGet200ResponseWithDefaults() *ApiSubscriptionsGetMonthlyCostPhpGet200Response {
	this := ApiSubscriptionsGetMonthlyCostPhpGet200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetMonthlyCost returns the MonthlyCost field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetMonthlyCost() float32 {
	if o == nil || IsNil(o.MonthlyCost) {
		var ret float32
		return ret
	}
	return *o.MonthlyCost
}

// GetMonthlyCostOk returns a tuple with the MonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetMonthlyCostOk() (*float32, bool) {
	if o == nil || IsNil(o.MonthlyCost) {
		return nil, false
	}
	return o.MonthlyCost, true
}

// HasMonthlyCost returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasMonthlyCost() bool {
	if o != nil && !IsNil(o.MonthlyCost) {
		return true
	}

	return false
}

// SetMonthlyCost gets a reference to the given float32 and assigns it to the MonthlyCost field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetMonthlyCost(v float32) {
	o.MonthlyCost = &v
}

// GetLocalizedMonthlyCost returns the LocalizedMonthlyCost field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetLocalizedMonthlyCost() string {
	if o == nil || IsNil(o.LocalizedMonthlyCost) {
		var ret string
		return ret
	}
	return *o.LocalizedMonthlyCost
}

// GetLocalizedMonthlyCostOk returns a tuple with the LocalizedMonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetLocalizedMonthlyCostOk() (*string, bool) {
	if o == nil || IsNil(o.LocalizedMonthlyCost) {
		return nil, false
	}
	return o.LocalizedMonthlyCost, true
}

// HasLocalizedMonthlyCost returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasLocalizedMonthlyCost() bool {
	if o != nil && !IsNil(o.LocalizedMonthlyCost) {
		return true
	}

	return false
}

// SetLocalizedMonthlyCost gets a reference to the given string and assigns it to the LocalizedMonthlyCost field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetLocalizedMonthlyCost(v string) {
	o.LocalizedMonthlyCost = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetNotes(v []string) {
	o.Notes = v
}

func (o ApiSubscriptionsGetMonthlyCostPhpGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubscriptionsGetMonthlyCostPhpGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.MonthlyCost) {
		toSerialize["monthly_cost"] = o.MonthlyCost
	}
	if !IsNil(o.LocalizedMonthlyCost) {
		toSerialize["localized_monthly_cost"] = o.LocalizedMonthlyCost
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencySymbol) {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableApiSubscriptionsGetMonthlyCostPhpGet200Response struct {
	value *ApiSubscriptionsGetMonthlyCostPhpGet200Response
	isSet bool
}

func (v NullableApiSubscriptionsGetMonthlyCostPhpGet200Response) Get() *ApiSubscriptionsGetMonthlyCostPhpGet200Response {
	return v.value
}

func (v *NullableApiSubscriptionsGetMonthlyCostPhpGet200Response) Set(val *ApiSubscriptionsGetMonthlyCostPhpGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubscriptionsGetMonthlyCostPhpGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubscriptionsGetMonthlyCostPhpGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubscriptionsGetMonthlyCostPhpGet200Response(val *ApiSubscriptionsGetMonthlyCostPhpGet200Response) *NullableApiSubscriptionsGetMonthlyCostPhpGet200Response {
	return &NullableApiSubscriptionsGetMonthlyCostPhpGet200Response{value: val, isSet: true}
}

func (v NullableApiSubscriptionsGetMonthlyCostPhpGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubscriptionsGetMonthlyCostPhpGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

