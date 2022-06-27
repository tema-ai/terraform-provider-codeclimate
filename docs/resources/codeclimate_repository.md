---
page_title: "codeclimate_repository Resource - terraform-provider-codeclimate"
subcategory: ""
description: |-

---

# Resource:  codeclimate_repository

## Schema

### Required

- **repository_url** (String, Required) GitHub url to the repository
- **organization_id** (String, Required) Organization id, you can compute this directly from the data-resource

### Read-only

- **codeclimate_id** (String, Read-only) id of the CodeClimate project created
- **test_reporter_id** (String, Read-only) Each repository added to Code Climate is automatically assigned a unique [Test Reporter ID](https://docs.codeclimate.com/docs/finding-your-test-coverage-token). When you run the test reporter, it submits your coverage data to Code Climate and includes this ID (which it reads from the environment) so we can map the data to the correct repository.
- **branch** (String, Read-only) Default branch for CodeClimate to run its analysis. Currently this is a read only and you cannot update it from terraform. Please update it directly from your CodeClimate Dashboard.
- **human_name** (String, Read-only) Name given to the repository in CodeClimate. urrently this is a read only and you cannot update it from terraform. Please update it directly from your CodeClimate Dashboard.
- **link_services** (String, Read-only) link to the respository on CodeClimate
- **link_self** (String, Read-only) link to the respository on CodeClimate
- **link_web_coverage** (String, Read-only) link to the coverage page on the CodeClimate project.
- **link_web_issues** (String, Read-only) link to the issues page on the CodeClimate project.
- **link_maintainability_badge** (String, Read-only) link to the maintainability badge.
- **link_test_coverage_badge** (String, Read-only) link to the test coverage badge.

## Example usage

```hcl
terraform {
  required_version = "~> 1.1"
  required_providers {
    codeclimate = {
      version = "~> 1.1.1"
      source  = "tema.ai/tema/codeclimate"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.12.0"
    }
  }
}

data "aws_secretsmanager_secret" "codeclimate_api_token" {
  name = "codeclimate_api_token"
}

data "aws_secretsmanager_secret_version" "codeclimate_api_token" {
  secret_id = data.aws_secretsmanager_secret.codeclimate_api_token.id
}

locals {
  code_climate_api_key = jsondecode(data.aws_secretsmanager_secret_version.codeclimate_api_token.secret_string)["token"]
}

provider "codeclimate" {
  api_key = local.code_climate_api_key
}

data "codeclimate_organization" "tema-ai" {
  name = "tema-ai"
}

resource "codeclimate_repository" "repo" {
  repository_url  = "https://github.com/tema-ai/terraform-provider-codeclimate"
  organization_id = data.codeclimate_organization.tema-ai.id
}

```
