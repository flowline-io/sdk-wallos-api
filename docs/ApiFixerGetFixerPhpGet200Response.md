# ApiFixerGetFixerPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates whether the request was successful. | [optional] 
**Title** | Pointer to **string** | The title of the response. | [optional] 
**Fixer** | Pointer to [**ApiFixerGetFixerPhpGet200ResponseFixer**](ApiFixerGetFixerPhpGet200ResponseFixer.md) |  | [optional] 
**Notes** | Pointer to **[]string** | Warning messages or additional information. | [optional] 

## Methods

### NewApiFixerGetFixerPhpGet200Response

`func NewApiFixerGetFixerPhpGet200Response() *ApiFixerGetFixerPhpGet200Response`

NewApiFixerGetFixerPhpGet200Response instantiates a new ApiFixerGetFixerPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFixerGetFixerPhpGet200ResponseWithDefaults

`func NewApiFixerGetFixerPhpGet200ResponseWithDefaults() *ApiFixerGetFixerPhpGet200Response`

NewApiFixerGetFixerPhpGet200ResponseWithDefaults instantiates a new ApiFixerGetFixerPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiFixerGetFixerPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiFixerGetFixerPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiFixerGetFixerPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiFixerGetFixerPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiFixerGetFixerPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiFixerGetFixerPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiFixerGetFixerPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiFixerGetFixerPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFixer

`func (o *ApiFixerGetFixerPhpGet200Response) GetFixer() ApiFixerGetFixerPhpGet200ResponseFixer`

GetFixer returns the Fixer field if non-nil, zero value otherwise.

### GetFixerOk

`func (o *ApiFixerGetFixerPhpGet200Response) GetFixerOk() (*ApiFixerGetFixerPhpGet200ResponseFixer, bool)`

GetFixerOk returns a tuple with the Fixer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixer

`func (o *ApiFixerGetFixerPhpGet200Response) SetFixer(v ApiFixerGetFixerPhpGet200ResponseFixer)`

SetFixer sets Fixer field to given value.

### HasFixer

`func (o *ApiFixerGetFixerPhpGet200Response) HasFixer() bool`

HasFixer returns a boolean if a field has been set.

### GetNotes

`func (o *ApiFixerGetFixerPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiFixerGetFixerPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiFixerGetFixerPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiFixerGetFixerPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


