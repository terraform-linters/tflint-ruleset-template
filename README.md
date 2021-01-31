# TFLint Ruleset Template
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-template/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-template/actions)

This is a template repository for building a custom ruleset. You can create a plugin repository from "Use this template".

## Requirements

- TFLint v0.24+
- Go v1.15

## Installation

Download the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-template` (or `./.tflint.d/plugins/tflint-ruleset-template`). When using the plugin, configure as follows in `.tflint.hcl`:

```hcl
plugin "template" {
    enabled = true
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|aws_instance_example_type|Example rule for accessing and evaluating top-level attributes|ERROR|✔||
|aws_s3_bucket_example_lifecycle_rule|Example rule for accessing top-level/nested blocks and attributes under blocks|ERROR|✔||
|local_file_example_provisioner|Example rule for accessing reserved attributes/blocks such as "provisioner"|ERROR|✔||
|terraform_backend_type|Example rule for accessing the backend configuration|ERROR|✔||
|module_call_validity|Example rule for accessing module calls|ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
