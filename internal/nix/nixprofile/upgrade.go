// Copyright 2024 Jetify Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package nixprofile

import (
	"os"

	"go.jetify.com/devbox/internal/devpkg"
	"go.jetify.com/devbox/internal/lock"
	"go.jetify.com/devbox/internal/nix"
)

func ProfileUpgrade(ProfileDir string, pkg *devpkg.Package, lock *lock.File) error {
	nameOrIndex, err := ProfileListNameOrIndex(
		&ProfileListNameOrIndexArgs{
			Lockfile:   lock,
			Writer:     os.Stderr,
			Package:    pkg,
			ProfileDir: ProfileDir,
		},
	)
	if err != nil {
		return err
	}

	return nix.ProfileUpgrade(ProfileDir, nameOrIndex)
}
