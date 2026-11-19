# Alibaba Cloud PrivateLink example

This example connects an OceanBase tenant to a VPC through Alibaba Cloud
PrivateLink. Terraform creates the OceanBase service, creates an Alibaba Cloud
endpoint and its zone, and then connects and confirms the endpoint through the
OceanBase API.

Before running it:

- Configure the OceanBase provider installation as described in the
  [repository README](../../README.md#using-the-provider). Set
  `TF_VAR_oceanbase_access_key` and `TF_VAR_oceanbase_secret_key` using your
  OceanBase API credentials.
- Configure Alibaba Cloud credentials using a supported
  [provider authentication method](https://registry.terraform.io/providers/aliyun/alicloud/latest/docs#authentication).
  When using an Alibaba Cloud CLI profile, set `ALIBABA_CLOUD_PROFILE` to the
  profile name, for example `default`. The account must be able to use
  PrivateLink and create endpoint resources.
- Provide an eligible OceanBase cluster and tenant, the Alibaba Cloud account
  ID as `user_account`, and a VPC and vSwitch in the endpoint region. The
  vSwitch zone must be supported by the OceanBase service.

Set the remaining input variables, then run `terraform init` and
`terraform apply`. The resulting endpoint and domain are available as outputs.
The Alibaba Cloud endpoint is a billable resource.
