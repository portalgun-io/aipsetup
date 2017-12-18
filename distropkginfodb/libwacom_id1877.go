package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_libwacom = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://sourceforge.net/projects/linuxwacom",

	TarballFileNameParser: "std",
	TarballName:           "libwacom",
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
		"'sf_project:linuxwacom", "gnome_project", "group:gnome"},

	TarballVersionTool: "gnome",

	TarballProvider: "sf",
	TarballProviderArguments: []string{
		"linuxwacom"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "gnome",
	TarballProviderVersionSyncDepth: 0,
}
