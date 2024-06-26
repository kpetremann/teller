//go:build !windows
// +build !windows

package appdir

import (
	"os"
	"path/filepath"
)

// UserConfig returns the users config dir
func UserConfig() string {
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		return filepath.Join(hd, ".config", Name)
	}

	base := os.Getenv("XDG_CONFIG_HOME")
	if base == "" {
		base = filepath.Join(os.Getenv("HOME"), ".config")
	}

	return filepath.Join(base, Name)
}

// UserCache returns the users cache dir
func UserCache() string {
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		return filepath.Join(hd, ".cache", Name)
	}

	base := os.Getenv("XDG_CACHE_HOME")
	if base == "" {
		base = filepath.Join(os.Getenv("HOME"), ".cache")
	}

	return filepath.Join(base, Name)
}

// UserData returns the users data dir
func UserData() string {
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		return filepath.Join(hd, ".local", "share", Name)
	}

	base := os.Getenv("XDG_DATA_HOME")
	if base == "" {
		base = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}

	return filepath.Join(base, Name)
}
