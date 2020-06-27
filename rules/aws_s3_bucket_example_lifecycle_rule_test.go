package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsS3BucketExampleLifecycleRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "aws_s3_bucket" "bucket" {
  lifecycle_rule {
	enabled = false

    transition {
      days          = 30
      storage_class = "STANDARD_IA"
    }
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsS3BucketExampleLifecycleRuleRule(),
					Message: "`lifecycle_rule` block found",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 17},
					},
				},
				{
					Rule:    NewAwsS3BucketExampleLifecycleRuleRule(),
					Message: "`enabled` attribute found",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 12},
						End:      hcl.Pos{Line: 4, Column: 17},
					},
				},
				{
					Rule:    NewAwsS3BucketExampleLifecycleRuleRule(),
					Message: "`transition` block found",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 5},
						End:      hcl.Pos{Line: 6, Column: 15},
					},
				},
			},
		},
	}

	rule := NewAwsS3BucketExampleLifecycleRuleRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
