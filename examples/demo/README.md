# OceanBase Terraform Provider 能力与演示手册

## 适用范围

本手册对应本地分支 `codex/terraform-cn-prod-validation-20261118`，功能基线为 `e534f58`。截至 2026-11-19，GitHub `main` 为 `85f5ec6`；本地分支中的修复和本演示目录尚未推送，不能把它们当作 `main` 已发布的能力。演示使用本地构建的 Provider，而不是从 Registry 下载的已发布版本。

Provider 通过 OceanBase Cloud OpenAPI 管理资源。以下生产验证针对国内 `test` 项目、阿里云杭州、按量集群；其他项目、规格和地域是否可用，以对应环境返回为准。本手册暂不覆盖 PrivateLink。

## 能力清单

| 对象 | Terraform 类型 | 当前分支的主要操作 | 验证范围 |
| --- | --- | --- | --- |
| 实例 | `oceanbase_cluster` | 创建、查询、修改名称/规格/磁盘/节点数、启停、删除、导入 | 按量创建、名称与磁盘组合修改、查询和删除已在国内生产验证；节点数变化和启停未做本轮生产实测 |
| 租户 | `oceanbase_tenant` | 创建、查询、修改名称/主区/Unit 规格及数量、删除、导入 | MySQL 租户创建、名称与主区组合修改、查询和删除已在国内生产验证；Unit 变更未做本轮生产实测 |
| 租户 IP 白名单组 | `oceanbase_tenant_security_group` | 创建、查询、修改 IP 列表、删除、导入 | 创建、更新、外部删除后的重建及删除已在国内生产验证 |
| 租户公网地址 | `oceanbase_tenant_connection` | 创建、查询、删除、导入 `INTERNET` 地址 | 创建、回读、删除及通过临时数据库账号执行 `SELECT 1` 已在国内生产验证 |
| 数据库 | `oceanbase_database` | 创建、查询、修改描述、删除、导入 | 已实现；本轮未对该 Terraform 资源做生产生命周期验证 |
| 数据库账号 | `oceanbase_tenant_user` | 创建、查询、修改描述/密码/权限/锁定状态、删除、导入 | 已实现；本轮未对该 Terraform 资源做生产生命周期验证 |

以上六类均有同名 `data` 数据源，可只读查询已有对象。`resource` 交由 Terraform 管理，`data` 只读取，不接管已有对象。

重要边界：

- 演示只使用 `PostPay`。包年包月实例创建不在当前支持范围。
- `enable_public_address` 和租户创建时的 `whitelist` 属于创建参数；后续公网地址和 IP 组分别通过独立资源管理。
- 公网地址资源当前只支持 `type = "INTERNET"`；创建时不要配置 `enable_ssl`、`enable_rpc` 或 `private_link`，已有连接也不支持通过该资源更新这些选项。
- `node_num` 必须是 3～60 之间的 3 的倍数；主备、存储分离和部分 K8s 实例类型不支持该配置。创建接口没有节点数字段，目标节点数与创建后的实际值不同时，Provider 在实例上线后再发起节点数调整。
- 实例的地域、可用区、版本、付费方式等不能通过当前资源原地修改；租户兼容模式、字符集等也不是本演示的原地修改项。账号资源一次只支持修改一种属性。
- 数据库和账号资源的删除会删除对应对象；数据库删除请求使用强制删除。执行销毁前务必检查目标和数据。

## 演示目录

```text
examples/demo/
├── setup.sh                 # 构建当前分支的 Provider，生成本目录专用配置
├── create/                  # 创建与管理演示资源，使用独立 Terraform state
│   ├── cluster.tf           # 按量实例
│   ├── tenant.tf            # MySQL 租户
│   ├── network.tf           # 可选的 IP 组和公网地址
│   ├── database.tf          # 可选的数据库和普通账号
│   ├── variables.tf
│   ├── outputs.tf
│   └── terraform.tfvars.example
└── read/                    # 只读查询已有资源，使用另一份独立 state
    ├── main.tf              # 六类 data source
    ├── variables.tf
    ├── outputs.tf
    └── terraform.tfvars.example
```

默认创建杭州双可用区、4C16G、50 GB 的按量实例及一个 1C4G MySQL 租户；公网、数据库和账号默认关闭，可在演示时逐项打开。当前工作区的 `create/terraform.tfvars` 已填入本次 `test` 项目；`read/terraform.tfvars` 已填入本次已创建的实例和租户 ID。这两份本机文件被 Git 忽略，不包含密钥；换环境时根据各自的 `.example` 文件填写。

## 1. 准备 Provider 和凭证

需要 Go 1.22 或更高版本，以及 Terraform；本次演示目录使用 Terraform 1.15.8 验证。从仓库根目录执行：

```bash
cd examples/demo
./setup.sh
export TF_CLI_CONFIG_FILE="$PWD/.terraformrc"
```

`setup.sh` 只在本目录构建 Provider 和生成 `.terraformrc`，不会修改用户主目录的 Terraform 配置。每次切换代码分支或重新修改 Provider 后，应再次运行它。Terraform 提示 `Provider development overrides are in effect` 是使用本地构建的提示。

在同一个终端输入 OceanBase Cloud OpenAPI 凭证：

