package tarballrepository

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/cachepresets"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/providers"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
)

type Repository struct {
	sys basictypes.SystemI
}

func NewRepository(sys basictypes.SystemI) (*Repository, error) {
	ret := new(Repository)
	ret.sys = sys
	return ret, nil
}

func (self *Repository) GetRepositoryPath() string {
	return self.sys.GetTarballRepoRootDir()
}

func (self *Repository) GetCachesDir() string {
	return path.Join(self.GetRepositoryPath(), "cache")
}

func (self *Repository) GetPackagePath(name string) string {
	return path.Join(self.GetRepositoryPath(), "packages", name)
}

func (self *Repository) GetPackageTarballsPath(name string) string {
	return path.Join(self.GetPackagePath(name), "tarballs")
}

func (self *Repository) GetPackageCachePath(name string) string {
	return path.Join(self.GetCachesDir(), name)
}

func (self *Repository) CreateCacheObjectForPackage(name string) (
	*cache01.CacheDir,
	error,
) {
	info, err := distropkginfodb.Get(name)
	if err != nil {
		return nil, err
	}

	var preset *cache01.Settings

	if i, err :=
		cachepresets.Get(info.TarballProviderCachePresetName); err != nil {
		return nil, err
	} else {
		preset = i
	}

	return cache01.NewCacheDir(self.GetPackageCachePath(name), preset)
}

func (self *Repository) PerformPackageTarballsUpdate(name string) error {

	info, err := distropkginfodb.Get(name)
	if err != nil {
		return err
	}

	cache, err := self.CreateCacheObjectForPackage(name)
	if err != nil {
		return err
	}

	prov, err := providers.Get(
		info.TarballProvider,
		self,
		name,
		info,
		self.sys,
		self.GetPackageTarballsPath(name),
		cache,
	)
	if err != nil {
		return err
	}

	err = prov.PerformUpdate()
	if err != nil {
		return err
	}

	return nil
}

func (self *Repository) ListLocalTarballs(package_name string) ([]string, error) {
	ret := make([]string, 0)

	res, err := self.ListLocalFiles(package_name)
	if err != nil {
		return ret, err
	}

	info, err := distropkginfodb.Get(package_name)
	if err != nil {
		return ret, err
	}

	parser, err := tarballnameparsers.Get(info.TarballFileNameParser)
	if err != nil {
		return ret, err
	}

	for _, i := range res {
		if parse_res, err := parser.ParseName(i); err != nil {
			continue
		} else {
			if parse_res.Name != info.TarballName {
				continue
			}
		}
		full_out_path_done := self.GetTarballDoneFilePath(package_name, i)
		_, err := os.Stat(full_out_path_done)
		if err != nil {
			continue
		}

		ret = append(ret, i)
	}

	return ret, nil
}

func (self *Repository) ListLocalFiles(package_name string) ([]string, error) {
	ret := make([]string, 0)

	pth := self.GetPackageTarballsPath(package_name)

	files, err := ioutil.ReadDir(pth)
	if err != nil {
		return ret, err
	}

	for _, i := range files {
		if !i.IsDir() {
			ret = append(ret, i.Name())
		}
	}

	return ret, nil
}

func (self *Repository) GetTarballDoneFilePath(
	package_name string,
	as_filename string,
) string {
	as_filename = path.Base(as_filename)
	return path.Join(
		self.GetPackageTarballsPath(package_name),
		as_filename,
	) + ".done"
}

func (self *Repository) PerformDownload(
	package_name string,
	as_filename string,
	uri string,
) error {
	as_filename = path.Base(as_filename)
	tarballs_dir := self.GetPackageTarballsPath(package_name)

	full_out_path := path.Join(tarballs_dir, as_filename)
	full_out_path_done := self.GetTarballDoneFilePath(package_name, as_filename)

	_, err := os.Stat(full_out_path_done)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(path.Dir(full_out_path), 0700)
	if err != nil {
		return err
	}
	c := exec.Command("wget", "--progress=dot", "-c", "-O", full_out_path, uri)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	ret := c.Run()
	if ret == nil {
		f, err := os.Create(full_out_path_done)
		if err != nil {
			return err
		}
		f.Close()
	}
	return ret
}

func (self *Repository) DeleteFile(
	package_name string,
	filename string,
) error {
	tarballs_dir := self.GetPackageTarballsPath(package_name)
	filename = path.Base(filename)
	full_path := path.Join(tarballs_dir, filename)
	return os.Remove(full_path)
}