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
			"portfolio_share_type": {
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
		PortfolioShareType: aws.String(d.Get("portfolio_share_type").(string)),
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
		PortfolioShareType: aws.String(d.Get("portfolio_share_type").(string)),
	}

	output, err := conn.ListAcceptedPortfolioShares(input)

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

func resourceAcceptPortfolioShareDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
