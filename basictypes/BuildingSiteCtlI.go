package basictypes

type BuildingSiteCtlI interface {
	GetDIR_TARBALL() string
	GetDIR_SOURCE() string
	GetDIR_PATCHES() string
	GetDIR_BUILDING() string
	GetDIR_DESTDIR() string
	GetDIR_BUILD_LOGS() string
	GetDIR_LISTS() string
	GetDIR_TEMP() string

	ReadInfo() (*BuildingSiteInfo, error)
	WriteInfo(*BuildingSiteInfo) error

	DetermineMainTarrball() (string, error)

	GetSystem() SystemI
	GetBuildingSiteValuesCalculator() BuildingSiteValuesCalculatorI

	GetSources() error
	GetTarballs() error
	GetPatches() error

	Run(targets []string) error

	PrePackager() PrePackagerI
	Packager() PackagerI
}
