package attributes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// FieldListenerARN is the Amazon Resource Name (ARN) of the listener.
	FieldListenerARN = "listener_arn"
	// FieldStrictTransportSecurity header value applied to responses.
	FieldStrictTransportSecurity = "strict_transport_security_header_value"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: Update,
		ReadContext:   Read,
		UpdateContext: Update,
		DeleteContext: Delete,

		Schema: map[string]*schema.Schema{
			FieldListenerARN: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldStrictTransportSecurity: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
