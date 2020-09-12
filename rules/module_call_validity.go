package rules

import (
	"github.com/hashicorp/go-version"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ModuleCallValidityRule checks whether ...
type ModuleCallValidityRule struct{}

// NewModuleCallValidityRule returns a new rule
func NewModuleCallValidityRule() *ModuleCallValidityRule {
	return &ModuleCallValidityRule{}
}

// Name returns the rule name
func (r *ModuleCallValidityRule) Name() string {
	return "module_call_validity"
}

// Enabled returns whether the rule is enabled by default
func (r *ModuleCallValidityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ModuleCallValidityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleCallValidityRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *ModuleCallValidityRule) Check(runner tflint.Runner) error {
	return runner.WalkModuleCalls(func(call *configs.ModuleCall) error {
		if call.SourceAddr != "acceptable/source" {
			return runner.EmitIssue(r, "unacceptable module source", call.SourceAddrRange)
		}

		if len(call.Providers) != 0 {
			return runner.EmitIssue(r, "must not pass providers", hcl.Range{})
		}

		expectedVersion, _ := version.NewVersion("0.1.0")
		if !call.Version.Required.Check(expectedVersion) {
			return runner.EmitIssue(r, "must accept version 0.1.0", call.Version.DeclRange)
		}

		return nil
	})
}
