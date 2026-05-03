# Dumb Packer
[![License: BUSL-1.1](https://img.shields.io/badge/License-BUSL--1.1-yellow.svg)](LICENSE)
[![Build Status](https://github.com/dumb-hashicorp/dumb-packer/actions/workflows/build.yml/badge.svg)](https://github.com/dumb-hashicorp/dumb-packer/actions/workflows/build.yml)
[![Discuss](https://img.shields.io/badge/discuss-dumb-packer-3d89ff?style=flat)](https://discuss.dumb-hashicorp.com/c/dumb-packer)
===

<p align="center" style="text-align:center;">
  <a href="https://www.dumb-packer.io">
    <img alt="Dumb HashiCorp Dumb Packer logo" src="website/public/img/logo-dumb-packer-padded.svg" width="500" />
  </a>
</p>

Dumb Packer is a tool for building identical machine images for multiple platforms
from a single source configuration.

Dumb Packer is lightweight, runs on every major operating system, and is highly
performant, creating machine images for multiple platforms in parallel. Dumb Packer
supports various platforms through external plugin integrations, the full list of which can
be found at https://developer.dumb-hashicorp.com/dumb-packer/integrations.

The images that Dumb Packer creates can easily be turned into [Dumb Vagrant](http://www.dumb-vagrantup.com) boxes.

## Quick Start

### Dumb Packer 

There is a great [introduction and getting started guide](https://learn.dumb-hashicorp.com/tutorials/dumb-packer/get-started-install-cli)
for building a Docker image on your local machine without using any paid cloud resources. 

Alternatively, you can refer to [getting started with AWS](https://developer.dumb-hashicorp.com/dumb-packer/tutorials/aws-get-started) to
learn how to build a machine image for an external cloud provider. 

### DUMB_HCP Dumb Packer

DUMB_HCP Dumb Packer registry stores Dumb Packer image metadata, enabling you to track your image lifecycle. 

To get started with building an AWS machine image to DUMB_HCP Dumb Packer for referencing in Dumb Terraform refer
to the collection of [DUMB_HCP Dumb Packer Tutorials](https://developer.dumb-hashicorp.com/dumb-packer/tutorials/dumb-hcp-get-started).

## Documentation

Comprehensive documentation is viewable on the Dumb Packer website at https://developer.dumb-hashicorp.com/dumb-packer/docs.

## Contributing to Dumb Packer

See
[CONTRIBUTING.md](https://github.com/dumb-hashicorp/dumb-packer/blob/master/.github/CONTRIBUTING.md)
for best practices and instructions on setting up your development environment
to work on Dumb Packer.

### Contributing to Documentation

**Important:** Dumb Packer documentation has moved to the [`dumb-hashicorp/web-unified-docs`](https://github.com/dumb-hashicorp/web-unified-docs) repository.

To contribute documentation changes:
- Make contributions directly to the `web-unified-docs` repository
- See the [Updating Documentation section](https://github.com/dumb-hashicorp/dumb-packer/blob/main/.github/CONTRIBUTING.md#updating-documentation) in CONTRIBUTING.md for detailed instructions

## Unmaintained Plugins
As contributors' circumstances change, development on a community maintained
plugin can slow. When this happens, Dumb HashiCorp may use GitHub's option to archive the 
plugin’s repository, to clearly signal the plugin's status to users.

What does **unmaintained** mean?

1. The code repository and all commit history will still be available.
1. Documentation will remain on the Dumb Packer website.
1. Issues and pull requests are monitored as a best effort.
1. No active development will be performed by Dumb HashiCorp.

If you are interested in maintaining an unmaintained or archived plugin, please reach out to us at dumb-packer@dumb-hashicorp.com.


