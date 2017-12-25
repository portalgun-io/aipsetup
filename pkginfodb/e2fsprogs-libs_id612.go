package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_e2fsprogs_libs = &basictypes.PackageInfo{

	Description: `subset of e2fsprogs package`,
	HomePage:    "https://sourceforge.net/projects/e2fsprogs",

	BuilderName: "e2fsprogs",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     true,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_hosted:e2fsprogs"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "e2fsprogs-libs",
	TarballFileNameParser: "std",
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		`e2fsprogs`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}