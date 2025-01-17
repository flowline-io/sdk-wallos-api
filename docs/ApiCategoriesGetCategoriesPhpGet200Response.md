# ApiCategoriesGetCategoriesPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates whether the request was successful. | [optional] 
**Title** | Pointer to **string** | The title of the response. | [optional] 
**Categories** | Pointer to [**[]ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner**](ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner.md) |  | [optional] 
**Notes** | Pointer to **[]string** | Warning messages or additional information. | [optional] 

## Methods

### NewApiCategoriesGetCategoriesPhpGet200Response

`func NewApiCategoriesGetCategoriesPhpGet200Response() *ApiCategoriesGetCategoriesPhpGet200Response`

NewApiCategoriesGetCategoriesPhpGet200Response instantiates a new ApiCategoriesGetCategoriesPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCategoriesGetCategoriesPhpGet200ResponseWithDefaults

`func NewApiCategoriesGetCategoriesPhpGet200ResponseWithDefaults() *ApiCategoriesGetCategoriesPhpGet200Response`

NewApiCategoriesGetCategoriesPhpGet200ResponseWithDefaults instantiates a new ApiCategoriesGetCategoriesPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategories

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetCategories() []ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetCategoriesOk() (*[]ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) SetCategories(v []ApiCategoriesGetCategoriesPhpGet200ResponseCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetNotes

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiCategoriesGetCategoriesPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


