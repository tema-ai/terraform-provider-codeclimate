<p align="center">
 <img width="100px" src="https://avatars.githubusercontent.com/u/96249131?s=400&u=b2257d99b282c443f80bbd4a8e9d015099c466ea&v=4" align="center" alt="GitHub Readme Stats" />
 <h2 align="center">Tema AI</h2>
 <p align="center">Terraform CodeClimate Provider</p>
</p>

<p align="center">

  <a href="#">
    <img src="https://img.shields.io/github/v/release/tema-ai/terraform-provider-codeclimate?label=last%20release&logo=github" />
  </a>
  <a href="https://www.terraform.io">
    <img src="https://img.shields.io/badge/terraform-1.1.6-%23623CE4" />
  </a>
  <a href="#">
    <img alt="Issues" src="https://github.com/tema-ai/terraform-provider-codeclimate/actions/workflows/test_and_release.yml/badge.svg" />
  </a>
  <a href="https://gitter.im/hashicorp-terraform/Lobby">
    <img alt="chat" src="https://badges.gitter.im/hashicorp-terraform/Lobby.png"/>
  </a>
  <a>
    <img alt="downloads" src="https://img.shields.io/github/downloads/tema-ai/terraform-provider-codeclimate/total?label=Total%20Downloads%20&logo=github&style=flat"/>
  </a>
  <a>
    <img alt="stars" src="https://img.shields.io/github/stars/tema-ai/terraform-provider-codeclimate?label=stars&logo=github"/>
  </a>
  <a>
    <img alt="last production commit" src="https://img.shields.io/github/last-commit/tema-ai/terraform-provider-codeclimate/production?logo=github"/>
  </a>
  <a>
    <img alt="license" src="https://img.shields.io/github/license/tema-ai/terraform-provider-codeclimate?logo=github"/>
  </a>
  <a>
    <img alt="license" src="https://img.shields.io/github/go-mod/go-version/tema-ai/terraform-provider-codeclimate/production?label=Go&logo=go"/>
  </a>
  <a href="https://github.com/tema-ai/terraform-provider-codeclimate/issues">
    <img alt="issues" src="https://img.shields.io/github/issues-raw/tema-ai/terraform-provider-codeclimate?logo=github"/>
  </a>
</p>

<p align="center">
  <a href="#demo">View Demo</a>
  ·
  <a href="https://github.com/tema-ai/terraform-provider-codeclimate/issues/new/choose">Report Bug</a>
  ·
  <a href="https://github.com/tema-ai/terraform-provider-codeclimate/issues/new/choose">Request Feature</a>
  ·
  <a href="https://github.com/tema-ai/terraform-provider-codeclimate/discussions">Ask Question</a>
</p>

-----------------------

## Introduction

The original providers **[travelaudience/terraform-provider-codeclimate](https://github.com/travelaudience/terraform-provider-codeclimate)** and **[babbel/terraform-provider-codeclimate](https://github.com/babbel/terraform-provider-codeclimate)** were using the GitHub slug of the CodeClimate repository as a unique id to store the resource in terraform. Meaning that you couldn't have two CodeClimate repositories connected to the same Github Repository. This was a problem for monorepo approaches where different services in a single repository have different CodeClimate projects.


This new provider **[tema-ai/terraform-provider-codeclimate]()** solves this by using the CodeClimate Repository id as the terraform resource id.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Using the provider](#using-the-provider)
- [Working on the provider](#working-on-the-provider)
  - [Github Releases](#github-releases)
  - [Adding to the Terraform registry](#adding-to-the-terraform-registry)
- [Acknowledgements](#acknowledgements)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Using the provider

Please refer to the Terraform registry documentation **[here](https://registry.terraform.io/providers/tema-ai/codeclimate/latest/docs)**

-----------
## Working on the provider
### Github Releases

In order to push a release to Github the feature branch has to be merged into production and then a tag needs to be created with the version name of the provider e.g. **v0.0.1** and pushed.

```sh
git checkout production
git pull origin production
git tag v<semver>
git push origin production --tags
```

Then a GitHub action will test the code and build a new release which will be automatically picked up by terraform.

### Adding to the Terraform registry
You can follow **[this](https://learn.hashicorp.com/tutorials/terraform/provider-release-publish?in=terraform/providers#gpg_private_key)** tutorial to add a provider to the terraform registry.

----------------------

## Acknowledgements
Original provider by:
  - [travelaudience/terraform-provider-codeclimate](https://github.com/travelaudience/terraform-provider-codeclimate)
  - [babbel/terraform-provider-codeclimate](https://github.com/babbel/terraform-provider-codeclimate)
