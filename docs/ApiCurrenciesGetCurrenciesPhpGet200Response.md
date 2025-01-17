# ApiCurrenciesGetCurrenciesPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates whether the request was successful. | [optional] 
**Title** | Pointer to **string** | The title of the response. | [optional] 
**MainCurrency** | Pointer to **int32** | The ID of the main currency of the user. | [optional] 
**Currencies** | Pointer to [**[]ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner**](ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner.md) |  | [optional] 
**Notes** | Pointer to **[]string** | Warning messages or additional information. | [optional] 

## Methods

### NewApiCurrenciesGetCurrenciesPhpGet200Response

`func NewApiCurrenciesGetCurrenciesPhpGet200Response() *ApiCurrenciesGetCurrenciesPhpGet200Response`

NewApiCurrenciesGetCurrenciesPhpGet200Response instantiates a new ApiCurrenciesGetCurrenciesPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCurrenciesGetCurrenciesPhpGet200ResponseWithDefaults

`func NewApiCurrenciesGetCurrenciesPhpGet200ResponseWithDefaults() *ApiCurrenciesGetCurrenciesPhpGet200Response`

NewApiCurrenciesGetCurrenciesPhpGet200ResponseWithDefaults instantiates a new ApiCurrenciesGetCurrenciesPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMainCurrency

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetMainCurrency() int32`

GetMainCurrency returns the MainCurrency field if non-nil, zero value otherwise.

### GetMainCurrencyOk

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetMainCurrencyOk() (*int32, bool)`

GetMainCurrencyOk returns a tuple with the MainCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainCurrency

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetMainCurrency(v int32)`

SetMainCurrency sets MainCurrency field to given value.

### HasMainCurrency

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasMainCurrency() bool`

HasMainCurrency returns a boolean if a field has been set.

### GetCurrencies

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetCurrencies() []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetCurrenciesOk() (*[]ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetCurrencies(v []ApiCurrenciesGetCurrenciesPhpGet200ResponseCurrenciesInner)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetNotes

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiCurrenciesGetCurrenciesPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


