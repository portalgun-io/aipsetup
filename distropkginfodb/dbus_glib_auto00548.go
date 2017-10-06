package distropkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_dbus_glib = &aipsetup.CompletePackageInfo{
  OveralPackageInfo: aipsetup.OveralPackageInfo{
    Description: `dbus-glib is a deprecated API for use of D-Bus from GLib applications. Do not use it in new code.
Since version 2.26, GLib's accompanying GIO library provides a high-level API for D-Bus, "GDBus", based on an independent reimplementation of the D-Bus protocol. The maintainers of D-Bus recommend that GLib applications should use GDBus instead of dbus-glib.

but some applications still using it (NetworkManager)`,
    HomePage: "http://www.freedesktop.org",

    Removable: true,
    Reducible: true,
    NonInstallable: false,
    Deprecated: false,
    PrimaryInstallOnly: false,

    BuildDeps   : []string{},
    SODeps      : []string{},
    RunTimeDeps : []string{},
  },

  TarballPackageInfo: aipsetup.TarballPackageInfo{
    Name : "dbus-glib",
    VersionTool: "std", //versiontools.Standard,
  },

  BuildingPackageInfo: aipsetup.BuildingPackageInfo{
    BuilderName : "std", //buildercollection.Builder_std,
  },


}


