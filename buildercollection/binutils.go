package buildercollection

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["binutils"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderBinutils(bs), nil
	}
}

type BuilderBinutils struct {
	bs basictypes.BuildingSiteCtlI

	std_builder *BuilderStdAutotools
}

func NewBuilderBinutils(bs basictypes.BuildingSiteCtlI) *BuilderBinutils {

	self := new(BuilderBinutils)
	self.bs = bs

	self.std_builder = NewBuilderStdAutotools(bs)

	self.std_builder.AfterExtractCB = self.AfterExtract
	self.std_builder.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *BuilderBinutils) DefineActions() (basictypes.BuilderActions, error) {
	return self.std_builder.DefineActions()
}

func (self *BuilderBinutils) AfterExtract(log *logger.Logger, err error) error {

	if err != nil {
		return err
	}

	a_tools := new(buildingtools.Autotools)
	tar_dir := self.bs.GetDIR_TARBALL()
	files, err := ioutil.ReadDir(tar_dir)
	if err != nil {
		return err
	}

	for _, i := range []string{
		"gmp", "mpc", "mpfr", "isl", "cloog",
	} {
		filename := ""
		for _, j := range files {
			b := path.Base(j.Name())
			if strings.HasPrefix(b, i) {
				filename = b
				break
			}
		}
		if filename == "" {
			return errors.New("not found tarball for " + i)
		}

		err = a_tools.Extract(
			path.Join(tar_dir, filename),
			path.Join(self.bs.GetDIR_SOURCE(), i),
			self.bs.GetDIR_TEMP(),
			true,
			false,
			true,
			log,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BuilderBinutils) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	// cb, err := calc.CalculateIsCrossbuilder()
	// if err != nil {
	// 	return nil, err
	// }

	if info.ThisIsCrossbuilder {

		host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
		if err != nil {
			return nil, err
		}

		prefix := path.Join(host_builders_dir, info.CrossbuilderTarget)

		hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
		if err != nil {
			return nil, err
		}

		ret = make([]string, 0)
		ret = append(
			ret,
			[]string{
				"--prefix=" + prefix,
				"--mandir=" + path.Join(prefix, basictypes.DIRNAME_SHARE, "man"),
				"--sysconfdir=/etc",
				"--localstatedir=/var",
				"--enable-shared",
			}...,
		)
		ret = append(
			ret,
			hbt_opts...,
		)

	}

	host_dir, err := calc.CalculateHostDir()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{

			"--enable-targets=all",

			"--enable-64-bit-bfd",
			"--disable-werror",
			"--enable-libada",
			"--enable-libssp",
			"--enable-objc-gc",

			"--enable-lto",
			"--enable-ld",

			// # NOTE: no google software in Lailalo
			"--disable-gold",
			"--without-gold",

			// # this is required. else libs will be searched in /lib and
			// # /usr/lib, but not in /multihost/xxx/lib!:
			"--with-sysroot=" + host_dir,

			// # more experiment:
			"--enable-multiarch",
			"--enable-multilib",
		}...,
	)

	if info.ThisIsCrossbuilder {
		ret = append(ret, "--with-sysroot")
	}

	return ret, nil
}
