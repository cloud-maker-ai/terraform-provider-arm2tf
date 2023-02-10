---
page_title: "Provider: ARM2TF"
description: |-
  The ARM2TF provider is used to support migrations from ARM / Bicep to Terraform.
---

# ARM2TF Provider

The `arm2tf` provider implements ARM template functions, allowing configurations written in ARM or Bicep to be reproduced in Terraform whilst preserving the results (e.g. resource names).

Currently, the implemented functions are:

- (Logical) Resource [arm2tf_guid](./docs/resources/guid.md) implements the ARM Template function [Guid()](https://learn.microsoft.com/en-us/azure/azure-resource-manager/templates/template-functions-string#guid).

- (Logical) Resource [arm2tf_unique_string](./docs/resources/unique_string.md) implements the ARM Template function [UniqueString()](https://learn.microsoft.com/en-us/azure/azure-resource-manager/templates/template-functions-string#uniquestring).

**N.B.** All resources marked (Logical) carry out their processing within the provider, *they don't call any network resources* such as the Azure Resource Manager APIs.