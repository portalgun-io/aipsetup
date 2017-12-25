package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_astyle = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://sourceforge.net/projects/astyle",

	BuilderName: "astyle",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_hosted:astyle"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "astyle",
	TarballFileNameParser: "std",
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		`astyle`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
