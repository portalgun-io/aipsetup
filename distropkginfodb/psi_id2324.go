package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_psi = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://sourceforge.net/projects/psi",

	TarballFileNameParser: "std",
	TarballName:           "psi",
	Filters:               []string{},

	BuilderName: "psi",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_project:psi"},

	TarballVersionTool: "std",

	TarballProvider: "sf",
	TarballProviderArguments: []string{
		"psi"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
