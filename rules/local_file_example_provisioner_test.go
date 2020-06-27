package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_LocalFileExampleProvisioner(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "local_file" "file" {
  provisioner "local-exec" {
    command = "echo 'Good morning!'"
  }

  provisioner "remote-exec" {
    command = "echo 'Hello!'"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewLocalFileExampleProvisionerRule(),
					Message: "local-exec provisioner command found",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 15},
						End:      hcl.Pos{Line: 4, Column: 37},
					},
				},
			},
		},
	}

	rule := NewLocalFileExampleProvisionerRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
