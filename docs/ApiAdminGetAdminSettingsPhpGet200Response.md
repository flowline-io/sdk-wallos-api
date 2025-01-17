# ApiAdminGetAdminSettingsPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates whether the request was successful. | [optional] 
**Title** | Pointer to **string** | The title of the response. | [optional] 
**AdminSettings** | Pointer to [**ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings**](ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings.md) |  | [optional] 
**Notes** | Pointer to **[]string** | Warning messages or additional information. | [optional] 

## Methods

### NewApiAdminGetAdminSettingsPhpGet200Response

`func NewApiAdminGetAdminSettingsPhpGet200Response() *ApiAdminGetAdminSettingsPhpGet200Response`

NewApiAdminGetAdminSettingsPhpGet200Response instantiates a new ApiAdminGetAdminSettingsPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAdminGetAdminSettingsPhpGet200ResponseWithDefaults

`func NewApiAdminGetAdminSettingsPhpGet200ResponseWithDefaults() *ApiAdminGetAdminSettingsPhpGet200Response`

NewApiAdminGetAdminSettingsPhpGet200ResponseWithDefaults instantiates a new ApiAdminGetAdminSettingsPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAdminSettings

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetAdminSettings() ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings`

GetAdminSettings returns the AdminSettings field if non-nil, zero value otherwise.

### GetAdminSettingsOk

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetAdminSettingsOk() (*ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings, bool)`

GetAdminSettingsOk returns a tuple with the AdminSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSettings

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetAdminSettings(v ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings)`

SetAdminSettings sets AdminSettings field to given value.

### HasAdminSettings

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasAdminSettings() bool`

HasAdminSettings returns a boolean if a field has been set.

### GetNotes

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiAdminGetAdminSettingsPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


