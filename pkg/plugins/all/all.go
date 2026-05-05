// Package all imports all RedMap plugins to trigger their init() registration.
// Import this package to load all available plugins into the registry.
package all

import (
	// CIDR plugins
	_ "github.com/praetorian-inc/redmap/pkg/plugins/cidrs"
	// Domain plugins
	_ "github.com/praetorian-inc/redmap/pkg/plugins/domains"
)
