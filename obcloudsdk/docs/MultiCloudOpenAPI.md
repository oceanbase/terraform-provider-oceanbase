# \MultiCloudOpenAPI

All URIs are relative to *http://api.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivateLinkServiceUser**](MultiCloudOpenAPI.md#AddPrivateLinkServiceUser) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink/user | 租户private link service增加白名单用户
[**ConfirmPrivatelinkConnection**](MultiCloudOpenAPI.md#ConfirmPrivatelinkConnection) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink/confirmation | 确认完成private link私网连接
[**ConnectPrivateLinkService**](MultiCloudOpenAPI.md#ConnectPrivateLinkService) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink/connection | private link endpoint连接service
[**CreateDatabase**](MultiCloudOpenAPI.md#CreateDatabase) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/databases | CreateDatabase
[**CreateInstance**](MultiCloudOpenAPI.md#CreateInstance) | **Post** /api/v2/instances | CreateInstance
[**CreatePrivateLinkService**](MultiCloudOpenAPI.md#CreatePrivateLinkService) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink/service | 创建租户的private link主地址
[**CreateReadonlyPrivatelink**](MultiCloudOpenAPI.md#CreateReadonlyPrivatelink) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink/readonly | 创建租户的私网只读地址
[**CreateTenant**](MultiCloudOpenAPI.md#CreateTenant) | **Post** /api/v2/instances/{instanceId}/tenants | 创建租户
[**CreateTenantAddress**](MultiCloudOpenAPI.md#CreateTenantAddress) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/address | 新建租户singleTunnelSLB地址
[**CreateTenantSecurityIpGroup**](MultiCloudOpenAPI.md#CreateTenantSecurityIpGroup) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/securityIpGroups | CreateTenantSecurityIpGroup
[**CreateTenantUser**](MultiCloudOpenAPI.md#CreateTenantUser) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers | 创建租户内的用户
[**DeleteDatabase**](MultiCloudOpenAPI.md#DeleteDatabase) | **Delete** /api/v2/instances/{instanceId}/tenants/{tenantId}/databases | 
[**DeleteInstance**](MultiCloudOpenAPI.md#DeleteInstance) | **Delete** /api/v2/instances | DeleteInstance
[**DeletePrivatelinkEndpoint**](MultiCloudOpenAPI.md#DeletePrivatelinkEndpoint) | **Delete** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink | 删除privatelink地址
[**DeleteTenant**](MultiCloudOpenAPI.md#DeleteTenant) | **Delete** /api/v2/instances/{instanceId}/tenants/{tenantId} | 删除租户
[**DeleteTenantAddress**](MultiCloudOpenAPI.md#DeleteTenantAddress) | **Delete** /api/v2/instances/{instanceId}/tenants/{tenantId}/address | 删除租户地址
[**DeleteTenantSecurityIpGroup**](MultiCloudOpenAPI.md#DeleteTenantSecurityIpGroup) | **Delete** /api/v2/instances/{instanceId}/tenants/{tenantId}/securityIpGroups | DeleteTenantSecurityIpGroup
[**DeleteTenantUsers**](MultiCloudOpenAPI.md#DeleteTenantUsers) | **Delete** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers | 删除租户内特定用户
[**DescribeDatabases**](MultiCloudOpenAPI.md#DescribeDatabases) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId}/databases | DescribeDatabases
[**DescribeInstance**](MultiCloudOpenAPI.md#DescribeInstance) | **Get** /api/v2/instances/{instanceId} | DescribeInstance
[**DescribeInstances**](MultiCloudOpenAPI.md#DescribeInstances) | **Get** /api/v2/instances | DescribeInstances
[**DescribeTenant**](MultiCloudOpenAPI.md#DescribeTenant) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId} | 查询特定租户
[**DescribeTenantAddress**](MultiCloudOpenAPI.md#DescribeTenantAddress) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId}/address | 查询租户地址
[**DescribeTenantPrivateLink**](MultiCloudOpenAPI.md#DescribeTenantPrivateLink) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink | 查询租户的privatelink地址信息
[**DescribeTenantPublicAddress**](MultiCloudOpenAPI.md#DescribeTenantPublicAddress) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId}/publicaddress | 查询租户的公网主地址
[**DescribeTenantSecurityIpGroups**](MultiCloudOpenAPI.md#DescribeTenantSecurityIpGroups) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId}/securityIpGroups | DescribeTenantSecurityIpGroups
[**DescribeTenantUsers**](MultiCloudOpenAPI.md#DescribeTenantUsers) | **Get** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers | 查询租户内的用户
[**ModifyDatabaseDescription**](MultiCloudOpenAPI.md#ModifyDatabaseDescription) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/databases/{databaseName}/description | ModifyDatabaseDescription
[**ModifyDatabaseUserRole**](MultiCloudOpenAPI.md#ModifyDatabaseUserRole) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/databases/{databaseId}/authorization | 修改角色列表
[**ModifyInstanceName**](MultiCloudOpenAPI.md#ModifyInstanceName) | **Put** /api/v2/instances/{instanceId}/instanceName | ModifyInstanceName
[**ModifyInstanceNodeNum**](MultiCloudOpenAPI.md#ModifyInstanceNodeNum) | **Put** /api/v2/instances/{instanceId}/nodeNum | ModifyInstanceNodeNum
[**ModifyInstanceSpec**](MultiCloudOpenAPI.md#ModifyInstanceSpec) | **Put** /api/v2/instances/{instanceId}/spec | ModifyInstanceSpec
[**ModifyTenantName**](MultiCloudOpenAPI.md#ModifyTenantName) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantName | 修改租户名称
[**ModifyTenantPrimaryZone**](MultiCloudOpenAPI.md#ModifyTenantPrimaryZone) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/primaryZone | 修改租户主可用区
[**ModifyTenantSecurityIpGroup**](MultiCloudOpenAPI.md#ModifyTenantSecurityIpGroup) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/securityIpGroups | ModifyTenantSecurityIpGroup
[**ModifyTenantUserDescription**](MultiCloudOpenAPI.md#ModifyTenantUserDescription) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers/{userName}/description | 修改租户内的用户备注信息
[**ModifyTenantUserPassword**](MultiCloudOpenAPI.md#ModifyTenantUserPassword) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers/{userName}/password | 修改租户内的用户密码
[**ModifyTenantUserRole**](MultiCloudOpenAPI.md#ModifyTenantUserRole) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers/{userName}/authorization | 修改用户权限
[**ModifyTenantUserStatus**](MultiCloudOpenAPI.md#ModifyTenantUserStatus) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId}/tenantUsers/{userName}/status | 修改租户内的用户锁定状态
[**QueryTenants**](MultiCloudOpenAPI.md#QueryTenants) | **Get** /api/v2/instances/{instanceId}/tenants | 条件查询租户
[**StartCluster**](MultiCloudOpenAPI.md#StartCluster) | **Put** /api/v2/instances/{instanceId}/startCluster | StartCluster
[**StopCluster**](MultiCloudOpenAPI.md#StopCluster) | **Put** /api/v2/instances/{instanceId}/stopCluster | StopCluster
[**UpdatePrivatelinkInformation**](MultiCloudOpenAPI.md#UpdatePrivatelinkInformation) | **Post** /api/v2/instances/{instanceId}/tenants/{tenantId}/privatelink/information | 更新租户private link信息
[**UpdateTenant**](MultiCloudOpenAPI.md#UpdateTenant) | **Put** /api/v2/instances/{instanceId}/tenants/{tenantId} | 修改租户



## AddPrivateLinkServiceUser

> OBCloudResultAddTenantPlServiceUserResponse AddPrivateLinkServiceUser(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

租户private link service增加白名单用户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewAddTenantPlServiceUserRequest() // AddTenantPlServiceUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.AddPrivateLinkServiceUser(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.AddPrivateLinkServiceUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPrivateLinkServiceUser`: OBCloudResultAddTenantPlServiceUserResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.AddPrivateLinkServiceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivateLinkServiceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**AddTenantPlServiceUserRequest**](AddTenantPlServiceUserRequest.md) |  | 

### Return type

[**OBCloudResultAddTenantPlServiceUserResponse**](OBCloudResultAddTenantPlServiceUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPrivatelinkConnection

> OBCloudResultBoolean ConfirmPrivatelinkConnection(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

确认完成private link私网连接

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewConfirmTenantPlServiceRequest() // ConfirmTenantPlServiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ConfirmPrivatelinkConnection(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ConfirmPrivatelinkConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmPrivatelinkConnection`: OBCloudResultBoolean
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ConfirmPrivatelinkConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPrivatelinkConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**ConfirmTenantPlServiceRequest**](ConfirmTenantPlServiceRequest.md) |  | 

### Return type

[**OBCloudResultBoolean**](OBCloudResultBoolean.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectPrivateLinkService

> OBCloudResultConnectTenantPlServiceResponse ConnectPrivateLinkService(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

private link endpoint连接service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewConnectTenantPlServiceRequest() // ConnectTenantPlServiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ConnectPrivateLinkService(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ConnectPrivateLinkService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectPrivateLinkService`: OBCloudResultConnectTenantPlServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ConnectPrivateLinkService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectPrivateLinkServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**ConnectTenantPlServiceRequest**](ConnectTenantPlServiceRequest.md) |  | 

### Return type

[**OBCloudResultConnectTenantPlServiceResponse**](OBCloudResultConnectTenantPlServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabase

> OBCloudResultMapStringString CreateDatabase(ctx, instanceId, tenantId).Body(body).Execute()

CreateDatabase

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewCreateDatabaseRequestV2() // CreateDatabaseRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateDatabase(context.Background(), instanceId, tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDatabase`: OBCloudResultMapStringString
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateDatabaseRequestV2**](CreateDatabaseRequestV2.md) |  | 

### Return type

[**OBCloudResultMapStringString**](OBCloudResultMapStringString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstance

> OBCloudResultCreateInstanceResponseV2OpenApi CreateInstance(ctx).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()

CreateInstance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestId := "requestId_example" // string | 
	xObProjectId := "xObProjectId_example" // string |  (optional)
	body := *openapiclient.NewMultiCloudCreateInstanceRequest("CloudProvider_example", int32(123), "Region_example", "ObVersion_example", "Zones_example", "InstanceClass_example") // MultiCloudCreateInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateInstance(context.Background()).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: OBCloudResultCreateInstanceResponseV2OpenApi
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** |  | 
 **xObProjectId** | **string** |  | 
 **body** | [**MultiCloudCreateInstanceRequest**](MultiCloudCreateInstanceRequest.md) |  | 

### Return type

[**OBCloudResultCreateInstanceResponseV2OpenApi**](OBCloudResultCreateInstanceResponseV2OpenApi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateLinkService

> OBCloudResultCreateTenantPlServiceResponse CreatePrivateLinkService(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

创建租户的private link主地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewCreateTenantPlServiceRequest() // CreateTenantPlServiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreatePrivateLinkService(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreatePrivateLinkService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateLinkService`: OBCloudResultCreateTenantPlServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreatePrivateLinkService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateLinkServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**CreateTenantPlServiceRequest**](CreateTenantPlServiceRequest.md) |  | 

### Return type

[**OBCloudResultCreateTenantPlServiceResponse**](OBCloudResultCreateTenantPlServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReadonlyPrivatelink

> OBCloudResultCreateReadonlyConnectionResponse CreateReadonlyPrivatelink(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

创建租户的私网只读地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewCreateReadonlyConnectionRequest() // CreateReadonlyConnectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateReadonlyPrivatelink(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateReadonlyPrivatelink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReadonlyPrivatelink`: OBCloudResultCreateReadonlyConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateReadonlyPrivatelink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReadonlyPrivatelinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**CreateReadonlyConnectionRequest**](CreateReadonlyConnectionRequest.md) |  | 

### Return type

[**OBCloudResultCreateReadonlyConnectionResponse**](OBCloudResultCreateReadonlyConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenant

> OBCloudResultCreateTenantResponse CreateTenant(ctx, instanceId).RequestId(requestId).Body(body).Execute()

创建租户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewCreateTenantRequest() // CreateTenantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateTenant(context.Background(), instanceId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: OBCloudResultCreateTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 
 **body** | [**CreateTenantRequest**](CreateTenantRequest.md) |  | 

### Return type

[**OBCloudResultCreateTenantResponse**](OBCloudResultCreateTenantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantAddress

> OBCloudResultCreateMcTenantAddressResponse CreateTenantAddress(ctx, instanceId, tenantId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()

新建租户singleTunnelSLB地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	xObProjectId := "xObProjectId_example" // string |  (optional)
	body := *openapiclient.NewCreateMcTenantAddressRequest() // CreateMcTenantAddressRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateTenantAddress(context.Background(), instanceId, tenantId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateTenantAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantAddress`: OBCloudResultCreateMcTenantAddressResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateTenantAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **xObProjectId** | **string** |  | 
 **body** | [**CreateMcTenantAddressRequest**](CreateMcTenantAddressRequest.md) |  | 

### Return type

[**OBCloudResultCreateMcTenantAddressResponse**](OBCloudResultCreateMcTenantAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantSecurityIpGroup

> OBCloudResultMcModifyTenantSecurityIpGroupResponse CreateTenantSecurityIpGroup(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

CreateTenantSecurityIpGroup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewMcModifyTenantSecurityIpGroupRequest() // McModifyTenantSecurityIpGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateTenantSecurityIpGroup(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateTenantSecurityIpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantSecurityIpGroup`: OBCloudResultMcModifyTenantSecurityIpGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateTenantSecurityIpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantSecurityIpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**McModifyTenantSecurityIpGroupRequest**](McModifyTenantSecurityIpGroupRequest.md) |  | 

### Return type

[**OBCloudResultMcModifyTenantSecurityIpGroupResponse**](OBCloudResultMcModifyTenantSecurityIpGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantUser

> OBCloudResultCreateUserResponseDtoV2 CreateTenantUser(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

创建租户内的用户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewCreateUserParamDo() // CreateUserParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.CreateTenantUser(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.CreateTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantUser`: OBCloudResultCreateUserResponseDtoV2
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.CreateTenantUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**CreateUserParamDo**](CreateUserParamDo.md) |  | 

### Return type

[**OBCloudResultCreateUserResponseDtoV2**](OBCloudResultCreateUserResponseDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> OBCloudResult DeleteDatabase(ctx, instanceId, tenantId).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewDeleteDatabasesParamDo() // DeleteDatabasesParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeleteDatabase(context.Background(), instanceId, tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeleteDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabase`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeleteDatabasesParamDo**](DeleteDatabasesParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> OBCloudResultDeleteInstanceResponse DeleteInstance(ctx).RequestId(requestId).Body(body).Execute()

DeleteInstance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewDeleteInstanceOpenRequest() // DeleteInstanceOpenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeleteInstance(context.Background()).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstance`: OBCloudResultDeleteInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** |  | 
 **body** | [**DeleteInstanceOpenRequest**](DeleteInstanceOpenRequest.md) |  | 

### Return type

[**OBCloudResultDeleteInstanceResponse**](OBCloudResultDeleteInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivatelinkEndpoint

> OBCloudResultBoolean DeletePrivatelinkEndpoint(ctx, instanceId, tenantId).RequestId(requestId).AddressId(addressId).EndpointId(endpointId).Status(status).Execute()

删除privatelink地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	addressId := "addressId_example" // string |  (optional)
	endpointId := "endpointId_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeletePrivatelinkEndpoint(context.Background(), instanceId, tenantId).RequestId(requestId).AddressId(addressId).EndpointId(endpointId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeletePrivatelinkEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePrivatelinkEndpoint`: OBCloudResultBoolean
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeletePrivatelinkEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivatelinkEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **addressId** | **string** |  | 
 **endpointId** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**OBCloudResultBoolean**](OBCloudResultBoolean.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> OBCloudResultSimpleTenantResponse DeleteTenant(ctx, instanceId, tenantId).RequestId(requestId).Execute()

删除租户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeleteTenant(context.Background(), instanceId, tenantId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenant`: OBCloudResultSimpleTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeleteTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 

### Return type

[**OBCloudResultSimpleTenantResponse**](OBCloudResultSimpleTenantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantAddress

> OBCloudResultBoolean DeleteTenantAddress(ctx, instanceId, tenantId).RequestId(requestId).AddressId(addressId).Execute()

删除租户地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	addressId := "addressId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeleteTenantAddress(context.Background(), instanceId, tenantId).RequestId(requestId).AddressId(addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeleteTenantAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantAddress`: OBCloudResultBoolean
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeleteTenantAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **addressId** | **string** |  | 

### Return type

[**OBCloudResultBoolean**](OBCloudResultBoolean.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantSecurityIpGroup

> OBCloudResultMcDeleteTenantSecurityIpGroupResponse DeleteTenantSecurityIpGroup(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

DeleteTenantSecurityIpGroup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewMcDeleteTenantSecurityIpGroupRequest() // McDeleteTenantSecurityIpGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeleteTenantSecurityIpGroup(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeleteTenantSecurityIpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantSecurityIpGroup`: OBCloudResultMcDeleteTenantSecurityIpGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeleteTenantSecurityIpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantSecurityIpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**McDeleteTenantSecurityIpGroupRequest**](McDeleteTenantSecurityIpGroupRequest.md) |  | 

### Return type

[**OBCloudResultMcDeleteTenantSecurityIpGroupResponse**](OBCloudResultMcDeleteTenantSecurityIpGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantUsers

> OBCloudResult DeleteTenantUsers(ctx, instanceId, tenantId).Body(body).Execute()

删除租户内特定用户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewDeleteTenantUserParamDo() // DeleteTenantUserParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DeleteTenantUsers(context.Background(), instanceId, tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DeleteTenantUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantUsers`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DeleteTenantUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeleteTenantUserParamDo**](DeleteTenantUserParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeDatabases

> OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2 DescribeDatabases(ctx, instanceId, tenantId).RequestId(requestId).DatabaseName(databaseName).PageNumber(pageNumber).PageSize(pageSize).Execute()

DescribeDatabases

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	databaseName := "databaseName_example" // string |  (optional)
	pageNumber := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeDatabases(context.Background(), instanceId, tenantId).RequestId(requestId).DatabaseName(databaseName).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeDatabases`: OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **databaseName** | **string** |  | 
 **pageNumber** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2**](OBCloudResultOBCloudPagingDataDescribeDatabasesResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeInstance

> OBCloudResultDescribeInstanceResponseV2OpenAPI DescribeInstance(ctx, instanceId).RequestId(requestId).Execute()

DescribeInstance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeInstance(context.Background(), instanceId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeInstance`: OBCloudResultDescribeInstanceResponseV2OpenAPI
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 

### Return type

[**OBCloudResultDescribeInstanceResponseV2OpenAPI**](OBCloudResultDescribeInstanceResponseV2OpenAPI.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeInstances

> OBCloudResultOBCloudPagingDataListDescribeInstancesResponseV2OpenAPI DescribeInstances(ctx).RequestId(requestId).InstanceName(instanceName).InstanceId(instanceId).XObProjectId(xObProjectId).TagList(tagList).OrganizationId(organizationId).PageNumber(pageNumber).PageSize(pageSize).Execute()

DescribeInstances

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestId := "requestId_example" // string | 
	instanceName := "instanceName_example" // string |  (optional)
	instanceId := "instanceId_example" // string |  (optional)
	xObProjectId := "xObProjectId_example" // string |  (optional)
	tagList := "tagList_example" // string |  (optional)
	organizationId := "organizationId_example" // string |  (optional)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeInstances(context.Background()).RequestId(requestId).InstanceName(instanceName).InstanceId(instanceId).XObProjectId(xObProjectId).TagList(tagList).OrganizationId(organizationId).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeInstances`: OBCloudResultOBCloudPagingDataListDescribeInstancesResponseV2OpenAPI
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** |  | 
 **instanceName** | **string** |  | 
 **instanceId** | **string** |  | 
 **xObProjectId** | **string** |  | 
 **tagList** | **string** |  | 
 **organizationId** | **string** |  | 
 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]

### Return type

[**OBCloudResultOBCloudPagingDataListDescribeInstancesResponseV2OpenAPI**](OBCloudResultOBCloudPagingDataListDescribeInstancesResponseV2OpenAPI.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenant

> OBCloudResultTenantDTO DescribeTenant(ctx, instanceId, tenantId).RequestId(requestId).Execute()

查询特定租户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeTenant(context.Background(), instanceId, tenantId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTenant`: OBCloudResultTenantDTO
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 

### Return type

[**OBCloudResultTenantDTO**](OBCloudResultTenantDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenantAddress

> OBCloudResultTenantConnectionDTO DescribeTenantAddress(ctx, instanceId, tenantId).RequestId(requestId).AddressId(addressId).Execute()

查询租户地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	addressId := "addressId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeTenantAddress(context.Background(), instanceId, tenantId).RequestId(requestId).AddressId(addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeTenantAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTenantAddress`: OBCloudResultTenantConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeTenantAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **addressId** | **string** |  | 

### Return type

[**OBCloudResultTenantConnectionDTO**](OBCloudResultTenantConnectionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenantPrivateLink

> OBCloudResultListTenantConnectionDTO DescribeTenantPrivateLink(ctx, instanceId, tenantId).RequestId(requestId).AddressId(addressId).Execute()

查询租户的privatelink地址信息

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	addressId := "addressId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeTenantPrivateLink(context.Background(), instanceId, tenantId).RequestId(requestId).AddressId(addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeTenantPrivateLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTenantPrivateLink`: OBCloudResultListTenantConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeTenantPrivateLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantPrivateLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **addressId** | **string** |  | 

### Return type

[**OBCloudResultListTenantConnectionDTO**](OBCloudResultListTenantConnectionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenantPublicAddress

> OBCloudResultTenantConnectionDTO DescribeTenantPublicAddress(ctx, instanceId, tenantId).RequestId(requestId).Execute()

查询租户的公网主地址

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeTenantPublicAddress(context.Background(), instanceId, tenantId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeTenantPublicAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTenantPublicAddress`: OBCloudResultTenantConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeTenantPublicAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantPublicAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 

### Return type

[**OBCloudResultTenantConnectionDTO**](OBCloudResultTenantConnectionDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenantSecurityIpGroups

> OBCloudResultOBCloudPagingDataMcTenantSecurityIpGroupsResponse DescribeTenantSecurityIpGroups(ctx, instanceId, tenantId).RequestId(requestId).Execute()

DescribeTenantSecurityIpGroups

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeTenantSecurityIpGroups(context.Background(), instanceId, tenantId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeTenantSecurityIpGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTenantSecurityIpGroups`: OBCloudResultOBCloudPagingDataMcTenantSecurityIpGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeTenantSecurityIpGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantSecurityIpGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 

### Return type

[**OBCloudResultOBCloudPagingDataMcTenantSecurityIpGroupsResponse**](OBCloudResultOBCloudPagingDataMcTenantSecurityIpGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeTenantUsers

> OBCloudResultOBCloudPagingDataOcpDbUserDtoV2 DescribeTenantUsers(ctx, instanceId, tenantId).RequestId(requestId).UserName(userName).PageNumber(pageNumber).PageSize(pageSize).SearchKey(searchKey).Execute()

查询租户内的用户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	userName := "userName_example" // string |  (optional)
	pageNumber := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	searchKey := "searchKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.DescribeTenantUsers(context.Background(), instanceId, tenantId).RequestId(requestId).UserName(userName).PageNumber(pageNumber).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.DescribeTenantUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTenantUsers`: OBCloudResultOBCloudPagingDataOcpDbUserDtoV2
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.DescribeTenantUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTenantUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **userName** | **string** |  | 
 **pageNumber** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **searchKey** | **string** |  | 

### Return type

[**OBCloudResultOBCloudPagingDataOcpDbUserDtoV2**](OBCloudResultOBCloudPagingDataOcpDbUserDtoV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDatabaseDescription

> OBCloudResultModifyDatabaseDescriptionResponseV2 ModifyDatabaseDescription(ctx, instanceId, tenantId, databaseName).RequestId(requestId).Body(body).Execute()

ModifyDatabaseDescription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	databaseName := "databaseName_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewModifyDatabaseDescriptionRequestV2() // ModifyDatabaseDescriptionRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyDatabaseDescription(context.Background(), instanceId, tenantId, databaseName).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyDatabaseDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDatabaseDescription`: OBCloudResultModifyDatabaseDescriptionResponseV2
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyDatabaseDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 
**databaseName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDatabaseDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestId** | **string** |  | 
 **body** | [**ModifyDatabaseDescriptionRequestV2**](ModifyDatabaseDescriptionRequestV2.md) |  | 

### Return type

[**OBCloudResultModifyDatabaseDescriptionResponseV2**](OBCloudResultModifyDatabaseDescriptionResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDatabaseUserRole

> OBCloudResult ModifyDatabaseUserRole(ctx, instanceId, tenantId, databaseId).Body(body).Execute()

修改角色列表

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	databaseId := "databaseId_example" // string | 
	body := *openapiclient.NewModifyDatabaseUserRolesParamDo() // ModifyDatabaseUserRolesParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyDatabaseUserRole(context.Background(), instanceId, tenantId, databaseId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyDatabaseUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDatabaseUserRole`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyDatabaseUserRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 
**databaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDatabaseUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ModifyDatabaseUserRolesParamDo**](ModifyDatabaseUserRolesParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyInstanceName

> OBCloudResultModifyInstanceNameResponseV2 ModifyInstanceName(ctx, instanceId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()

ModifyInstanceName

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 
	xObProjectId := "xObProjectId_example" // string |  (optional)
	body := *openapiclient.NewModifyInstanceNameRequestV2() // ModifyInstanceNameRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyInstanceName(context.Background(), instanceId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyInstanceName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInstanceName`: OBCloudResultModifyInstanceNameResponseV2
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyInstanceName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInstanceNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 
 **xObProjectId** | **string** |  | 
 **body** | [**ModifyInstanceNameRequestV2**](ModifyInstanceNameRequestV2.md) |  | 

### Return type

[**OBCloudResultModifyInstanceNameResponseV2**](OBCloudResultModifyInstanceNameResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyInstanceNodeNum

> OBCloudResultModifyInstanceResponseV2OpenApi ModifyInstanceNodeNum(ctx, instanceId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()

ModifyInstanceNodeNum

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 
	xObProjectId := "xObProjectId_example" // string |  (optional)
	body := *openapiclient.NewMultiCloudModifyInstanceNodeNumRequest(int32(123)) // MultiCloudModifyInstanceNodeNumRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyInstanceNodeNum(context.Background(), instanceId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyInstanceNodeNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInstanceNodeNum`: OBCloudResultModifyInstanceResponseV2OpenApi
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyInstanceNodeNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInstanceNodeNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 
 **xObProjectId** | **string** |  | 
 **body** | [**MultiCloudModifyInstanceNodeNumRequest**](MultiCloudModifyInstanceNodeNumRequest.md) |  | 

### Return type

[**OBCloudResultModifyInstanceResponseV2OpenApi**](OBCloudResultModifyInstanceResponseV2OpenApi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyInstanceSpec

> OBCloudResultModifyInstanceResponseV2OpenApi ModifyInstanceSpec(ctx, instanceId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()

ModifyInstanceSpec

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 
	xObProjectId := "xObProjectId_example" // string |  (optional)
	body := *openapiclient.NewMultiCloudModifyInstanceSpecRequest() // MultiCloudModifyInstanceSpecRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyInstanceSpec(context.Background(), instanceId).RequestId(requestId).XObProjectId(xObProjectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyInstanceSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyInstanceSpec`: OBCloudResultModifyInstanceResponseV2OpenApi
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyInstanceSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyInstanceSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 
 **xObProjectId** | **string** |  | 
 **body** | [**MultiCloudModifyInstanceSpecRequest**](MultiCloudModifyInstanceSpecRequest.md) |  | 

### Return type

[**OBCloudResultModifyInstanceResponseV2OpenApi**](OBCloudResultModifyInstanceResponseV2OpenApi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantName

> OBCloudResultModifyTenantNameResponseV2 ModifyTenantName(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

修改租户名称

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewModifyTenantNameRequestV2() // ModifyTenantNameRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantName(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantName`: OBCloudResultModifyTenantNameResponseV2
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**ModifyTenantNameRequestV2**](ModifyTenantNameRequestV2.md) |  | 

### Return type

[**OBCloudResultModifyTenantNameResponseV2**](OBCloudResultModifyTenantNameResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantPrimaryZone

> OBCloudResultModifyTenantPrimaryZoneResponse ModifyTenantPrimaryZone(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

修改租户主可用区

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewModifyTenantPrimaryZoneRequest() // ModifyTenantPrimaryZoneRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantPrimaryZone(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantPrimaryZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantPrimaryZone`: OBCloudResultModifyTenantPrimaryZoneResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantPrimaryZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantPrimaryZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**ModifyTenantPrimaryZoneRequest**](ModifyTenantPrimaryZoneRequest.md) |  | 

### Return type

[**OBCloudResultModifyTenantPrimaryZoneResponse**](OBCloudResultModifyTenantPrimaryZoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantSecurityIpGroup

> OBCloudResultMcModifyTenantSecurityIpGroupResponse ModifyTenantSecurityIpGroup(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

ModifyTenantSecurityIpGroup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewMcModifyTenantSecurityIpGroupRequest() // McModifyTenantSecurityIpGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantSecurityIpGroup(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantSecurityIpGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantSecurityIpGroup`: OBCloudResultMcModifyTenantSecurityIpGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantSecurityIpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantSecurityIpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**McModifyTenantSecurityIpGroupRequest**](McModifyTenantSecurityIpGroupRequest.md) |  | 

### Return type

[**OBCloudResultMcModifyTenantSecurityIpGroupResponse**](OBCloudResultMcModifyTenantSecurityIpGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantUserDescription

> OBCloudResult ModifyTenantUserDescription(ctx, instanceId, tenantId, userName).RequestId(requestId).Body(body).Execute()

修改租户内的用户备注信息

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	userName := "userName_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewModifyTenantUserDescriptionParamDo() // ModifyTenantUserDescriptionParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantUserDescription(context.Background(), instanceId, tenantId, userName).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantUserDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantUserDescription`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantUserDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 
**userName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantUserDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestId** | **string** |  | 
 **body** | [**ModifyTenantUserDescriptionParamDo**](ModifyTenantUserDescriptionParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantUserPassword

> OBCloudResult ModifyTenantUserPassword(ctx, instanceId, tenantId, userName).RequestId(requestId).Body(body).Execute()

修改租户内的用户密码

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	userName := "userName_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewModifyUserPasswordParamDo() // ModifyUserPasswordParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantUserPassword(context.Background(), instanceId, tenantId, userName).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantUserPassword`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 
**userName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestId** | **string** |  | 
 **body** | [**ModifyUserPasswordParamDo**](ModifyUserPasswordParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantUserRole

> OBCloudResult ModifyTenantUserRole(ctx, instanceId, tenantId, userName).Body(body).Execute()

修改用户权限

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	userName := "userName_example" // string | 
	body := *openapiclient.NewModifyUserRolesParamDo() // ModifyUserRolesParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantUserRole(context.Background(), instanceId, tenantId, userName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantUserRole`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantUserRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 
**userName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ModifyUserRolesParamDo**](ModifyUserRolesParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantUserStatus

> OBCloudResult ModifyTenantUserStatus(ctx, instanceId, tenantId, userName).RequestId(requestId).Body(body).Execute()

修改租户内的用户锁定状态

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	userName := "userName_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewModifyUserStatusParamDo() // ModifyUserStatusParamDo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.ModifyTenantUserStatus(context.Background(), instanceId, tenantId, userName).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.ModifyTenantUserStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantUserStatus`: OBCloudResult
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.ModifyTenantUserStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 
**userName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantUserStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestId** | **string** |  | 
 **body** | [**ModifyUserStatusParamDo**](ModifyUserStatusParamDo.md) |  | 

### Return type

[**OBCloudResult**](OBCloudResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTenants

> OBCloudResultOBCloudPagingDataTenantDTO QueryTenants(ctx, instanceId).RequestId(requestId).PageNumber(pageNumber).PageSize(pageSize).TenantName(tenantName).TagList(tagList).Execute()

条件查询租户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 
	pageNumber := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	tenantName := "tenantName_example" // string |  (optional)
	tagList := "tagList_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.QueryTenants(context.Background(), instanceId).RequestId(requestId).PageNumber(pageNumber).PageSize(pageSize).TenantName(tenantName).TagList(tagList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.QueryTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryTenants`: OBCloudResultOBCloudPagingDataTenantDTO
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.QueryTenants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 
 **pageNumber** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **tenantName** | **string** |  | 
 **tagList** | **string** |  | 

### Return type

[**OBCloudResultOBCloudPagingDataTenantDTO**](OBCloudResultOBCloudPagingDataTenantDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCluster

> OBCloudResultString StartCluster(ctx, instanceId).RequestId(requestId).Execute()

StartCluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.StartCluster(context.Background(), instanceId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.StartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCluster`: OBCloudResultString
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.StartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 

### Return type

[**OBCloudResultString**](OBCloudResultString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopCluster

> OBCloudResultString StopCluster(ctx, instanceId).RequestId(requestId).Execute()

StopCluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	requestId := "requestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.StopCluster(context.Background(), instanceId).RequestId(requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.StopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopCluster`: OBCloudResultString
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.StopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestId** | **string** |  | 

### Return type

[**OBCloudResultString**](OBCloudResultString.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivatelinkInformation

> OBCloudResultUpdateTenantPrivatelinkInformationResponse UpdatePrivatelinkInformation(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

更新租户private link信息

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewUpdateTenantPrivatelinkInformationRequest() // UpdateTenantPrivatelinkInformationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.UpdatePrivatelinkInformation(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.UpdatePrivatelinkInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrivatelinkInformation`: OBCloudResultUpdateTenantPrivatelinkInformationResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.UpdatePrivatelinkInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivatelinkInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**UpdateTenantPrivatelinkInformationRequest**](UpdateTenantPrivatelinkInformationRequest.md) |  | 

### Return type

[**OBCloudResultUpdateTenantPrivatelinkInformationResponse**](OBCloudResultUpdateTenantPrivatelinkInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> OBCloudResultSimpleTenantResponse UpdateTenant(ctx, instanceId, tenantId).RequestId(requestId).Body(body).Execute()

修改租户

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	instanceId := "instanceId_example" // string | 
	tenantId := "tenantId_example" // string | 
	requestId := "requestId_example" // string | 
	body := *openapiclient.NewUpdateTenantRequest() // UpdateTenantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiCloudOpenAPI.UpdateTenant(context.Background(), instanceId, tenantId).RequestId(requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiCloudOpenAPI.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenant`: OBCloudResultSimpleTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiCloudOpenAPI.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **string** |  | 
 **body** | [**UpdateTenantRequest**](UpdateTenantRequest.md) |  | 

### Return type

[**OBCloudResultSimpleTenantResponse**](OBCloudResultSimpleTenantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

