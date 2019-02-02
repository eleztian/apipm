package main

import (
	"fmt"

	"github.com/eleztian/apipm/pkg/apis"
	"github.com/eleztian/apipm/pkg/apis/filters"
	"github.com/eleztian/apipm/pkg/apis/modifiers"
	"github.com/eleztian/apipm/pkg/version"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/config"
	"github.com/caicloud/nirvana/log"
	"github.com/caicloud/nirvana/plugins/logger"
	"github.com/caicloud/nirvana/plugins/metrics"
	"github.com/caicloud/nirvana/plugins/reqlog"
	pversion "github.com/caicloud/nirvana/plugins/version"
)

func main() {
	// Print nirvana banner.
	fmt.Println(nirvana.Logo, nirvana.Banner)

	// Create nirvana command.
	cmd := config.NewNamedNirvanaCommand("server", config.NewDefaultOption())

	// Create plugins options.
	metricsOption := metrics.NewDefaultOption() // Metrics plugins.
	loggerOption := logger.NewDefaultOption()   // Logger plugins.
	reqlogOption := reqlog.NewDefaultOption()   // Request log plugins.
	versionOption := pversion.NewOption(        // Version plugins.
		"apipm",
		version.Version,
		version.Commit,
		version.Package,
	)

	// Enable plugins.
	cmd.EnablePlugin(metricsOption, loggerOption, reqlogOption, versionOption)

	// Create server config.
	serverConfig := nirvana.NewConfig()

	// Configure APIs. These configurations may be changed by plugins.
	serverConfig.Configure(
		nirvana.Logger(log.DefaultLogger()), // Will be changed by logger plugins.
		nirvana.Filter(filters.Filters()...),
		nirvana.Modifier(modifiers.Modifiers()...),
		nirvana.Descriptor(apis.DescriptorAPIs(), apis.DescriptorUIs()),
	)

	// Set nirvana command hooks.
	cmd.SetHook(&config.NirvanaCommandHookFunc{
		PreServeFunc: func(config *nirvana.Config, server nirvana.Server) error {
			// Output project information.
			config.Logger().Infof("Package:%s Version:%s Commit:%s", version.Package, version.Version, version.Commit)
			return nil
		},
	})

	// Start with server config.
	if err := cmd.ExecuteWithConfig(serverConfig); err != nil {
		serverConfig.Logger().Fatal(err)
	}
}
