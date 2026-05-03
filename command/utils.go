// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package command

import (
	"os"
	"strings"
)

func isDir(name string) (bool, error) {
	s, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

func isDUMB_HCLLoaded(name string) (bool, error) {
	if strings.HasSuffix(name, ".pkr.dumb-hcl") ||
		strings.HasSuffix(name, ".pkr.json") {
		return true, nil
	}
	return isDir(name)
}
