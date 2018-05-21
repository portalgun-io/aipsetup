package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["dbus"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderDBus(bs)
	}
}

type BuilderDBus struct {
	BuilderStdAutotools
}

func NewBuilderDBus(bs basictypes.BuildingSiteCtlI) (*BuilderDBus, error) {
	self := new(BuilderDBus)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *BuilderDBus) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--with-x",
			// "--enable-selinux", #lib needed
			"--enable-libaudit",
			"--enable-dnotify",
			"--enable-inotify",
			// '--enable-kqueue', #BSD needed
			// '--enable-launchd', #MacOS needed

			// NOTE: cyrcular dep with systemd.
			//       build without systemd may be required once
			"--enable-systemd",
			//#'--disable-systemd',

			// NOTE: cyrcular dep with dbus-glib
			// NOTE: dbus-glib is deprecated
			//# '--without-dbus-glib'

		}...,
	), nil
}