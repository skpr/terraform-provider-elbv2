package attributes

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// We do not modify any of the current headers.
	return nil
}
