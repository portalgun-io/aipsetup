package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	// "github.com/AnimusPEXUS/aipsetup/buildercollection"
	// "github.com/AnimusPEXUS/aipsetup/versiontools"
)

var DistroPackageInfo_xf86_input_tek4957 = &basictypes.PackageInfo{
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
	TarballName:           "xf86-input-tek4957",

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