```bash
printf 'OceanBase AccessKey ID: '
read -r TF_VAR_oceanbase_access_key
printf 'OceanBase AccessKey Secret: '
read -rs TF_VAR_oceanbase_secret_key
printf '\n'
export TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key
```

凭证仅在当前 shell 会话使用，不要写进 `.tf` 或提交到 Git。若启用数据库账号，Terraform state 可能保存账号密码；`sensitive` 只隐藏常规终端输出，应保护 state 文件。

## 2. 先演示只读查询

`read` 目录和 `create` 目录的 state 相互独立。当前工作区已填入现有测试实例和租户，可直接执行：

```bash
cd read
[ -f terraform.tfvars ] || cp terraform.tfvars.example terraform.tfvars
```

首次复制示例文件时，先将 `cluster_id`、`tenant_id` 替换为实际 ID，再执行：

```bash
terraform init
terraform validate
terraform plan
```

计划会读取实例和租户并显示状态；出现 `Changes to Outputs` 表示尚未把查询结果写入本地 state，不表示要修改云资源。如需之后使用 `terraform output`，可在 `read` 目录执行 `terraform apply`，它只保存数据源输出到本地 state。

要查询其他对象，在 `read/terraform.tfvars` 中填写已存在对象对应的 `database_name`、`database_user`、`security_group_name` 或 `internet_address_id`；留空的对象不会查询。对象必须存在并在当前凭证可访问范围内。

## 3. 演示创建实例和租户

回到演示目录的 `create` 子目录，确认 `terraform.tfvars` 中的项目、实例名和租户名：

```bash
cd ../create
[ -f terraform.tfvars ] || cp terraform.tfvars.example terraform.tfvars
```

首次复制示例文件时，先将 `project_id` 替换为实际项目 ID，然后查看计划：

```bash
terraform init
terraform validate
terraform plan
```

确认计划中的项目、地域、规格及资源数量后执行：

```bash
terraform apply
```

`apply` 会创建新的按量实例并产生费用；实例通常需要数分钟开通，Terraform 等待实例和租户就绪。完成后：

```bash
terraform output
terraform plan
```

再次 `plan` 可检查是否仍有待执行变更。`create` 有自己的 state，不会接管 `read` 中的现有实例；因此执行本节会新建实例。

## 4. 可选：数据库和账号

将 `create/terraform.tfvars` 的 `enable_database_user` 设为 `true`，再在当前终端输入一个符合目标环境规则的数据库账号密码：

```bash
printf 'Demo database user password: '
read -rs TF_VAR_database_user_password
printf '\n'
export TF_VAR_database_user_password
terraform plan
terraform apply
```

配置会创建 `tf_demo_db` 和普通账号 `tf_demo_user`，并授予该库 `readwrite` 权限。此组 Terraform 资源尚未做本轮生产生命周期验证，演示前请确认目标项目允许创建。

## 5. 可选：公网地址和 IP 白名单组

在 `create/terraform.tfvars` 中设置：

```hcl
enable_internet = true
allowed_ips     = ["<演示客户端的实际公网 IP>"]
```

替换占位值后执行 `terraform plan` 和 `terraform apply`。Terraform 先创建 IP 白名单组，再创建公网地址；创建和释放地址均可能需要数分钟。完成后可查看：

```bash
terraform output internet_address
terraform output internet_port
```

实际连接还需要数据库账号、密码及客户端网络连通。生产验证中单个 IP 可直接填写，例如 `192.0.2.1`；该地址仅为文档示例，不能替代你的实际出口 IP。

## 6. 变更、导入与清理

- **变更**：修改配置后先 `terraform plan`，确认差异再 `terraform apply`。已验证实例名与磁盘、租户名与主区可以在同一次 apply 中变更；账号资源应一次只改一种属性。
- **导入**：先写好与现有资源匹配的配置，再使用 `terraform import`，之后执行 `terraform plan` 核对。实例的导入 ID 是实例 ID；租户为 `<实例 ID>.<租户 ID>`；IP 组、数据库、账号为 `<实例 ID>.<租户 ID>.<名称>`；公网地址为 `<实例 ID>.<租户 ID>.INTERNET.<地址 ID>`。导入格式来自源码，本轮未逐项做生产导入验证。
- **清理**：在 `create` 目录执行 `terraform plan -destroy` 审核，再由操作者决定是否执行 `terraform destroy`。它会删除该目录 state 管理的实例、租户和已启用的子资源；`read` 目录不会管理或删除云资源。
- **超时或报错**：先根据返回的资源 ID 在控制台/OpenAPI 核对实际状态，再执行 `terraform plan`；不要直接重复创建。实例外部删除后，如果 API 返回权限错误而不能证明不存在，Provider 会保留 state 并报错。

演示结束可在终端执行 `unset TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key TF_VAR_database_user_password`。之前单独创建的测试实例不在本目录的 `create` state 中，需单独决定是否保留或清理。

## 本次目录验证

- `create` 和 `read` 均通过 `terraform validate`；默认创建计划为 2 个资源，打开可选项后为 6 个资源。
- 使用 Terraform 模拟 Provider 验证了基础资源、可选资源依赖和六类数据源配置；模拟测试不等同于云端创建。
- `read` 目录已用真实 OpenAPI 只读查询当前测试实例和租户，两者均返回 `ONLINE`。本次没有通过演示目录创建或删除云资源。
