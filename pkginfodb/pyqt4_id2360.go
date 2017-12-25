package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_pyqt4 = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://sourceforge.net/projects/pyqt",

	BuilderName: "pyqt",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_hosted:pyqt"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "PyQt-x11-gpl",
	TarballFileNameParser: "std",
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		`pyqt`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}