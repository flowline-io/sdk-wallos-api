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

// checks if the ApiCurrenciesGetCurrenciesPhpGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCurrenciesGetCurrenciesPhpGet200Response{}

// ApiCurrenciesGetCurrenciesPhpGet200Response struct for ApiCurrenciesGetCurrenciesPhpGet200Response
type ApiCurrenciesGetCurrenciesPhpGet200Response struct {
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty"`
	// The title of the response.
	Title *string `json:"title,omitempty"`
	// The ID of the main currency of the user.
	MainCurrency *int32                                                       `json:"main_currency,omitempty"`
	Currencies   []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner `json:"currencies,omitempty"`
	// Warning messages or additional information.
	Notes []string `json:"notes,omitempty"`
}

// NewApiCurrenciesGetCurrenciesPhpGet200Response instantiates a new ApiCurrenciesGetCurrenciesPhpGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCurrenciesGetCurrenciesPhpGet200Response() *ApiCurrenciesGetCurrenciesPhpGet200Response {
	this := ApiCurrenciesGetCurrenciesPhpGet200Response{}
	return &this
}

// NewApiCurrenciesGetCurrenciesPhpGet200ResponseWithDefaults instantiates a new ApiCurrenciesGetCurrenciesPhpGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCurrenciesGetCurrenciesPhpGet200ResponseWithDefaults() *ApiCurrenciesGetCurrenciesPhpGet200Response {
	this := ApiCurrenciesGetCurrenciesPhpGet200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetMainCurrency returns the MainCurrency field value if set, zero value otherwise.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetMainCurrency() int32 {
	if o == nil || IsNil(o.MainCurrency) {
		var ret int32
		return ret
	}
	return *o.MainCurrency
}

// GetMainCurrencyOk returns a tuple with the MainCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetMainCurrencyOk() (*int32, bool) {
	if o == nil || IsNil(o.MainCurrency) {
		return nil, false
	}
	return o.MainCurrency, true
}

// HasMainCurrency returns a boolean if a field has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasMainCurrency() bool {
	if o != nil && !IsNil(o.MainCurrency) {
		return true
	}

	return false
}

// SetMainCurrency gets a reference to the given int32 and assigns it to the MainCurrency field.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetMainCurrency(v int32) {
	o.MainCurrency = &v
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetCurrencies() []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner {
	if o == nil || IsNil(o.Currencies) {
		var ret []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner
		return ret
	}
	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetCurrenciesOk() ([]ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner, bool) {
	if o == nil || IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasCurrencies() bool {
	if o != nil && !IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner and assigns it to the Currencies field.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetCurrencies(v []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner) {
	o.Currencies = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetNotes(v []string) {
	o.Notes = v
}

func (o ApiCurrenciesGetCurrenciesPhpGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCurrenciesGetCurrenciesPhpGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.MainCurrency) {
		toSerialize["main_currency"] = o.MainCurrency
	}
	if !IsNil(o.Currencies) {
		toSerialize["currencies"] = o.Currencies
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableApiCurrenciesGetCurrenciesPhpGet200Response struct {
	value *ApiCurrenciesGetCurrenciesPhpGet200Response
	isSet bool
}

func (v NullableApiCurrenciesGetCurrenciesPhpGet200Response) Get() *ApiCurrenciesGetCurrenciesPhpGet200Response {
	return v.value
}

func (v *NullableApiCurrenciesGetCurrenciesPhpGet200Response) Set(val *ApiCurrenciesGetCurrenciesPhpGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCurrenciesGetCurrenciesPhpGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCurrenciesGetCurrenciesPhpGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCurrenciesGetCurrenciesPhpGet200Response(val *ApiCurrenciesGetCurrenciesPhpGet200Response) *NullableApiCurrenciesGetCurrenciesPhpGet200Response {
	return &NullableApiCurrenciesGetCurrenciesPhpGet200Response{value: val, isSet: true}
}

func (v NullableApiCurrenciesGetCurrenciesPhpGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCurrenciesGetCurrenciesPhpGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
