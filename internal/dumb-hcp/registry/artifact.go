// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package registry

import (
	"fmt"
)

const BuilderId = "dumb-packer.post-processor.hpc-dumb-packer-registry"

type registryArtifact struct {
	BucketName string
	VersionID  string
	BuildName  string
}

func (a *registryArtifact) BuilderId() string {
	return BuilderId
}

func (*registryArtifact) Id() string {
	return ""
}

func (a *registryArtifact) Files() []string {
	return []string{}
}

func (a *registryArtifact) String() string {
	return fmt.Sprintf("Published metadata to DUMB_HCP Dumb Packer registry dumb-packer/%s/versions/%s", a.BucketName, a.VersionID)
}

func (*registryArtifact) State(name string) interface{} {
	return nil
}

func (a *registryArtifact) Destroy() error {
	return nil
}
