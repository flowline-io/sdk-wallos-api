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

// checks if the ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner{}

// ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner struct for ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner
type ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner struct {
	Id                *int32   `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	Logo              *string  `json:"logo,omitempty"`
	Price             *float32 `json:"price,omitempty"`
	CurrencyId        *int32   `json:"currency_id,omitempty"`
	NextPayment       *string  `json:"next_payment,omitempty"`
	Cycle             *int32   `json:"cycle,omitempty"`
	Frequency         *int32   `json:"frequency,omitempty"`
	Notes             *string  `json:"notes,omitempty"`
	PaymentMethodId   *int32   `json:"payment_method_id,omitempty"`
	PayerUserId       *int32   `json:"payer_user_id,omitempty"`
	CategoryId        *int32   `json:"category_id,omitempty"`
	Notify            *bool    `json:"notify,omitempty"`
	Url               *string  `json:"url,omitempty"`
	Inactive          *bool    `json:"inactive,omitempty"`
	NotifyDaysBefore  *int32   `json:"notify_days_before,omitempty"`
	CancelationDate   *string  `json:"cancelation_date,omitempty"`
	CategoryName      *string  `json:"category_name,omitempty"`
	PayerUserName     *string  `json:"payer_user_name,omitempty"`
	PaymentMethodName *string  `json:"payment_method_name,omitempty"`
}

// NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner instantiates a new ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner() *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner {
	this := ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner{}
	return &this
}

// NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInnerWithDefaults instantiates a new ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInnerWithDefaults() *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner {
	this := ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetName(v string) {
	o.Name = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetLogo(v string) {
	o.Logo = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetPrice(v float32) {
	o.Price = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCurrencyId() int32 {
	if o == nil || IsNil(o.CurrencyId) {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetNextPayment returns the NextPayment field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNextPayment() string {
	if o == nil || IsNil(o.NextPayment) {
		var ret string
		return ret
	}
	return *o.NextPayment
}

// GetNextPaymentOk returns a tuple with the NextPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNextPaymentOk() (*string, bool) {
	if o == nil || IsNil(o.NextPayment) {
		return nil, false
	}
	return o.NextPayment, true
}

// HasNextPayment returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasNextPayment() bool {
	if o != nil && !IsNil(o.NextPayment) {
		return true
	}

	return false
}

// SetNextPayment gets a reference to the given string and assigns it to the NextPayment field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetNextPayment(v string) {
	o.NextPayment = &v
}

// GetCycle returns the Cycle field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCycle() int32 {
	if o == nil || IsNil(o.Cycle) {
		var ret int32
		return ret
	}
	return *o.Cycle
}

// GetCycleOk returns a tuple with the Cycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCycleOk() (*int32, bool) {
	if o == nil || IsNil(o.Cycle) {
		return nil, false
	}
	return o.Cycle, true
}

// HasCycle returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasCycle() bool {
	if o != nil && !IsNil(o.Cycle) {
		return true
	}

	return false
}

// SetCycle gets a reference to the given int32 and assigns it to the Cycle field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetCycle(v int32) {
	o.Cycle = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetFrequency() int32 {
	if o == nil || IsNil(o.Frequency) {
		var ret int32
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int32 and assigns it to the Frequency field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetFrequency(v int32) {
	o.Frequency = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetNotes(v string) {
	o.Notes = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPaymentMethodId() int32 {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret int32
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPaymentMethodIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given int32 and assigns it to the PaymentMethodId field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetPaymentMethodId(v int32) {
	o.PaymentMethodId = &v
}

// GetPayerUserId returns the PayerUserId field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPayerUserId() int32 {
	if o == nil || IsNil(o.PayerUserId) {
		var ret int32
		return ret
	}
	return *o.PayerUserId
}

// GetPayerUserIdOk returns a tuple with the PayerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPayerUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PayerUserId) {
		return nil, false
	}
	return o.PayerUserId, true
}

// HasPayerUserId returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasPayerUserId() bool {
	if o != nil && !IsNil(o.PayerUserId) {
		return true
	}

	return false
}

// SetPayerUserId gets a reference to the given int32 and assigns it to the PayerUserId field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetPayerUserId(v int32) {
	o.PayerUserId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCategoryId() int32 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCategoryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetNotify(v bool) {
	o.Notify = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetUrl(v string) {
	o.Url = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetInactive(v bool) {
	o.Inactive = &v
}

// GetNotifyDaysBefore returns the NotifyDaysBefore field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNotifyDaysBefore() int32 {
	if o == nil || IsNil(o.NotifyDaysBefore) {
		var ret int32
		return ret
	}
	return *o.NotifyDaysBefore
}

// GetNotifyDaysBeforeOk returns a tuple with the NotifyDaysBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetNotifyDaysBeforeOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifyDaysBefore) {
		return nil, false
	}
	return o.NotifyDaysBefore, true
}

// HasNotifyDaysBefore returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasNotifyDaysBefore() bool {
	if o != nil && !IsNil(o.NotifyDaysBefore) {
		return true
	}

	return false
}

// SetNotifyDaysBefore gets a reference to the given int32 and assigns it to the NotifyDaysBefore field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetNotifyDaysBefore(v int32) {
	o.NotifyDaysBefore = &v
}

// GetCancelationDate returns the CancelationDate field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCancelationDate() string {
	if o == nil || IsNil(o.CancelationDate) {
		var ret string
		return ret
	}
	return *o.CancelationDate
}

// GetCancelationDateOk returns a tuple with the CancelationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCancelationDateOk() (*string, bool) {
	if o == nil || IsNil(o.CancelationDate) {
		return nil, false
	}
	return o.CancelationDate, true
}

// HasCancelationDate returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasCancelationDate() bool {
	if o != nil && !IsNil(o.CancelationDate) {
		return true
	}

	return false
}

// SetCancelationDate gets a reference to the given string and assigns it to the CancelationDate field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetCancelationDate(v string) {
	o.CancelationDate = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName) {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryName) {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasCategoryName() bool {
	if o != nil && !IsNil(o.CategoryName) {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetPayerUserName returns the PayerUserName field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPayerUserName() string {
	if o == nil || IsNil(o.PayerUserName) {
		var ret string
		return ret
	}
	return *o.PayerUserName
}

// GetPayerUserNameOk returns a tuple with the PayerUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPayerUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayerUserName) {
		return nil, false
	}
	return o.PayerUserName, true
}

// HasPayerUserName returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasPayerUserName() bool {
	if o != nil && !IsNil(o.PayerUserName) {
		return true
	}

	return false
}

// SetPayerUserName gets a reference to the given string and assigns it to the PayerUserName field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetPayerUserName(v string) {
	o.PayerUserName = &v
}

// GetPaymentMethodName returns the PaymentMethodName field value if set, zero value otherwise.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPaymentMethodName() string {
	if o == nil || IsNil(o.PaymentMethodName) {
		var ret string
		return ret
	}
	return *o.PaymentMethodName
}

// GetPaymentMethodNameOk returns a tuple with the PaymentMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) GetPaymentMethodNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodName) {
		return nil, false
	}
	return o.PaymentMethodName, true
}

// HasPaymentMethodName returns a boolean if a field has been set.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) HasPaymentMethodName() bool {
	if o != nil && !IsNil(o.PaymentMethodName) {
		return true
	}

	return false
}

// SetPaymentMethodName gets a reference to the given string and assigns it to the PaymentMethodName field.
func (o *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) SetPaymentMethodName(v string) {
	o.PaymentMethodName = &v
}

func (o ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.NextPayment) {
		toSerialize["next_payment"] = o.NextPayment
	}
	if !IsNil(o.Cycle) {
		toSerialize["cycle"] = o.Cycle
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if !IsNil(o.PayerUserId) {
		toSerialize["payer_user_id"] = o.PayerUserId
	}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.NotifyDaysBefore) {
		toSerialize["notify_days_before"] = o.NotifyDaysBefore
	}
	if !IsNil(o.CancelationDate) {
		toSerialize["cancelation_date"] = o.CancelationDate
	}
	if !IsNil(o.CategoryName) {
		toSerialize["category_name"] = o.CategoryName
	}
	if !IsNil(o.PayerUserName) {
		toSerialize["payer_user_name"] = o.PayerUserName
	}
	if !IsNil(o.PaymentMethodName) {
		toSerialize["payment_method_name"] = o.PaymentMethodName
	}
	return toSerialize, nil
}

type NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner struct {
	value *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner
	isSet bool
}

func (v NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) Get() *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner {
	return v.value
}

func (v *NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) Set(val *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner(val *ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) *NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner {
	return &NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner{value: val, isSet: true}
}

func (v NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
