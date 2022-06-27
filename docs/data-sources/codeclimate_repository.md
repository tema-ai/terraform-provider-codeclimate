---
page_title: "codeclimate_repository Data Source - terraform-provider-codeclimate"
subcategory: ""
description: |-

---

# Data Source:  codeclimate_repository

## Schema

### Required

- **repository_slug** (String, Required) repository slug

### Read-only

- **test_reporter_id** (String, Read-only) Test reporter id

## Example usage

```hcl
data "codeclimate_repository" "tema-ai" {
  repository_slug = "tema-ai"
}
```
