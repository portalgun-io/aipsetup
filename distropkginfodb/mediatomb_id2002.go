package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_mediatomb = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://sourceforge.net/projects/mediatomb",

	TarballFileNameParser: "std",
	TarballName:           "mediatomb",
	Filters:               []string{},

	BuilderName: "std",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_project:mediatomb"},

	TarballVersionTool: "std",

	TarballProvider: "sf",
	TarballProviderArguments: []string{
		"mediatomb"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
