package horizon

import (
	"encoding/json"
	"testing"

	"github.com/fonero-project/fonero-horizon/resource"
	"github.com/fonero-project/fonero-horizon/test"
)

func TestRootAction(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	server := test.NewStaticMockServer(`{
			"info": {
				"network": "test",
				"build": "test-core",
				"protocol_version": 4
			}
		}`)
	defer server.Close()

	ht.App.horizonVersion = "test-horizon"
	ht.App.config.FoneroCoreURL = server.URL

	w := ht.Get("/")
	if ht.Assert.Equal(200, w.Code) {
		var actual resource.Root
		err := json.Unmarshal(w.Body.Bytes(), &actual)
		ht.Require.NoError(err)
		ht.Assert.Equal("test-horizon", actual.HorizonVersion)
		ht.Assert.Equal("test-core", actual.FoneroCoreVersion)
	}
}
