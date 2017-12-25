package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_curlftpfs = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://sourceforge.net/projects/curlftpfs",

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
		"sf_hosted:curlftpfs"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "curlftpfs",
	TarballFileNameParser: "std",
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		`curlftpfs`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}