// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceDataCellsFilter,
			Name:    "Data Cells Filter",
		},
		{
			Factory: newResourceResourceLFTag,
			Name:    "Resource LF Tag",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDataLakeSettings,
			TypeName: "aws_lakeformation_data_lake_settings",
		},
		{
			Factory:  DataSourcePermissions,
			TypeName: "aws_lakeformation_permissions",
		},
		{
			Factory:  DataSourceResource,
			TypeName: "aws_lakeformation_resource",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDataLakeSettings,
			TypeName: "aws_lakeformation_data_lake_settings",
		},
		{
			Factory:  ResourceLFTag,
			TypeName: "aws_lakeformation_lf_tag",
		},
		{
			Factory:  ResourcePermissions,
			TypeName: "aws_lakeformation_permissions",
		},
		{
			Factory:  ResourceResource,
			TypeName: "aws_lakeformation_resource",
			Name:     "Resource",
		},
		{
			Factory:  ResourceResourceLFTags,
			TypeName: "aws_lakeformation_resource_lf_tags",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LakeFormation
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*lakeformation.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return lakeformation.NewFromConfig(cfg,
		lakeformation.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
