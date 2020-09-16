package test

import (
	env "istio.io/proxy/test/envoye2e/env"
)

var ExtensionE2ETests *env.TestInventory

func init() {
	ExtensionE2ETests = &env.TestInventory{
		Tests: []string{
			"TestHelloWorldExtension",
		},
	}
}
