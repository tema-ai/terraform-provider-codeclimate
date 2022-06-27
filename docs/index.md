---
page_title: "CodeClimate Provider"
subcategory: ""
description: |-

---

# CodeClimate Provider

## Schema

### Required

- **api_key** (String, Required) Code Climate uses API access tokens to allow access to the API. You can generate a new Code Climate personal access token on codeclimate.com in the [token settings](https://codeclimate.com/profile/tokens) area of your Code Climate user profile.

## Example usage

Simple configuration by providing the CodeClimate token directly
```terraform
terraform {
  required_providers {
    codeclimate = {
      source  = "tema.ai/codeclimate"
      version = "~> 1.0.0"
    }
  }
}

provider "codeclimate" {
  api_key = "234324324...34258282"
}
```

We strongly discourage hardcoding the token on your terraform files or variables. The following example shows how you could use [AWS SecretsManager](https://aws.amazon.com/secrets-manager/) to store the token.
Assuming a Secret has been stored under the name **codeclimate_api_token** and containing the following data:
```json
{"token":"234324324...34258282"}
```

```terraform
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    codeclimate = {
      source  = "tema.ai/codeclimate"
      version = "~> 1.0.0"
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
```
