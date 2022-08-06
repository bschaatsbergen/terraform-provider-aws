package servicecatalog

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func ResourceAcceptPortfolioShare() *schema.Resource {
	return &schema.Resource{
		Create: resourceAcceptPortfolioShareCreate,
		Read:   resourceAcceptPortfolioShareRead,
		Update: resourceAcceptPortfolioShareUpdate,
		Delete: resourceAcceptPortfolioShareDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"accept_language": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      AcceptLanguageEnglish,
				ValidateFunc: validation.StringInSlice(AcceptLanguage_Values(), false),
			},
			"portfolio_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(servicecatalog.PortfolioShareType_Values(), false),
			},
		},
	}
}

func resourceAcceptPortfolioShareCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ServiceCatalogConn

	portfolioId := d.Get("portfolio_id").(string)

	input := &servicecatalog.AcceptPortfolioShareInput{
		PortfolioId:        aws.String(portfolioId),
		PortfolioShareType: aws.String(d.Get("type").(string)),
	}

	if v, ok := d.GetOk("accept_language"); ok {
		input.AcceptLanguage = aws.String(v.(string))
	}

	var output *servicecatalog.AcceptPortfolioShareOutput

	err := resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var err error

		output, err = conn.AcceptPortfolioShare(input)

		if tfawserr.ErrMessageContains(err, servicecatalog.ErrCodeInvalidParametersException, "profile does not exist") {
			return resource.RetryableError(err)
		}

		if err != nil {
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if tfresource.TimedOut(err) {
		output, err = conn.AcceptPortfolioShare(input)
	}

	if err != nil {
		return fmt.Errorf("error accepting Service Catalog Portfolio Share: %w", err)
	}

	if output == nil {
		return fmt.Errorf("error accepting Service Catalog Portfolio Share: empty response")
	}

	d.SetId(portfolioId)

	return resourceAcceptPortfolioShareRead(d, meta)
}

func resourceAcceptPortfolioShareRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ServiceCatalogConn

	input := &servicecatalog.ListAcceptedPortfolioSharesInput{
		AcceptLanguage:     aws.String(d.Get("accept_language").(string)),
		PortfolioShareType: aws.String(d.Get("type").(string)),
	}

	//TODO: Iterate over results to find the one with the matching portfolio id
	if err != nil {
		return fmt.Errorf("error describing accepted Service Catalog Portfolio Share (%s): %w", d.Id(), err)
	}

	if output == nil {
		return fmt.Errorf("error getting accepted Service Catalog Portfolio Share (%s): empty response", d.Id())
	}

	d.Set("portfolio_id", d.Id())
	d.Set("type", output.Type)
	d.Set("accept_language", output.AcceptLanguage)

	return nil
}

func resourceAcceptPortfolioShareUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ServiceCatalogConn

	if d.HasChanges("accept_language", "share_tag_options") {
		input := &servicecatalog.UpdateAcceptPortfolioShareInput{
			PortfolioId: aws.String(d.Get("portfolio_id").(string)),
		}

		if v, ok := d.GetOk("accept_language"); ok {
			input.AcceptLanguage = aws.String(v.(string))
		}

		if v, ok := d.GetOk("share_tag_options"); ok {
			input.ShareTagOptions = aws.Bool(v.(bool))
		}

		err := resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			_, err := conn.UpdateAcceptPortfolioShare(input)

			if tfawserr.ErrMessageContains(err, servicecatalog.ErrCodeInvalidParametersException, "profile does not exist") {
				return resource.RetryableError(err)
			}

			if err != nil {
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if tfresource.TimedOut(err) {
			_, err = conn.UpdateAcceptPortfolioShare(input)
		}

		if err != nil {
			return fmt.Errorf("error updating Service Catalog Portfolio Share (%s): %w", d.Id(), err)
		}
	}

	return resourceAcceptPortfolioShareRead(d, meta)
}

func resourceAcceptPortfolioShareDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ServiceCatalogConn

	input := &servicecatalog.DeleteAcceptPortfolioShareInput{
		PortfolioId: aws.String(d.Get("portfolio_id").(string)),
	}

	if v, ok := d.GetOk("accept_language"); ok {
		input.AcceptLanguage = aws.String(v.(string))
	}

	if v, ok := d.GetOk("type"); ok && v.(string) == servicecatalog.DescribeAcceptPortfolioShareTypeAccount {
		input.AccountId = aws.String(d.Get("principal_id").(string))
	} else {
		orgNode := &servicecatalog.OrganizationNode{}
		orgNode.Value = aws.String(d.Get("principal_id").(string))

		if v.(string) == servicecatalog.DescribeAcceptPortfolioShareTypeOrganizationMemberAccount {
			// portfolio_share type ORGANIZATION_MEMBER_ACCOUNT = org node type ACCOUNT
			orgNode.Type = aws.String(servicecatalog.OrganizationNodeTypeAccount)
		} else {
			orgNode.Type = aws.String(d.Get("type").(string))
		}

		input.OrganizationNode = orgNode
	}

	output, err := conn.DeleteAcceptPortfolioShare(input)

	if tfawserr.ErrCodeEquals(err, servicecatalog.ErrCodeResourceNotFoundException) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error deleting Service Catalog Portfolio Share (%s): %w", d.Id(), err)
	}

	// only get a token if organization node, otherwise check without token
	if output.AcceptPortfolioShareToken != nil {
		if _, err := WaitAcceptPortfolioShareDeletedWithToken(conn, aws.StringValue(output.AcceptPortfolioShareToken), d.Timeout(schema.TimeoutDelete)); err != nil {
			return fmt.Errorf("error waiting for Service Catalog Portfolio Share (%s) to be deleted: %w", d.Id(), err)
		}
	} else {
		if _, err := WaitAcceptPortfolioShareDeleted(conn, d.Get("portfolio_id").(string), d.Get("type").(string), d.Get("principal_id").(string), d.Timeout(schema.TimeoutDelete)); err != nil {
			return fmt.Errorf("error waiting for Service Catalog Portfolio Share (%s) to be deleted: %w", d.Id(), err)
		}
	}

	return nil
}
