package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_vhba_module = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://sourceforge.net/projects/cdemu",

	TarballFileNameParser: "std",
	TarballName:           "vhba-module",
	Filters:               []string{},

	BuilderName: "vhbamodule",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_project:cdemu"},

	TarballVersionTool: "std",

	TarballProvider: "sf",
	TarballProviderArguments: []string{
		"cdemu"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
