# ApiSubscriptionsGetMonthlyCostPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Whether the request was successful. | [optional] 
**Title** | Pointer to **string** | A string representing the month and year (e.g., March 2025). | [optional] 
**MonthlyCost** | Pointer to **float32** | The total cost for the specified month. | [optional] 
**LocalizedMonthlyCost** | Pointer to **string** | The total cost formatted according to the user&#39;s locale and currency. | [optional] 
**CurrencyCode** | Pointer to **string** | The currency code of the user&#39;s main currency. | [optional] 
**CurrencySymbol** | Pointer to **string** | The currency symbol of the user&#39;s main currency. | [optional] 
**Notes** | Pointer to **[]string** | Warning messages or additional information. | [optional] 

## Methods

### NewApiSubscriptionsGetMonthlyCostPhpGet200Response

`func NewApiSubscriptionsGetMonthlyCostPhpGet200Response() *ApiSubscriptionsGetMonthlyCostPhpGet200Response`

NewApiSubscriptionsGetMonthlyCostPhpGet200Response instantiates a new ApiSubscriptionsGetMonthlyCostPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubscriptionsGetMonthlyCostPhpGet200ResponseWithDefaults

`func NewApiSubscriptionsGetMonthlyCostPhpGet200ResponseWithDefaults() *ApiSubscriptionsGetMonthlyCostPhpGet200Response`

NewApiSubscriptionsGetMonthlyCostPhpGet200ResponseWithDefaults instantiates a new ApiSubscriptionsGetMonthlyCostPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMonthlyCost

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetMonthlyCost() float32`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetMonthlyCostOk() (*float32, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetMonthlyCost(v float32)`

SetMonthlyCost sets MonthlyCost field to given value.

### HasMonthlyCost

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasMonthlyCost() bool`

HasMonthlyCost returns a boolean if a field has been set.

### GetLocalizedMonthlyCost

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetLocalizedMonthlyCost() string`

GetLocalizedMonthlyCost returns the LocalizedMonthlyCost field if non-nil, zero value otherwise.

### GetLocalizedMonthlyCostOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetLocalizedMonthlyCostOk() (*string, bool)`

GetLocalizedMonthlyCostOk returns a tuple with the LocalizedMonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedMonthlyCost

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetLocalizedMonthlyCost(v string)`

SetLocalizedMonthlyCost sets LocalizedMonthlyCost field to given value.

### HasLocalizedMonthlyCost

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasLocalizedMonthlyCost() bool`

HasLocalizedMonthlyCost returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetNotes

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiSubscriptionsGetMonthlyCostPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


