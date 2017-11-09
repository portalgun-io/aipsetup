package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	// "github.com/AnimusPEXUS/aipsetup/buildercollection"
	// "github.com/AnimusPEXUS/aipsetup/versiontools"
)

var DistroPackageInfo_lcms2 = &basictypes.PackageInfo{
	Description: `write something here, please`,
	HomePage:    "http://www.littlecms.com",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	TarballFileNameParser: "std",
	TarballName:           "lcms2",

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
