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

// checks if the ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner{}

// ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner struct for ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner
type ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner struct {
	// The unique identifier for the category.
	Id *int32 `json:"id,omitempty"`
	// The name of the category.
	Name *string `json:"name,omitempty"`
	// The display order of the category.
	Order *int32 `json:"order,omitempty"`
	// Indicates if the category is currently in use.
	InUse *bool `json:"in_use,omitempty"`
}

// NewApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner instantiates a new ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner() *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner {
	this := ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner{}
	return &this
}

// NewApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInnerWithDefaults instantiates a new ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInnerWithDefaults() *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner {
	this := ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) SetOrder(v int32) {
	o.Order = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetInUse() bool {
	if o == nil || IsNil(o.InUse) {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) GetInUseOk() (*bool, bool) {
	if o == nil || IsNil(o.InUse) {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) HasInUse() bool {
	if o != nil && !IsNil(o.InUse) {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) SetInUse(v bool) {
	o.InUse = &v
}

func (o ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.InUse) {
		toSerialize["in_use"] = o.InUse
	}
	return toSerialize, nil
}

type NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner struct {
	value *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner
	isSet bool
}

func (v NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) Get() *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner {
	return v.value
}

func (v *NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) Set(val *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner(val *ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) *NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner {
	return &NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner{value: val, isSet: true}
}

func (v NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
