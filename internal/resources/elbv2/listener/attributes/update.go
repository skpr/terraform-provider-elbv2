package attributes

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Update modifies the attributes of an ELBv2 listener.
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cfg := m.(aws.Config)

	svc := elasticloadbalancingv2.NewFromConfig(cfg)

	var (
		arn    = d.Get(FieldListenerARN).(string)
		strict = d.Get(FieldStrictTransportSecurity).(string)
	)

	_, err := svc.ModifyListenerAttributes(ctx, &elasticloadbalancingv2.ModifyListenerAttributesInput{
		ListenerArn: aws.String(arn),
		Attributes: []types.ListenerAttribute{
			{
				Key:   aws.String(KeyStrictTransportSecurity),
				Value: aws.String(strict),
			},
		},
	})
	if err != nil {
		return diag.Diagnostics{
			{
				Severity: diag.Error,
				Summary:  "Failed to modify listener attributes",
				Detail:   err.Error(),
			},
		}
	}

	d.SetId(arn)

	return nil
}
