package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_ffmpeg = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://github.com/FFmpeg/FFmpeg",

	BuilderName: "ffmpeg",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"github_hosted"},

	TarballVersionTool: "std",

	TarballName:           "ffmpeg",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/FFmpeg/FFmpeg.git`, `ffmpeg`, `TagName:n`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
