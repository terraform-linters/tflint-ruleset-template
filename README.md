# TFLint ruleset template

This is a template repository for building a custom ruleset. You can create a plugin repository from "Use this template".

## Requirements

- TFLint v0.13.0
- Go v1.13.x

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|aws_instance_example_type|Show instance type|ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
