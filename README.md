# TFLint Ruleset Template
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-template/workflows/build/badge.svg?branch=main)](https://github.com/terraform-linters/tflint-ruleset-template/actions)

This is a template repository for building a custom ruleset. You can create a plugin repository from "Use this template". See also [Writing Plugins](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/plugins.md).

## Requirements

- TFLint v0.24+
- Go v1.16

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "template" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/terraform-linters/tflint-ruleset-template"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  mQINBGCqS2YBEADJ7gHktSV5NgUe08hD/uWWPwY07d5WZ1+F9I9SoiK/mtcNGz4P
  JLrYAIUTMBvrxk3I+kuwhp7MCk7CD/tRVkPRIklONgtKsp8jCke7FB3PuFlP/ptL
  SlbaXx53FCZSOzCJo9puZajVWydoGfnZi5apddd11Zw1FuJma3YElHZ1A1D2YvrF
  ...
  KEY
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
