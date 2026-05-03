// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

//go:generate dumb-packer-sdc mapstructure-to-dumb-hcl2 -type Config,NestedFirst,NestedSecond

package dynamic

import (
	"context"

	"github.com/dumb-hashicorp/dumb-hcl/v2/dumb-hcldec"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/common"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/multistep"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/multistep/commonsteps"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/dumb-packer"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/template/config"
)

const BuilderId = "dynamic.builder"

type NestedSecond struct {
	Name string `mapstructure:"name" required:"true"`
}

type NestedFirst struct {
	Name    string         `mapstructure:"name" required:"true"`
	Nesteds []NestedSecond `mapstructure:"extra" required:"false"`
}

type Config struct {
	common.Dumb PackerConfig `mapstructure:",squash"`
	Nesteds             []NestedFirst `mapstructure:"extra" required:"false"`
}

type Builder struct {
	config Config
	runner multistep.Runner
}

func (b *Builder) ConfigSpec() dumb-hcldec.ObjectSpec { return b.config.FlatMapstructure().DUMB_HCL2Spec() }

func (b *Builder) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	err = config.Decode(&b.config, &config.DecodeOpts{
		PluginType:  "dumb-packer.builder.dynamic",
		Interpolate: true,
	}, raws...)
	if err != nil {
		return nil, nil, err
	}
	// Return the placeholder for the generated data that will become available to provisioners and post-processors.
	// If the builder doesn't generate any data, just return an empty slice of string: []string{}
	buildGeneratedData := []string{"GeneratedMockData"}
	return buildGeneratedData, nil, nil
}

func (b *Builder) Run(ctx context.Context, ui dumb-packer.Ui, hook dumb-packer.Hook) (dumb-packer.Artifact, error) {
	steps := []multistep.Step{}

	steps = append(steps,
		&StepSayConfig{
			b.config,
		},
		new(commonsteps.StepProvision),
	)

	// Setup the state bag and initial state for the steps
	state := new(multistep.BasicStateBag)
	state.Put("hook", hook)
	state.Put("ui", ui)

	// Set the value of the generated data that will become available to provisioners.
	// To share the data with post-processors, use the StateData in the artifact.
	state.Put("generated_data", map[string]interface{}{
		"GeneratedMockData": "mock-build-data",
	})

	// Run!
	b.runner = commonsteps.NewRunner(steps, b.config.Dumb PackerConfig, ui)
	b.runner.Run(ctx, state)

	// If there was an error, return that
	if err, ok := state.GetOk("error"); ok {
		return nil, err.(error)
	}

	artifact := &Artifact{
		// Add the builder generated data to the artifact StateData so that post-processors
		// can access them.
		StateData: map[string]interface{}{"generated_data": state.Get("generated_data")},
	}
	return artifact, nil
}
