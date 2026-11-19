# OceanBase Terraform Provider：配置与演示手册

## 写给第一次使用 Terraform 的你

你不需要先手写一整套配置。这个目录已经把文件准备好，你主要做四件事：**选环境 → 填参数 → 看计划 → 决定是否执行**。后面的“文件说明”和“参数说明”是字典，不需要第一次就记住。

```text
你填写参数文件  →  terraform plan 看清单  →  terraform apply 执行
                                            ↓
                                   OceanBase OpenAPI 创建云资源
```

先认识几个词：

| 名称 | 用大白话解释 |
| --- | --- |
| Terraform | 你电脑上的命令行工具。它读取配置，计算需要创建或修改什么，然后调用云服务执行。 |
| Provider | Terraform 的插件。这里的 OceanBase Provider 负责把 Terraform 操作转换成 OceanBase Cloud OpenAPI 调用。 |
| Resource / 资源 | 你希望 Terraform 创建和管理的东西，例如实例、租户。配置在 `create/` 目录。 |
| Data Source / 数据源 | 只查询一个已经存在的资源，不创建、不接管它。配置在 `read/` 目录。 |
| `.tf` 文件 | 描述资源和变量的配置文件；演示目录已经写好，通常不用逐个修改。 |
| `.tfvars` 文件 | 你填写具体值的文件，例如项目 ID、地域、实例名。切环境时主要修改它。 |
| State / 状态 | Terraform 保存“这份配置管理了哪些实际资源”的记录。不能把生产和预发共用同一份 state。 |
| Workspace / 工作区 | 同一个 Terraform 目录里隔离不同 state 的方式。本手册用它区分环境；它不会自动切换地址、凭证或变量文件。 |
| `plan` / 计划 | 先展示 Terraform 打算做什么，不创建云资源。 |
| `apply` / 执行 | 根据配置实际执行操作；创建实例会产生费用。 |

## 先看这里

这套目录用于演示 OceanBase Cloud OpenAPI 对应的 Terraform Provider。**默认值是国内生产的阿里云杭州配置**：当前本机的 `create/terraform.tfvars` 还预填了本次生产测试项目，`read/terraform.tfvars` 预填了本次生产测试实例。演示其他环境前，必须切换 OpenAPI 地址、凭证、项目 ID、实例/租户 ID 以及对应的地域和规格。

本手册针对本地分支 `codex/terraform-cn-prod-validation-20261118`。截至 2026-11-19，GitHub `main` 为 `85f5ec6`，该分支的修复和演示目录尚未推送；操作时通过 `setup.sh` 构建本地 Provider。本手册不演示包年包月实例创建。

默认创建计划只有一个按量实例和一个 MySQL 租户。公网、IP 白名单组、数据库和账号由开关控制，默认都不创建。实际 `terraform apply` 会创建云资源并可能产生费用；本手册中的环境示例不是对其他环境规格可售或接口可用的保证。

## 新手快速上手：先只读，再按需创建

下面以 macOS 终端和**国内 PRE** 为例。这里的 `cn-pre` 只是文件/工作区名字；如果使用别的环境，把域名、`site`、凭证、项目和资源参数一起换成目标环境的值。国内生产、CC PRE 和海外地址见“二、环境由哪几项共同决定”。

### 第 1 步：准备工具和目标环境的信息

