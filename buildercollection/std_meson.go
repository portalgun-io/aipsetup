package buildercollection

import (
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_meson"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_meson(bs)
	}
}

type Builder_std_meson struct {
	meson string
	ninja string

	builder_std *Builder_std

	// TODO: it's possible, what action editing possibility is not neeeded at all for
	//       meson based projects. but I'm not sure yet.. maybe after-distribution
	//       actions will find use
	EditActionsCB       func(ret basictypes.BuilderActions) (basictypes.BuilderActions, error)
	EditConfigureEnvCB  func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditConfigureArgsCB func(log *logger.Logger, ret []string) ([]string, error)
	EditBuildEnvCB      func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditBuildArgsCB     func(log *logger.Logger, ret []string) ([]string, error)
}

func NewBuilder_std_meson(bs basictypes.BuildingSiteCtlI) (*Builder_std_meson, error) {

	self := new(Builder_std_meson)

	self.builder_std = NewBuilder_std(bs)

	if t, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("meson"); err != nil {
		return nil, err
	} else {
		self.meson = t
	}

	if t, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("ninja"); err != nil {
		return nil, err
	} else {
		self.ninja = t
	}

	return self, nil
}

func (self *Builder_std_meson) GetBuildingSiteCtl() basictypes.BuildingSiteCtlI {
	return self.builder_std.GetBuildingSiteCtl()
}

func (self *Builder_std_meson) DefineActions() (basictypes.BuilderActions, error) {

	ret := basictypes.BuilderActions{

		&basictypes.BuilderAction{"dst_cleanup", self.builder_std.BuilderActionDstCleanup},
		&basictypes.BuilderAction{"src_cleanup", self.builder_std.BuilderActionSrcCleanup},
		&basictypes.BuilderAction{"bld_cleanup", self.builder_std.BuilderActionBldCleanup},
		&basictypes.BuilderAction{"extract", self.builder_std.BuilderActionExtract},
		&basictypes.BuilderAction{"configure", self.BuilderActionConfigure},
		&basictypes.BuilderAction{"build", self.BuilderActionBuild},
		&basictypes.BuilderAction{"distribute", self.BuilderActionDistribute},
		&basictypes.BuilderAction{"prepack", self.builder_std.BuilderActionPrePack},
		&basictypes.BuilderAction{"pack", self.builder_std.BuilderActionPack},
	}

	if self.EditActionsCB != nil {
		var err error
		ret, err = self.EditActionsCB(ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_std_meson) BuilderActionConfigureEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {

	ret, err := self.builder_std.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return nil, err
	}

	if self.EditConfigureEnvCB != nil {
		ret, err = self.EditConfigureEnvCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil

}

func (self *Builder_std_meson) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {

	ret, err := self.builder_std.BuilderActionConfigureArgsDef(log)
	if err != nil {
		return nil, err
	}

	if self.EditConfigureArgsCB != nil {
		ret, err = self.EditConfigureArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_std_meson) BuilderActionConfigure(log *logger.Logger) error {

	env := environ.New()

	nenv, err := self.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return err
	}

	env.UpdateWith(nenv)

	args, err := self.BuilderActionConfigureArgsDef(log)
	if err != nil {
		return err
	}

	args2 := make([]string, 0)
	args2 = append(args2, []string{"../01.SOURCE"}...)
	args2 = append(args2, args...)

	c := exec.Command(self.meson, args2...)
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_BUILDING()
	c.Env = env.Strings()

	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_std_meson) BuilderActionBuildEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {
	log.Info(
		"this builder uses same environment variables for make as for configure",
	)

	ret, err := self.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return nil, err
	}

	if self.EditBuildEnvCB != nil {

		ret, err = self.EditBuildEnvCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_std_meson) BuilderActionBuildArgsDef(
	log *logger.Logger,
) ([]string, error) {
	ret := make([]string, 0)

	if self.EditBuildArgsCB != nil {
		var err error
		ret, err = self.EditBuildArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_std_meson) BuilderActionBuild(log *logger.Logger) error {

	env := environ.New()

	nenv, err := self.BuilderActionBuildEnvDef(log)
	if err != nil {
		return err
	}

	env.UpdateWith(nenv)

	args, err := self.BuilderActionBuildArgsDef(log)
	if err != nil {
		return err
	}

	args2 := make([]string, 0)
	args2 = append(args2, []string{"../01.SOURCE"}...)
	args2 = append(args2, args...)

	c := exec.Command(self.ninja, "build")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_BUILDING()
	c.Env = env.Strings()

	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_std_meson) BuilderActionDistribute(log *logger.Logger) error {

	c := exec.Command(
		self.ninja,
		"install",
		"DESTDIR="+self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
	)
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_BUILDING()

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
