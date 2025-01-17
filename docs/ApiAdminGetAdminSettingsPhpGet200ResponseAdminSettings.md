# ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationsOpen** | Pointer to **int32** | Indicates if registrations are open (1 for true, 0 for false). | [optional] 
**MaxUsers** | Pointer to **int32** | The maximum number of users allowed. | [optional] 
**RequireEmailVerification** | Pointer to **int32** | Indicates if email verification is required (1 for true, 0 for false). | [optional] 
**ServerUrl** | Pointer to **string** | The server URL. | [optional] 
**SmtpAddress** | Pointer to **string** | The SMTP address for sending emails. | [optional] 
**SmtpPort** | Pointer to **int32** | The port for the SMTP server. | [optional] 
**SmtpUsername** | Pointer to **string** | The username for SMTP authentication. | [optional] 
**SmtpPassword** | Pointer to **string** | The password for SMTP authentication. | [optional] 
**FromEmail** | Pointer to **string** | The email address used as the sender. | [optional] 
**Encryption** | Pointer to **string** | The encryption method for SMTP (e.g., tls, ssl). | [optional] 
**LoginDisabled** | Pointer to **int32** | Indicates if login is disabled (1 for true, 0 for false). | [optional] 
**LatestVersion** | Pointer to **string** | The latest version of the application. | [optional] 
**UpdateNotification** | Pointer to **int32** | Indicates if update notifications are enabled (1 for true, 0 for false). | [optional] 

## Methods

### NewApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings

`func NewApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings() *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings`

NewApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings instantiates a new ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAdminGetAdminSettingsPhpGet200ResponseAdminSettingsWithDefaults

`func NewApiAdminGetAdminSettingsPhpGet200ResponseAdminSettingsWithDefaults() *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings`

NewApiAdminGetAdminSettingsPhpGet200ResponseAdminSettingsWithDefaults instantiates a new ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationsOpen

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetRegistrationsOpen() int32`

GetRegistrationsOpen returns the RegistrationsOpen field if non-nil, zero value otherwise.

### GetRegistrationsOpenOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetRegistrationsOpenOk() (*int32, bool)`

GetRegistrationsOpenOk returns a tuple with the RegistrationsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationsOpen

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetRegistrationsOpen(v int32)`

SetRegistrationsOpen sets RegistrationsOpen field to given value.

### HasRegistrationsOpen

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasRegistrationsOpen() bool`

HasRegistrationsOpen returns a boolean if a field has been set.

### GetMaxUsers

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetRequireEmailVerification

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetRequireEmailVerification() int32`

GetRequireEmailVerification returns the RequireEmailVerification field if non-nil, zero value otherwise.

### GetRequireEmailVerificationOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetRequireEmailVerificationOk() (*int32, bool)`

GetRequireEmailVerificationOk returns a tuple with the RequireEmailVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEmailVerification

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetRequireEmailVerification(v int32)`

SetRequireEmailVerification sets RequireEmailVerification field to given value.

### HasRequireEmailVerification

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasRequireEmailVerification() bool`

HasRequireEmailVerification returns a boolean if a field has been set.

### GetServerUrl

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetSmtpAddress

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpAddress() string`

GetSmtpAddress returns the SmtpAddress field if non-nil, zero value otherwise.

### GetSmtpAddressOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpAddressOk() (*string, bool)`

GetSmtpAddressOk returns a tuple with the SmtpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAddress

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetSmtpAddress(v string)`

SetSmtpAddress sets SmtpAddress field to given value.

### HasSmtpAddress

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasSmtpAddress() bool`

HasSmtpAddress returns a boolean if a field has been set.

### GetSmtpPort

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpUsername

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpUsername() string`

GetSmtpUsername returns the SmtpUsername field if non-nil, zero value otherwise.

### GetSmtpUsernameOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpUsernameOk() (*string, bool)`

GetSmtpUsernameOk returns a tuple with the SmtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUsername

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetSmtpUsername(v string)`

SetSmtpUsername sets SmtpUsername field to given value.

### HasSmtpUsername

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasSmtpUsername() bool`

HasSmtpUsername returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetFromEmail

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.

### HasFromEmail

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasFromEmail() bool`

HasFromEmail returns a boolean if a field has been set.

### GetEncryption

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetLoginDisabled

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetLoginDisabled() int32`

GetLoginDisabled returns the LoginDisabled field if non-nil, zero value otherwise.

### GetLoginDisabledOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetLoginDisabledOk() (*int32, bool)`

GetLoginDisabledOk returns a tuple with the LoginDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisabled

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetLoginDisabled(v int32)`

SetLoginDisabled sets LoginDisabled field to given value.

### HasLoginDisabled

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasLoginDisabled() bool`

HasLoginDisabled returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetUpdateNotification

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetUpdateNotification() int32`

GetUpdateNotification returns the UpdateNotification field if non-nil, zero value otherwise.

### GetUpdateNotificationOk

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) GetUpdateNotificationOk() (*int32, bool)`

GetUpdateNotificationOk returns a tuple with the UpdateNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateNotification

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) SetUpdateNotification(v int32)`

SetUpdateNotification sets UpdateNotification field to given value.

### HasUpdateNotification

`func (o *ApiAdminGetAdminSettingsPhpGet200ResponseAdminSettings) HasUpdateNotification() bool`

HasUpdateNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


