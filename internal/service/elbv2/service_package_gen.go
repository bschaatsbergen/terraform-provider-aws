// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package elbv2

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	elasticloadbalancingv2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	elbv2_sdkv1 "github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceLoadBalancer,
			TypeName: "aws_alb",
		},
		{
			Factory:  DataSourceListener,
			TypeName: "aws_alb_listener",
		},
		{
			Factory:  DataSourceTargetGroup,
			TypeName: "aws_alb_target_group",
		},
		{
			Factory:  DataSourceLoadBalancer,
			TypeName: "aws_lb",
		},
		{
			Factory:  DataSourceHostedZoneID,
			TypeName: "aws_lb_hosted_zone_id",
		},
		{
			Factory:  DataSourceListener,
			TypeName: "aws_lb_listener",
		},
		{
			Factory:  DataSourceTargetGroup,
			TypeName: "aws_lb_target_group",
		},
		{
			Factory:  DataSourceTrustStore,
			TypeName: "aws_lb_trust_store",
			Name:     "Trust Store",
		},
		{
			Factory:  DataSourceLoadBalancers,
			TypeName: "aws_lbs",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_alb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceListener,
			TypeName: "aws_alb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceListenerCertificate,
			TypeName: "aws_alb_listener_certificate",
		},
		{
			Factory:  ResourceListenerRule,
			TypeName: "aws_alb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceTargetGroup,
			TypeName: "aws_alb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceTargetGroupAttachment,
			TypeName: "aws_alb_target_group_attachment",
		},
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_lb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceListener,
			TypeName: "aws_lb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceListenerCertificate,
			TypeName: "aws_lb_listener_certificate",
		},
		{
			Factory:  ResourceListenerRule,
			TypeName: "aws_lb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceTargetGroup,
			TypeName: "aws_lb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceTargetGroupAttachment,
			TypeName: "aws_lb_target_group_attachment",
		},
		{
			Factory:  ResourceTrustStore,
			TypeName: "aws_lb_trust_store",
			Name:     "Trust Store",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceTrustStoreRevocation,
			TypeName: "aws_lb_trust_store_revocation",
			Name:     "Trust Store Revocation",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELBV2
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*elbv2_sdkv1.ELBV2, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return elbv2_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticloadbalancingv2_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return elasticloadbalancingv2_sdkv2.NewFromConfig(cfg, func(o *elasticloadbalancingv2_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
