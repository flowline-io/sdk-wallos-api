# ApiPaymentMethodsGetPaymentMethodsPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**PaymentMethods** | Pointer to [**[]ApiPaymentMethodsGetPaymentMethodsPhpGet200ResponsePaymentMethodsInner**](ApiPaymentMethodsGetPaymentMethodsPhpGet200ResponsePaymentMethodsInner.md) |  | [optional] 
**Notes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiPaymentMethodsGetPaymentMethodsPhpGet200Response

`func NewApiPaymentMethodsGetPaymentMethodsPhpGet200Response() *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response`

NewApiPaymentMethodsGetPaymentMethodsPhpGet200Response instantiates a new ApiPaymentMethodsGetPaymentMethodsPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPaymentMethodsGetPaymentMethodsPhpGet200ResponseWithDefaults

`func NewApiPaymentMethodsGetPaymentMethodsPhpGet200ResponseWithDefaults() *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response`

NewApiPaymentMethodsGetPaymentMethodsPhpGet200ResponseWithDefaults instantiates a new ApiPaymentMethodsGetPaymentMethodsPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetPaymentMethods() []ApiPaymentMethodsGetPaymentMethodsPhpGet200ResponsePaymentMethodsInner`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetPaymentMethodsOk() (*[]ApiPaymentMethodsGetPaymentMethodsPhpGet200ResponsePaymentMethodsInner, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) SetPaymentMethods(v []ApiPaymentMethodsGetPaymentMethodsPhpGet200ResponsePaymentMethodsInner)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetNotes

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiPaymentMethodsGetPaymentMethodsPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


