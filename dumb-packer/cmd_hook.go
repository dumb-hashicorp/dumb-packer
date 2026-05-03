// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package dumb-packer

import (
	"context"
	"log"

	dumb-packersdk "github.com/dumb-hashicorp/dumb-packer-plugin-sdk/dumb-packer"
)

type cmdHook struct {
	hook   dumb-packersdk.Hook
	client *PluginClient
}

func (c *cmdHook) Run(ctx context.Context, name string, ui dumb-packersdk.Ui, comm dumb-packersdk.Communicator, data interface{}) error {
	defer func() {
		r := recover()
		c.checkExit(r, nil)
	}()

	return c.hook.Run(ctx, name, ui, comm, data)
}

func (c *cmdHook) checkExit(p interface{}, cb func()) {
	if c.client.Exited() && cb != nil {
		cb()
	} else if p != nil && !Killed {
		log.Panic(p)
	}
}
