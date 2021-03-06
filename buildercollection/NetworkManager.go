package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["NetworkManager"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_NetworkManager(bs)
	}
}

type Builder_NetworkManager struct {
	*Builder_std
}

func NewBuilder_NetworkManager(bs basictypes.BuildingSiteCtlI) (*Builder_NetworkManager, error) {
	self := new(Builder_NetworkManager)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	// NOTE: python version of NetworkManager had this option enabled, so
	//       I live it here for now in disabled state
	//	self.EditBuildConcurentJobsCountCB = func(log *logger.Logger, ret int) int {
	//		return 1
	//	}

	return self, nil
}

func (self *Builder_NetworkManager) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	pkgconfig, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	if err != nil {
		return nil, err
	}

	nss_cflags, err := pkgconfig.CommandOutput("--cflags", "nss", "nspr")
	if err != nil {
		return nil, err
	}

	nss_libs, err := pkgconfig.CommandOutput("--libs", "nss", "nspr")
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"NSS_CFLAGS=" + nss_cflags,
			"NSS_LDFLAGS=" + nss_libs,

			"--without-libnm-glib",
			"--with-suspend-resume=systemd",
			"--with-session-tracking=systemd",
		}...,
	)

	return ret, nil
}
