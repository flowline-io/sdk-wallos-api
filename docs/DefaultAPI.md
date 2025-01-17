# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAdminGetAdminSettingsPhpGet**](DefaultAPI.md#ApiAdminGetAdminSettingsPhpGet) | **Get** /api/admin/get_admin_settings.php | Get Admin Settings
[**ApiCategoriesGetCategoriesPhpGet**](DefaultAPI.md#ApiCategoriesGetCategoriesPhpGet) | **Get** /api/categories/get_categories.php | Get Categories
[**ApiCurrenciesGetCurrenciesPhpGet**](DefaultAPI.md#ApiCurrenciesGetCurrenciesPhpGet) | **Get** /api/currencies/get_currencies.php | Get Currencies
[**ApiFixerGetFixerPhpGet**](DefaultAPI.md#ApiFixerGetFixerPhpGet) | **Get** /api/fixer/get_fixer.php | Get Fixer Settings
[**ApiHouseholdGetHouseholdPhpGet**](DefaultAPI.md#ApiHouseholdGetHouseholdPhpGet) | **Get** /api/household/get_household.php | Get Household
[**ApiNotificationsGetNotificationSettingsPhpGet**](DefaultAPI.md#ApiNotificationsGetNotificationSettingsPhpGet) | **Get** /api/notifications/get_notification_settings.php | Get Notification Settings
[**ApiPaymentMethodsGetPaymentMethodsPhpGet**](DefaultAPI.md#ApiPaymentMethodsGetPaymentMethodsPhpGet) | **Get** /api/payment_methods/get_payment_methods.php | Get Payment Methods
[**ApiSettingsGetSettingsPhpGet**](DefaultAPI.md#ApiSettingsGetSettingsPhpGet) | **Get** /api/settings/get_settings.php | Get User Settings
[**ApiSubscriptionsGetMonthlyCostPhpGet**](DefaultAPI.md#ApiSubscriptionsGetMonthlyCostPhpGet) | **Get** /api/subscriptions/get_monthly_cost.php | Get Monthly Cost
[**ApiSubscriptionsGetSubscriptionsPhpGet**](DefaultAPI.md#ApiSubscriptionsGetSubscriptionsPhpGet) | **Get** /api/subscriptions/get_subscriptions.php | Get Subscriptions



## ApiAdminGetAdminSettingsPhpGet

> ApiAdminGetAdminSettingsPhpGet200Response ApiAdminGetAdminSettingsPhpGet(ctx).ApiKey(apiKey).Execute()

Get Admin Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | API key of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiAdminGetAdminSettingsPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiAdminGetAdminSettingsPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiAdminGetAdminSettingsPhpGet`: ApiAdminGetAdminSettingsPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiAdminGetAdminSettingsPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAdminGetAdminSettingsPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key of the user. | 

### Return type

[**ApiAdminGetAdminSettingsPhpGet200Response**](ApiAdminGetAdminSettingsPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCategoriesGetCategoriesPhpGet

> ApiCategoriesGetCategoriesPhpGet200Response ApiCategoriesGetCategoriesPhpGet(ctx).ApiKey(apiKey).Execute()

Get Categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | API key of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiCategoriesGetCategoriesPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiCategoriesGetCategoriesPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCategoriesGetCategoriesPhpGet`: ApiCategoriesGetCategoriesPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiCategoriesGetCategoriesPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCategoriesGetCategoriesPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key of the user. | 

### Return type

[**ApiCategoriesGetCategoriesPhpGet200Response**](ApiCategoriesGetCategoriesPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiCurrenciesGetCurrenciesPhpGet

> ApiCurrenciesGetCurrenciesPhpGet200Response ApiCurrenciesGetCurrenciesPhpGet(ctx).ApiKey(apiKey).Execute()

Get Currencies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | API key of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiCurrenciesGetCurrenciesPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiCurrenciesGetCurrenciesPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCurrenciesGetCurrenciesPhpGet`: ApiCurrenciesGetCurrenciesPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiCurrenciesGetCurrenciesPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCurrenciesGetCurrenciesPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key of the user. | 

### Return type

[**ApiCurrenciesGetCurrenciesPhpGet200Response**](ApiCurrenciesGetCurrenciesPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiFixerGetFixerPhpGet

> ApiFixerGetFixerPhpGet200Response ApiFixerGetFixerPhpGet(ctx).ApiKey(apiKey).Execute()

Get Fixer Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | API key of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiFixerGetFixerPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiFixerGetFixerPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiFixerGetFixerPhpGet`: ApiFixerGetFixerPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiFixerGetFixerPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiFixerGetFixerPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | API key of the user. | 

### Return type

[**ApiFixerGetFixerPhpGet200Response**](ApiFixerGetFixerPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHouseholdGetHouseholdPhpGet

> ApiHouseholdGetHouseholdPhpGet200Response ApiHouseholdGetHouseholdPhpGet(ctx).ApiKey(apiKey).Execute()

Get Household



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiHouseholdGetHouseholdPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiHouseholdGetHouseholdPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHouseholdGetHouseholdPhpGet`: ApiHouseholdGetHouseholdPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiHouseholdGetHouseholdPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHouseholdGetHouseholdPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 

### Return type

[**ApiHouseholdGetHouseholdPhpGet200Response**](ApiHouseholdGetHouseholdPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiNotificationsGetNotificationSettingsPhpGet

> ApiNotificationsGetNotificationSettingsPhpGet200Response ApiNotificationsGetNotificationSettingsPhpGet(ctx).ApiKey(apiKey).Execute()

Get Notification Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiNotificationsGetNotificationSettingsPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiNotificationsGetNotificationSettingsPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiNotificationsGetNotificationSettingsPhpGet`: ApiNotificationsGetNotificationSettingsPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiNotificationsGetNotificationSettingsPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiNotificationsGetNotificationSettingsPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 

### Return type

[**ApiNotificationsGetNotificationSettingsPhpGet200Response**](ApiNotificationsGetNotificationSettingsPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPaymentMethodsGetPaymentMethodsPhpGet

> ApiPaymentMethodsGetPaymentMethodsPhpGet200Response ApiPaymentMethodsGetPaymentMethodsPhpGet(ctx).ApiKey(apiKey).Execute()

Get Payment Methods



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiPaymentMethodsGetPaymentMethodsPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiPaymentMethodsGetPaymentMethodsPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiPaymentMethodsGetPaymentMethodsPhpGet`: ApiPaymentMethodsGetPaymentMethodsPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiPaymentMethodsGetPaymentMethodsPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPaymentMethodsGetPaymentMethodsPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 

### Return type

[**ApiPaymentMethodsGetPaymentMethodsPhpGet200Response**](ApiPaymentMethodsGetPaymentMethodsPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSettingsGetSettingsPhpGet

> ApiSettingsGetSettingsPhpGet200Response ApiSettingsGetSettingsPhpGet(ctx).ApiKey(apiKey).Execute()

Get User Settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiSettingsGetSettingsPhpGet(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiSettingsGetSettingsPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSettingsGetSettingsPhpGet`: ApiSettingsGetSettingsPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiSettingsGetSettingsPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSettingsGetSettingsPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 

### Return type

[**ApiSettingsGetSettingsPhpGet200Response**](ApiSettingsGetSettingsPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubscriptionsGetMonthlyCostPhpGet

> ApiSubscriptionsGetMonthlyCostPhpGet200Response ApiSubscriptionsGetMonthlyCostPhpGet(ctx).ApiKey(apiKey).Month(month).Year(year).Execute()

Get Monthly Cost



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | The API key of the user.
	month := int32(56) // int32 | The month for which the cost is to be calculated.
	year := int32(56) // int32 | The year for which the cost is to be calculated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiSubscriptionsGetMonthlyCostPhpGet(context.Background()).ApiKey(apiKey).Month(month).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiSubscriptionsGetMonthlyCostPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubscriptionsGetMonthlyCostPhpGet`: ApiSubscriptionsGetMonthlyCostPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiSubscriptionsGetMonthlyCostPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsGetMonthlyCostPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | The API key of the user. | 
 **month** | **int32** | The month for which the cost is to be calculated. | 
 **year** | **int32** | The year for which the cost is to be calculated. | 

### Return type

[**ApiSubscriptionsGetMonthlyCostPhpGet200Response**](ApiSubscriptionsGetMonthlyCostPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSubscriptionsGetSubscriptionsPhpGet

> ApiSubscriptionsGetSubscriptionsPhpGet200Response ApiSubscriptionsGetSubscriptionsPhpGet(ctx).ApiKey(apiKey).Member(member).Category(category).PaymentMethod(paymentMethod).State(state).DisabledToBottom(disabledToBottom).Sort(sort).ConvertCurrency(convertCurrency).Execute()

Get Subscriptions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flowline-io/sdk-wallos-api"
)

func main() {
	apiKey := "apiKey_example" // string | The API key of the user.
	member := "member_example" // string | Comma-separated IDs of the members to filter. (optional)
	category := int32(56) // int32 | The ID of the category to filter. (optional)
	paymentMethod := int32(56) // int32 | The ID of the payment method to filter. (optional)
	state := true // bool | The state of the subscription to filter [0 - active, 1 - inactive]. (optional)
	disabledToBottom := true // bool | Whether to sort inactive subscriptions to the bottom. (optional)
	sort := "sort_example" // string | The sorting method ['name', 'id', 'next_payment', 'price', etc.]. (optional)
	convertCurrency := true // bool | Whether to convert the prices to the main currency. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiSubscriptionsGetSubscriptionsPhpGet(context.Background()).ApiKey(apiKey).Member(member).Category(category).PaymentMethod(paymentMethod).State(state).DisabledToBottom(disabledToBottom).Sort(sort).ConvertCurrency(convertCurrency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiSubscriptionsGetSubscriptionsPhpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSubscriptionsGetSubscriptionsPhpGet`: ApiSubscriptionsGetSubscriptionsPhpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiSubscriptionsGetSubscriptionsPhpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSubscriptionsGetSubscriptionsPhpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | The API key of the user. | 
 **member** | **string** | Comma-separated IDs of the members to filter. | 
 **category** | **int32** | The ID of the category to filter. | 
 **paymentMethod** | **int32** | The ID of the payment method to filter. | 
 **state** | **bool** | The state of the subscription to filter [0 - active, 1 - inactive]. | 
 **disabledToBottom** | **bool** | Whether to sort inactive subscriptions to the bottom. | 
 **sort** | **string** | The sorting method [&#39;name&#39;, &#39;id&#39;, &#39;next_payment&#39;, &#39;price&#39;, etc.]. | 
 **convertCurrency** | **bool** | Whether to convert the prices to the main currency. | 

### Return type

[**ApiSubscriptionsGetSubscriptionsPhpGet200Response**](ApiSubscriptionsGetSubscriptionsPhpGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

