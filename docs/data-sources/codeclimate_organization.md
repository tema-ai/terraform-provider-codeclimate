---
page_title: "codeclimate_organization Data Source - terraform-provider-codeclimate"
subcategory: ""
description: |-

---

# Data Source:  codeclimate_organization

## Schema

### Required

- **name** (String, Required) Name of the organization in CodeClimate

### Read-only

- **id** (String, Read-only) Unique identifier for the codeclimate organization

## Example usage

```hcl
data "codeclimate_organization" "tema-ai" {
  name = "tema-ai"
}
```
