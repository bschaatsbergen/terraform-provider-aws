// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

// Exports for use in tests only.
var (
	ResourceAMICopy                                  = resourceAMICopy
	ResourceAMIFromInstance                          = resourceAMIFromInstance
	ResourceAMILaunchPermission                      = resourceAMILaunchPermission
	ResourceAvailabilityZoneGroup                    = resourceAvailabilityZoneGroup
	ResourceCapacityReservation                      = resourceCapacityReservation
	ResourceCarrierGateway                           = resourceCarrierGateway
	ResourceClientVPNAuthorizationRule               = resourceClientVPNAuthorizationRule
	ResourceClientVPNEndpoint                        = resourceClientVPNEndpoint
	ResourceClientVPNNetworkAssociation              = resourceClientVPNNetworkAssociation
	ResourceClientVPNRoute                           = resourceClientVPNRoute
	ResourceCustomerGateway                          = resourceCustomerGateway
	ResourceDefaultNetworkACL                        = resourceDefaultNetworkACL
	ResourceDefaultRouteTable                        = resourceDefaultRouteTable
	ResourceEBSDefaultKMSKey                         = resourceEBSDefaultKMSKey
	ResourceEBSEncryptionByDefault                   = resourceEBSEncryptionByDefault
	ResourceEBSFastSnapshotRestore                   = newEBSFastSnapshotRestoreResource
	ResourceEBSSnapshot                              = resourceEBSSnapshot
	ResourceEBSSnapshotCopy                          = resourceEBSSnapshotCopy
	ResourceEBSSnapshotImport                        = resourceEBSSnapshotImport
	ResourceEBSVolume                                = resourceEBSVolume
	ResourceEIP                                      = resourceEIP
	ResourceEIPAssociation                           = resourceEIPAssociation
	ResourceEIPDomainName                            = newEIPDomainNameResource
	ResourceFleet                                    = resourceFleet
	ResourceHost                                     = resourceHost
	ResourceIPAM                                     = resourceIPAM
	ResourceIPAMOrganizationAdminAccount             = resourceIPAMOrganizationAdminAccount
	ResourceIPAMPool                                 = resourceIPAMPool
	ResourceIPAMPoolCIDR                             = resourceIPAMPoolCIDR
	ResourceIPAMPoolCIDRAllocation                   = resourceIPAMPoolCIDRAllocation
	ResourceIPAMPreviewNextCIDR                      = resourceIPAMPreviewNextCIDR
	ResourceIPAMResourceDiscovery                    = resourceIPAMResourceDiscovery
	ResourceIPAMResourceDiscoveryAssociation         = resourceIPAMResourceDiscoveryAssociation
	ResourceIPAMScope                                = resourceIPAMScope
	ResourceImageBlockPublicAccess                   = resourceImageBlockPublicAccess
	ResourceInstance                                 = resourceInstance
	ResourceInstanceConnectEndpoint                  = newInstanceConnectEndpointResource
	ResourceInstanceMetadataDefaults                 = newInstanceMetadataDefaultsResource
	ResourceInstanceState                            = resourceInstanceState
	ResourceKeyPair                                  = resourceKeyPair
	ResourceLaunchTemplate                           = resourceLaunchTemplate
	ResourceLocalGatewayRoute                        = resourceLocalGatewayRoute
	ResourceLocalGatewayRouteTableVPCAssociation     = resourceLocalGatewayRouteTableVPCAssociation
	ResourceMainRouteTableAssociation                = resourceMainRouteTableAssociation
	ResourceNetworkACL                               = resourceNetworkACL
	ResourceNetworkACLRule                           = resourceNetworkACLRule
	ResourceNetworkInsightsAnalysis                  = resourceNetworkInsightsAnalysis
	ResourceNetworkInsightsPath                      = resourceNetworkInsightsPath
	ResourceNetworkInterface                         = resourceNetworkInterface
	ResourcePlacementGroup                           = resourcePlacementGroup
	ResourceRoute                                    = resourceRoute
	ResourceRouteTable                               = resourceRouteTable
	ResourceSecurityGroupEgressRule                  = newSecurityGroupEgressRuleResource
	ResourceSecurityGroupIngressRule                 = newSecurityGroupIngressRuleResource
	ResourceSnapshotCreateVolumePermission           = resourceSnapshotCreateVolumePermission
	ResourceSpotDataFeedSubscription                 = resourceSpotDataFeedSubscription
	ResourceSpotFleetRequest                         = resourceSpotFleetRequest
	ResourceSpotInstanceRequest                      = resourceSpotInstanceRequest
	ResourceTag                                      = resourceTag
	ResourceTrafficMirrorFilter                      = resourceTrafficMirrorFilter
	ResourceTrafficMirrorFilterRule                  = resourceTrafficMirrorFilterRule
	ResourceTrafficMirrorSession                     = resourceTrafficMirrorSession
	ResourceTrafficMirrorTarget                      = resourceTrafficMirrorTarget
	ResourceTransitGatewayConnect                    = resourceTransitGatewayConnect
	ResourceTransitGatewayMulticastDomain            = resourceTransitGatewayMulticastDomain
	ResourceTransitGatewayMulticastDomainAssociation = resourceTransitGatewayMulticastDomainAssociation
	ResourceTransitGatewayMulticastGroupMember       = resourceTransitGatewayMulticastGroupMember
	ResourceTransitGatewayMulticastGroupSource       = resourceTransitGatewayMulticastGroupSource
	ResourceTransitGatewayPeeringAttachment          = resourceTransitGatewayPeeringAttachment
	ResourceTransitGatewayPeeringAttachmentAccepter  = resourceTransitGatewayPeeringAttachmentAccepter
	ResourceTransitGatewayPolicyTable                = resourceTransitGatewayPolicyTable
	ResourceTransitGatewayPolicyTableAssociation     = resourceTransitGatewayPolicyTableAssociation
	ResourceTransitGatewayPrefixListReference        = resourceTransitGatewayPrefixListReference
	ResourceTransitGatewayRoute                      = resourceTransitGatewayRoute
	ResourceTransitGatewayRouteTable                 = resourceTransitGatewayRouteTable
	ResourceTransitGatewayRouteTableAssociation      = resourceTransitGatewayRouteTableAssociation
	ResourceTransitGatewayRouteTablePropagation      = resourceTransitGatewayRouteTablePropagation
	ResourceTransitGatewayVPCAttachment              = resourceTransitGatewayVPCAttachment
	ResourceTransitGatewayVPCAttachmentAccepter      = resourceTransitGatewayVPCAttachmentAccepter
	ResourceVPCEndpoint                              = resourceVPCEndpoint
	ResourceVPNConnection                            = resourceVPNConnection
	ResourceVPNConnectionRoute                       = resourceVPNConnectionRoute
	ResourceVPNGateway                               = resourceVPNGateway
	ResourceVPNGatewayAttachment                     = resourceVPNGatewayAttachment
	ResourceVPNGatewayRoutePropagation               = resourceVPNGatewayRoutePropagation
	ResourceVolumeAttachment                         = resourceVolumeAttachment

	CustomFiltersSchema                                        = customFiltersSchema
	ErrCodeDefaultSubnetAlreadyExistsInAvailabilityZone        = errCodeDefaultSubnetAlreadyExistsInAvailabilityZone
	ErrCodeInvalidSpotDatafeedNotFound                         = errCodeInvalidSpotDatafeedNotFound
	FindAvailabilityZones                                      = findAvailabilityZones
	FindCapacityReservationByID                                = findCapacityReservationByID
	FindCarrierGatewayByID                                     = findCarrierGatewayByID
	FindClientVPNAuthorizationRuleByThreePartKey               = findClientVPNAuthorizationRuleByThreePartKey
	FindClientVPNEndpointByID                                  = findClientVPNEndpointByID
	FindClientVPNNetworkAssociationByTwoPartKey                = findClientVPNNetworkAssociationByTwoPartKey
	FindClientVPNRouteByThreePartKey                           = findClientVPNRouteByThreePartKey
	FindCreateSnapshotCreateVolumePermissionByTwoPartKey       = findCreateSnapshotCreateVolumePermissionByTwoPartKey
	FindCustomerGatewayByID                                    = findCustomerGatewayByID
	FindEBSVolumeAttachment                                    = findVolumeAttachment
	FindEBSVolumeByID                                          = findEBSVolumeByID
	FindEIPByAllocationID                                      = findEIPByAllocationID
	FindEIPByAssociationID                                     = findEIPByAssociationID
	FindEIPDomainNameAttributeByAllocationID                   = findEIPDomainNameAttributeByAllocationID
	FindFastSnapshotRestoreByTwoPartKey                        = findFastSnapshotRestoreByTwoPartKey
	FindFleetByID                                              = findFleetByID
	FindHostByID                                               = findHostByID
	FindIPAMByID                                               = findIPAMByID
	FindIPAMPoolAllocationByTwoPartKey                         = findIPAMPoolAllocationByTwoPartKey
	FindIPAMPoolByID                                           = findIPAMPoolByID
	FindIPAMPoolCIDRByTwoPartKey                               = findIPAMPoolCIDRByTwoPartKey
	FindIPAMResourceDiscoveryAssociationByID                   = findIPAMResourceDiscoveryAssociationByID
	FindIPAMResourceDiscoveryByID                              = findIPAMResourceDiscoveryByID
	FindIPAMScopeByID                                          = findIPAMScopeByID
	FindImageLaunchPermission                                  = findImageLaunchPermission
	FindInstanceConnectEndpointByID                            = findInstanceConnectEndpointByID
	FindInstanceMetadataDefaults                               = findInstanceMetadataDefaults
	FindInstanceStateByID                                      = findInstanceStateByID
	FindKeyPairByName                                          = findKeyPairByName
	FindLaunchTemplateByID                                     = findLaunchTemplateByID
	FindMainRouteTableAssociationByID                          = findMainRouteTableAssociationByID
	FindNetworkACLByIDV2                                       = findNetworkACLByID
	FindNetworkInsightsAnalysisByID                            = findNetworkInsightsAnalysisByID
	FindNetworkInsightsPathByID                                = findNetworkInsightsPathByID
	FindNetworkInterfaceByIDV2                                 = findNetworkInterfaceByID
	FindNetworkPerformanceMetricSubscriptionByFourPartKey      = findNetworkPerformanceMetricSubscriptionByFourPartKey
	FindPlacementGroupByName                                   = findPlacementGroupByName
	FindPublicIPv4Pools                                        = findPublicIPv4Pools
	FindRouteByIPv4Destination                                 = findRouteByIPv4Destination
	FindRouteByIPv6Destination                                 = findRouteByIPv6Destination
	FindRouteByPrefixListIDDestination                         = findRouteByPrefixListIDDestination
	FindRouteTableAssociationByID                              = findRouteTableAssociationByID
	FindRouteTableByID                                         = findRouteTableByID
	FindSnapshot                                               = findSnapshot
	FindSnapshotByID                                           = findSnapshotByID
	FindSpotDatafeedSubscription                               = findSpotDatafeedSubscription
	FindSpotFleetRequestByID                                   = findSpotFleetRequestByID
	FindSpotFleetRequests                                      = findSpotFleetRequests
	FindSpotInstanceRequestByID                                = findSpotInstanceRequestByID
	FindSubnetsV2                                              = findSubnets
	FindTag                                                    = findTag
	FindTrafficMirrorFilterByID                                = findTrafficMirrorFilterByID
	FindTrafficMirrorFilterRuleByTwoPartKey                    = findTrafficMirrorFilterRuleByTwoPartKey
	FindTrafficMirrorSessionByID                               = findTrafficMirrorSessionByID
	FindTrafficMirrorTargetByID                                = findTrafficMirrorTargetByID
	FindTransitGatewayByID                                     = findTransitGatewayByID
	FindTransitGatewayConnectByID                              = findTransitGatewayConnectByID
	FindTransitGatewayConnectPeerByID                          = findTransitGatewayConnectPeerByID
	FindTransitGatewayMulticastDomainAssociationByThreePartKey = findTransitGatewayMulticastDomainAssociationByThreePartKey
	FindTransitGatewayMulticastDomainByID                      = findTransitGatewayMulticastDomainByID
	FindTransitGatewayMulticastGroupMemberByThreePartKey       = findTransitGatewayMulticastGroupMemberByThreePartKey
	FindTransitGatewayMulticastGroupSourceByThreePartKey       = findTransitGatewayMulticastGroupSourceByThreePartKey
	FindTransitGatewayPeeringAttachmentByID                    = findTransitGatewayPeeringAttachmentByID
	FindTransitGatewayPolicyTableAssociationByTwoPartKey       = findTransitGatewayPolicyTableAssociationByTwoPartKey
	FindTransitGatewayPolicyTableByID                          = findTransitGatewayPolicyTableByID
	FindTransitGatewayPrefixListReferenceByTwoPartKey          = findTransitGatewayPrefixListReferenceByTwoPartKey
	FindTransitGatewayRouteTableAssociationByTwoPartKey        = findTransitGatewayRouteTableAssociationByTwoPartKey
	FindTransitGatewayRouteTableByID                           = findTransitGatewayRouteTableByID
	FindTransitGatewayRouteTablePropagationByTwoPartKey        = findTransitGatewayRouteTablePropagationByTwoPartKey
	FindTransitGatewayStaticRoute                              = findTransitGatewayStaticRoute
	FindTransitGatewayVPCAttachmentByID                        = findTransitGatewayVPCAttachmentByID
	FindVPCEndpointConnectionByServiceIDAndVPCEndpointID       = findVPCEndpointConnectionByServiceIDAndVPCEndpointID
	FindVPCEndpointConnectionNotificationByID                  = findVPCEndpointConnectionNotificationByID
	FindVPCEndpointRouteTableAssociationExists                 = findVPCEndpointRouteTableAssociationExists
	FindVPCEndpointSecurityGroupAssociationExists              = findVPCEndpointSecurityGroupAssociationExists
	FindVPCEndpointServiceConfigurationByID                    = findVPCEndpointServiceConfigurationByID
	FindVPCEndpointServicePermission                           = findVPCEndpointServicePermission
	FindVPCEndpointSubnetAssociationExists                     = findVPCEndpointSubnetAssociationExists
	FindVPCIPv6CIDRBlockAssociationByIDV2                      = findVPCIPv6CIDRBlockAssociationByID
	FindVPNConnectionByID                                      = findVPNConnectionByID
	FindVPNConnectionRouteByTwoPartKey                         = findVPNConnectionRouteByTwoPartKey
	FindVPNGatewayByID                                         = findVPNGatewayByID
	FindVPNGatewayRoutePropagationExists                       = findVPNGatewayRoutePropagationExists
	FindVPNGatewayVPCAttachmentByTwoPartKey                    = findVPNGatewayVPCAttachmentByTwoPartKey
	FindVerifiedAccessEndpointByID                             = findVerifiedAccessEndpointByID
	FindVerifiedAccessGroupByID                                = findVerifiedAccessGroupByID
	FindVerifiedAccessInstanceByID                             = findVerifiedAccessInstanceByID
	FindVerifiedAccessInstanceLoggingConfigurationByInstanceID = findVerifiedAccessInstanceLoggingConfigurationByInstanceID
	FindVerifiedAccessInstanceTrustProviderAttachmentExists    = findVerifiedAccessInstanceTrustProviderAttachmentExists
	FindVerifiedAccessTrustProviderByID                        = findVerifiedAccessTrustProviderByID
	FindVolumeAttachmentInstanceByID                           = findVolumeAttachmentInstanceByID
	FlattenNetworkInterfacePrivateIPAddresses                  = flattenNetworkInterfacePrivateIPAddresses
	IPAMServicePrincipal                                       = ipamServicePrincipal
	NewAttributeFilterList                                     = newAttributeFilterList
	NewAttributeFilterListV2                                   = newAttributeFilterListV2
	NewCustomFilterList                                        = newCustomFilterList
	NewTagFilterList                                           = newTagFilterList
	ProtocolForValue                                           = protocolForValue
	StopEBSVolumeAttachmentInstance                            = stopVolumeAttachmentInstance
	StopInstance                                               = stopInstance
	UpdateTags                                                 = updateTags
	UpdateTagsV2                                               = updateTagsV2
	WaitVolumeAttachmentCreated                                = waitVolumeAttachmentCreated
)

type (
	IPProtocol = ipProtocol
)
