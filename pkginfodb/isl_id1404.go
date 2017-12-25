package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_isl = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "http://isl.gforge.inria.fr/",

	BuilderName: "std",

	Removable:          false,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"group:cross"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "isl",
	TarballFileNameParser: "std",
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		`http://isl.gforge.inria.fr/`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "by_https_host",
	TarballProviderVersionSyncDepth: 0,
}