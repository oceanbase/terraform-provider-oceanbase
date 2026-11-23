# Terraform Provider for OceanBase

Manage OceanBase Cloud resources with Terraform and the OceanBase Cloud OpenAPI.
This repository provides a provider for instances, tenants, databases, users,
network access, and PrivateLink connections.

使用 Terraform 和 OceanBase Cloud OpenAPI 管理实例、租户、数据库、账号及网络资源。

## Requirements / 前置条件

- Terraform 1.5 or later. The published binaries support Linux AMD64 (`x86_64`)
  and ARM64 (`aarch64`). Other platforms can build the provider from source.
- An OceanBase Cloud OpenAPI AccessKey and permission to manage the target
  project. An Alibaba Cloud account is also required for the Alibaba Cloud
  PrivateLink example.
- Confirm the target region, zones, instance specifications, and available
  features in your OceanBase Cloud environment before creating resources.

需要 Terraform 1.5 或以上版本、对应环境的 OceanBase Cloud OpenAPI 凭证及项目权限。
创建资源前请确认目标地域、可用区、规格和功能可用性。

## Install the released provider / 安装已发布的 Provider

[Release 1.1.0](https://github.com/oceanbase/terraform-provider-oceanbase/releases/tag/1.1.0)
contains Linux AMD64 and ARM64 binaries. The provider is distributed through
GitHub Releases; it is not currently available from the Terraform Registry.
The commands below install it into a local filesystem mirror so that
`terraform init` can use the released binary.

以下命令在 Linux 的 Bash 终端中执行。根据 `uname -m` 选择架构：`x86_64` 使用
`amd64`，`aarch64` 使用 `arm64`。

```bash
git clone https://github.com/oceanbase/terraform-provider-oceanbase.git
cd terraform-provider-oceanbase
VERSION=1.1.0
ARCH=amd64 # Set to arm64 for aarch64.
MIRROR="$HOME/.terraform.d/oceanbase-providers"
PROVIDER_DIR="$MIRROR/oceanbase.github.io/oceanbase/oceanbase/$VERSION/linux_$ARCH"
mkdir -p "$PROVIDER_DIR"
curl -fL \
  "https://github.com/oceanbase/terraform-provider-oceanbase/releases/download/$VERSION/terraform-provider-oceanbase-$ARCH" \
  -o "$PROVIDER_DIR/terraform-provider-oceanbase_v$VERSION"
chmod 755 "$PROVIDER_DIR/terraform-provider-oceanbase_v$VERSION"
sha256sum "$PROVIDER_DIR/terraform-provider-oceanbase_v$VERSION"
```

For 1.1.0, the SHA-256 values are:

| Architecture | SHA-256 |
| --- | --- |
| AMD64 | `88ef9ce68d88bdc4d090401a44e7e80bbd728e0fe0ba6350eb4a744d55eff985` |
| ARM64 | `98fc9858986b5105fe20c7520cc364c7ef9256375afa31320ad95c030cce6a74` |

On another platform, replace the download commands with a local build using
Go 1.22 or later. Set `VERSION` and `MIRROR` as above, then run:

```bash
PLATFORM="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
PROVIDER_DIR="$MIRROR/oceanbase.github.io/oceanbase/oceanbase/$VERSION/$PLATFORM"
mkdir -p "$PROVIDER_DIR"
go build -trimpath -ldflags "-X main.version=$VERSION" \
  -o "$PROVIDER_DIR/terraform-provider-oceanbase_v$VERSION" .
```

Create a dedicated Terraform CLI configuration and use it in the current
terminal. If you already use a custom CLI configuration, merge the
`provider_installation` block into that configuration instead of replacing it.

创建独立的 Terraform CLI 配置，并在当前终端启用。已有自定义配置时，请合并配置块，
避免覆盖原有设置。

```bash
cat > "$HOME/.terraform.d/oceanbase-release.tfrc" <<EOF
provider_installation {
  filesystem_mirror {
    path    = "$MIRROR"
    include = ["oceanbase.github.io/oceanbase/oceanbase"]
  }
  direct {
    exclude = ["oceanbase.github.io/oceanbase/oceanbase"]
  }
}
EOF
export TF_CLI_CONFIG_FILE="$HOME/.terraform.d/oceanbase-release.tfrc"
```

The setting applies only to this terminal and its child processes. Repeat the
`export` in a new terminal.

Terraform may show `unauthenticated` and an incomplete lock-file warning for
the local mirror. If the same `.terraform.lock.hcl` is shared across Linux
architectures, install both binaries into the mirror and run the following
from the example directory to record both checksums:

```bash
MIRROR="$HOME/.terraform.d/oceanbase-providers"
terraform providers lock -fs-mirror="$MIRROR" \
  -platform=linux_amd64 -platform=linux_arm64 \
  oceanbase.github.io/oceanbase/oceanbase
```

## Authentication / 配置凭证

Use the OceanBase Cloud OpenAPI AccessKey, not the Alibaba Cloud AccessKey.
The examples take credentials through sensitive Terraform variables. In Bash:

使用 OceanBase Cloud 的 OpenAPI 凭证；不要将真实 AK/SK 写入示例文件或提交到 Git。

```bash
read -r -p 'OceanBase AccessKey ID: ' TF_VAR_oceanbase_access_key
read -r -s -p 'OceanBase AccessKey Secret: ' TF_VAR_oceanbase_secret_key
printf '\n'
export TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key
```

The Secret is not displayed while typing. Set `site` to `CHINA` for the China
site or `GLOBAL` for the global site; the PrivateLink example calls this input
`oceanbase_site`. Credentials, site, project, and resource IDs must belong to
the same target environment. The optional
`TF_OCEANBASE_API_HOST_CHINA` and `TF_OCEANBASE_API_HOST_GLOBAL` environment
variables override the default API hosts; check any existing overrides before
running an example.

Terraform state and saved plans can contain sensitive data, including database
user passwords. Restrict access to them and do not commit them to Git.

## Examples / 示例

| Directory | Purpose |
| --- | --- |
| [`examples/demo/read`](examples/demo/read) | Read an existing instance and tenant, with optional related resources. |
| [`examples/demo/create`](examples/demo/create) | Create a pay-as-you-go instance and tenant, with optional database and network resources. |
| [`examples/private_link`](examples/private_link) | Create an OceanBase PrivateLink service and an Alibaba Cloud endpoint for an existing tenant. |
| [`examples/cluster`](examples/cluster), [`cluster_data_source`](examples/cluster_data_source), [`cluster_import`](examples/cluster_import) | Minimal instance creation, lookup, and import examples; creation and import use Alibaba Cloud. |
| [`examples/provider`](examples/provider) | Minimal provider configuration. |

### Read an existing instance and tenant / 查询已有资源

From the repository root:

```bash
cd examples/demo/read
test -f terraform.tfvars || cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars and set site, cluster_id, and tenant_id.
terraform init
terraform validate
terraform plan
```

`plan` reads the selected resources and shows the output values; this example
does not create cloud resources. Fill optional names or address IDs in the
template only when those resources already exist.

在 `terraform.tfvars` 中填写已有实例 ID 和租户 ID；不需要查询的可选项保持为空。

For the minimal instance lookup or creation examples, use the
`terraform.tfvars.example` in the corresponding directory in the same way.
To adopt an existing instance, first match the configuration in
`examples/cluster_import` to the instance, then run
`terraform import oceanbase_cluster.example <instance-id>` and review the plan.

### Create an instance and tenant / 创建实例和租户

```bash
cd examples/demo/create
test -f terraform.tfvars || cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars using the project and available specifications.
terraform init
terraform validate
terraform plan
terraform apply
terraform output
```

Replace every `REPLACE_...` value in the template and check all other settings
against the target environment. Review the plan before accepting `apply`.
Creating resources can incur charges. Public access, database users, and
PrivateLink are disabled by default. When enabling a database user, supply its
password through `TF_VAR_database_user_password` rather than a `.tfvars` file.

将模板中的 `REPLACE_...` 占位符替换为目标项目的可用配置。启用数据库账号时，可在
当前终端安全输入密码：

```bash
read -r -s -p 'Database user password: ' TF_VAR_database_user_password
printf '\n'
export TF_VAR_database_user_password
```

执行 `apply` 会创建资源并可能产生费用；请先核对计划。`terraform destroy` 会删除
该目录的 state 管理的资源，包括实例和租户。

### Connect an existing tenant with PrivateLink / 为已有租户接入 PrivateLink

The Alibaba Cloud example manages the OceanBase service, Alibaba Cloud
endpoint and zone, and OceanBase connection. Prepare an existing instance and
tenant, an eligible cloud account, and an existing VPC, vSwitch, and security
group in a compatible region and zone. Configure Alibaba Cloud credentials
using a [supported authentication method](https://registry.terraform.io/providers/aliyun/alicloud/latest/docs#authentication).

此示例会创建 OceanBase 私网服务和阿里云终端节点。请先准备已有 VPC、交换机和安全组，
并确认功能开通、地域、可用区及费用。

```bash
cd examples/private_link
test -f terraform.tfvars || cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars with the existing resource IDs and cloud account ID.
terraform init
terraform validate
terraform plan
terraform apply
terraform output
```

If the target environment requires a new commercialized ODP, fill
`proxy_setting` with supported values as shown in the template. The security
group must allow the intended client traffic to the database port. Configure
the OceanBase IP allowlist as required, and test the resulting domain from a
client that can reach the endpoint VPC. PrivateLink database connectivity has
not yet been verified end to end in this repository's test environment.

如果需要新建商业化 ODP，请按模板填写 `proxy_setting`。安全组需放通客户端到数据库
端口的访问，并按 OceanBase 控制台要求配置白名单；连接测试应从可访问该 VPC 的客户端进行。

For existing services or connections, import them rather than creating
duplicates; see the [service](docs/resources/tenant_private_link_service.md)
and [connection](docs/resources/tenant_private_link_connection.md) resource
documentation. Deleting the last connection may also remove its service.

## Resource behavior / 资源行为

| Resource | Configuration |
| --- | --- |
| `oceanbase_cluster` | Set `payment_type`, `disk_type`, and `storage_architecture` when creating the instance. `node_num`, when set, must be a multiple of 3 between 3 and 60; scaling starts after creation reaches the online state. |
| `oceanbase_tenant` | `enable_public_address` and `whitelist` are creation settings. Use connection and security group resources for individual addresses and IP groups afterwards. |
| `oceanbase_tenant_connection` | Create a public address with `type = "INTERNET"`. Omit `private_link`, `enable_ssl`, and `enable_rpc`; SSL and RPC settings are read from the service. |
| `oceanbase_tenant_private_link_service` | Creates the tenant's primary PrivateLink service and exposes its cloud-facing service name. |
| `oceanbase_tenant_private_link_connection` | Connects an endpoint created by a cloud provider to the tenant's service. |
| `oceanbase_tenant_security_group` | A single IPv4 address may be written as `192.0.2.1` or `192.0.2.1/32`. |

Changes to an instance's name and size, or to a tenant's name, unit
configuration, and primary zone, are applied sequentially. If a cloud operation
times out, inspect the recorded resource ID and run `terraform plan` before
retrying. When finished, clear credentials from the current terminal with:

```bash
unset TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key TF_VAR_database_user_password
```
