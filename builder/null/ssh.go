// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package null

import (
	"github.com/dumb-hashicorp/dumb-packer-plugin-sdk/multistep"
)

func CommHost(host string) func(multistep.StateBag) (string, error) {
	return func(state multistep.StateBag) (string, error) {
		return host, nil
	}
}
