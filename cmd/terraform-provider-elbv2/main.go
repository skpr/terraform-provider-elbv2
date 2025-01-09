package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/skpr/terraform-provider-elbv2/internal/resources/elbv2/listener/attributes"
)

const (
	// ResourceListenerAttributes provides a resource for configuring ELBv2 listener attributes.
	ResourceListenerAttributes = "aws_elbv2_listener_attributes"

	// FieldRegion identifier for region field.
	FieldRegion = "region"
	// FieldProfile identifier for profile field.
	FieldProfile = "profile"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider,
	})
}

func provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			FieldProfile: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_PROFILE", ""),
				Description: "AWS Profile for authentication.",
			},
			FieldRegion: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", ""),
				Description: "AWS Profile for authentication.",
			},
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			var (
				region  string
				profile string
			)

			if v, ok := d.GetOk(FieldRegion); ok {
				region = v.(string)
			}

			if v, ok := d.GetOk(FieldProfile); ok {
				profile = v.(string)
			}

			return config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile))
		},
		ResourcesMap: map[string]*schema.Resource{
			ResourceListenerAttributes: attributes.Resource(),
		},
	}
}
