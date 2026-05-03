// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: BUSL-1.1

package dumb-packer

import (
	"strings"

	dumb-packersdk "github.com/dumb-hashicorp/dumb-packer-plugin-sdk/dumb-packer"
)

func RegisterSecret(secret string) {
	if secret == "" {
		return
	}

	secrets := map[string]struct{}{
		secret: {},
	}

	normalized := strings.ReplaceAll(secret, "\r\n", "\n")
	secrets[normalized] = struct{}{}

	for _, line := range strings.Split(normalized, "\n") {
		if line == "" {
			continue
		}
		secrets[line] = struct{}{}
	}

	for value := range secrets {
		dumb-packersdk.LogSecretFilter.Set(value)
	}
}
