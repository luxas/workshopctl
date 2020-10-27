package cmd

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/cloud-native-nordics/workshopctl/pkg/config"
	"github.com/cloud-native-nordics/workshopctl/pkg/constants"
	"github.com/cloud-native-nordics/workshopctl/pkg/gen"
	"github.com/cloud-native-nordics/workshopctl/pkg/provider/providers"
	"github.com/cloud-native-nordics/workshopctl/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type GenFlags struct {
	*RootFlags

	SkipLocalCharts bool
}

// NewGenCommand returns the "gen" command
func NewGenCommand(rf *RootFlags) *cobra.Command {
	gf := &GenFlags{
		RootFlags:       rf,
		SkipLocalCharts: false,
	}
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate a set of manifests based on the configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if err := RunGen(gf); err != nil {
				log.Fatal(err)
			}
		},
	}

	addGenFlags(cmd.Flags(), gf)
	return cmd
}

func addGenFlags(fs *pflag.FlagSet, gf *GenFlags) {
	fs.BoolVar(&gf.SkipLocalCharts, "skip-local-charts", gf.SkipLocalCharts, "Don't consider the local directory's charts/ directory")
}

func loadConfig(configPath string) (*config.Config, error) {
	cfg := &config.Config{}
	if err := util.ReadYAMLFile(configPath, cfg); err != nil {
		return nil, err
	}
	if err := initConfig(cfg); err != nil {
		return nil, err
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func RunGen(gf *GenFlags) error {
	ctx := context.Background()
	cfg, err := loadConfig(gf.ConfigPath)
	if err != nil {
		return err
	}

	charts, err := gen.SetupInternalChartCache(cfg.RootDir)
	if err != nil {
		return err
	}

	if !gf.SkipLocalCharts {
		chartsDir := filepath.Join(cfg.RootDir, constants.ChartsDir)
		chartInfos, err := ioutil.ReadDir(chartsDir)
		if err != nil {
			return err
		}
		for _, chartInfo := range chartInfos {
			if !chartInfo.IsDir() {
				continue
			}
			chart, err := gen.SetupExternalChartCache(cfg.RootDir, chartInfo.Name())
			if err != nil {
				return err
			}
			charts = append(charts, chart)
		}
	}

	// dry-run can be always true here as we're not gonna use the provider for requests
	dnsProvider, err := providers.DNSProviders().NewDNSProvider(ctx, cfg.DNSProvider, cfg.RootDomain, true)
	if err != nil {
		return err
	}

	return config.ForCluster(ctx, cfg.Clusters, cfg, func(clusterCtx context.Context, clusterInfo *config.ClusterInfo) error {
		for _, chart := range charts {
			logger := util.Logger(ctx)
			logger.Infof("Generating chart %q...", chart.Name)
			if err := gen.GenerateChart(clusterCtx, chart, clusterInfo, dnsProvider.ValuesProcessors(), dnsProvider.ChartProcessors()); err != nil {
				return err
			}
		}
		return nil
	})
}
