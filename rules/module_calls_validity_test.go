package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ModuleCallValidity(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "module source issue",
			Content: `
module "foo" {
  source = "in/correct"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleCallValidityRule(),
					Message: "unacceptable module source",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 24},
					},
				},
			},
		},

		{
			Name: "providers issue",
			Content: `
module "foo" {
  source    = "acceptable/source"
  providers = { aws.dns = aws.east }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleCallValidityRule(),
					Message: "must not pass providers",
					Range:   hcl.Range{},
				},
			},
		},

		{
			Name: "version constraint",
			Content: `
module "foo" {
  source  = "acceptable/source"
  version = ">= 1.0.0"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewModuleCallValidityRule(),
					Message: "must accept version 0.1.0",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 13},
						End:      hcl.Pos{Line: 4, Column: 23},
					},
				},
			},
		},
	}

	rule := NewModuleCallValidityRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
