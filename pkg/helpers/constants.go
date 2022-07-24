package helpers

const (
	AccountUsername  = "rhn-support-tiwillia"
	DefaultAWSRegion = "us-east-1"

	// Cluster Auhtorization
	AWSComputeNodeResourceType = "compute.node.aws"
	MinClusterCount            = 4
	MaxClusterCount            = 10
	StandardBillingModel       = "standard"
	OsdProductID               = "osd"
	AWSCloudProvider           = "aws"
	ClusterAuthAccountUsername = "rh-perfscale"
	ClusterAuthManaged         = true
	ClusterAuthReserve         = true
	ClusterAuthBYOC            = true
	SingleAvailabilityZone     = "single"
)

var (
	AWSResources = [5]string{"m3.2xlarge", "m4.large", "m4.xlarge", "m5.large", "m5.xlarge"}
)
