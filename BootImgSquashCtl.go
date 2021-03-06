package aipsetup

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/shadowusers"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

type BootImgSquashCtl struct {
	src_root    string
	wd_path     string
	os_files    string
	squashed_fs string
	log         *logger.Logger
}

func NewBootImgSquashCtl(src_root string, wd_path string, log *logger.Logger) (*BootImgSquashCtl, error) {

	self := new(BootImgSquashCtl)

	self.src_root = src_root

	wd_path, err := filepath.Abs(wd_path)
	if err != nil {
		return nil, err
	}

	self.wd_path = wd_path
	self.os_files = path.Join(wd_path, "osfiles")
	self.squashed_fs = path.Join(wd_path, "root.squash")
	self.log = log

	return self, nil
}

func (self *BootImgSquashCtl) CopyOSFiles() error {

	{
		root_files_to_copy := []string{
			"bin", "sbin", "lib", "lib64", "usr",
			"var", "etc", "daemons", "multihost",
		}

		{
			root_files, err := ioutil.ReadDir(self.src_root)
			if err != nil {
				return err
			}
			for _, i := range root_files {
				for _, j := range []string{"etc.", "var."} {
					if strings.HasPrefix(i.Name(), j) {
						root_files_to_copy = append(root_files_to_copy, i.Name())
					}
				}
			}
		}

		for _, i := range root_files_to_copy {
			err := filetools.CopyTree(
				path.Join(self.src_root, i),
				path.Join(self.os_files, i),
				false,
				true,
				true,
				true,
				self.log,
				true,
				true,
				func(f, t string, log logger.LoggerI) error {
					fstat, err := os.Lstat(f)
					if err != nil {
						return err
					}

					if !fstat.Mode().IsRegular() && !filetools.Is(fstat.Mode()).Symlink() {
						log.Error("skipping irregular file " + f)
						return nil
					}

					err = filetools.CopyWithInfo(f, t, log)
					if err != nil {
						return err
					}

					return nil
				},
			)
			if err != nil {
				return err
			}
		}
	}

	for _, i := range []string{
		"mnt", "run", "tmp", "root", "dev",
		"proc", "sys", "overlay", "overlay_merged", "boot", "root_old",
	} {
		err := os.MkdirAll(path.Join(self.os_files, i), 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BootImgSquashCtl) InstallAipSetup() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	err = filetools.CopyWithInfo(
		exe,
		path.Join(self.os_files, "bin", "aipsetup5"),
		self.log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgSquashCtl) RemoveUsers() error {

	susers := shadowusers.NewCtl(path.Join(self.os_files, "etc"))

	err := susers.ReadAll()
	if err != nil {
		return err
	}

	simple_users := make([]int, 0)
	simple_users_names := make([]string, 0)

	for _, i := range susers.Passwds.Passwds {
		if i.UserId > basictypes.SYS_UID_MAX {
			simple_users = append(simple_users, i.UserId)
			simple_users_names = append(simple_users_names, i.Login)
		}
	}

	for i := len(susers.Passwds.Passwds) - 1; i != -1; i-- {
		if susers.Passwds.Passwds[i].UserId > basictypes.SYS_UID_MAX {
			susers.Passwds.Passwds = append(
				susers.Passwds.Passwds[:i],
				susers.Passwds.Passwds[i+1:]...,
			)
		}
	}

	for i := len(susers.Groups.Groups) - 1; i != -1; i-- {
		del := false

		for _, j := range simple_users {
			if susers.Groups.Groups[i].GID == j {
				del = true
				break
			}
		}

		if del {
			susers.Groups.Groups = append(
				susers.Groups.Groups[:i],
				susers.Groups.Groups[i+1:]...,
			)
		}
	}

	for i := len(susers.Shadows.Shadows) - 1; i != -1; i-- {
		del := false

		for _, j := range simple_users_names {
			if susers.Shadows.Shadows[i].Login == j {
				del = true
				break
			}
		}

		if del {
			susers.Shadows.Shadows = append(
				susers.Shadows.Shadows[:i],
				susers.Shadows.Shadows[i+1:]...,
			)
		}
	}

	for i := len(susers.GShadows.GShadows) - 1; i != -1; i-- {
		del := false

		for _, j := range simple_users_names {
			if susers.GShadows.GShadows[i].Name == j {
				del = true
				break
			}
		}

		if del {
			susers.GShadows.GShadows = append(
				susers.GShadows.GShadows[:i],
				susers.GShadows.GShadows[i+1:]...,
			)
		}
	}

	err = susers.WriteAll()
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgSquashCtl) ResetRootPasswd() error {
	susers := shadowusers.NewCtl(path.Join(self.os_files, "etc"))

	err := susers.ReadAll()
	if err != nil {
		return err
	}

	r, err := susers.Shadows.GetByLogin("root")
	if err != nil {
		return err
	}

	r.Password =
		`$6$cLewkeecW2A4.SvS$t0ckgHfOu8jPtPPDelZwH6binT7sLigIhyz0Pou6Kz9lb.Xy/qMWA6bNvWTOfSGwNsssTTWiKc2bmo10GEjVP.`

	err = susers.WriteAll()
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgSquashCtl) CleanupOSFS() error {

	targets_to_remove := []string{
		"/etc/passwd-",
		"/etc/shadow-",
		"/etc/groups-",
		"/etc/gshadow-",
		"/etc/machine-id",
		"/etc/dhcpcd.secret",
		"/etc/dhcpcd.secret",
		"/etc/tor",
	}

	for _, i := range targets_to_remove {
		err := os.RemoveAll(path.Join(self.os_files, i))
		if err != nil {
			return err
		}
	}

	{
		sd_journals_pth := path.Join(self.os_files, "var", "log", "journal")

		err := os.MkdirAll(sd_journals_pth, 0700)
		if err != nil {
			return err
		}

		sd_journals_pth_files, err := ioutil.ReadDir(sd_journals_pth)
		if err != nil {
			return err
		}

		for _, i := range sd_journals_pth_files {
			err = os.RemoveAll(i.Name())
			if err != nil {
				return err
			}
		}
	}

	{
		etc_ssh := path.Join(self.os_files, "etc", "ssh")
		etc_ssh_files, err := ioutil.ReadDir(etc_ssh)
		if err != nil {
			return err
		}

		for _, i := range etc_ssh_files {
			if strings.HasPrefix(i.Name(), "ssh_host") {
				err = os.Remove(path.Join(etc_ssh, i.Name()))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (self *BootImgSquashCtl) CleanupLinuxSrc() error {
	pth := path.Join(self.os_files, "usr", "src", "linux")
	c := exec.Command("make", "clean")
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = pth
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgSquashCtl) InstallOverlayInit() error {
	overlay_init_file := path.Join(self.os_files, "overlay_init.sh")
	overlay_init_sh := `#!/bin/bash

set -x

umount /root_old

mount -t tmpfs tmpfs /overlay

mkdir /overlay/upper
mkdir /overlay/work

mount -t overlay overlay -olowerdir=/,upperdir=/overlay/upper,workdir=/overlay/work /overlay_merged

mount --move /boot /overlay_merged/boot

exec chroot /overlay_merged /bin/bash
`
	err := ioutil.WriteFile(
		overlay_init_file,
		[]byte(overlay_init_sh),
		0700,
	)
	if err != nil {
		return err
	}

	err = os.Chmod(overlay_init_file, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgSquashCtl) DoEverythingBeforeSquash() error {
	for _, i := range [](func() error){
		self.CopyOSFiles,
		self.InstallAipSetup,
		self.RemoveUsers,
		self.ResetRootPasswd,
		self.CleanupOSFS,
		self.CleanupLinuxSrc,
		self.InstallOverlayInit,
	} {
		err := i()
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BootImgSquashCtl) SquashOSFiles() error {
	os.Remove(self.squashed_fs)
	c := exec.Command("mksquashfs", self.os_files, self.squashed_fs, "-comp", "xz")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
