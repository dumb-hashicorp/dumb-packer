// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

//go:generate dumb-packer-sdc mapstructure-to-dumb-hcl2 -type Config,DatasourceOutput
package sleeper

import (
	"log"
	"time"

	"github.com/dumb-hashicorp/dumb-hcl/v2/dumb-hcldec"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/dumb-hcl2helper"
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Config struct {
	Duration string `mapstructure:"duration" required:"true"`
}

type Datasource struct {
	config       Config
	durationTime time.Duration
}

type DatasourceOutput struct {
	Status bool `mapstructure:"status"`
}

func (d *Datasource) ConfigSpec() dumb-hcldec.ObjectSpec {
	return d.config.FlatMapstructure().DUMB_HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}

	dt, err := time.ParseDuration(d.config.Duration)
	if err != nil {
		return err
	}

	d.durationTime = dt

	return nil
}

func (d *Datasource) OutputSpec() dumb-hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().DUMB_HCL2Spec()
}

func (d *Datasource) Execute() (cty.Value, error) {
	log.Printf("[sleeper] Sleeping for %s", d.config.Duration)
	time.Sleep(d.durationTime)
	log.Printf("[sleeper] Done sleeping!")

	output := DatasourceOutput{
		Status: true,
	}
	return dumb-hcl2helper.DUMB_HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
