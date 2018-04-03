package basictypes

type BuildingSiteInfo struct {
	SystemTitle   string `json:"system_title"`
	SystemVersion string `json:"system_version"`

	// System which going to run this package
	Host     string `json:"host"`
	HostArch string `json:"hostarch"`

	// This package is crosscompiler to build other packages.
	// This package is going to run under Host
	// CrossbuilderTarget is System which going to run packages built by this crosscompiler
	ThisIsCrossbuilder bool   `json:"this_is_crossbuilder"`
	CrossbuilderTarget string `json:"crossbuilder_target"`

	// This package is crossbuilding
	// This package is being built to run under Host
	// This package uses CrossbuildersHost's crossbuilder to be built
	ThisIsCrossbuilding bool   `json:"this_is_crossbuilding"`
	CrossbuildersHost   string `json:"crossbuilder_s_host"`

	ThisIsSubarchBuilding bool `json:"this_is_subarchbuilding"`

	// RunningByHost is host which started this building site
	RunningByHost bool `json:"this_bs_starter_host"`

	PackageName      string `json:"package_name"`
	PackageVersion   string `json:"package_version"`
	PackageStatus    string `json:"package_status"`
	PackageTimestamp string `json:"package_timestamp"`

	Sources []string `json:"sources"`

	// MainTarballInfo *PackageInfo `json:"main_tarball_info"`
}

func (self *BuildingSiteInfo) SetInfoLailalo50() {
	self.SystemTitle = "LAILALO"
	self.SystemVersion = "5.0"
}