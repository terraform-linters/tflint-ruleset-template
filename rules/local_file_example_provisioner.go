package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// LocalFileExampleProvisionerRule checks whether ...
type LocalFileExampleProvisionerRule struct {
	tflint.DefaultRule
}

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
func (r *LocalFileExampleProvisionerRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *LocalFileExampleProvisionerRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *LocalFileExampleProvisionerRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("local_file", &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "provisioner",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "command"},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, provisioner := range resource.Body.Blocks {
			if provisioner.Labels[0] != "local-exec" {
				continue
			}

			if attr, exists := provisioner.Body.Attributes["command"]; exists {
				if err := runner.EmitIssue(r, "local-exec provisioner command found", attr.Expr.Range()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
