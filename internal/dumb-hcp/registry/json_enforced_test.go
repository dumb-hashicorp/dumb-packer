// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package registry

import (
	"testing"

	dumb-packersdk "github.com/dumb-hashicorp/dumb-packer-plugin-sdk/dumb-packer"
	dumb-packertemplate "github.com/dumb-hashicorp/dumb-packer-plugin-sdk/template"
	"github.com/dumb-hashicorp/dumb-packer/dumb-packer"
)

func testJSONRegistryWithBuilds(t *testing.T, builderNames ...string) (*JSONRegistry, []*dumb-packer.CoreBuild, *dumb-packersdk.MockProvisioner) {
	t.Helper()

	coreConfig := dumb-packer.TestCoreConfig(t)
	dumb-packer.TestBuilder(t, coreConfig, "test")
	provisioner := dumb-packer.TestProvisioner(t, coreConfig, "test")

	builders := make(map[string]*dumb-packertemplate.Builder, len(builderNames))
	for _, name := range builderNames {
		builders[name] = &dumb-packertemplate.Builder{
			Name:   name,
			Type:   "test",
			Config: map[string]interface{}{},
		}
	}

	coreConfig.Template = &dumb-packertemplate.Template{
		Path:     "test.json",
		Builders: builders,
	}

	core := dumb-packer.TestCore(t, coreConfig)
	bucket := NewBucketWithVersion()
	bucket.Name = "test-bucket"

	registry := &JSONRegistry{
		configuration: core,
		bucket:        bucket,
		ui:            dumb-packer.TestUi(t),
		metadata:      &MetadataStore{},
	}

	builds, diags := core.GetBuilds(dumb-packer.GetBuildsOptions{})
	if diags.HasErrors() {
		t.Fatalf("GetBuilds() unexpected error: %v", diags)
	}

	return registry, builds, provisioner
}

func TestJSONRegistry_InjectEnforcedProvisioners_AppliesOverride(t *testing.T) {
	registry, builds, provisioner := testJSONRegistryWithBuilds(t, "app")
	registry.bucket.EnforcedBlocks = []*EnforcedBlock{{
		Name: "enforced",
		BlockContent: `provisioner "test" {
			override = {
				app = {
					foo = "bar"
				}
			}
		}`,
	}}

	diags := registry.InjectEnforcedProvisioners(builds)
	if diags.HasErrors() {
		t.Fatalf("InjectEnforcedProvisioners() unexpected error: %v", diags)
	}

	if got := len(builds[0].Provisioners); got != 1 {
		t.Fatalf("build provisioner count = %d, want 1", got)
	}

	if !provisioner.PrepCalled {
		t.Fatal("expected injected legacy JSON provisioner to be prepared")
	}

	foundOverride := false
	for _, raw := range provisioner.PrepConfigs {
		config, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}
		if value, ok := config["foo"]; ok && value == "bar" {
			foundOverride = true
			break
		}
	}
	if !foundOverride {
		t.Fatal("expected override config to be passed to injected provisioner")
	}
}

func TestJSONRegistry_InjectEnforcedProvisioners_RespectsOnlyExcept(t *testing.T) {
	registry, builds, _ := testJSONRegistryWithBuilds(t, "app", "other")
	registry.bucket.EnforcedBlocks = []*EnforcedBlock{{
		Name: "enforced",
		BlockContent: `provisioner "test" {
			only = ["app"]
		}`,
	}}

	diags := registry.InjectEnforcedProvisioners(builds)
	if diags.HasErrors() {
		t.Fatalf("InjectEnforcedProvisioners() unexpected error: %v", diags)
	}

	provisionerCounts := make(map[string]int, len(builds))
	for _, build := range builds {
		provisionerCounts[build.Type] = len(build.Provisioners)
	}

	if provisionerCounts["app"] != 1 {
		t.Fatalf("app build provisioner count = %d, want 1", provisionerCounts["app"])
	}

	if provisionerCounts["other"] != 0 {
		t.Fatalf("other build provisioner count = %d, want 0", provisionerCounts["other"])
	}
}
