package google

import "github.com/snyk/driftctl/pkg/resource"

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initGoogleStorageBucketMetadata(resourceSchemaRepository)
	initGoogleComputeFirewallMetadata(resourceSchemaRepository)
	initGoogleComputeRouterMetadata(resourceSchemaRepository)
	initGoogleComputeNetworkMetadata(resourceSchemaRepository)
	initGoogleStorageBucketIamBMemberMetadata(resourceSchemaRepository)
	initGoogleComputeInstanceGroupMetadata(resourceSchemaRepository)
	initGoogleBigqueryDatasetMetadata(resourceSchemaRepository)
	initGoogleBigqueryTableMetadata(resourceSchemaRepository)
	initGoogleProjectIAMMemberMetadata(resourceSchemaRepository)
	initGoogleComputeAddressMetadata(resourceSchemaRepository)
	initGoogleComputeGlobalAddressMetadata(resourceSchemaRepository)
	initGoogleComputeSubnetworkMetadata(resourceSchemaRepository)
	initGoogleComputeDiskMetadata(resourceSchemaRepository)
	initGoogleComputeImageMetadata(resourceSchemaRepository)
	initGoogleComputeHealthCheckMetadata(resourceSchemaRepository)
	initComputeInstanceGroupManagerMetadata(resourceSchemaRepository)
}
