package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	// "github.com/AnimusPEXUS/aipsetup/buildercollection"
	// "github.com/AnimusPEXUS/aipsetup/versiontools"
)

var DistroPackageInfo_gnunet_java = &basictypes.PackageInfo{
	Description: ``,
	HomePage:    "http://www.gnunet.org",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     true,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	TarballFileNameParser: "std",
	TarballName:           "gnunet-java",

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
