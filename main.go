package main

import (
	"github.com/terraform-linters/tflint-ruleset-template/rules"
	"github.com/wata727/tflint/plugin"
)

// Name returns the plugin name
func Name() string {
	return "template"
}

// Version returns the plugin version
func Version() string {
	return "0.1.0"
}

// NewRules returns a ruleset of the plugin
func NewRules() []plugin.Rule {
	return []plugin.Rule{
		rules.NewAwsInstanceExampleTypeRule(),
	}
}
