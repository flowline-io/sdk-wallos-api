# ApiSubscriptionsGetSubscriptionsPhpGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Whether the request was successful. | [optional] 
**Title** | Pointer to **string** | The title of the response, typically &#39;subscriptions&#39;. | [optional] 
**Subscriptions** | Pointer to [**[]ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner**](ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner.md) |  | [optional] 
**Notes** | Pointer to **[]string** | Warning messages or additional information. | [optional] 

## Methods

### NewApiSubscriptionsGetSubscriptionsPhpGet200Response

`func NewApiSubscriptionsGetSubscriptionsPhpGet200Response() *ApiSubscriptionsGetSubscriptionsPhpGet200Response`

NewApiSubscriptionsGetSubscriptionsPhpGet200Response instantiates a new ApiSubscriptionsGetSubscriptionsPhpGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseWithDefaults

`func NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseWithDefaults() *ApiSubscriptionsGetSubscriptionsPhpGet200Response`

NewApiSubscriptionsGetSubscriptionsPhpGet200ResponseWithDefaults instantiates a new ApiSubscriptionsGetSubscriptionsPhpGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTitle

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetSubscriptions() []ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetSubscriptionsOk() (*[]ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) SetSubscriptions(v []ApiSubscriptionsGetSubscriptionsPhpGet200ResponseSubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetNotes

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiSubscriptionsGetSubscriptionsPhpGet200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


