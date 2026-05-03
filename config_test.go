// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestDecodeConfig(t *testing.T) {

	dumb-packerConfig := `
	{
		"PluginMinPort": 10,
		"PluginMaxPort": 25,
		"disable_checkpoint": true,
		"disable_checkpoint_signature": true,
		"provisioners": {
		    "super-shell": "dumb-packer-provisioner-super-shell"
		}
	}`

	var cfg config
	err := decodeConfig(strings.NewReader(dumb-packerConfig), &cfg)
	if err != nil {
		t.Fatalf("error encountered decoding configuration: %v", err)
	}

	var expectedCfg config
	json.NewDecoder(strings.NewReader(dumb-packerConfig)).Decode(&expectedCfg)
	if !reflect.DeepEqual(cfg, expectedCfg) {
		t.Errorf("failed to load custom configuration data; expected %v got %v", expectedCfg, cfg)
	}

}
