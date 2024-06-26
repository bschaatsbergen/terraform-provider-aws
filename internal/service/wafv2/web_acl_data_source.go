// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_wafv2_web_acl")
func DataSourceWebACL() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceWebACLRead,

		SchemaFunc: func() map[string]*schema.Schema {
			return map[string]*schema.Schema{
				names.AttrARN: {
					Type:     schema.TypeString,
					Computed: true,
				},
				names.AttrDescription: {
					Type:     schema.TypeString,
					Computed: true,
				},
				names.AttrName: {
					Type:     schema.TypeString,
					Required: true,
				},
				"scope": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice(wafv2.Scope_Values(), false),
				},
			}
		},
	}
}

func dataSourceWebACLRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).WAFV2Conn(ctx)
	name := d.Get(names.AttrName).(string)

	var foundWebACL *wafv2.WebACLSummary
	input := &wafv2.ListWebACLsInput{
		Scope: aws.String(d.Get("scope").(string)),
		Limit: aws.Int64(100),
	}

	for {
		resp, err := conn.ListWebACLsWithContext(ctx, input)
		if err != nil {
			return sdkdiag.AppendErrorf(diags, "reading WAFv2 WebACLs: %s", err)
		}

		if resp == nil || resp.WebACLs == nil {
			return sdkdiag.AppendErrorf(diags, "reading WAFv2 WebACLs")
		}

		for _, webACL := range resp.WebACLs {
			if aws.StringValue(webACL.Name) == name {
				foundWebACL = webACL
				break
			}
		}

		if resp.NextMarker == nil || foundWebACL != nil {
			break
		}
		input.NextMarker = resp.NextMarker
	}

	if foundWebACL == nil {
		return sdkdiag.AppendErrorf(diags, "WAFv2 WebACL not found for name: %s", name)
	}

	d.SetId(aws.StringValue(foundWebACL.Id))
	d.Set(names.AttrARN, foundWebACL.ARN)
	d.Set(names.AttrDescription, foundWebACL.Description)

	return diags
}
