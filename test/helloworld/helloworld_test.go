package helloworld

import (
	"path/filepath"
	"testing"
	"time"

	"istio.io/proxy/test/envoye2e/driver"
	"istio.io/proxy/test/envoye2e/env"
	"istio.io/proxy/testdata"

	"github.com/istio-ecosystem/wasm-extensions/test"
)

func TestHelloWorldExtension(t *testing.T) {
	params := driver.NewTestParams(t, map[string]string{
		"HelloWorldWasmFile": filepath.Join(env.GetBazelBin(), "extensions/helloworld/helloworld.wasm"),
	}, test.ExtensionE2ETests)
	params.Vars["ServerHTTPFilters"] = params.LoadTestData("testdata/helloworld/server_filter.yaml.tmpl")
	if err := (&driver.Scenario{
		Steps: []driver.Step{
			&driver.XDS{},
			&driver.Update{
				Node: "server", Version: "0", Listeners: []string{string(testdata.MustAsset("listener/server.yaml.tmpl"))},
			},
			&driver.Envoy{
				Bootstrap:       params.FillTestData(string(testdata.MustAsset("bootstrap/server.yaml.tmpl"))),
				DownloadVersion: "1.7.0",
			},
			&driver.Sleep{1 * time.Second},
			&driver.HTTPCall{
				Port: params.Ports.ServerPort,
				Body: "hello, world!",
				ResponseHeaders: map[string]string{
					"hello": "world",
				},
			},
		},
	}).Run(params); err != nil {
		t.Fatal(err)
	}
}
