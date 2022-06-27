Terraform CodeClimate Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x+
- [Go](https://golang.org/doc/install) 1.13.x+ (to build the provider plugin)

Building The Provider
---------------------
Build for linux (default) or darwin with make.

```sh
make build
```

Using the provider
----------------------

Currently the provider supports just retreaving the repository as data source.

```hcl
provider "codeclimate" {
  api_key = "${var.api_key}"        # Will fallback to CODECLIMATE_TOKEN environment variable if not explicitly specified.
}

data "codeclimate_repository" "test" {
  repository_slug = "babbel/test"
}
```

Get organization information

```hcl
provider "codeclimate" {
  api_key = "${var.api_key}"
}

data "codeclimate_organization" "babbel" {
  name = "babbel"
}
```

Create codeclimate repository

```hcl
provider "codeclimate" {
  api_key = "${var.api_key}"
}

data "codeclimate_organization" "babbel" {
  name = "babbel"
}

resource "codeclimate_repository" "codeclimate_terraform_test" {
  repository_url  = "https://github.com/babbel/codeclimate_terraform_test"
  organization_id = data.codeclimate_organization.babbel.id
}
```


Importing repository

```
terraform import codeclimate_repository.codeclimate_terraform_test babbel/codeclimate_terraform_test
```


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13.x+ is *required*). This provider works using Go Modules.

In order to test the provider, you can simply run `make test`.

```sh
$ go test ./...
```

Github Releases
---------------------------
In order to push a release to Github the feature branch has to be merged into master and then a tag needs to be created with the version name of the provider e.g. **v0.0.1** and pushed.

```sh
git checkout master
git pull origin master
git tag v<semver>
git push origin master --tags
```
