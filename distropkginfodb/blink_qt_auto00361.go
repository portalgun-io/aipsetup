package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	// "github.com/AnimusPEXUS/aipsetup/buildercollection"
	// "github.com/AnimusPEXUS/aipsetup/versiontools"
)

var DistroPackageInfo_blink_qt = &basictypes.PackageInfo{
	Description: `write something here, please`,
	HomePage:    "http://None",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	TarballFileNameParser: "std",
	TarballName:           "blink-qt",

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
