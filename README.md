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

Please refer to the Terraform registry documentation **[here](https://registry.terraform.io/providers/tema-ai/codeclimate/latest/docs)**


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13.x+ is *required*). This provider works using Go Modules.

In order to test the provider, you can simply run `make test`.

```sh
$ go test ./...
```

Github Releases
---------------------------
In order to push a release to Github the feature branch has to be merged into production and then a tag needs to be created with the version name of the provider e.g. **v0.0.1** and pushed.

```sh
git checkout production
git pull origin production
git tag v<semver>
git push origin production --tags
```

Adding to the Terraform registry
--------------------------------
You can follow [this](https://learn.hashicorp.com/tutorials/terraform/provider-release-publish?in=terraform/providers#gpg_private_key) tutorial to add the provider to the terraform registry.

Acknowledgements
----------------
Original provider module by:
  - [travelaudience/terraform-provider-codeclimate](https://github.com/travelaudience/terraform-provider-codeclimate)
  - [babbel/terraform-provider-codeclimate](https://github.com/babbel/terraform-provider-codeclimate)
