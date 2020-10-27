package provider

import (
	"context"
	"net"
	"net/url"
	"time"

	"github.com/cloud-native-nordics/workshopctl/pkg/config"
	"github.com/cloud-native-nordics/workshopctl/pkg/gen"
)

type Cluster struct {
	Spec   ClusterSpec
	Status ClusterStatus
}

type ClusterSpec struct {
	Name       string
	Version    string
	NodeGroups []config.NodeGroup
}

type ClusterStatus struct {
	ID              string
	Index           config.ClusterNumber
	ProvisionStart  *time.Time
	ProvisionDone   *time.Time
	EndpointURL     *url.URL
	EndpointIP      net.IP
	KubeconfigBytes []byte
}

type CloudProviderFactory interface {
	NewCloudProvider(ctx context.Context, p *config.Provider, dryrun bool) (CloudProvider, error)
}

type CloudProvider interface {
	// CreateCluster creates a cluster. This call is _blocking_ until the cluster is properly provisioned
	CreateCluster(ctx context.Context, index config.ClusterNumber, c ClusterSpec) (*Cluster, error)
	// DeleteCluster deletes a cluster and its associated load balancers
	DeleteCluster(ctx context.Context, index config.ClusterNumber) error
}

type DNSProviderFactory interface {
	NewDNSProvider(ctx context.Context, p *config.Provider, rootDomain string, dryrun bool) (DNSProvider, error)
}

type DNSProvider interface {
	ChartProcessors() []gen.Processor
	ValuesProcessors() []gen.Processor
	// CleanupRecords deletes records associated with a cluster
	CleanupRecords(ctx context.Context, index config.ClusterNumber) error
}
