package rules

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty/gocty"
)

// ModuleCallValidityRule checks whether ...
type ModuleCallValidityRule struct {
	tflint.DefaultRule
}

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
func (r *ModuleCallValidityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ModuleCallValidityRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *ModuleCallValidityRule) Check(runner tflint.Runner) error {
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "module",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "source"},
						{Name: "providers"},
						{Name: "version"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, module := range content.Blocks {
		if sourceAttr, exists := module.Body.Attributes["source"]; exists {
			var source string
			val, diags := sourceAttr.Expr.Value(nil)
			if diags.HasErrors() {
				return diags
			}
			if err := gocty.FromCtyValue(val, &source); err != nil {
				return err
			}

			if source != "acceptable/source" {
				if err := runner.EmitIssue(r, "unacceptable module source", sourceAttr.Expr.Range()); err != nil {
					return err
				}
			}
		}

		if _, exists := module.Body.Attributes["providers"]; exists {
			if err := runner.EmitIssue(r, "must not pass providers", hcl.Range{}); err != nil {
				return err
			}
		}

		if versionAttr, exists := module.Body.Attributes["version"]; exists {
			var versionConstraint string
			val, diags := versionAttr.Expr.Value(nil)
			if diags.HasErrors() {
				return diags
			}
			if err := gocty.FromCtyValue(val, &versionConstraint); err != nil {
				return err
			}

			expectedVersion, _ := version.NewVersion("0.1.0")
			constraint, err := version.NewConstraint(versionConstraint)
			if err != nil {
				return err
			}
			if !constraint.Check(expectedVersion) {
				if err := runner.EmitIssue(r, "must accept version 0.1.0", versionAttr.Expr.Range()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
