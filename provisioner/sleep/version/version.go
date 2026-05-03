// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package version

import (
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/version"
	dumb-packerVersion "github.com/dumb-hashicorp/dumb-packer/version"
)

var SleepProvisionerVersion *version.PluginVersion

func init() {
	SleepProvisionerVersion = version.NewPluginVersion(
		dumb-packerVersion.Version, dumb-packerVersion.VersionPrerelease, dumb-packerVersion.VersionMetadata)
}