1. 打开 macOS 的“终端”，分别执行：

   ```bash
   terraform version
   go version
   ```

   能看到版本号即可继续；本示例使用 Terraform 1.15.8 和 Go 1.22 及以上版本。如果提示 `command not found`，先按 [Terraform 安装说明](https://developer.hashicorp.com/terraform/install) 或 [Go 安装说明](https://go.dev/doc/install) 安装。macOS 使用 Homebrew 时，Terraform 官方命令为 `brew install hashicorp/tap/terraform`。
2. 确认你有目标 OceanBase 环境的 OpenAPI AccessKey ID/Secret，以及项目权限。国内生产的入口可参考 [AccessKey 管理](https://www.oceanbase.com/docs/common-oceanbase-cloud-0000000002295194)：控制台右上角“账号设置”→“AccessKey”。其他环境使用对应环境账号和入口。**这是 OceanBase 的密钥，不是阿里云的密钥。**
3. 准备目标项目 ID。如果只是查询，还需要已有实例 ID 和租户 ID；如果要新建，还需要该环境支持的地域、可用区、规格、版本和渠道。项目 ID 可在目标环境控制台的“项目设置 → 项目信息”查看，详见 [项目管理](https://www.oceanbase.com/docs/common-oceanbase-cloud-1000000000703404)。

### 第 2 步：进入演示目录并构建 Provider

这台机器上的演示目录是下面这个位置；如果在其他机器使用，替换成实际仓库路径。先不要关闭终端：

```bash
cd /Users/sam/Documents/projects/EX/terraform-provider-oceanbase/examples/demo
./setup.sh
export TF_CLI_CONFIG_FILE="$PWD/.terraformrc"
```

看到 `Provider built` 表示本地插件构建完成。最后一行的环境变量只告诉 Terraform 从哪里加载插件，**不会**选择 PRE 或生产环境。新开一个终端后需要重新设置；代码分支改变后需要重新运行 `./setup.sh`。

### 第 3 步：选择 API 环境并输入密钥

仍在刚才的终端。下面设置国内 PRE；若要使用国内生产，执行第一行后不要执行第二行：

```bash
unset TF_OCEANBASE_API_HOST_CHINA TF_OCEANBASE_API_HOST_GLOBAL
export TF_OCEANBASE_API_HOST_CHINA="pre-api-cloud-cn.oceanbase.com"
```

然后输入**国内 PRE 的 OceanBase 密钥**：

```bash
printf 'OceanBase AccessKey ID: '
read -r TF_VAR_oceanbase_access_key
printf 'OceanBase AccessKey Secret: '
read -rs TF_VAR_oceanbase_secret_key
printf '\n'
export TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key
```

输入 Secret 时屏幕没有字符是正常的，输入完成按回车。密钥不需要放进文件或发给其他人。切换环境或新开终端时，重新设置对应的地址和凭证。

### 第 4 步：先查询已有实例（不创建资源）

1. 进入只读目录，准备参数文件：

   ```bash
   cd read
   [ -f cn-pre-read.tfvars ] || cp terraform.tfvars.example cn-pre-read.tfvars
   nano cn-pre-read.tfvars
   ```

   如果该文件之前已存在，第二条命令会保留它；请检查内容确实属于本次环境。
2. 在编辑器里，把最上面三行改成实际值；保留英文双引号。示例中的文字需要替换：

   ```hcl
   site       = "CHINA"
   cluster_id = "<PRE 环境中已有的实例 ID>"
   tenant_id  = "<这个实例下已有的租户 ID>"
   ```

   其他 `database_name`、`database_user`、`security_group_name`、`internet_address_id` 可以先留空。`nano` 中按 **Control+O** 保存、回车确认，再按 **Control+X** 退出。
3. 执行以下命令：

   ```bash
   terraform init
   terraform workspace new cn-pre-read
   terraform validate
   terraform plan -var-file=cn-pre-read.tfvars
   ```

   如果提示 workspace 已存在，改用 `terraform workspace select cn-pre-read`，再执行后面的两条。`validate` 成功会显示 `Success!`；`plan` 会查询现有实例和租户。看到 `Changes to Outputs` 是准备显示查询结果，**不是创建实例**。只看结果可以停在这里。

如果目标环境还没有实例或租户，可以跳过第 4 步，直接从演示目录进入下一步的 `create` 目录。

### 第 5 步：需要创建时，配置新实例和租户

这一步会创建新的按量实例，不会复用第 4 步查询的实例。

1. 如果刚完成第 4 步，当前位于 `read` 目录，先执行 `cd ../create`。如果跳过了第 4 步、当前仍位于 `examples/demo`，执行 `cd create`。接着准备另一个参数文件：

   ```bash
   [ -f cn-pre.tfvars ] || cp terraform.tfvars.example cn-pre.tfvars
   nano cn-pre.tfvars
   ```

   如果 `cn-pre.tfvars` 已经存在，复制命令不会覆盖它；请确认它是本次演示想使用的文件。

2. 修改文件中的值。下面是一组**结构示例**，其中项目 ID 必须替换；地域、可用区、规格、版本和渠道也要改为目标环境实际支持的值：

   ```hcl
   site                 = "CHINA"
   project_id           = "<PRE 项目 ID>"
   instance_name        = "tf_demo_pre"
   cloud_provider       = "ALIYUN"
   region               = "cn-hangzhou"
   zones                = ["cn-hangzhou-j", "cn-hangzhou-k"]
   primary_zone         = "cn-hangzhou-j"
   instance_class       = "4C16G"
   ob_version           = "4.4.2"
   disk_size            = 50
   replica_mode         = "2"
   instance_type        = "CLUSTER"
   sale_channel         = "domestic"
   tenant_name          = "tf_demo_tenant"
   enable_internet      = false
   enable_database_user = false
   ```

   逐行修改模板中已有的同名参数，不要把示例整段追加到文件末尾，否则会出现重复定义。上面只列出首次需要核对的部分；模板中其他行保留或按“参数说明”修改。`primary_zone` 要与 `zones` 对应；`site` 只写 `CHINA` 或 `GLOBAL`，不写 `PRE`。修改完按 **Control+O**、回车、**Control+X** 保存退出。
3. 初始化并查看计划：

   ```bash
   terraform init
   terraform workspace new cn-pre-demo
   terraform validate
   terraform plan -var-file=cn-pre.tfvars
   ```

   工作区已存在时使用 `terraform workspace select cn-pre-demo`。默认不开可选项时，预期计划包含 `oceanbase_cluster.demo` 和 `oceanbase_tenant.demo`，并显示 `Plan: 2 to add, 0 to change, 0 to destroy`。如果看到其他数量、要删除资源或目标参数不对，先停下核对。
4. 确认计划和费用后，才执行：

   ```bash
   terraform apply -var-file=cn-pre.tfvars
   ```

   Terraform 会再次展示计划并要求输入 `yes`；输入后开始创建，实例和租户上线可能需要数分钟。成功后：

   ```bash
   terraform output
   terraform plan -var-file=cn-pre.tfvars
   ```

   `output` 展示新实例、租户 ID 等；再次 `plan` 用于确认没有意外的待执行变更。

### 第 6 步：需要演示更多资源时再开启

- **数据库和账号**：在 `cn-pre.tfvars` 中设置 `enable_database_user = true`，按后文修改数据库名、账号名和角色。在终端运行 `printf 'Database password: '; read -rs TF_VAR_database_user_password; printf '\n'; export TF_VAR_database_user_password`，再用相同的 `-var-file` 执行 `plan` 和 `apply`。这两个 Terraform 资源本轮未做生产生命周期验证。
- **公网和 IP 白名单组**：设置 `enable_internet = true`，并把 `allowed_ips` 改成演示客户端的真实出口 IP 列表，例如 `allowed_ips = ["你的公网 IP"]`。先检查计划再执行；创建和释放公网地址可能需要几分钟。
- **PrivateLink 私网连接**：OceanBase 侧先创建服务，云厂商侧创建 Endpoint 后再关联，具体步骤见下方“第 7 步”。
- **清理本次新建资源**：先确认仍在 `create` 目录、目标环境和正确 workspace，执行 `terraform plan -destroy -var-file=cn-pre.tfvars`。确认后才由操作者决定是否执行 `terraform destroy -var-file=cn-pre.tfvars`。这会删除该 workspace 中的实例、租户和受管子资源，数据库数据也会被删除；不会删除 `read` 目录只是查询的资源。

### 第 7 步：按两阶段演示 PrivateLink

PrivateLink 把云厂商 VPC 中的客户端与 OceanBase 租户连接起来。第一阶段由 OceanBase Provider 创建**服务**；第二阶段由云厂商账号创建 **Endpoint（终端节点）**，再将 Endpoint ID 提供给 OceanBase Provider 完成关联。仅完成第一阶段时，还不能证明网络或数据库已经连通。云厂商资源可能产生费用，地域和可用区支持情况以实际服务为准。

以下命令假设前面已在国内 PRE 的 `create` 目录创建实例与租户，当前 workspace 是 `cn-pre-demo`，终端仍配置了 PRE 地址、凭证和本地 Provider。其他环境要一并切换环境、凭证、变量文件和 workspace。

1. 运行 `nano cn-pre-pvl.tfvars`，打开或新建 PrivateLink 变量文件。写入下面的内容，将账号 ID 换成**创建云厂商 Endpoint 的主账号 ID**，保留引号；先将 Endpoint ID 留空：

   ```hcl
   enable_private_link        = true
   private_link_user_account  = "<云厂商主账号 ID>"
   private_link_endpoint_id   = ""
   ```

   按 **Control+O** 保存、回车确认，再按 **Control+X** 退出。`enable_private_link` 控制是否创建服务；`private_link_user_account` 是获准接入服务的云厂商账号，不是 OceanBase 项目 ID；`private_link_endpoint_id` 在第二阶段填写。服务创建后更改 `private_link_user_account` 会使 Terraform 计划替换服务；需要允许其他账号时，先确认服务端的加白方式，不要直接改成另一个账号并执行。
2. 查看工作区和计划。两个变量文件都要带上，后面的文件覆盖前面同名的值：

   ```bash
   terraform workspace show
   terraform plan -var-file=cn-pre.tfvars -var-file=cn-pre-pvl.tfvars
   ```

   确认工作区为 `cn-pre-demo`，计划只新增 `oceanbase_tenant_private_link_service.demo[0]`，没有修改或删除实例、租户。如果服务已经在云端存在，应先核对并按[资源文档](../../docs/resources/tenant_private_link_service.md)导入，不能再创建一份。
3. 确认计划后执行，并查看服务信息：

   ```bash
   terraform apply -var-file=cn-pre.tfvars -var-file=cn-pre-pvl.tfvars
   terraform output private_link_service_name
   terraform output private_link_service_account
   terraform output private_link_service_status
   terraform output private_link_address_id
   ```

   Terraform 提示时输入 `yes`。把服务名称、地域、获准接入的账号以及按需返回的服务所属账号交给负责云网络的同事；地址 ID 用于 OceanBase 侧关联，不是云厂商 Endpoint ID。服务状态和名称以实际输出为准；服务所属账号仅在接口返回时才有输出。
4. 云网络同事在相应地域、获准账号下，使用服务名称和已有 VPC / 交换机创建云厂商 Endpoint。取得 Endpoint ID 后，把 `cn-pre-pvl.tfvars` 的 `private_link_endpoint_id` 改为该 ID，再执行：

   ```bash
   terraform plan -var-file=cn-pre.tfvars -var-file=cn-pre-pvl.tfvars
   terraform apply -var-file=cn-pre.tfvars -var-file=cn-pre-pvl.tfvars
   terraform output private_link_connection_status
   terraform output private_link_domain
   ```

   执行前确认计划只新增 `oceanbase_tenant_private_link_connection.demo[0]`，没有替换服务或修改现有资源。关联完成后仍需从 VPC 内检查 DNS、网络访问和数据库连接。

启用 PrivateLink 后，后续对这个工作区执行 `plan` 或 `apply` 时都要带两个 `-var-file`；漏掉 `cn-pre-pvl.tfvars` 会使 Terraform 认为应删除 PrivateLink。删除最后一个 Endpoint 关联时，OceanBase 也可能删除对应服务，详见[资源文档](../../docs/resources/tenant_private_link_connection.md)。

### 遇到问题先看这里

| 现象 | 先检查什么 |
| --- | --- |
| `terraform: command not found` 或 `go: command not found` | 第 1 步对应工具没有安装，或安装后终端尚未刷新。 |
| `site 必须是 CHINA、GLOBAL 或 MOCK` | 检查当前使用的 `.tfvars`；不要填写 `PRE`、`PROD` 或内部 Provider 的 `CC_CHINA_PRE`。 |
| 认证失败或没有权限 | 确认当前终端的 OpenAPI 域名、OceanBase AccessKey 和项目属于同一个环境；必要时请该环境管理员检查权限。 |
| 提示找不到变量文件 | 先用 `pwd` 确认自己在 `create` 还是 `read`；两个目录的 `.tfvars` 不共用。 |
| `plan` 意外出现修改或删除 | 暂停执行 `apply`，核对 `terraform workspace show`、环境域名、`-var-file` 和文件中所有资源 ID/参数。 |
| 创建超时或失败 | 先按返回的资源 ID 在同一环境的控制台/OpenAPI 查询实际状态，再运行 `plan`；不要直接重复创建。 |

如果终端已关闭，需要重新设置 `TF_CLI_CONFIG_FILE`、目标环境域名和凭证；这些是终端环境变量，不会自动写进 `.tfvars`。

## 一、每个文件负责什么

所有路径相对于仓库根目录的 `examples/demo/`。

| 文件或目录 | 作用 | 演示时怎么处理 |
| --- | --- | --- |
| `setup.sh` | 从当前仓库代码构建 Provider，并生成本目录专用的 `.terraformrc` | 第一次使用、切换分支或修改 Provider 后运行；通常不改脚本 |
| `.provider/` | `setup.sh` 生成的 Provider 可执行文件 | 自动生成，不用手工编辑 |
| `.terraformrc` | 告诉 Terraform 从 `.provider/` 加载本地 Provider | 自动生成；它不决定国内、海外或 PRE 环境 |
| `create/versions.tf` | 声明 OceanBase Provider，并将凭证和 `site` 变量传给它 | 通常不改；变量值通过终端和 `.tfvars` 提供 |
| `create/variables.tf` | 定义创建演示可用的变量、类型和默认值 | 通常不改；需要改值时优先写入 `.tfvars` |
| `create/terraform.tfvars.example` | 创建演示的完整参数模板，不含密钥 | 复制成自己的 `.tfvars`，按目标环境修改 |
| `create/terraform.tfvars` | 当前工作区的默认创建参数文件 | Terraform 自动读取；本机已填生产测试项目，切环境时不要直接沿用 |
| `create/cluster.tf` | 定义实例资源，并引用 `.tfvars` 中的实例参数 | 常用参数已变量化；演示只使用 `PostPay` |
| `create/tenant.tf` | 定义租户，关联上面的实例 | 常用租户参数已变量化；创建时公网开关固定为 `false` |
| `create/network.tf` | 可选的租户 IP 白名单组和 `INTERNET` 公网地址 | 修改 `enable_internet`、`security_group_name`、`allowed_ips`，通常不改文件 |
| `create/private_link.tf` | 可选的 OceanBase PrivateLink 服务及 Endpoint 关联 | 按第 7 步填写三个 PrivateLink 参数；云厂商 Endpoint 单独创建 |
| `create/database.tf` | 可选的数据库、普通账号及数据库权限 | 修改数据库相关变量，账号密码通过终端环境变量提供 |
| `create/outputs.tf` | 输出新建实例、租户、数据库和公网地址的标识 | 需要新增演示输出时编辑；不参与目标环境选择 |
| `read/versions.tf`、`read/variables.tf` | 声明只读查询使用的 Provider 和输入变量 | 通常不改；输入放在 `read` 自己的 `.tfvars` |
| `read/main.tf`、`read/outputs.tf` | 只读查询实例、租户及可选对象并输出结果 | 不会接管或创建这些云资源；通常不改 |
| `read/terraform.tfvars.example` | 已有资源查询参数模板 | 复制后填写目标环境已有资源的 ID 或名称 |
| `.terraform/`、`terraform.tfstate*` | Terraform 初始化信息和本地状态 | 自动生成；不要把不同环境的 state 混用或提交到 Git |

`create` 和 `read` 是两个独立的 Terraform 根目录，各自有自己的变量文件、workspace 和 state。上次单独创建的生产测试实例由另一份 state 管理，`create` 目录不会自动接管它。

## 二、环境由哪几项共同决定

| 需要选择的内容 | 在哪里配置 | 作用 |
| --- | --- | --- |
| Provider 加载位置 | 终端变量 `TF_CLI_CONFIG_FILE`，指向 `setup.sh` 生成的 `.terraformrc` | 只决定使用哪一个本地 Provider 二进制，不决定云环境 |
| 国内或海外站点 | `create` 或 `read` 的 `.tfvars` 中的 `site` | `CHINA` 或 `GLOBAL`；不是 `DEV`、`SIT` 或 `PRE` |
| 具体 OpenAPI 环境 | 终端变量 `TF_OCEANBASE_API_HOST_CHINA` 或 `TF_OCEANBASE_API_HOST_GLOBAL` | 覆盖对应站点的 OpenAPI 域名 |
| 账号和项目 | 终端中的 OceanBase AccessKey；`create` 的 `.tfvars` 中的 `project_id` | 必须属于目标环境并有相应权限 |
| 云上资源位置和规格 | `create` 的 `.tfvars` 中的 `cloud_provider`、`region`、`zones`、规格和渠道等 | 决定在哪里创建什么实例，必须是该环境实际支持的组合 |
| 已有资源身份 | `read` 的 `.tfvars` 中的 `cluster_id`、`tenant_id` 等 | 必须是目标环境中真实存在且当前账号可访问的资源 |
| Terraform state | 每个根目录下的 Terraform workspace | 隔离不同环境的受管资源记录，不会自动切换凭证或变量文件 |

### 2.1 选择 OpenAPI 地址

| 目标环境 | `site` | 对应覆盖变量的值 |
| --- | --- | --- |
| 国内生产 | `CHINA` | 不设置 `TF_OCEANBASE_API_HOST_CHINA`，默认 `api-cloud-cn.oceanbase.com` |
| 国内 DEV | `CHINA` | `dev-api-cloud-cn.oceanbase.com` |
| 国内 SIT | `CHINA` | `sit-api-cloud-cn.oceanbase.com` |
| 国内 PRE | `CHINA` | `pre-api-cloud-cn.oceanbase.com` |
| 国内 CC PRE | `CHINA` | `alibaba-api-cloud-cn-pre.oceanbase.com` |
| 国内 CC 生产 | `CHINA` | `alibaba-api-cloud-cn.oceanbase.com` |
| 海外 DEV | `GLOBAL` | `g-dev-api-cloud.oceanbase.com` |
| 海外 SIT | `GLOBAL` | `g-sit-api-cloud.oceanbase.com` |
| 海外 PRE | `GLOBAL` | `g-pre-api-cloud.oceanbase.com` |
| 海外生产 | `GLOBAL` | 显式填写 `api-cloud.oceanbase.com` |

当前 Provider 的 `GLOBAL` 默认值是 `g-api-cloud.oceanbase.com`；2026-11-19 从本机访问该默认域名时 TLS 未连通。上表的海外生产地址来自 OpenAPI 项目文档，并在同日返回匿名访问 401。其他表中地址也只做过匿名入口检查：401 表示入口可达，不代表凭证有效、权限足够或创建接口已验证。地址和目标环境配置可能变化。

例如，切到国内 PRE：

```bash
unset TF_OCEANBASE_API_HOST_CHINA TF_OCEANBASE_API_HOST_GLOBAL
export TF_OCEANBASE_API_HOST_CHINA="pre-api-cloud-cn.oceanbase.com"
```

再在目标目录的 `.tfvars` 中设置 `site = "CHINA"`。切回国内生产时，执行 `unset TF_OCEANBASE_API_HOST_CHINA`，同时切换生产凭证、项目和 workspace。

覆盖变量**只填域名**，不要填 `https://` 或 API 路径。`CHINA` 和 `GLOBAL` 使用 HTTPS；`MOCK` 仅连接本机 `http://127.0.0.1:8080` 模拟服务。当前开源 Provider 不识别内部 Provider 的 `CC_CHINA_PRE` 作为 `site`；CC PRE 要用 `CHINA` 加对应域名覆盖。

### 2.2 配置凭证

在运行 Terraform 的同一个终端输入**目标环境的 OceanBase Cloud OpenAPI** AccessKey：

```bash
printf 'OceanBase AccessKey ID: '
read -r TF_VAR_oceanbase_access_key
printf 'OceanBase AccessKey Secret: '
read -rs TF_VAR_oceanbase_secret_key
printf '\n'
export TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key
```

这不是阿里云 AccessKey。不要把密钥写进 `.tfvars`、命令行参数或 Git。切换环境时重新输入该环境的凭证；演示结束后执行：

```bash
unset TF_VAR_oceanbase_access_key TF_VAR_oceanbase_secret_key TF_VAR_database_user_password
```

### 2.3 分开环境的变量文件和 state

`terraform.tfvars` 会被 Terraform **自动读取**。例如运行 `-var-file=cn-pre.tfvars` 时，指定文件中同名的值会覆盖自动读取文件，但未在指定文件中写出的值仍可能来自默认文件或变量默认值。因此切环境时，应在目标文件中完整填写环境相关参数，并检查 `terraform plan`。

`.gitignore` 会忽略 `*.tfvars`；`terraform.tfvars.example` 是可提交的模板。不要在任何变量文件里保存 AccessKey 或账号密码。

每个环境使用独立 workspace，例如在 `create` 目录：

```bash
terraform init
terraform workspace new cn-pre-demo
terraform workspace show
```

如果 workspace 已存在，使用 `terraform workspace select cn-pre-demo`，不要重复 `new`。workspace 只隔离 state；仍须分别选择 OpenAPI 域名、凭证和 `-var-file`。`read` 目录是另一个根目录，也要单独初始化和选择自己的 workspace。

## 三、创建配置：`create/terraform.tfvars`

从 `create/terraform.tfvars.example` 复制，例如：

```bash
cp terraform.tfvars.example cn-pre.tfvars
```

用编辑器打开 `cn-pre.tfvars`，根据下表逐项填写。模板中的杭州值是一套示例，切到 PRE、CC 或海外时需要核对该环境的售卖配置。`project_id` 的占位符必须替换成真实值。下面的名称都对应 `create/variables.tf` 中的变量。

### 3.1 站点、项目和实例参数

| 参数 | 默认/示例 | 作用与怎么填写 |
| --- | --- | --- |
| `site` | `"CHINA"` | 国内或 CC 填 `CHINA`，海外填 `GLOBAL`；具体 PRE/PROD 地址按第二节的终端变量设置。 |
| `project_id` | 无默认值 | 目标环境的 OceanBase 项目 ID，可在该环境控制台的项目页面确认；不要沿用其他环境的 ID。 |
| `instance_name` | `"tf_demo_cluster"` | 新建实例的名称；多次演示建议改成容易识别的名称。创建后可通过 Terraform 修改。 |
| `cloud_provider` | `"ALIYUN"` | 云厂商标识，不是 PRE/PROD 环境名；填写目标环境实际支持的值。创建后不能原地修改。 |
| `region` | `"cn-hangzhou"` | 云厂商地域 ID；必须是目标环境和项目可使用的地域。创建后不能原地修改。 |
| `zones` | `["cn-hangzhou-j", "cn-hangzhou-k"]` | 实例使用的可用区列表；必须属于 `region` 并满足所选副本模式。创建后不能原地修改。 |
| `instance_class` | `"4C16G"` | 实例计算规格；从目标环境可售规格中选择。支持变更，但可用规格以服务端校验为准。 |
| `ob_version` | `"4.4.2"` | OceanBase 数据库版本，不是 Terraform 或 Provider 版本。创建后不能通过本资源原地修改。 |
| `disk_size` | `50` | 存储空间，单位 GB；有效范围和步长以目标环境规格为准。支持调整。 |
| `disk_type` | `null` | 云盘类型；`null` 表示请求中不指定，由目标环境决定默认值。指定时使用该环境支持的盘型。创建后不能原地修改。 |
| `node_num` | `3` | 目标节点数；必须是 3～60 的 3 的倍数。主备、存储分离和部分 K8s 实例类型不支持该参数，可设为 `null`。创建接口不直接接收节点数，目标与实际值不同时，Provider 在实例上线后调整。 |
| `replica_mode` | `"2"` | 副本模式，**字符串**；根据目标环境的实例类型和可用区选择，不要把数字 `2` 不加引号写入。创建后不能原地修改。 |
| `instance_type` | `"CLUSTER"` | 实例类型；具体支持值由目标环境决定。创建后不能原地修改。 |
| `storage_architecture` | `"standardStorage"` | 存储架构；`null` 表示不指定，后台决定。创建后不能原地修改。 |
| `sale_channel` | `"domestic"` | 售卖渠道；必须与目标项目及环境匹配，`null` 表示不指定。创建后不能原地修改。 |

`payment_type` 在 `cluster.tf` 中固定为 `PostPay`，因为本演示不支持包年包月创建。`series` 和 `cpu_arch` 当前 Provider 不接受在创建时指定。

举例：个人知识库中的 CC PRE 历史实测使用过 `ALIYUN`、`ap-southeast-1`、可用区 `ap-southeast-1a/b` 和 `normandy_domestic` 渠道；那是 2026-11-03 内部 Provider 分支的结果，不能直接当作当前开源分支或当前库存的验证结果。实际演示前请用目标环境确认。

### 3.2 租户参数

| 参数 | 默认/示例 | 作用与怎么填写 |
| --- | --- | --- |
| `tenant_name` | `"tf_demo_tenant"` | 新建租户名称；创建后可修改。 |
| `tenant_mode` | `"MySQL"` | 租户兼容模式；目标实例必须支持。创建后不能通过本资源原地修改。 |
| `primary_zone` | `"cn-hangzhou-j"` | 租户主可用区，应从 `zones` 对应的可用区中选择。支持变更。 |
| `tenant_charset` | `"utf8mb4"` | 创建租户使用的字符集；创建后不能通过本资源原地修改。 |
| `tenant_timezone` | `"Asia/Shanghai"` | 创建租户使用的时区；创建后不能通过本资源原地修改。 |
| `tenant_unit_num` | `1` | 租户 Unit 数量；可用范围由实例资源和服务端决定。 |
| `tenant_unit_cpu` | `1` | 每个 Unit 的 CPU 核数。 |
| `tenant_unit_memory` | `4` | 每个 Unit 的内存，单位 GB；和 CPU、Unit 数量一起受实例容量约束。 |

`tenant.tf` 固定 `enable_public_address = false`，因此默认创建租户时不申请公网地址。后续是否创建公网，由下一节的 `enable_internet` 控制。

### 3.3 可选公网和 IP 白名单组

| 参数 | 默认/示例 | 作用与怎么填写 |
| --- | --- | --- |
| `enable_internet` | `false` | `false` 不创建公网和演示 IP 组；设为 `true` 时创建两者。 |
| `security_group_name` | `"tf_demo_ips"` | 新建 IP 白名单组名称。 |
| `allowed_ips` | `[]` | 客户端出口 IP 或 CIDR 列表。开启公网时填写目标环境允许的真实地址，例如 `["203.0.113.10"]`；这个地址只是文档示例。 |

`network.tf` 使用 `type = "INTERNET"`。当前公网地址资源不支持在创建时设置 SSL、RPC 或 PrivateLink 配置，也不支持通过该资源更新它们。Terraform 会先创建 IP 组，再创建公网地址；具体地址和端口通过 `terraform output` 查看。PrivateLink 使用单独的资源，见第 7 步。

### 3.4 可选数据库和账号

| 参数 | 默认/示例 | 作用与怎么填写 |
| --- | --- | --- |
| `enable_database_user` | `false` | 设为 `true` 后，创建一个数据库和普通账号。 |
| `database_name` | `"tf_demo_db"` | 新建数据库名称。 |
| `database_charset` | `"utf8mb4"` | 数据库字符集，必须与目标租户/环境支持范围匹配。 |
| `database_user_name` | `"tf_demo_user"` | 新建普通数据库账号名称。 |
| `database_user_role` | `"readwrite"` | 该账号对演示数据库的角色，例如 `readwrite` 或 `readonly`。 |
| `database_user_password` | `null` | 开启账号时必填；通过终端变量 `TF_VAR_database_user_password` 输入，不写进 `.tfvars`。 |

启用后，`database.tf` 把账号关联到演示数据库。账号类型固定为 `Normal`，数据库描述为 `Terraform demo database`；确需改变时可编辑该文件。数据库和账号的 Terraform 资源已实现，但本轮没有做它们的生产生命周期验证。账号密码可能保存在 Terraform state 中，`sensitive` 只隐藏常规输出。

### 3.5 只读查询文件中的参数

这些参数写在 `read` 目录自己的 `.tfvars` 文件里，和 `create` 的文件分开。`read/main.tf` 会根据这些值选择已有对象；留空的可选项不会查询。

| 参数 | 是否必填 | 作用与怎么填写 |
| --- | --- | --- |
| `site` | 默认 `CHINA` | 国内/CC 使用 `CHINA`，海外使用 `GLOBAL`；还需在终端设置对应的 OpenAPI 地址。 |
| `cluster_id` | 必填 | 已有实例 ID，可在目标环境的实例列表或实例详情中查看。 |
| `tenant_id` | 必填 | 上述实例下的已有租户 ID，可在租户列表或详情中查看。 |
| `database_name` | 可选，默认空 | 要查询的已有数据库名称；不是 ID。 |
| `database_user` | 可选，默认空 | 要查询的已有数据库账号名称；不是密码。 |
| `security_group_name` | 可选，默认空 | 要查询的已有 IP 白名单组名称；不是 IP 地址。 |
| `internet_address_id` | 可选，默认空 | 要查询的已有公网地址 ID；不是公网 IP 或域名。通过本演示创建时可用 `terraform output internet_address_id` 获取。 |

## 四、当前能力与验证边界

| 对象 | Terraform 类型 | 能力与验证 |
| --- | --- | --- |
| 实例 | `oceanbase_cluster` / 同名 `data` | 创建、查询、名称/规格/磁盘/节点数变更、启停、删除、导入。按量创建、名称与磁盘组合修改、查询和删除已在国内生产验证；节点数变化和启停未做本轮生产实测。 |
| 租户 | `oceanbase_tenant` / 同名 `data` | 创建、查询、名称/主区/Unit 变更、删除、导入。MySQL 创建、名称与主区组合修改、查询和删除已在国内生产验证；Unit 变更未做本轮生产实测。 |
| IP 白名单组 | `oceanbase_tenant_security_group` / 同名 `data` | 创建、查询、修改、删除、导入；创建、更新、外部删除后重建及删除已在国内生产验证。 |
| 公网地址 | `oceanbase_tenant_connection` / 同名 `data` | `INTERNET` 地址创建、查询、删除、导入；创建、回读、删除及临时账号 `SELECT 1` 已在国内生产验证。 |
| PrivateLink | `oceanbase_tenant_private_link_service`、`oceanbase_tenant_private_link_connection` | 管理 OceanBase 侧服务及 Endpoint 关联；云厂商 Endpoint 由云厂商侧另行创建。具体环境是否可用以及端到端连通性需要实际验证。 |
| 数据库 | `oceanbase_database` / 同名 `data` | 创建、查询、修改描述、删除、导入；本轮未做 Terraform 资源生产生命周期验证。 |
| 数据库账号 | `oceanbase_tenant_user` / 同名 `data` | 创建、查询、修改描述/密码/权限/锁定状态、删除、导入；本轮未做 Terraform 资源生产生命周期验证。 |

源码中的导入格式：实例使用实例 ID；租户使用 `<实例 ID>.<租户 ID>`；IP 组、数据库、账号使用 `<实例 ID>.<租户 ID>.<名称>`；公网地址使用 `<实例 ID>.<租户 ID>.INTERNET.<地址 ID>`。导入前需要写好匹配的配置，并确认该对象没有同时由另一份 state 管理；本轮未逐项做生产导入验证。

本次演示目录的验证：`create` 和 `read` 通过 `terraform validate`；默认及可选资源的计划和模拟依赖测试通过；`read` 目录已用真实 OpenAPI 只读查询当前测试实例和租户，两者均返回 `ONLINE`。模拟测试不等同于云端创建，也不代表其他环境已验证。
