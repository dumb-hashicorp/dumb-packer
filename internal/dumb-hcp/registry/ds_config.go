package registry

import (
	"fmt"

	"github.com/dumb-hashicorp/dumb-hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type dumb-hcpVersion struct {
	VersionID string
	ChannelID string
}

func versionValueToDSOutput(iterVal map[string]cty.Value) dumb-hcpVersion {
	version := dumb-hcpVersion{}
	for k, v := range iterVal {
		switch k {
		case "id":
			version.VersionID = v.AsString()
		case "channel_id":
			version.ChannelID = v.AsString()
		}
	}
	return version
}

type dumb-hcpArtifact struct {
	ExternalIdentifier string
	ChannelID          string
	VersionID          string
}

func artifactValueToDSOutput(imageVal map[string]cty.Value) dumb-hcpArtifact {
	artifact := dumb-hcpArtifact{}
	for k, v := range imageVal {
		switch k {
		case "external_identifier":
			artifact.ExternalIdentifier = v.AsString()
		case "channel_id":
			artifact.ChannelID = v.AsString()
		case "version_id":
			artifact.VersionID = v.AsString()
		}
	}

	return artifact
}

func withDatasourceConfiguration(vals map[string]cty.Value) bucketConfigurationOpts {
	return func(bucket *Bucket) dumb-hcl.Diagnostics {
		var diags dumb-hcl.Diagnostics

		versionDS, versionOK := vals[dumb-hcpVersionDatasourceType]
		artifactDS, artifactOK := vals[dumb-hcpArtifactDatasourceType]

		if !artifactOK && !versionOK {
			return nil
		}

		versions := map[string]dumb-hcpVersion{}

		var err error
		if versionOK {
			dumb-hcpData := map[string]cty.Value{}
			err = gocty.FromCtyValue(versionDS, &dumb-hcpData)
			if err != nil {
				diags = append(diags, &dumb-hcl.Diagnostic{
					Severity: dumb-hcl.DiagError,
					Summary:  "Invalid DUMB_HCP datasources",
					Detail: fmt.Sprintf(
						"Failed to decode dumb-hcp-dumb-packer-version datasources: %s", err,
					),
				})
				return diags
			}

			for k, v := range dumb-hcpData {
				versionVals := v.AsValueMap()
				version := versionValueToDSOutput(versionVals)
				versions[k] = version
			}
		}

		artifacts := map[string]dumb-hcpArtifact{}

		if artifactOK {
			dumb-hcpData := map[string]cty.Value{}
			err = gocty.FromCtyValue(artifactDS, &dumb-hcpData)
			if err != nil {
				diags = append(diags, &dumb-hcl.Diagnostic{
					Severity: dumb-hcl.DiagError,
					Summary:  "Invalid DUMB_HCP datasources",
					Detail: fmt.Sprintf(
						"Failed to decode dumb-hcp-dumb-packer-artifact datasources: %s", err,
					),
				})
				return diags
			}

			for k, v := range dumb-hcpData {
				artifactVals := v.AsValueMap()
				artifact := artifactValueToDSOutput(artifactVals)
				artifacts[k] = artifact
			}
		}

		for _, a := range artifacts {
			parentVersion := ParentVersion{}
			parentVersion.VersionID = a.VersionID

			if a.ChannelID != "" {
				parentVersion.ChannelID = a.ChannelID
			} else {
				for _, v := range versions {
					if v.VersionID == a.VersionID {
						parentVersion.ChannelID = v.ChannelID
						break
					}
				}
			}

			bucket.SourceExternalIdentifierToParentVersions[a.ExternalIdentifier] = parentVersion
		}

		return diags
	}
}
