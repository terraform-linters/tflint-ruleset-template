package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformBackendTypeRule checks whether ...
type TerraformBackendTypeRule struct{}

// NewTerraformBackendTypeRule returns a new rule
func NewTerraformBackendTypeRule() *TerraformBackendTypeRule {
	return &TerraformBackendTypeRule{}
}

// Name returns the rule name
func (r *TerraformBackendTypeRule) Name() string {
	return "terraform_backend_type"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformBackendTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformBackendTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *TerraformBackendTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *TerraformBackendTypeRule) Check(runner tflint.Runner) error {
	backend, err := runner.Backend()
	if err != nil {
		return err
	}
	if backend == nil {
		return nil
	}

	return runner.EmitIssue(
		r,
		fmt.Sprintf("backend type is %s", backend.Type),
		backend.DeclRange,
	)
}
