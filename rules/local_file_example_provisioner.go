package rules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// LocalFileExampleProvisionerRule checks whether ...
type LocalFileExampleProvisionerRule struct{}

// NewLocalFileExampleProvisionerRule returns a new rule
func NewLocalFileExampleProvisionerRule() *LocalFileExampleProvisionerRule {
	return &LocalFileExampleProvisionerRule{}
}

// Name returns the rule name
func (r *LocalFileExampleProvisionerRule) Name() string {
	return "local_file_example_provisioner"
}

// Enabled returns whether the rule is enabled by default
func (r *LocalFileExampleProvisionerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *LocalFileExampleProvisionerRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *LocalFileExampleProvisionerRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *LocalFileExampleProvisionerRule) Check(runner tflint.Runner) error {
	return runner.WalkResources("local_file", func(resource *configs.Resource) error {

		for _, provisioner := range resource.Managed.Provisioners {
			if provisioner.Type != "local-exec" {
				continue
			}

			content, _, diags := provisioner.Config.PartialContent(&hcl.BodySchema{
				Attributes: []hcl.AttributeSchema{
					{Name: "command"},
				},
			})
			if diags.HasErrors() {
				return diags
			}

			if attr, exists := content.Attributes["command"]; exists {
				if err := runner.EmitIssueOnExpr(r, "local-exec provisioner command found", attr.Expr); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
