package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_miniupnpd = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://github.com/miniupnp/miniupnp",

	BuilderName: "miniupnpd",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"github_hosted", "group:core1"},

	TarballVersionTool: "std",

	TarballName:           "miniupnpd",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/miniupnp/miniupnp.git`, `miniupnpc`, `TagName:miniupnpd`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
