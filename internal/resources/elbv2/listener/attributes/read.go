package attributes

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Read retrieves the attributes of an ELBv2 listener.
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cfg := m.(aws.Config)

	svc := elasticloadbalancingv2.NewFromConfig(cfg)

	arn := d.Id()

	resp, err := svc.DescribeListenerAttributes(ctx, &elasticloadbalancingv2.DescribeListenerAttributesInput{
		ListenerArn: aws.String(arn),
	})
	if err != nil {
		return diag.Diagnostics{
			{
				Severity: diag.Error,
				Summary:  "Failed to get listener attributes",
				Detail:   err.Error(),
			},
		}
	}

	var strict string

	for _, attribute := range resp.Attributes {
		if attribute.Key == nil {
			continue
		}

		if attribute.Value == nil {
			continue
		}

		if *attribute.Key != KeyStrictTransportSecurity {
			continue
		}

		strict = *attribute.Value
	}

	d.Set(FieldListenerARN, arn)
	d.Set(FieldStrictTransportSecurity, strict)

	return nil
}
